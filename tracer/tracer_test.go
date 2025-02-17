package tracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	t.Run("creates a tuple with W=1", func(t *testing.T) {
		point := Point(1.2, 2.3, 3.4)

		assert.Equal(t, Tuple{1.2, 2.3, 3.4, 1.0}, point)
	})
}

func TestVector(t *testing.T) {
	t.Run("creates a tuple with W=0", func(t *testing.T) {
		point := Vector(1.2, 2.3, 3.4)

		assert.Equal(t, Tuple{1.2, 2.3, 3.4, 0.0}, point)
	})
}
