// Code generated by pkg/capabilities/cli, DO NOT EDIT.

package ocr3cap

import (
    "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
    "github.com/smartcontractkit/chainlink-common/pkg/workflows"

    "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3"
    streams "github.com/smartcontractkit/chainlink-common/pkg/capabilities/triggers/streams"
)

func NewConsensus(w *workflows.Workflow,ref string, input ConsensusInput, cfg ocr3.ConsensusConfig)Consensus {
    def := workflows.StepDefinition{
       ID: "offchain_reporting@1.0.0",Ref: ref,
       Inputs: workflows.StepInputs{
           Mapping: map[string]any{
               "observations": input.Observations.Ref(),
           },
       },
       Config: map[string]any{
           "aggregation_config": cfg.AggregationConfig,
           "aggregation_method": cfg.AggregationMethod,
           "encoder": cfg.Encoder,
           "encoder_config": cfg.EncoderConfig,
           "report_id": cfg.ReportId,
       },
       CapabilityType: capabilities.CapabilityTypeConsensus,
   }
    step := workflows.Step[ocr3.SignedReport]{Definition: def}
     raw := workflows.AddStep(w, step)
    return &consensus{CapDefinition: raw}
}


type FeedValue interface {
    workflows.CapDefinition[ocr3.FeedValue]
    Deviation() workflows.CapDefinition[string]
    Heartbeat() workflows.CapDefinition[int]
    RemappedID() workflows.CapDefinition[*string]
    private()
}

type feedValue struct {
    workflows.CapDefinition[ocr3.FeedValue]
}


func (*feedValue) private() {}
func (c *feedValue) Deviation() workflows.CapDefinition[string] {
    return workflows.AccessField[ocr3.FeedValue, string](c.CapDefinition, "Deviation")
}
func (c *feedValue) Heartbeat() workflows.CapDefinition[int] {
    return workflows.AccessField[ocr3.FeedValue, int](c.CapDefinition, "Heartbeat")
}
func (c *feedValue) RemappedID() workflows.CapDefinition[*string] {
    return workflows.AccessField[ocr3.FeedValue, *string](c.CapDefinition, "RemappedID")
}

func NewFeedValueFromFields(
                                                                        deviation workflows.CapDefinition[string],
                                                                        heartbeat workflows.CapDefinition[int],
                                                                        remappedID workflows.CapDefinition[*string],) FeedValue {
    return &simpleFeedValue{
        CapDefinition: workflows.ComponentCapDefinition[ocr3.FeedValue]{
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
    workflows.CapDefinition[ocr3.FeedValue]
    deviation workflows.CapDefinition[string]
    heartbeat workflows.CapDefinition[int]
    remappedID workflows.CapDefinition[*string]
}
func (c *simpleFeedValue) Deviation() workflows.CapDefinition[string] {
    return c.deviation
}
func (c *simpleFeedValue) Heartbeat() workflows.CapDefinition[int] {
    return c.heartbeat
}
func (c *simpleFeedValue) RemappedID() workflows.CapDefinition[*string] {
    return c.remappedID
}

func (c *simpleFeedValue) private() {}


type Consensus interface {
    workflows.CapDefinition[ocr3.SignedReport]
    Err() workflows.CapDefinition[bool]
    Value() SignedReportValue
    WorkflowExecutionID() workflows.CapDefinition[string]
    private()
}

type consensus struct {
    workflows.CapDefinition[ocr3.SignedReport]
}


func (*consensus) private() {}
func (c *consensus) Err() workflows.CapDefinition[bool] {
    return workflows.AccessField[ocr3.SignedReport, bool](c.CapDefinition, "Err")
}
func (c *consensus) Value() SignedReportValue {
     return &signedReportValue{ CapDefinition: workflows.AccessField[ocr3.SignedReport, ocr3.SignedReportValue](c.CapDefinition, "Value")}
}
func (c *consensus) WorkflowExecutionID() workflows.CapDefinition[string] {
    return workflows.AccessField[ocr3.SignedReport, string](c.CapDefinition, "WorkflowExecutionID")
}

func NewConsensusFromFields(
                                                                        err workflows.CapDefinition[bool],
                                                                        value SignedReportValue,
                                                                        workflowExecutionID workflows.CapDefinition[string],) Consensus {
    return &simpleConsensus{
        CapDefinition: workflows.ComponentCapDefinition[ocr3.SignedReport]{
        "err": err.Ref(),
        "value": value.Ref(),
        "workflowExecutionID": workflowExecutionID.Ref(),
        },
        err: err,
        value: value,
        workflowExecutionID: workflowExecutionID,
    }
}

type simpleConsensus struct {
    workflows.CapDefinition[ocr3.SignedReport]
    err workflows.CapDefinition[bool]
    value SignedReportValue
    workflowExecutionID workflows.CapDefinition[string]
}
func (c *simpleConsensus) Err() workflows.CapDefinition[bool] {
    return c.err
}
func (c *simpleConsensus) Value() SignedReportValue {
    return c.value
}
func (c *simpleConsensus) WorkflowExecutionID() workflows.CapDefinition[string] {
    return c.workflowExecutionID
}

func (c *simpleConsensus) private() {}


type SignedReportValue interface {
    workflows.CapDefinition[ocr3.SignedReportValue]
    Underlying() SignedReportValueUnderlying
    private()
}

type signedReportValue struct {
    workflows.CapDefinition[ocr3.SignedReportValue]
}


func (*signedReportValue) private() {}
func (c *signedReportValue) Underlying() SignedReportValueUnderlying {
     return SignedReportValueUnderlying(workflows.AccessField[ocr3.SignedReportValue, ocr3.SignedReportValueUnderlying](c.CapDefinition, "Underlying"))
}

func NewSignedReportValueFromFields(
                                                                        underlying SignedReportValueUnderlying,) SignedReportValue {
    return &simpleSignedReportValue{
        CapDefinition: workflows.ComponentCapDefinition[ocr3.SignedReportValue]{
        "underlying": underlying.Ref(),
        },
        underlying: underlying,
    }
}

type simpleSignedReportValue struct {
    workflows.CapDefinition[ocr3.SignedReportValue]
    underlying SignedReportValueUnderlying
}
func (c *simpleSignedReportValue) Underlying() SignedReportValueUnderlying {
    return c.underlying
}

func (c *simpleSignedReportValue) private() {}


type SignedReportValueUnderlying workflows.CapDefinition[ocr3.SignedReportValueUnderlying]


type ConsensusInput struct {
    Observations workflows.CapDefinition[[]streams.Feed]
}