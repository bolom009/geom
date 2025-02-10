package geom

import "math"

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (v Vector3) YZX() Vector3 {
	return Vector3{
		X: v.Y,
		Y: v.Z,
		Z: v.X,
	}
}

func (v Vector3) XZ() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Z,
	}
}

func (v Vector3) MulV(x Vector3) Vector3 {
	return Vector3{X: v.X * x.X, Y: v.Y * x.Y, Z: v.Z * x.Z}
}

func (v Vector3) Len() float32 {
	return float32(math.Hypot(float64(v.X), float64(v.Y)))
}

func (v Vector3) Div(scalar float32) Vector3 {
	return Vector3{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v Vector3) Mul(x float32) Vector3 {
	return Vector3{
		X: v.X * x,
		Y: v.Y * x,
		Z: v.Z * x,
	}
}

func (v Vector3) Add(x Vector3) Vector3 {
	return Vector3{X: v.X + x.X, Y: v.Y + x.Y, Z: v.Z + x.Z}
}

func (v Vector3) Sub(x Vector3) Vector3 {
	return Vector3{X: v.X - x.X, Y: v.Y - x.Y, Z: v.Z - x.Z}
}

func (v Vector3) Dot(b Vector3) float32 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

func (v Vector3) Distance(y Vector3) float32 {
	dx := y.X - v.X
	dy := y.Y - v.Y
	dz := y.Z - v.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

func (v Vector3) Lerp(b Vector3, t float32) Vector3 {
	return Vector3{X: v.X + (b.X-v.X)*t, Y: v.Y + (b.Y-v.Y)*t, Z: v.Z + (b.Z-v.Z)*t}
}

// Normalize the vector (make its length 1)
func (v Vector3) Normalize() Vector3 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector3{0, 0, 0}
	}
	return v.Div(magnitude)
}

func (v Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}
