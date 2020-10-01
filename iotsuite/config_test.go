package iotsuite

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigFile(t *testing.T) {
	var config = ReadConfig()
	assert.Equal(t, "https://accounts.bosch-iot-suite.com", config.BaseUrl)
}

