package cmd

import (
	"errors"

	"github.com/smartcontractkit/chainlink-common/observability-lib/capabilities"

	atlasdon "github.com/smartcontractkit/chainlink-common/observability-lib/atlas-don"
	corenode "github.com/smartcontractkit/chainlink-common/observability-lib/core-node"
	corenodecomponents "github.com/smartcontractkit/chainlink-common/observability-lib/core-node-components"
	"github.com/smartcontractkit/chainlink-common/observability-lib/grafana"
	k8sresources "github.com/smartcontractkit/chainlink-common/observability-lib/k8s-resources"
	nopocr "github.com/smartcontractkit/chainlink-common/observability-lib/nop-ocr"
)

type TypeDashboard string

const (
	TypeDashboardCoreNode           TypeDashboard = "core-node"
	TypeDashboardCoreNodeComponents TypeDashboard = "core-node-components"
	TypeDashboardCoreNodeResources  TypeDashboard = "core-node-resources"
	TypeDashboardDONOCR             TypeDashboard = "don-ocr"
	TypeDashboardDONOCR2            TypeDashboard = "don-ocr2"
	TypeDashboardDONOCR3            TypeDashboard = "don-ocr3"
	TypeDashboardNOPOCR2            TypeDashboard = "nop-ocr2"
	TypeDashboardNOPOCR3            TypeDashboard = "nop-ocr3"
	TypeDashboardCapabilities       TypeDashboard = "capabilities"
)

type OCRVersion string

const (
	OCRVersionOCR  OCRVersion = "ocr"
	OCRVersionOCR2 OCRVersion = "ocr2"
	OCRVersionOCR3 OCRVersion = "ocr3"
)

type BuildOptions struct {
	Name              string
	Platform          grafana.TypePlatform
	TypeDashboard     TypeDashboard
	MetricsDataSource *grafana.DataSource
	LogsDataSource    *grafana.DataSource
	SlackChannel      string
	SlackWebhookURL   string
	AlertsTags        map[string]string
}

func BuildDashboardWithType(options *BuildOptions) (*grafana.Dashboard, error) {
	switch options.TypeDashboard {
	case TypeDashboardCoreNode:
		return corenode.NewDashboard(&corenode.Props{
			Name:              options.Name,
			Platform:          options.Platform,
			MetricsDataSource: options.MetricsDataSource,
			SlackChannel:      options.SlackChannel,
			SlackWebhookURL:   options.SlackWebhookURL,
			AlertsTags:        options.AlertsTags,
		})
	case TypeDashboardCoreNodeComponents:
		return corenodecomponents.NewDashboard(&corenodecomponents.Props{
			Name:              options.Name,
			Platform:          options.Platform,
			MetricsDataSource: options.MetricsDataSource,
			LogsDataSource:    options.LogsDataSource,
		})
	case TypeDashboardCoreNodeResources:
		if options.Platform != grafana.TypePlatformKubernetes {
			return nil, errors.New("core-node-resources dashboard is only available for kubernetes")
		}
		return k8sresources.NewDashboard(&k8sresources.Props{
			Name:              options.Name,
			MetricsDataSource: options.MetricsDataSource,
		})
	case TypeDashboardDONOCR:
		return atlasdon.NewDashboard(&atlasdon.Props{
			Name:              options.Name,
			Platform:          options.Platform,
			MetricsDataSource: options.MetricsDataSource,
			OCRVersion:        string(OCRVersionOCR),
		})
	case TypeDashboardDONOCR2:
		return atlasdon.NewDashboard(&atlasdon.Props{
			Name:              options.Name,
			Platform:          options.Platform,
			MetricsDataSource: options.MetricsDataSource,
			OCRVersion:        string(OCRVersionOCR2),
		})
	case TypeDashboardDONOCR3:
		return atlasdon.NewDashboard(&atlasdon.Props{
			Name:              options.Name,
			Platform:          options.Platform,
			MetricsDataSource: options.MetricsDataSource,
			OCRVersion:        string(OCRVersionOCR3),
		})
	case TypeDashboardNOPOCR2:
		return nopocr.NewDashboard(&nopocr.Props{
			Name:              options.Name,
			MetricsDataSource: options.MetricsDataSource,
			OCRVersion:        string(OCRVersionOCR2),
		})
	case TypeDashboardNOPOCR3:
		return nopocr.NewDashboard(&nopocr.Props{
			Name:              options.Name,
			MetricsDataSource: options.MetricsDataSource,
			OCRVersion:        string(OCRVersionOCR3),
		})
	case TypeDashboardCapabilities:
		return capabilities.NewDashboard(&capabilities.Props{
			Name:              options.Name,
			MetricsDataSource: options.MetricsDataSource,
		})
	default:
		return nil, errors.New("invalid dashboard type")
	}
}
