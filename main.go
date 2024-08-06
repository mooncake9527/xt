package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/go-xt/xt/cmd/xt/v2/internal/change"
	"github.com/go-xt/xt/cmd/xt/v2/internal/project"
	"github.com/go-xt/xt/cmd/xt/v2/internal/proto"
	"github.com/go-xt/xt/cmd/xt/v2/internal/run"
	"github.com/go-xt/xt/cmd/xt/v2/internal/upgrade"
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
