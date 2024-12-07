package cmd

import (
	"context"

	"gabe565.com/utils/cobrax"
	"gabe565.com/webos-dev-mode/cmd/check"
	"gabe565.com/webos-dev-mode/cmd/cron"
	"gabe565.com/webos-dev-mode/cmd/extend"
	"gabe565.com/webos-dev-mode/internal/config"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "webos-dev-mode",
		Short:         "webOS dev mode tools",
		SilenceErrors: true,

		DisableAutoGenTag: true,
	}
	cmd.AddCommand(extend.New(), cron.New(), check.New())

	conf := config.New()
	conf.RegisterFlags(cmd)
	cmd.SetContext(config.NewContext(context.Background(), conf))

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}
