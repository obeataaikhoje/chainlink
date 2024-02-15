package ocr3

import (
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/test"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type HelperProcessCommand test.HelperProcessCommand

func (h *HelperProcessCommand) New() *exec.Cmd {
	h.CommandLocation = "../../internal/test/cmd/main.go"
	return (test.HelperProcessCommand)(*h).New()
}

func NewHelperProcessCommand(command string) *exec.Cmd {
	h := HelperProcessCommand{
		Command: command,
	}
	return h.New()
}

func TestLOOPPService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Plugin string
	}{
		// A generic plugin with a median provider
		{Plugin: test.OCR3ReportingPluginWithMedianProviderName},
		// A generic plugin with a plugin provider
		{Plugin: PluginServiceName},
	}
	for _, ts := range tests {
		looppSvc := NewLOOPPService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
			return NewHelperProcessCommand(ts.Plugin)
		}, types.ReportingPluginServiceConfig{}, test.MockConn{}, &test.StaticPipelineRunnerService{}, &test.StaticTelemetry{}, &test.StaticErrorLog{}, &test.StaticCapabilitiesRegistry{})
		hook := looppSvc.XXXTestHook()
		servicetest.Run(t, looppSvc)

		t.Run("control", func(t *testing.T) {
			test.OCR3ReportingPluginFactory(t, looppSvc)
		})

		t.Run("Kill", func(t *testing.T) {
			hook.Kill()

			// wait for relaunch
			time.Sleep(2 * internal.KeepAliveTickDuration)

			test.OCR3ReportingPluginFactory(t, looppSvc)
		})

		t.Run("Reset", func(t *testing.T) {
			hook.Reset()

			// wait for relaunch
			time.Sleep(2 * internal.KeepAliveTickDuration)

			test.OCR3ReportingPluginFactory(t, looppSvc)
		})
	}
}

func TestLOOPPService_recovery(t *testing.T) {
	t.Parallel()
	var limit atomic.Int32
	looppSvc := NewLOOPPService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
		h := HelperProcessCommand{
			Command: test.OCR3ReportingPluginWithMedianProviderName,
			Limit:   int(limit.Add(1)),
		}
		return h.New()
	}, types.ReportingPluginServiceConfig{}, test.MockConn{}, &test.StaticPipelineRunnerService{}, &test.StaticTelemetry{}, &test.StaticErrorLog{}, &test.StaticCapabilitiesRegistry{})
	servicetest.Run(t, looppSvc)

	test.OCR3ReportingPluginFactory(t, looppSvc)
}
