package check

import (
	"fmt"
	"time"

	"gabe565.com/utils/must"
	"gabe565.com/webos-dev-mode/pkg/webosdev"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check the current dev mode session expiration",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	cmd.SilenceUsage = true
	token := must.Must2(cmd.Flags().GetString("token"))
	client := webosdev.NewClient(webosdev.WithSessionToken(token))

	expiresIn, _, err := client.CheckExpiration(cmd.Context())
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Expires in", expiresIn)
	_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Expires at", time.Now().Add(expiresIn).Format(time.RFC3339))
	return nil
}
