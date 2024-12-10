package config

import "github.com/spf13/cobra"

const (
	FlagBaseURL        = "base-url"
	FlagInsecure       = "insecure"
	FlagToken          = "token"
	FlagInterval       = "interval"
	FlagRequestTimeout = "request-timeout"
	FlagJSON           = "json"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVar(&c.BaseURL, FlagBaseURL, c.BaseURL, "Base URL of the API")
	fs.BoolVarP(&c.Insecure, FlagInsecure, "k", c.Insecure, "Skip TLS verification")
	fs.StringVarP(&c.Token, FlagToken, "t", c.Token, "Session token")
	fs.DurationVar(&c.RequestTimeout, FlagRequestTimeout, c.RequestTimeout, "HTTP request timeout")

	if cmd, _, err := cmd.Find([]string{"cron"}); err == nil {
		cmd.Flags().DurationVarP(&c.CronInterval, FlagInterval, "i", c.CronInterval, "Extend cron interval")
	} else {
		panic(err)
	}

	if cmd, _, err := cmd.Find([]string{"check"}); err == nil {
		cmd.Flags().BoolVarP(&c.JSON, FlagJSON, "j", c.JSON, "Show output as JSON")
	} else {
		panic(err)
	}
}
