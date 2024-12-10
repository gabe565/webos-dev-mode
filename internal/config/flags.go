package config

import (
	"gabe565.com/utils/must"
	"github.com/spf13/cobra"
)

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
	must.Must(cmd.RegisterFlagCompletionFunc(FlagBaseURL, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"https://"}, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	}))

	fs.BoolVarP(&c.Insecure, FlagInsecure, "k", c.Insecure, "Skip TLS verification")
	must.Must(cmd.RegisterFlagCompletionFunc(FlagInsecure, boolCompletion))

	fs.StringVarP(&c.Token, FlagToken, "t", c.Token, "Session token")
	must.Must(cmd.RegisterFlagCompletionFunc(FlagToken, cobra.NoFileCompletions))

	fs.DurationVar(&c.RequestTimeout, FlagRequestTimeout, c.RequestTimeout, "HTTP request timeout")
	must.Must(cmd.RegisterFlagCompletionFunc(FlagRequestTimeout, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"1m", "30s"}, cobra.ShellCompDirectiveNoFileComp
	}))

	if cmd, _, err := cmd.Find([]string{"cron"}); err == nil {
		cmd.Flags().DurationVarP(&c.CronInterval, FlagInterval, "i", c.CronInterval, "Extend cron interval")
		must.Must(cmd.RegisterFlagCompletionFunc(FlagInterval, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return []string{"1h\thourly", "24h\tdaily", "168h\tweekly", "720h\tmonthly"}, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveKeepOrder
		}))
	} else {
		panic(err)
	}

	if cmd, _, err := cmd.Find([]string{"check"}); err == nil {
		cmd.Flags().BoolVarP(&c.JSON, FlagJSON, "j", c.JSON, "Show output as JSON")
		must.Must(cmd.RegisterFlagCompletionFunc(FlagJSON, boolCompletion))
	} else {
		panic(err)
	}
}

func boolCompletion(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return []string{"true", "false"}, cobra.ShellCompDirectiveNoFileComp
}
