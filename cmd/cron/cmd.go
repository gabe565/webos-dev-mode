package cron

import (
	"context"
	"log/slog"
	"time"

	"gabe565.com/utils/must"
	"gabe565.com/webos-dev-mode/cmd/extend"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cron",
		Short: "Extends dev mode session on a given interval",
		RunE:  run,
	}
	cmd.Flags().Duration("interval", 24*time.Hour, "Extend cron interval")
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true
	interval := must.Must2(cmd.Flags().GetDuration("interval"))
	token := must.Must2(cmd.Flags().GetString("token"))

	if err := extend.Extend(cmd.Context(), token); err != nil {
		return err
	}

	ticker := time.NewTicker(interval)
	for range ticker.C {
		ctx, cancel := context.WithTimeout(cmd.Context(), time.Minute)
		if err := extend.Extend(ctx, token); err != nil {
			slog.Error("Extend failed", "error", err)
		}
		cancel()
	}
	return nil
}
