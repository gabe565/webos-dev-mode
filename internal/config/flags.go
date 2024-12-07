package config

import "github.com/spf13/cobra"

const (
	FlagToken          = "token"
	FlagInterval       = "interval"
	FlagRequestTimeout = "request-timeout"
	FlagJSON           = "json"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVarP(&c.Token, FlagToken, "t", c.Token, "Session token")
	fs.DurationVar(&c.RequestTimeout, FlagRequestTimeout, c.RequestTimeout, "HTTP request timeout")

	if cmd, _, err := cmd.Find([]string{"cron"}); err == nil {
		cmd.Flags().DurationVar(&c.CronInterval, FlagInterval, c.CronInterval, "Extend cron interval")
	} else {
		panic(err)
	}

	if cmd, _, err := cmd.Find([]string{"check"}); err == nil {
		cmd.Flags().BoolVarP(&c.JSON, FlagJSON, "j", c.JSON, "Show output as JSON")
	} else {
		panic(err)
	}
}
