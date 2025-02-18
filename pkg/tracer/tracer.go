package tracer

import (
	"errors"
	"math"
)

const (
	pointW  = float64(1)
	vectorW = float64(0)
)

var (
	ErrCannotAddPoints      = errors.New("cannot add two points")
	ErrCannotDotPoint       = errors.New("cannot dot a point")
	ErrCannotMeasurePoint   = errors.New("cannot measure a point")
	ErrCannotNegatePoint    = errors.New("cannot negate a point")
	ErrCannotNormalizePoint = errors.New("cannot normalize a point")
	ErrCannotSubtractPoint  = errors.New("cannot subtract a point")
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, pointW}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, vectorW}
}

func (t Tuple) IsPoint() bool {
	return t.W == pointW
}

func (t Tuple) IsVector() bool {
	return t.W == vectorW
}

func (t Tuple) Add(other Tuple) (Tuple, error) {
	if t.IsPoint() && other.IsPoint() {
		return Tuple{}, ErrCannotAddPoints
	}

	return Tuple{
		X: t.X + other.X,
		Y: t.Y + other.Y,
		Z: t.Z + other.Z,
		W: t.W + other.W,
	}, nil
}

func (t Tuple) Subtract(other Tuple) (Tuple, error) {
	if other.IsPoint() {
		return Tuple{}, ErrCannotSubtractPoint
	}

	return Tuple{
		X: t.X - other.X,
		Y: t.Y - other.Y,
		Z: t.Z - other.Z,
		W: t.W - other.W,
	}, nil
}

func (t Tuple) Negate() (Tuple, error) {
	if t.IsPoint() {
		return Tuple{}, ErrCannotNegatePoint
	}

	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}, nil
}

func (t Tuple) Multiply(scalar float64) (Tuple, error) {
	return Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W,
	}, nil
}

func (t Tuple) Divide(scalar float64) (Tuple, error) {
	return Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W,
	}, nil
}

func (t Tuple) Magnitude() (float64, error) {
	if t.IsPoint() {
		return 0, ErrCannotMeasurePoint
	}

	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2)), nil
}

func (t Tuple) Normalize() (Tuple, error) {
	if t.IsPoint() {
		return Tuple{}, ErrCannotNormalizePoint
	}

	magnitude, _ := t.Magnitude()

	return t.Divide(magnitude)
}

func (t Tuple) Dot(other Tuple) (float64, error) {
	if t.IsPoint() || other.IsPoint() {
		return 0, ErrCannotDotPoint
	}

	return t.X*other.X + t.Y*other.Y + t.Z*other.Z, nil
}
