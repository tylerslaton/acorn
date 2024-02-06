package cli

import (
	cli "github.com/acorn-io/runtime/pkg/cli/builder"
	"github.com/acorn-io/runtime/pkg/local"
	"github.com/spf13/cobra"
)

func NewLocalLogs() *cobra.Command {
	cmd := cli.Command(&LocalLogs{}, cobra.Command{
		Use:          "logs [flags]",
		Aliases:      []string{"log"},
		SilenceUsage: true,
		Short:        "Show logs of local development server",
	})
	return cmd
}

type LocalLogs struct {
	local.LogOptions
}

func (a *LocalLogs) Run(cmd *cobra.Command, _ []string) error {
	c, err := local.NewContainer()
	if err != nil {
		return err
	}

	return c.Logs(cmd.Context(), a.LogOptions)
}
