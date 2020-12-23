package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonAddOne(t *testing.T) {
	s := GetInstance()

	assert.Equal(t, 1, s.AddOne())
	assert.Equal(t, 2, s.AddOne())
	assert.Equal(t, 3, s.AddOne())
	assert.Equal(t, 4, s.AddOne())
	assert.Equal(t, 5, s.AddOne())
}
