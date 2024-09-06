package testutils

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows"
)

type computeCapability struct {
	sdk      workflows.SDK
	callback func(sdk workflows.SDK, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error)
}

func (c *computeCapability) Info(ctx context.Context) (capabilities.CapabilityInfo, error) {
	info := capabilities.MustNewCapabilityInfo(
		"__internal__custom_compute@1.0.0", capabilities.CapabilityTypeAction, "Custom compute capability",
	)
	info.IsLocal = true
	return info, nil
}

func (c *computeCapability) RegisterToWorkflow(ctx context.Context, request capabilities.RegisterToWorkflowRequest) error {
	return nil
}

func (c *computeCapability) UnregisterFromWorkflow(ctx context.Context, request capabilities.UnregisterFromWorkflowRequest) error {
	return nil
}

func (c *computeCapability) Execute(ctx context.Context, request capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
	return c.callback(c.sdk, request)
}

var _ capabilities.ActionCapability = &computeCapability{}
