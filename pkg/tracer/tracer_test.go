package tracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	epsilon = 0.00001
)

func TestNewPoint(t *testing.T) {
	t.Run("creates a tuple with W=1", func(t *testing.T) {
		point := NewPoint(1.2, 2.3, 3.4)

		assertTupleEqual(t, Tuple{1.2, 2.3, 3.4, 1.0}, point)
	})
}

func TestNewVector(t *testing.T) {
	t.Run("creates a tuple with W=0", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)

		assertTupleEqual(t, Tuple{1.2, 2.3, 3.4, 0.0}, vector)
	})
}

func TestIsPoint(t *testing.T) {
	t.Run("returns true for a point", func(t *testing.T) {
		point := NewPoint(1.2, 2.3, 3.4)

		assert.True(t, point.IsPoint())
	})

	t.Run("returns false for a vector", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)

		assert.False(t, vector.IsPoint())
	})
}

func TestIsVector(t *testing.T) {
	t.Run("returns true for a vector", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)

		assert.True(t, vector.IsVector())
	})

	t.Run("returns false for a point", func(t *testing.T) {
		point := NewPoint(1.2, 2.3, 3.4)

		assert.False(t, point.IsVector())
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds a vector and a point", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)
		point := NewPoint(4.5, 5.6, 6.7)

		result, err := vector.Add(point)
		require.NoError(t, err)

		assertTupleEqual(t, NewPoint(5.7, 7.9, 10.1), result)
	})

	t.Run("adds a point and a vector", func(t *testing.T) {
		point := NewVector(1.2, 2.3, 3.4)
		vector := NewPoint(4.5, 5.6, 6.7)

		result, err := point.Add(vector)
		require.NoError(t, err)

		assertTupleEqual(t, NewPoint(5.7, 7.9, 10.1), result)
	})

	t.Run("adds two vectors", func(t *testing.T) {
		vector1 := NewVector(1.2, 2.3, 3.4)
		vector2 := NewVector(4.5, 5.6, 6.7)

		result, err := vector1.Add(vector2)
		require.NoError(t, err)

		assertTupleEqual(t, NewVector(5.7, 7.9, 10.1), result)
	})

	t.Run("cannot add two points", func(t *testing.T) {
		point1 := NewPoint(1.2, 2.3, 3.4)
		point2 := NewPoint(4.5, 5.6, 6.7)

		result, err := point1.Add(point2)

		assert.ErrorIs(t, err, ErrCannotAddPoints)
		assert.Zero(t, result)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtracts a vector from a point", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)
		point := NewPoint(6.7, 5.6, 4.5)

		result, err := point.Subtract(vector)
		require.NoError(t, err)

		assertTupleEqual(t, NewPoint(5.5, 3.3, 1.1), result)
	})

	t.Run("subtracts a vector from a vector", func(t *testing.T) {
		vector1 := NewVector(1.2, 2.3, 3.4)
		vector2 := NewVector(6.7, 5.6, 4.5)

		result, err := vector2.Subtract(vector1)
		require.NoError(t, err)

		assertTupleEqual(t, NewVector(5.5, 3.3, 1.1), result)
	})

	t.Run("cannot subtract a point from a point", func(t *testing.T) {
		point1 := NewPoint(1.2, 2.3, 3.4)
		point2 := NewPoint(6.7, 5.6, 4.5)

		result, err := point2.Subtract(point1)

		assert.ErrorIs(t, err, ErrCannotSubtractPoint)
		assert.Zero(t, result)
	})

	t.Run("cannot subtract a point from a vector", func(t *testing.T) {
		point := NewPoint(1.2, 2.3, 3.4)
		vector := NewVector(6.7, 5.6, 4.5)

		result, err := vector.Subtract(point)

		assert.ErrorIs(t, err, ErrCannotSubtractPoint)
		assert.Zero(t, result)
	})
}

func TestNegate(t *testing.T) {
	t.Run("negates a vector", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)

		result, err := vector.Negate()
		require.NoError(t, err)

		assertTupleEqual(t, NewVector(-1.2, -2.3, -3.4), result)
	})

	t.Run("cannot negate a point", func(t *testing.T) {
		point := NewPoint(1.2, 2.3, 3.4)

		result, err := point.Negate()

		assert.ErrorIs(t, err, ErrCannotNegatePoint)
		assert.Zero(t, result)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("multiplies a vector by a scalar", func(t *testing.T) {
		vector := NewVector(1.2, 2.3, 3.4)

		result, err := vector.Multiply(2.0)
		require.NoError(t, err)

		assertTupleEqual(t, NewVector(2.4, 4.6, 6.8), result)
	})
}

func TestDivide(t *testing.T) {
	t.Run("divides a vector by a scalar", func(t *testing.T) {
		vector := NewVector(2.4, 4.6, 6.8)

		result, err := vector.Divide(2.0)
		require.NoError(t, err)

		assertTupleEqual(t, NewVector(1.2, 2.3, 3.4), result)
	})
}

func TestMagnitude(t *testing.T) {
	t.Run("calculates the magnitude of a vector", func(t *testing.T) {
		vector := NewVector(2.0, 3.0, 6.0)

		result, err := vector.Magnitude()
		require.NoError(t, err)

		assert.InDelta(t, 7.0, result, epsilon)
	})

	t.Run("cannot calculate the magnitude of a point", func(t *testing.T) {
	})
}

func assertTupleEqual(t *testing.T, expected, actual Tuple) {
	t.Helper()

	assert.InDelta(t, expected.X, actual.X, epsilon)
	assert.InDelta(t, expected.Y, actual.Y, epsilon)
	assert.InDelta(t, expected.Z, actual.Z, epsilon)
	assert.Equal(t, expected.W, actual.W)
}
