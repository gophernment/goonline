package exer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopularOscar(t *testing.T) {
	given := `10,1937,41,"Paul Muni","The Story of Louis Pasteur"
11,1938,37,"Spencer Tracy","Captains Courageous"
12,1939,38,"Spencer Tracy","Boys Town"
`
	want := map[string]int{
		"Spencer Tracy": 2,
	}

	get := popularOscar(given)

	assert.Equal(t, want, get, "want %#v but get %#v", want, get)
}
