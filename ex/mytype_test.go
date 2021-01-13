package ex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyTypeWithSet(t *testing.T) {
	var i Int = 10

	i.Set(15)
	want := "15"

	get := i.String()

	assert.Equal(t, want, get, "Int should set a new value")
}
