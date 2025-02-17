package tracer

const (
	PointW  = float64(1.0)
	VectorW = float64(0.0)
)

type Tuple struct {
	X, Y, Z, W float64
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, PointW}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, VectorW}
}

func (t *Tuple) IsPoint() bool {
	return t.W == PointW
}

func (t *Tuple) IsVector() bool {
	return t.W == VectorW
}

func (t *Tuple) Add(other Tuple) Tuple {
	return Tuple{
		X: t.X + other.X,
		Y: t.Y + other.Y,
		Z: t.Z + other.Z,
		W: t.W + other.W,
	}
}

func (t *Tuple) Subtract(other Tuple) Tuple {
	return Tuple{
		X: t.X - other.X,
		Y: t.Y - other.Y,
		Z: t.Z - other.Z,
		W: t.W - other.W,
	}
}

func (t *Tuple) Negate() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func (t *Tuple) Multiply(scalar float64) Tuple {
	return Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W * scalar,
	}
}

func (t *Tuple) Divide(scalar float64) Tuple {
	return Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W / scalar,
	}
}
