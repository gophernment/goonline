// +build integration

package exer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopularOscarFromCSV(t *testing.T) {
	given := "./oscar_age_male.csv"
	want := map[string]int{
		"Marlon Brando":    2,
		"Daniel Day-Lewis": 3,
		"Sean Penn":        2,
		"Tom Hanks":        2,
		"Fredric March":    2,
		"Spencer Tracy":    2,
		"Gary Cooper":      2,
		"Jack Nicholson":   2,
		"Dustin Hoffman":   2,
	}

	get, err := exer.PopularOscar(given)
	if err != nil {
		t.Errorf(err)
		return
	}

	assert.Equal(t, want, get, "")
}
