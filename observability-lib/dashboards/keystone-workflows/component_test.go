package keystone_workflows

import (
	"flag"
	"os"
	"testing"

	"github.com/smartcontractkit/chainlink-common/observability-lib/grafana"

	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update golden test files")

const fileOutput = "test-output.json"

func TestGenerateFile(t *testing.T) {
	if *update == false {
		t.Skip("skipping test")
	}

	testDashboard, err := NewDashboard(&Props{
		Name:              "Workflows",
		MetricsDataSource: grafana.NewDataSource("Prometheus", "1"),
		LogsDataSource:    grafana.NewDataSource("Loki", "2"),
		Tested:            true,
	})
	if err != nil {
		t.Errorf("Error creating dashboard: %v", err)
	}
	json, errJSON := testDashboard.GenerateJSON()
	if errJSON != nil {
		t.Errorf("Error generating JSON: %v", errJSON)
	}
	if _, errExists := os.Stat(fileOutput); errExists == nil {
		errRemove := os.Remove(fileOutput)
		if errRemove != nil {
			t.Errorf("Error removing file: %v", errRemove)
		}
	}
	file, errFile := os.Create(fileOutput)
	if errFile != nil {
		panic(errFile)
	}
	writeString, err := file.WriteString(string(json))
	if err != nil {
		t.Errorf("Error writing to file: %v", writeString)
	}
	t.Cleanup(func() {
		file.Close()
	})
}

func TestNewDashboard(t *testing.T) {
	t.Run("NewDashboard creates a dashboard", func(t *testing.T) {
		testDashboard, err := NewDashboard(&Props{
			Name:              "Workflows",
			MetricsDataSource: grafana.NewDataSource("Prometheus", "1"),
			LogsDataSource:    grafana.NewDataSource("Loki", "2"),
			Tested:            true,
		})
		if err != nil {
			t.Errorf("Error creating dashboard: %v", err)
		}
		require.IsType(t, grafana.Observability{}, *testDashboard)
		require.Equal(t, "Workflows", *testDashboard.Dashboard.Title)
		json, errJSON := testDashboard.GenerateJSON()
		if errJSON != nil {
			t.Errorf("Error generating JSON: %v", errJSON)
		}

		jsonCompared, errCompared := os.ReadFile(fileOutput)
		if errCompared != nil {
			t.Errorf("Error reading file: %v", errCompared)
		}

		require.JSONEq(t, string(jsonCompared), string(json))
	})
}
