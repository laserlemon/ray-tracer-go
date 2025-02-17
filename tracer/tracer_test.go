package tracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	t.Run("creates a tuple with W=1", func(t *testing.T) {
		point := Point(1.2, 2.3, 3.4)

		assertTupleEqual(t, Tuple{1.2, 2.3, 3.4, 1.0}, point)
	})
}

func TestVector(t *testing.T) {
	t.Run("creates a tuple with W=0", func(t *testing.T) {
		point := Vector(1.2, 2.3, 3.4)

		assertTupleEqual(t, Tuple{1.2, 2.3, 3.4, 0.0}, point)
	})
}

func TestIsPoint(t *testing.T) {
	t.Run("returns true for a point", func(t *testing.T) {
		point := Point(1.2, 2.3, 3.4)

		assert.True(t, point.IsPoint())
	})

	t.Run("returns false for a vector", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)

		assert.False(t, vector.IsPoint())
	})
}

func TestIsVector(t *testing.T) {
	t.Run("returns true for a vector", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)

		assert.True(t, vector.IsVector())
	})

	t.Run("returns false for a point", func(t *testing.T) {
		point := Point(1.2, 2.3, 3.4)

		assert.False(t, point.IsVector())
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds a vector and a point", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)
		point := Point(4.5, 5.6, 6.7)

		result := vector.Add(point)

		assertTupleEqual(t, Point(5.7, 7.9, 10.1), result)
	})

	t.Run("adds two vectors", func(t *testing.T) {
		vector1 := Vector(1.2, 2.3, 3.4)
		vector2 := Vector(4.5, 5.6, 6.7)

		result := vector1.Add(vector2)

		assertTupleEqual(t, Vector(5.7, 7.9, 10.1), result)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtracts a vector from a point", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)
		point := Point(6.7, 5.6, 4.5)

		result := point.Subtract(vector)

		assertTupleEqual(t, Point(5.5, 3.3, 1.1), result)
	})

	t.Run("subtracts a vector from a vector", func(t *testing.T) {
		vector1 := Vector(1.2, 2.3, 3.4)
		vector2 := Vector(6.7, 5.6, 4.5)

		result := vector2.Subtract(vector1)

		assertTupleEqual(t, Vector(5.5, 3.3, 1.1), result)
	})
}

func TestNegate(t *testing.T) {
	t.Run("negates a vector", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)

		result := vector.Negate()

		assertTupleEqual(t, Vector(-1.2, -2.3, -3.4), result)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("multiplies a vector by a scalar", func(t *testing.T) {
		vector := Vector(1.2, 2.3, 3.4)

		result := vector.Multiply(2.0)

		assertTupleEqual(t, Vector(2.4, 4.6, 6.8), result)
	})
}

func TestDivide(t *testing.T) {
	t.Run("divides a vector by a scalar", func(t *testing.T) {
		vector := Vector(2.4, 4.6, 6.8)

		result := vector.Divide(2.0)

		assertTupleEqual(t, Vector(1.2, 2.3, 3.4), result)
	})
}

const epsilon = 0.00001

func assertTupleEqual(t *testing.T, expected, actual Tuple) {
	t.Helper()

	assert.InDelta(t, expected.X, actual.X, epsilon)
	assert.InDelta(t, expected.Y, actual.Y, epsilon)
	assert.InDelta(t, expected.Z, actual.Z, epsilon)
	assert.Equal(t, expected.W, actual.W)
}
