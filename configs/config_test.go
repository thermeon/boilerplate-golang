package configs

import (
	"github.com/stretchr/testify/assert"
	"github.com/thermeon/boilerplate-golang/utils"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	// Set environment variables for testing.
	t.Setenv("DEBUG_ENABLE", "true")
	t.Setenv("DEBUG_MEMORYSTORE", "true")
	t.Setenv("DEBUG_SUPPRESS_METRICS", "true")

	c := Config{}
	c.Init()
	utils.Stringify(c)
	assert.Equal(t, true, c.DebugEnable)
}
