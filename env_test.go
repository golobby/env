package env_test

import (
	"github.com/golobby/env/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Sample struct {
	NoEnv   float64 // A private field
	private string  `env:"private"`  // A private field
	Name    string  `env:"NAME"`     // A public field
	Number  int32   `env:"NUMBER"`   // A numeric field
	Pi      float32 `env:"PI"`       // A float field
	IsAdmin bool    `env:"IS_ADMIN"` // A boolean field
	IsUser  bool    `env:"IS_USER"`  // A boolean field
}

func TestLoad(t *testing.T) {
	_ = os.Setenv("NAME", "Milad")
	_ = os.Setenv("private", "secret")
	_ = os.Setenv("NUMBER", "666")
	_ = os.Setenv("PI", "3.14")
	_ = os.Setenv("IS_ADMIN", "1")
	_ = os.Setenv("IS_USER", "false")

	sample := Sample{}

	err := env.Load(&sample)
	assert.NoError(t, err)

	assert.Equal(t, "Milad", sample.Name)
	assert.Equal(t, "secret", sample.private)
	assert.Equal(t, int32(666), sample.Number)
	assert.Equal(t, float32(3.14), sample.Pi)
	assert.Equal(t, true, sample.IsAdmin)
	assert.Equal(t, false, sample.IsUser)
}
