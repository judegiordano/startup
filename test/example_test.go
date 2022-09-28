package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Run("assert equal", func(t *testing.T) {
		assert.Equal(t, 1, 2)
	})
}
