package env_test

import (
	"env"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Load_It_Should_Read_Empty_File(t *testing.T) {
	vs, err := env.Load("test/.env1")
	assert.NoErrorf(t, err, "Expected no error got %v", err)
	assert.Empty(t, vs, "Expected empty env")
}

func Test_Load_It_Should_Read_Env_File_And_Load_To_OS(t *testing.T) {
	vs, err := env.Load("test/.env2")
	l := len(vs)

	assert.NoErrorf(t, err, "Expected no error got %v", err)
	assert.Equalf(t, 8, l, "Expected %v got %v", 8, l)

	assert.Equal(t, "127.0.0.1", vs["DB_HOST"])
	assert.Equal(t, "App", vs["DB_NAME"])
	assert.Equal(t, "3306", vs["DB_PORT"])
	assert.Equal(t, "MySQL", vs["DB_TYPE"])
	assert.Equal(t, "", vs["APP_NAME"])
	assert.Equal(t, "https://example.com", vs["APP_URL"])
	assert.Equal(t, "true", vs["DEBUG"])
	assert.Equal(t, "#VALUE!", vs["NOT_COMMENT"])

	assert.Equal(t, "127.0.0.1", os.Getenv("DB_HOST"))
	assert.Equal(t, "App", os.Getenv("DB_NAME"))
	assert.Equal(t, "3306", os.Getenv("DB_PORT"))
	assert.Equal(t, "MySQL", os.Getenv("DB_TYPE"))
	assert.Equal(t, "", os.Getenv("APP_NAME"))
	assert.Equal(t, "https://example.com", os.Getenv("APP_URL"))
	assert.Equal(t, "#VALUE!", os.Getenv("NOT_COMMENT"))
}

func Test_Load_It_Should_Not_Overwrite_OS_Variable(t *testing.T) {
	err := os.Setenv("DB_NAME", "Not App!")
	assert.NoErrorf(t, err, "Expected no error got %v", err)

	_, err = env.Load("test/.env2")
	assert.NoErrorf(t, err, "Expected no error got %v", err)

	assert.Equal(t, "Not App!", os.Getenv("DB_NAME"))
}

func Test_Overload_It_Should_Overwrite_OS_Variable(t *testing.T) {
	err := os.Setenv("DB_NAME", "Not App!")
	assert.NoErrorf(t, err, "Expected no error got %v", err)

	_, err = env.Overload("test/.env2")
	assert.NoErrorf(t, err, "Expected no error got %v", err)

	assert.Equal(t, "App", os.Getenv("DB_NAME"))
}
