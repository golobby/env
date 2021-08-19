package env_test

import (
	"github.com/golobby/env/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Database struct {
	Name string `env:"DB_NAME"`
	Port int16  `env:"DB_PORT"`
}

type Sample struct {
	NoEnv   float64 // A private field
	private string  `env:"private"`  // A private string field
	Name    string  `env:"NAME"`     // A string field
	Number  int32   `env:"NUMBER"`   // A numeric field
	Pi      float32 `env:"PI"`       // A float field
	IsAdmin bool    `env:"IS_ADMIN"` // A boolean field
	IsUser  bool    `env:"IS_USER"`  // A boolean field
	MySQL   *Database
	Nested  struct {
		Number int32 `env:"NESTED_NUMBER"` // A numeric field
	}
}

func TestLoad(t *testing.T) {
	_ = os.Setenv("NAME", "Milad")
	_ = os.Setenv("private", "secret")
	_ = os.Setenv("NUMBER", "666")
	_ = os.Setenv("PI", "3.14")
	_ = os.Setenv("IS_ADMIN", "1")
	_ = os.Setenv("IS_USER", "false")

	_ = os.Setenv("NESTED_NUMBER", "33")

	_ = os.Setenv("DB_NAME", "app")
	_ = os.Setenv("DB_PORT", "666")

	sample := Sample{}
	sample.MySQL = &Database{}

	err := env.Load(&sample)
	assert.NoError(t, err)

	assert.Equal(t, "Milad", sample.Name)
	assert.Equal(t, "secret", sample.private)
	assert.Equal(t, int32(666), sample.Number)
	assert.Equal(t, float32(3.14), sample.Pi)
	assert.Equal(t, true, sample.IsAdmin)
	assert.Equal(t, false, sample.IsUser)

	assert.Equal(t, int32(33), sample.Nested.Number)

	assert.Equal(t, "app", sample.MySQL.Name)
	assert.Equal(t, int16(666), sample.MySQL.Port)
}

func TestLoad_With_Invalid_Structure_It_Should_Fail(t *testing.T) {
	var number int

	err := env.Load(&number)
	assert.Error(t, err, "env: invalid structure")
}

func TestLoad_With_Invalid_Field_It_Should_Fail(t *testing.T) {
	type structure struct {
		Number int32 `env:"NUMBER"`
	}

	_ = os.Setenv("NUMBER", "invalid")

	sample := structure{}
	err := env.Load(&sample)
	assert.Error(t, err)
}

func TestLoad_With_Invalid_Structure_Field_It_Should_Fail(t *testing.T) {
	type structure struct {
		Nested struct {
			Number bool `env:"NUMBER"`
		}
	}

	_ = os.Setenv("NUMBER", "invalid")

	sample := structure{}
	err := env.Load(&sample)
	assert.Error(t, err)
}

func TestLoad_With_Invalid_Structure_Ptr_Field_It_Should_Fail(t *testing.T) {
	type Nested struct {
		Number bool `env:"NUMBER"`
	}

	type Structure struct {
		Nested *Nested
	}

	_ = os.Setenv("NUMBER", "invalid")

	sample := Structure{}
	sample.Nested = &Nested{}

	err := env.Load(&sample)
	assert.Error(t, err)
}
