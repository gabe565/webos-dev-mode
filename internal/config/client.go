package config

import (
	"gabe565.com/utils/cobrax"
	"gabe565.com/webos-dev-mode/pkg/webosdev"
	"github.com/spf13/cobra"
)

func (c *Config) NewClient(cmd *cobra.Command) *webosdev.Client {
	return webosdev.NewClient(
		webosdev.WithSessionToken(c.Token),
		webosdev.WithTimeout(c.RequestTimeout),
		webosdev.WithUserAgent(cobrax.BuildUserAgent(cmd)),
	)
}
