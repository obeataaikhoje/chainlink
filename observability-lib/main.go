package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/smartcontractkit/chainlink-common/observability-lib/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "observability-lib",
	Short: "observability-lib is a library for creating and deploying Grafana dashboards and alerts",
}

func init() {
	rootCmd.AddCommand(cmd.DeployCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
