package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelToString(t *testing.T) {
	assert.Equal(t, "DEBUG", LevelToString(DEBUG))
	assert.Equal(t, "INFO", LevelToString(INFO))
	assert.Equal(t, "WARN", LevelToString(WARN))
	assert.Equal(t, "TRACE", LevelToString(TRACE))
	assert.Equal(t, "ERROR", LevelToString(ERROR))
	assert.Equal(t, "FATAL", LevelToString(FATAL))
}

func TestStringToLevel(t *testing.T) {
	assert.Equal(t, DEBUG, StringToLevel("DEBUG"))
	assert.Equal(t, INFO, StringToLevel("INFO"))
	assert.Equal(t, WARN, StringToLevel("WARN"))
	assert.Equal(t, TRACE, StringToLevel("TRACE"))
	assert.Equal(t, ERROR, StringToLevel("ERROR"))
	assert.Equal(t, FATAL, StringToLevel("FATAL"))
}
