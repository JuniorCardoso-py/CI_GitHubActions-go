package main_test

import (
	"testing"

	"github.com/golangci_test"
	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	expected := 100

	v1, v2 := 50, 50

	result := main.Sum(v1, v2)

	assert.Equal(t, result, expected)
}
