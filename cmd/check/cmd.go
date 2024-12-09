package check

import (
	"encoding/json"
	"io"
	"time"

	"gabe565.com/webos-dev-mode/internal/config"
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

type Output struct {
	ExpiresIn string    `json:"expires_in"`
	ExpiresAt time.Time `json:"expires_at"`
}

func run(cmd *cobra.Command, _ []string) error {
	conf, err := config.Load(cmd)
	if err != nil {
		return err
	}
	cmd.SilenceUsage = true

	client := webosdev.NewClient(
		webosdev.WithSessionToken(conf.Token),
		webosdev.WithTimeout(conf.RequestTimeout),
	)

	expiresIn, err := client.CheckExpiration(cmd.Context())
	if err != nil {
		return err
	}

	output := Output{
		ExpiresIn: expiresIn.String(),
		ExpiresAt: time.Now().Add(expiresIn),
	}

	if conf.JSON {
		if err := json.NewEncoder(cmd.OutOrStdout()).Encode(output); err != nil {
			return err
		}
	} else {
		_, _ = io.WriteString(cmd.OutOrStdout(),
			"Expires in "+output.ExpiresIn+"\n"+
				"Expires at "+output.ExpiresAt.Format(time.RFC3339)+"\n",
		)
	}
	return nil
}
