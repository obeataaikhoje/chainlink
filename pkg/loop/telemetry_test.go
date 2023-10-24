package loop_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-relay/pkg/loop/internal/test"
)

func TestTelemetry(t *testing.T) {
	test.TelemetryClient(t)
	test.TelemetryServer(t)
}
