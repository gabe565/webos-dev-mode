package extend

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"gabe565.com/utils/cobrax"
	"gabe565.com/webos-dev-mode/internal/config"
	"gabe565.com/webos-dev-mode/pkg/webosdev"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend",
		Short: "Extend a dev mode session",
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

	return Extend(cmd.Context(),
		webosdev.WithSessionToken(conf.Token),
		webosdev.WithTimeout(conf.RequestTimeout),
		webosdev.WithUserAgent(cobrax.BuildUserAgent(cmd)),
	)
}

var ErrShortExpiration = errors.New("expiration time is too short")

func Extend(ctx context.Context, opts ...webosdev.Option) error {
	client := webosdev.NewClient(opts...)

	if err := client.ExtendSession(ctx); err != nil {
		return fmt.Errorf("failed to extend dev session: %w", err)
	}

	expiration, err := client.CheckExpiration(ctx)
	if err != nil {
		return fmt.Errorf("failed to check expiration: %w", err)
	}

	if expiration < 999*time.Hour {
		slog.Error("Expiration did not update", "value", expiration)
		return ErrShortExpiration
	}

	slog.Info("Dev mode session extended successfully.",
		"expires", time.Now().Add(expiration).Format(time.RFC3339),
	)
	return nil
}
