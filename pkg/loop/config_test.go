package loop

import (
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestEnvConfig_parse(t *testing.T) {
	cases := []struct {
		name                           string
		envVars                        map[string]string
		expectError                    bool
		expectedDatabaseURL            string
		expectedPrometheusPort         int
		expectedTracingEnabled         bool
		expectedTracingCollectorTarget string
		expectedTracingSamplingRatio   float64
		expectedTracingTLSCertPath     string
	}{
		{
			name: "All variables set correctly",
			envVars: map[string]string{
				envDatabaseURL:              "postgres://user:password@localhost:5432/db",
				envPromPort:                 "8080",
				envTracingEnabled:           "true",
				envTracingCollectorTarget:   "some:target",
				envTracingSamplingRatio:     "1.0",
				envTracingTLSCertPath:       "internal/test/fixtures/client.pem",
				envTracingAttribute + "XYZ": "value",
			},
			expectError:                    false,
			expectedDatabaseURL:            "postgres://user:password@localhost:5432/db",
			expectedPrometheusPort:         8080,
			expectedTracingEnabled:         true,
			expectedTracingCollectorTarget: "some:target",
			expectedTracingSamplingRatio:   1.0,
			expectedTracingTLSCertPath:     "internal/test/fixtures/client.pem",
		},
		{
			name: "CL_DATABASE_URL parse error",
			envVars: map[string]string{
				envDatabaseURL: "wrong-db-url",
			},
			expectError: true,
		},
		{
			name: "CL_PROMETHEUS_PORT parse error",
			envVars: map[string]string{
				envPromPort: "abc",
			},
			expectError: true,
		},
		{
			name: "TRACING_ENABLED parse error",
			envVars: map[string]string{
				envPromPort:       "8080",
				envTracingEnabled: "invalid_bool",
			},
			expectError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			var config EnvConfig
			err := config.parse()

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else {
					if config.DatabaseURL.String() != tc.expectedDatabaseURL {
						t.Errorf("Expected Database URL %s, got %s", tc.expectedDatabaseURL, config.DatabaseURL)
					}
					if config.PrometheusPort != tc.expectedPrometheusPort {
						t.Errorf("Expected Prometheus port %d, got %d", tc.expectedPrometheusPort, config.PrometheusPort)
					}
					if config.TracingEnabled != tc.expectedTracingEnabled {
						t.Errorf("Expected tracingEnabled %v, got %v", tc.expectedTracingEnabled, config.TracingEnabled)
					}
					if config.TracingCollectorTarget != tc.expectedTracingCollectorTarget {
						t.Errorf("Expected tracingCollectorTarget %s, got %s", tc.expectedTracingCollectorTarget, config.TracingCollectorTarget)
					}
					if config.TracingSamplingRatio != tc.expectedTracingSamplingRatio {
						t.Errorf("Expected tracingSamplingRatio %f, got %f", tc.expectedTracingSamplingRatio, config.TracingSamplingRatio)
					}
					if config.TracingTLSCertPath != tc.expectedTracingTLSCertPath {
						t.Errorf("Expected tracingTLSCertPath %s, got %s", tc.expectedTracingTLSCertPath, config.TracingTLSCertPath)
					}
				}
			}
		})
	}
}

func TestEnvConfig_AsCmdEnv(t *testing.T) {
	envCfg := EnvConfig{
		DatabaseURL:    &url.URL{Scheme: "postgres", Host: "localhost:5432", User: url.UserPassword("user", "password"), Path: "/db"},
		PrometheusPort: 9090,

		TracingEnabled:         true,
		TracingCollectorTarget: "http://localhost:9000",
		TracingSamplingRatio:   0.1,
		TracingTLSCertPath:     "some/path",
		TracingAttributes:      map[string]string{"key": "value"},

		TelemetryEnabled:            true,
		TelemetryEndpoint:           "example.com/beholder",
		TelemetryInsecureConnection: true,
		TelemetryCACertFile:         "foo/bar",
		TelemetryAttributes:         OtelAttributes{"foo": "bar", "baz": "42"},
		TelemetryTraceSampleRatio:   0.42,
	}
	got := map[string]string{}
	for _, kv := range envCfg.AsCmdEnv() {
		pair := strings.SplitN(kv, "=", 2)
		require.Len(t, pair, 2)
		got[pair[0]] = pair[1]
	}

	assert.Equal(t, "postgres://user:password@localhost:5432/db", got[envDatabaseURL])
	assert.Equal(t, strconv.Itoa(9090), got[envPromPort])

	assert.Equal(t, "true", got[envTracingEnabled])
	assert.Equal(t, "http://localhost:9000", got[envTracingCollectorTarget])
	assert.Equal(t, "0.1", got[envTracingSamplingRatio])
	assert.Equal(t, "some/path", got[envTracingTLSCertPath])
	assert.Equal(t, "value", got[envTracingAttribute+"key"])

	assert.Equal(t, "true", got[envTelemetryEnabled])
	assert.Equal(t, "example.com/beholder", got[envTelemetryEndpoint])
	assert.Equal(t, "true", got[envTelemetryInsecureConn])
	assert.Equal(t, "foo/bar", got[envTelemetryCACertFile])
	assert.Equal(t, "0.42", got[envTelemetryTraceSampleRatio])
	assert.Equal(t, "bar", got[envTelemetryAttribute+"foo"])
	assert.Equal(t, "42", got[envTelemetryAttribute+"baz"])
}

func TestManagedGRPCClientConfig(t *testing.T) {
	t.Parallel()

	t.Run("returns a new client config with the provided broker config", func(t *testing.T) {
		t.Parallel()

		brokerConfig := BrokerConfig{
			Logger: logger.Test(t),
			GRPCOpts: GRPCOpts{
				DialOpts: []grpc.DialOption{
					grpc.WithNoProxy(), // any grpc.DialOption will do
				},
			},
		}

		clientConfig := ManagedGRPCClientConfig(&plugin.ClientConfig{}, brokerConfig)

		assert.NotNil(t, clientConfig.Logger)
		assert.Equal(t, []plugin.Protocol{plugin.ProtocolGRPC}, clientConfig.AllowedProtocols)
		assert.Equal(t, brokerConfig.GRPCOpts.DialOpts, clientConfig.GRPCDialOptions)
		assert.True(t, clientConfig.Managed)
	})
}
