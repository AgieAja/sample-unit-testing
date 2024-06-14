package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddWithAssert(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Add(2, 3) should return 5")
	assert.NotEqual(t, 6, result, "Add(2, 3) should not return 6")
}

func TestAddWithRequire(t *testing.T) {
	result := Add(2, 3)
	require.Equal(t, 5, result, "Add(2, 3) should return 5")
	require.NotEqual(t, 6, result, "Add(2, 3) should not return 6")
}
