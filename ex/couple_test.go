package ex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoupleEven(t *testing.T) {
	given := "abcdef"
	want := []string{"ab", "cd", "ef"}

	get := couple(given)

	assert.Equal(t, want, get, fmt.Sprintf("given %q want %#v but get %#v", given, want, get))
}

func TestCoupleOdd(t *testing.T) {
	given := "abcdefg"
	want := []string{"ab", "cd", "ef", "g*"}

	get := couple(given)

	assert.Equal(t, want, get, fmt.Sprintf("given %q want %#v but get %#v", given, want, get))
}
