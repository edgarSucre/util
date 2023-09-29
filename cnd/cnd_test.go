package cnd_test

import (
	"testing"

	"github.com/edgarSucre/util/cnd"
	"github.com/stretchr/testify/assert"
)

func Test_NotNil(t *testing.T) {
	var a *int

	r := cnd.NotNil(a, 10)
	assert.Equal(t, 10, r)

	b := 20
	r = cnd.NotNil(&b, 10)
	assert.Equal(t, 20, r)
}

func Test_Or(t *testing.T) {
	var r string

	r = cnd.Or(true, "true", "false")
	assert.Equal(t, "true", r)

	r = cnd.Or(false, "true", "false")
	assert.Equal(t, "false", r)
}
