package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestHate(t *testing.T) {
	const hated = "Tom"

	m := Mouse{
		Name:  "Jerry",
		Age:   3,
		Hates: hated,
	}

	// Will this end in an error?
	_, err := m.IsHating("")
	if err == nil {
		t.Error("Mice cannot hate empty name")
	}

	// Hate gone strong
	hate, err := m.IsHating(hated)
	if err != nil {
		fmt.Println(err)
	}

	assert.Assert(t, hate)
}
