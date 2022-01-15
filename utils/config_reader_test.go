package utils

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type Sample struct {
	Test int `env:"TEST-FIELD" yaml:"test"`
}

func TestEnvConfigReader_Read(t *testing.T) {
	t.Setenv("TEST-FIELD", strconv.Itoa(122))
	s := Sample{}
	err := NewEnvConfigReader().Read(&s)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 122, s.Test)
}

func TestYamlConfigReader_Read(t *testing.T) {
	s := Sample{}
	if err := NewYamlConfigReader("./", "sample.yaml").Read(&s); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 122, s.Test)
}
