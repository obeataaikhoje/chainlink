package pb

import (
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

const (
	CapabilityTypeUnknown   = CapabilityType_CAPABILITY_TYPE_UNKNOWN
	CapabilityTypeTrigger   = CapabilityType_CAPABILITY_TYPE_TRIGGER
	CapabilityTypeAction    = CapabilityType_CAPABILITY_TYPE_ACTION
	CapabilityTypeConsensus = CapabilityType_CAPABILITY_TYPE_CONSENSUS
	CapabilityTypeTarget    = CapabilityType_CAPABILITY_TYPE_TARGET
)

func MarshalCapabilityRequest(req capabilities.CapabilityRequest) ([]byte, error) {
	return proto.MarshalOptions{Deterministic: true}.Marshal(CapabilityRequestToProto(req))
}

func MarshalCapabilityResponse(resp capabilities.CapabilityResponse) ([]byte, error) {
	return proto.MarshalOptions{Deterministic: true}.Marshal(CapabilityResponseToProto(resp))
}

func UnmarshalCapabilityRequest(raw []byte) (capabilities.CapabilityRequest, error) {
	var cr CapabilityRequest
	if err := proto.Unmarshal(raw, &cr); err != nil {
		return capabilities.CapabilityRequest{}, err
	}
	return CapabilityRequestFromProto(&cr), nil
}

func UnmarshalCapabilityResponse(raw []byte) (capabilities.CapabilityResponse, error) {
	var cr CapabilityResponse
	if err := proto.Unmarshal(raw, &cr); err != nil {
		return capabilities.CapabilityResponse{}, err
	}
	return CapabilityResponseFromProto(&cr), nil
}

func CapabilityRequestToProto(req capabilities.CapabilityRequest) *CapabilityRequest {
	inputs := &values.Map{Underlying: map[string]values.Value{}}
	if req.Inputs != nil {
		inputs = req.Inputs
	}
	config := &values.Map{Underlying: map[string]values.Value{}}
	if req.Config != nil {
		config = req.Config
	}
	return &CapabilityRequest{
		Metadata: &RequestMetadata{
			WorkflowId:          req.Metadata.WorkflowID,
			WorkflowExecutionId: req.Metadata.WorkflowExecutionID,
		},
		Inputs: values.Proto(inputs),
		Config: values.Proto(config),
	}
}

func TriggerRequestToProto(tr capabilities.TriggerRequest) *TriggerRequest {
	inputs := &values.Map{Underlying: map[string]values.Value{}}
	if req.Inputs != nil {
		inputs = req.Inputs
	}
	config := &values.Map{Underlying: map[string]values.Value{}}
	if req.Config != nil {
		config = req.Config
	}
	
	return &TriggerRequest{
		Id: tr.ID,
		Metadata: &RequestMetadata{
			WorkflowId:          tr.Metadata.WorkflowID,
			WorkflowExecutionId: tr.Metadata.WorkflowExecutionID,
		},
		Type:   tr.Type,
		Config: values.Proto(tr.Config),
		Inputs: values.Proto(tr.Inputs),
	}

}

func TriggerEventToProto(te capabilities.TriggerEvent) *TriggerEvent {
	return &TriggerEvent{
		Id:        te.ID,
		Timestamp: te.Timestamp,
		Payload:   values.Proto(te.Payload),
	}

}

func CapabilityResponseToProto(resp capabilities.CapabilityResponse) *CapabilityResponse {
	errStr := ""
	if resp.Err != nil {
		errStr = resp.Err.Error()
	}

	return &CapabilityResponse{
		Error: errStr,
		Value: values.Proto(resp.Value),
	}
}

func CapabilityRequestFromProto(pr *CapabilityRequest) capabilities.CapabilityRequest {
	md := pr.Metadata
	config := values.FromProto(pr.Config)
	inputs := values.FromProto(pr.Inputs)

	return capabilities.CapabilityRequest{
		Metadata: capabilities.RequestMetadata{
			WorkflowID:          md.WorkflowId,
			WorkflowExecutionID: md.WorkflowExecutionId,
		},
		Config: config.(*values.Map),
		Inputs: inputs.(*values.Map),
	}
}

func TriggerRequestFromProto(tr *TriggerRequest) capabilities.TriggerRequest {
	return capabilities.TriggerRequest{
		ID:   tr.Id,
		Type: tr.Type,
		Metadata: capabilities.RequestMetadata{
			WorkflowID:          tr.Metadata.WorkflowId,
			WorkflowExecutionID: tr.Metadata.WorkflowExecutionId,
		},
		Config: values.FromProto(tr.Config).(*values.Map),
		Inputs: values.FromProto(tr.Inputs).(*values.Map),
	}
}

func CapabilityResponseFromProto(pr *CapabilityResponse) capabilities.CapabilityResponse {
	val := values.FromProto(pr.Value)

	var err error
	if pr.Error != "" {
		err = errors.New(pr.Error)
	}
	return capabilities.CapabilityResponse{
		Value: val,
		Err:   err,
	}
}
