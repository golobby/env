package env_test

import (
	"github.com/golobby/env/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Database struct {
	Name string `env:"DB_NAME"` // A string field
	Port int16  `env:"DB_PORT"` // A numeric field
}

type Config struct {
	NoEnv   float64   // A private field
	private string    `env:"private"`  // A private string field
	Name    string    `env:"NAME"`     // A string field
	Number  int32     `env:"NUMBER"`   // A numeric field
	Pi      float32   `env:"PI"`       // A float field
	IsAdmin bool      `env:"IS_ADMIN"` // A boolean field
	IsUser  bool      `env:"IS_USER"`  // A boolean field
	IPs     []string  `env:"IPS"`      // A []string field
	Nums    []int8    `env:"NUMS"`     // A []int8 field
	MySQL   *Database // A nested struct pointer
	Nested  struct {  // A nested struct
		Number int32 `env:"NESTED_NUMBER"` // A numeric field
	}
}

func setDefaults() {
	_ = os.Setenv("NAME", "Milad")
	_ = os.Setenv("private", "secret")
	_ = os.Setenv("NUMBER", "666")
	_ = os.Setenv("PI", "3.14")
	_ = os.Setenv("IS_ADMIN", "1")
	_ = os.Setenv("IS_USER", "false")
	_ = os.Setenv("IPS", "192.168.0.1, 192.168.0.2")
	_ = os.Setenv("NUMS", "10, 11, 12 , 13")

	_ = os.Setenv("NESTED_NUMBER", "33")

	_ = os.Setenv("DB_NAME", "app")
	_ = os.Setenv("DB_PORT", "666")
}

func TestFeed(t *testing.T) {
	setDefaults()

	sample := Config{}
	sample.MySQL = &Database{}

	err := env.Feed(&sample)
	assert.NoError(t, err)

	assert.Equal(t, "Milad", sample.Name)
	assert.Equal(t, "secret", sample.private)
	assert.Equal(t, int32(666), sample.Number)
	assert.Equal(t, float32(3.14), sample.Pi)
	assert.Equal(t, true, sample.IsAdmin)
	assert.Equal(t, false, sample.IsUser)
	assert.Equal(t, []string{"192.168.0.1", "192.168.0.2"}, sample.IPs)
	assert.Equal(t, []int8{10, 11, 12, 13}, sample.Nums)

	assert.Equal(t, int32(33), sample.Nested.Number)

	assert.Equal(t, "app", sample.MySQL.Name)
	assert.Equal(t, int16(666), sample.MySQL.Port)
}

func TestFeed_With_Default_Values(t *testing.T) {
	_ = os.Setenv("DB_NAME", "mysql")
	_ = os.Setenv("DB_PORT", "")

	sample := Database{}
	sample.Port = 666

	err := env.Feed(&sample)
	assert.NoError(t, err)

	assert.Equal(t, "mysql", sample.Name)
	assert.Equal(t, int16(666), sample.Port)
}

func TestFeed_With_Invalid_Structure_It_Should_Fail(t *testing.T) {
	err := env.Feed(nil)
	assert.Error(t, err, "env: invalid struct")

	err = env.Feed(666)
	assert.Error(t, err, "env: invalid struct")

	err = env.Feed(struct{}{})
	assert.Error(t, err, "env: invalid struct")

	var pointer bool
	err = env.Feed(&pointer)
	assert.Error(t, err, "env: invalid struct")
}

func TestFeed_With_Invalid_Field_It_Should_Fail(t *testing.T) {
	_ = os.Setenv("DB_PORT", "invalid")

	sample := Database{}
	err := env.Feed(&sample)
	assert.Error(t, err)
}

func TestFeed_With_Invalid_Nested_Structure_Field_It_Should_Fail(t *testing.T) {
	_ = os.Setenv("NESTED_NUMBER", "invalid")

	sample := &Config{}
	err := env.Feed(sample)
	assert.Error(t, err)
}

func TestFeed_With_Invalid_Nested_Structure_Ptr_Field_It_Should_Fail(t *testing.T) {
	_ = os.Setenv("DB_PORT", "invalid")

	sample := Config{}
	sample.MySQL = &Database{}

	err := env.Feed(&sample)
	assert.Error(t, err)
}
