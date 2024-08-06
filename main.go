package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/mooncake9527/xt/internal/change"
	"github.com/mooncake9527/xt/internal/project"
	"github.com/mooncake9527/xt/internal/proto"
	"github.com/mooncake9527/xt/internal/run"
	"github.com/mooncake9527/xt/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "xt",
	Short:   "XT: An elegant toolkit for Go microservices.",
	Long:    `XT: An elegant toolkit for Go microservices.`,
	Version: "v1.0.0",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
