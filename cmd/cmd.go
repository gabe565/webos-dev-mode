package cmd

import (
	"errors"
	"os"

	"gabe565.com/lg-dev-mode/cmd/check"
	"gabe565.com/lg-dev-mode/cmd/cron"
	"gabe565.com/lg-dev-mode/cmd/extend"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/must"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "lg-dev-mode",
		Short:             "LG dev mode tools",
		SilenceErrors:     true,
		PersistentPreRunE: preRun,

		DisableAutoGenTag: true,
	}
	cmd.AddCommand(extend.New(), cron.New(), check.New())
	cmd.PersistentFlags().String("token", os.Getenv("LG_TOKEN"), "Session token")
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

var ErrTokenNotSet = errors.New("session token not set")

func preRun(cmd *cobra.Command, _ []string) error {
	token := must.Must2(cmd.Flags().GetString("token"))
	if token == "" {
		return ErrTokenNotSet
	}
	return nil
}
