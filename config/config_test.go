package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySQL(t *testing.T) {
	mc := MySQL()
	assert.Equal(t, "localhost", mc.Host)
	assert.Equal(t, "root", mc.User)
	assert.Equal(t, "password", mc.Password)
	assert.Equal(t, "database", mc.Database)
}
