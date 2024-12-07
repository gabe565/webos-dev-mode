package config

import "github.com/spf13/cobra"

const (
	FlagToken    = "token"
	FlagInterval = "interval"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.StringVar(&c.Token, FlagToken, c.Token, "Session token")

	if cmd, _, err := cmd.Find([]string{"cron"}); err == nil {
		cmd.Flags().DurationVar(&c.CronInterval, FlagInterval, c.CronInterval, "Extend cron interval")
	} else {
		panic(err)
	}
}
