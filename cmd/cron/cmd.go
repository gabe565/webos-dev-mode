package cron

import (
	"log/slog"
	"time"

	"gabe565.com/utils/cobrax"
	"gabe565.com/webos-dev-mode/cmd/extend"
	"gabe565.com/webos-dev-mode/internal/config"
	"gabe565.com/webos-dev-mode/pkg/webosdev"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cron",
		Short: "Extends dev mode session on a given interval",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	conf, err := config.Load(cmd)
	if err != nil {
		return err
	}
	cmd.SilenceUsage = true

	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-cmd.Context().Done():
			return cmd.Context().Err()
		case <-timer.C:
			timer.Reset(conf.CronInterval)

			if err := extend.Extend(cmd.Context(),
				webosdev.WithSessionToken(conf.Token),
				webosdev.WithTimeout(conf.RequestTimeout),
				webosdev.WithUserAgent(cobrax.BuildUserAgent(cmd)),
			); err != nil {
				slog.Error("Extend failed", "error", err)
			}
		}
	}
}
