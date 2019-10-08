package env_test

import (
	"env"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_It_Should_Read_Empty_File(t *testing.T) {
	e, err := env.New("test/.env1")
	assert.NoErrorf(t, err, "Expected no error got %v", err)
	assert.Empty(t, e, "Expected empty env")
}

func Test_New_It_Should_Read_Env_File(t *testing.T) {
	e, err := env.New("test/.env2")
	l := len(e)

	assert.NoErrorf(t, err, "Expected no error got %v", err)

	assert.Equalf(t, 7, l, "Expected %v got %v", 7, l)

	fmt.Println(e)
}
