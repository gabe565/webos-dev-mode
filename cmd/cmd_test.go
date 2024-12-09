package cmd

import (
	"testing"

	"gabe565.com/webos-dev-mode/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("has config", func(t *testing.T) {
		cmd := New()
		conf, ok := config.FromContext(cmd.Context())
		assert.NotNil(t, conf)
		assert.True(t, ok)
	})

	t.Run("opts are called", func(t *testing.T) {
		cmd := New(func(cmd *cobra.Command) {
			cmd.Annotations = map[string]string{
				"worked": "true",
			}
		})
		assert.Equal(t, "true", cmd.Annotations["worked"])
	})
}
