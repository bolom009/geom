package geom

import (
	"math"
)

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

func NewQuaternionFromMatrix3x3(m Matrix3x3) Quaternion {
	var q Quaternion

	m00, m01, m02 := m.c0.X, m.c0.Y, m.c0.Z
	m10, m11, m12 := m.c1.X, m.c1.Y, m.c1.Z
	m20, m21, m22 := m.c2.X, m.c2.Y, m.c2.Z

	trace := m00 + m11 + m22

	if trace > 0 {
		s := 0.5 / float32(math.Sqrt(float64(trace+1.0)))
		q.W = 0.25 / s
		q.X = (m21 - m12) * s
		q.Y = (m02 - m20) * s
		q.Z = (m10 - m01) * s
	} else {
		switch {
		case m00 > m11 && m00 > m22:
			s := 2.0 * float32(math.Sqrt(float64(1.0+m00-m11-m22)))
			q.W = (m21 - m12) / s
			q.X = 0.25 * s
			q.Y = (m01 + m10) / s
			q.Z = (m02 + m20) / s
		case m11 > m22:
			s := 2.0 * float32(math.Sqrt(float64(1.0+m11-m00-m22)))
			q.W = (m02 - m20) / s
			q.X = (m01 + m10) / s
			q.Y = 0.25 * s
			q.Z = (m12 + m21) / s
		default:
			s := 2.0 * float32(math.Sqrt(float64(1.0+m22-m00-m11)))
			q.W = (m10 - m01) / s
			q.X = (m02 + m20) / s
			q.Y = (m12 + m21) / s
			q.Z = 0.25 * s
		}
	}

	// Ensure unit quaternion (compensate for potential precision loss)
	l := float32(math.Sqrt(float64(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z)))
	if l > 0 {
		lenInv := 1.0 / l
		q.W *= lenInv
		q.X *= lenInv
		q.Y *= lenInv
		q.Z *= lenInv
	}

	q.W *= -1

	return q
}

func (q Quaternion) Set(w, x, y, z float32) Quaternion {
	return Quaternion{x, y, z, w}
}

func (q Quaternion) MulV3(v Vector3) Vector3 {
	t := CrossV3(q.XYZ(), v).Mul(2)
	return v.Add(t.Mul(q.W)).Add(CrossV3(q.XYZ(), t))
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{q.X * -1, q.Y * -1, q.Z * -1, q.W * 1}
}

func (q Quaternion) XYZ() Vector3 {
	return Vector3{
		X: q.X,
		Y: q.Y,
		Z: q.Z,
	}
}
