package extend

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"gabe565.com/lg-dev-mode/pkg/lgdevmode"
	"gabe565.com/utils/must"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend",
		Short: "Extend an LG dev mode session",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true
	token := must.Must2(cmd.Flags().GetString("token"))
	return Extend(cmd.Context(), token)
}

var ErrShortExpiration = errors.New("expiration time is too short")

func Extend(ctx context.Context, token string) error {
	client := lgdevmode.New(lgdevmode.WithSessionToken(token))

	if _, _, err := client.ExtendSession(ctx); err != nil {
		return fmt.Errorf("failed to extend dev session: %w", err)
	}

	expiration, _, err := client.CheckExpiration(ctx)
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
