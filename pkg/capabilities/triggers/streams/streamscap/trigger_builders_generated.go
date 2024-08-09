// Code generated by pkg/capabilities/cli, DO NOT EDIT.

package streamscap

import (
    "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
    "github.com/smartcontractkit/chainlink-common/pkg/workflows"

    "github.com/smartcontractkit/chainlink-common/pkg/capabilities/triggers/streams"
)

func NewTrigger(w *workflows.Workflow,ref string, cfg streams.TriggerConfig)Trigger {
    def := workflows.StepDefinition{
       ID: "streams-trigger@1.0.0",Ref: ref,
       Inputs: workflows.StepInputs{
       },
       Config: map[string]any{
           "feedIds": cfg.FeedIds,
           "maxFrequencyMs": cfg.MaxFrequencyMs,
       },
       CapabilityType: capabilities.CapabilityTypeTrigger,
   }
    step := workflows.Step[streams.Feed]{Definition: def}
     raw := workflows.AddStep(w, step)
    return &trigger{CapDefinition: raw}
}


type Trigger interface {
    workflows.CapDefinition[streams.Feed]
    BenchmarkPrice() workflows.CapDefinition[string]
    FeedId() FeedId
    FullReport() workflows.CapDefinition[string]
    ObservationTimestamp() workflows.CapDefinition[int]
    ReportContext() workflows.CapDefinition[string]
    Signatures() workflows.CapDefinition[[]string]
    private()
}

type trigger struct {
    workflows.CapDefinition[streams.Feed]
}


func (*trigger) private() {}
func (c *trigger) BenchmarkPrice() workflows.CapDefinition[string] {
    return workflows.AccessField[streams.Feed, string](c.CapDefinition, "BenchmarkPrice")
}
func (c *trigger) FeedId() FeedId {
     return FeedId(workflows.AccessField[streams.Feed, streams.FeedId](c.CapDefinition, "FeedId"))
}
func (c *trigger) FullReport() workflows.CapDefinition[string] {
    return workflows.AccessField[streams.Feed, string](c.CapDefinition, "FullReport")
}
func (c *trigger) ObservationTimestamp() workflows.CapDefinition[int] {
    return workflows.AccessField[streams.Feed, int](c.CapDefinition, "ObservationTimestamp")
}
func (c *trigger) ReportContext() workflows.CapDefinition[string] {
    return workflows.AccessField[streams.Feed, string](c.CapDefinition, "ReportContext")
}
func (c *trigger) Signatures() workflows.CapDefinition[[]string] {
    return workflows.AccessField[streams.Feed, []string](c.CapDefinition, "Signatures")
}

func NewTriggerFromFields(
                                                                        benchmarkPrice workflows.CapDefinition[string],
                                                                        feedId FeedId,
                                                                        fullReport workflows.CapDefinition[string],
                                                                        observationTimestamp workflows.CapDefinition[int],
                                                                        reportContext workflows.CapDefinition[string],
                                                                        signatures workflows.CapDefinition[[]string],) Trigger {
    return &simpleTrigger{
        CapDefinition: workflows.ComponentCapDefinition[streams.Feed]{
        "benchmarkPrice": benchmarkPrice.Ref(),
        "feedId": feedId.Ref(),
        "fullReport": fullReport.Ref(),
        "observationTimestamp": observationTimestamp.Ref(),
        "reportContext": reportContext.Ref(),
        "signatures": signatures.Ref(),
        },
        benchmarkPrice: benchmarkPrice,
        feedId: feedId,
        fullReport: fullReport,
        observationTimestamp: observationTimestamp,
        reportContext: reportContext,
        signatures: signatures,
    }
}

type simpleTrigger struct {
    workflows.CapDefinition[streams.Feed]
    benchmarkPrice workflows.CapDefinition[string]
    feedId FeedId
    fullReport workflows.CapDefinition[string]
    observationTimestamp workflows.CapDefinition[int]
    reportContext workflows.CapDefinition[string]
    signatures workflows.CapDefinition[[]string]
}
func (c *simpleTrigger) BenchmarkPrice() workflows.CapDefinition[string] {
    return c.benchmarkPrice
}
func (c *simpleTrigger) FeedId() FeedId {
    return c.feedId
}
func (c *simpleTrigger) FullReport() workflows.CapDefinition[string] {
    return c.fullReport
}
func (c *simpleTrigger) ObservationTimestamp() workflows.CapDefinition[int] {
    return c.observationTimestamp
}
func (c *simpleTrigger) ReportContext() workflows.CapDefinition[string] {
    return c.reportContext
}
func (c *simpleTrigger) Signatures() workflows.CapDefinition[[]string] {
    return c.signatures
}

func (c *simpleTrigger) private() {}


type FeedId workflows.CapDefinition[streams.FeedId]

