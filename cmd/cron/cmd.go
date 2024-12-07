package cron

import (
	"log/slog"
	"time"

	"gabe565.com/webos-dev-mode/cmd/extend"
	"gabe565.com/webos-dev-mode/internal/config"
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

	if err := extend.Extend(cmd.Context(), conf); err != nil {
		return err
	}

	ticker := time.NewTicker(conf.CronInterval)
	for range ticker.C {
		if err := extend.Extend(cmd.Context(), conf); err != nil {
			slog.Error("Extend failed", "error", err)
		}
	}
	return nil
}
