package ocr3

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

type request struct {
	Observations values.Value `mapstructure:"-"`
	ExpiresAt    time.Time

	CallbackCh chan<- capabilities.CapabilityResponse
	CleanupCh  chan struct{}
	RequestCtx context.Context

	WorkflowExecutionID string
	WorkflowID          string
}

type response struct {
	Value               values.Value
	Err                 error
	WorkflowExecutionID string
}
