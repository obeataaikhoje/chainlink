// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package ocr3

import (
    "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
    "github.com/smartcontractkit/chainlink-common/pkg/workflows"
    streams "github.com/smartcontractkit/chainlink-common/pkg/capabilities/triggers/streams"
)



func (cfg ConsensusConfig) New(w *workflows.WorkflowSpecFactory,ref string, input ConsensusInput)SignedReportCap {
    
    def := workflows.StepDefinition{
       ID: "offchain_reporting@1.0.0",Ref: ref,
       Inputs: input.ToSteps(),
       Config: map[string]any{
           "aggregation_config": cfg.AggregationConfig,
           "aggregation_method": cfg.AggregationMethod,
           "encoder": cfg.Encoder,
           "encoder_config": cfg.EncoderConfig,
           "report_id": cfg.ReportId,
       },
       CapabilityType: capabilities.CapabilityTypeConsensus,
   }


    step := workflows.Step[SignedReport]{Definition: def}
    return SignedReportCapFromStep(w, step)
}


type FeedValueCap interface {
    workflows.CapDefinition[FeedValue]
    Deviation() workflows.CapDefinition[string]
    Heartbeat() workflows.CapDefinition[int]
    RemappedID() workflows.CapDefinition[string]
    private()
}


// FeedValueCapFromStep should only be called from generated code to assure type safety
func FeedValueCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[FeedValue]) FeedValueCap {
    raw :=  step.AddTo(w)
    return &feedValue{CapDefinition: raw}
}


type feedValue struct {
    workflows.CapDefinition[FeedValue]
}

func (*feedValue) private() {}
func (c *feedValue) Deviation() workflows.CapDefinition[string] {
    return workflows.AccessField[FeedValue, string](c.CapDefinition, "Deviation")
}
func (c *feedValue) Heartbeat() workflows.CapDefinition[int] {
    return workflows.AccessField[FeedValue, int](c.CapDefinition, "Heartbeat")
}
func (c *feedValue) RemappedID() workflows.CapDefinition[string] {
    return workflows.AccessField[FeedValue, string](c.CapDefinition, "RemappedID")
}

func NewFeedValueFromFields(
                                                                        deviation workflows.CapDefinition[string],
                                                                        heartbeat workflows.CapDefinition[int],
                                                                        remappedID workflows.CapDefinition[string],) FeedValueCap {
    return &simpleFeedValue{
        CapDefinition: workflows.ComponentCapDefinition[FeedValue]{
        "deviation": deviation.Ref(),
        "heartbeat": heartbeat.Ref(),
        "remappedID": remappedID.Ref(),
        },
        deviation: deviation,
        heartbeat: heartbeat,
        remappedID: remappedID,
    }
}

type simpleFeedValue struct {
    workflows.CapDefinition[FeedValue]
    deviation workflows.CapDefinition[string]
    heartbeat workflows.CapDefinition[int]
    remappedID workflows.CapDefinition[string]
}
func (c *simpleFeedValue) Deviation() workflows.CapDefinition[string] {
    return c.deviation
}
func (c *simpleFeedValue) Heartbeat() workflows.CapDefinition[int] {
    return c.heartbeat
}
func (c *simpleFeedValue) RemappedID() workflows.CapDefinition[string] {
    return c.remappedID
}

func (c *simpleFeedValue) private() {}


type SignedReportCap interface {
    workflows.CapDefinition[SignedReport]
    Err() workflows.CapDefinition[bool]
    Value() SignedReportValueCap
    WorkflowExecutionID() workflows.CapDefinition[string]
    private()
}


// SignedReportCapFromStep should only be called from generated code to assure type safety
func SignedReportCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[SignedReport]) SignedReportCap {
    raw :=  step.AddTo(w)
    return &signedReport{CapDefinition: raw}
}


type signedReport struct {
    workflows.CapDefinition[SignedReport]
}

func (*signedReport) private() {}
func (c *signedReport) Err() workflows.CapDefinition[bool] {
    return workflows.AccessField[SignedReport, bool](c.CapDefinition, "Err")
}
func (c *signedReport) Value() SignedReportValueCap {
     return &signedReportValue{ CapDefinition: workflows.AccessField[SignedReport, SignedReportValue](c.CapDefinition, "Value")}
}
func (c *signedReport) WorkflowExecutionID() workflows.CapDefinition[string] {
    return workflows.AccessField[SignedReport, string](c.CapDefinition, "WorkflowExecutionID")
}

func NewSignedReportFromFields(
                                                                        err workflows.CapDefinition[bool],
                                                                        value SignedReportValueCap,
                                                                        workflowExecutionID workflows.CapDefinition[string],) SignedReportCap {
    return &simpleSignedReport{
        CapDefinition: workflows.ComponentCapDefinition[SignedReport]{
        "err": err.Ref(),
        "value": value.Ref(),
        "workflowExecutionID": workflowExecutionID.Ref(),
        },
        err: err,
        value: value,
        workflowExecutionID: workflowExecutionID,
    }
}

type simpleSignedReport struct {
    workflows.CapDefinition[SignedReport]
    err workflows.CapDefinition[bool]
    value SignedReportValueCap
    workflowExecutionID workflows.CapDefinition[string]
}
func (c *simpleSignedReport) Err() workflows.CapDefinition[bool] {
    return c.err
}
func (c *simpleSignedReport) Value() SignedReportValueCap {
    return c.value
}
func (c *simpleSignedReport) WorkflowExecutionID() workflows.CapDefinition[string] {
    return c.workflowExecutionID
}

func (c *simpleSignedReport) private() {}


type SignedReportValueCap interface {
    workflows.CapDefinition[SignedReportValue]
    Underlying() SignedReportValueUnderlyingCap
    private()
}


// SignedReportValueCapFromStep should only be called from generated code to assure type safety
func SignedReportValueCapFromStep(w *workflows.WorkflowSpecFactory, step workflows.Step[SignedReportValue]) SignedReportValueCap {
    raw :=  step.AddTo(w)
    return &signedReportValue{CapDefinition: raw}
}


type signedReportValue struct {
    workflows.CapDefinition[SignedReportValue]
}

func (*signedReportValue) private() {}
func (c *signedReportValue) Underlying() SignedReportValueUnderlyingCap {
     return SignedReportValueUnderlyingCap(workflows.AccessField[SignedReportValue, SignedReportValueUnderlying](c.CapDefinition, "Underlying"))
}

func NewSignedReportValueFromFields(
                                                                        underlying SignedReportValueUnderlyingCap,) SignedReportValueCap {
    return &simpleSignedReportValue{
        CapDefinition: workflows.ComponentCapDefinition[SignedReportValue]{
        "underlying": underlying.Ref(),
        },
        underlying: underlying,
    }
}

type simpleSignedReportValue struct {
    workflows.CapDefinition[SignedReportValue]
    underlying SignedReportValueUnderlyingCap
}
func (c *simpleSignedReportValue) Underlying() SignedReportValueUnderlyingCap {
    return c.underlying
}

func (c *simpleSignedReportValue) private() {}


type SignedReportValueUnderlyingCap workflows.CapDefinition[SignedReportValueUnderlying]


type ConsensusInput struct {
    Observations workflows.CapDefinition[[][]streams.Feed]
}

func (input ConsensusInput) ToSteps() workflows.StepInputs {
    return workflows.StepInputs{
       Mapping: map[string]any{
        "observations": input.Observations.Ref(),
       },
   }
}