package config

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const EnvPrefix = "WEBOS_"

var ErrTokenNotSet = errors.New("session token not set")

func Load(cmd *cobra.Command) (*Config, error) {
	conf, ok := FromContext(cmd.Context())
	if !ok {
		panic("command context does not contain config")
	}

	var errs []error
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			if val, ok := os.LookupEnv(EnvName(f.Name)); ok {
				if err := f.Value.Set(val); err != nil {
					errs = append(errs, err)
				}
			}
		}
	})

	if conf.Token == "" {
		return conf, ErrTokenNotSet
	}

	return conf, errors.Join(errs...)
}

func EnvName(name string) string {
	name = strings.ToUpper(name)
	name = strings.ReplaceAll(name, "-", "_")
	return EnvPrefix + name
}
