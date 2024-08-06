package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mooncake9527/xt/internal/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the xt tools",
	Long:  "Upgrade the xt tools. Example: xt upgrade",
	Run:   Run,
}

// Run upgrade the xt tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/mooncake9527/xt@latest",
		"github.com/go-xt/xt/cmd/protoc-gen-go-http/v2@latest",
		"github.com/go-xt/xt/cmd/protoc-gen-go-errors/v2@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
