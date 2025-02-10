package geom

import "math"

const epsilon = 1e-5

type Vector2 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func (v Vector2) Sub(q Vector2) Vector2 {
	return Vector2{v.X - q.X, v.Y - q.Y}
}

func (v Vector2) Scale(s float32) Vector2 {
	return Vector2{X: v.X * s, Y: v.Y * s}
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vector2) Div(s float32) Vector2 {
	return Vector2{X: v.X / s, Y: v.Y / s}
}

func (v Vector2) Mul(s float32) Vector2 {
	return Vector2{X: v.X * s, Y: v.Y * s}
}

func (v Vector2) MulV(v2 Vector2) Vector2 {
	return Vector2{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vector2) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Perpendicular() Vector2 {
	return Vector2{X: -v.Y, Y: v.X} // 90-degree rotation
}

func (v Vector2) Len() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector2) Unit() Vector2 {
	if v.X == 0 && v.Y == 0 {
		return Vector2{X: 0, Y: 0}
	}
	return v.Scale(1 / v.NormEuclidean())
}

func (v Vector2) Normalize() Vector2 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2{0, 0}
	}
	return Vector2{v.X / mag, v.Y / mag}
}

func (v Vector2) NormEuclidean() float32 {
	return float32(math.Hypot(float64(v.X), float64(v.Y)))
}

func (v Vector2) Dot(b Vector2) float32 {
	return v.X*b.X + v.Y*b.Y
}

func (v Vector2) Lerp(b Vector2, t float32) Vector2 {
	return Vector2{X: v.X + (b.X-v.X)*t, Y: v.Y + (b.Y-v.Y)*t}
}

func Distance(a, b Vector2) float32 {
	return float32(math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2)))
}
