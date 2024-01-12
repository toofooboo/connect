package config_test

import (
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/skip-mev/slinky/oracle/config"
)

var (
	goodConfigContent = `
name = "testname"
path = "testpath"
timeout = "5s"
interval = "10s"
max_queries = 10
`

	badConfigContent = `
name = "testname"
path = "testpath"
timeout = "10"
interval = "60"
max_queries = 10
`

	missingFieldConfigContent = `
name = "testname"
path = "testpath"
timeout = "10s"
`

	invalidIntervalConfigContent = `
name = "testname"
path = "testpath"
timeout = "10s"
interval = "5s"
`

	invalidMaxQueriesConfigContent = `
name = "testname"
path = "testpath"
timeout = "10s"
interval = "60s"
max_queries = -1
`
)

func TestReadProviderConfigFromFile(t *testing.T) {
	testCases := []struct {
		name        string
		config      string
		expectedErr bool
	}{
		{
			name:        "good config",
			config:      goodConfigContent,
			expectedErr: false,
		},
		{
			name:        "bad config",
			config:      badConfigContent,
			expectedErr: true,
		},
		{
			name:        "missing field config",
			config:      missingFieldConfigContent,
			expectedErr: true,
		},
		{
			name:        "invalid interval config",
			config:      invalidIntervalConfigContent,
			expectedErr: true,
		},
		{
			name:        "invalid max queries config",
			config:      invalidMaxQueriesConfigContent,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create temp file
			f, err := os.CreateTemp("", "provider_config")
			assert.NoError(t, err)
			defer os.Remove(f.Name())

			// Write the config as a toml file
			_, err = f.WriteString(tc.config)
			assert.NoError(t, err)

			// Read config from file
			_, err = config.ReadProviderConfigFromFile(f.Name())
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}