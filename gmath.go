// Package geom (gmath.go) it's a small appendix for working with vector and other math
package geom

import (
	"math"
)

const (
	ToRadians    = 0.017453292519943296
	FltMinNormal = 1.175494351e-38
	Epsilon      = 1.1920928955078125e-7
)

func IsFinite(v float64) bool {
	return !math.IsInf(v, 0)
}

func Radians(x float32) float32 {
	return x * float32(ToRadians)
}

func LengthV3(v Vector3) float32 {
	return float32(math.Sqrt(float64(DotV3(v, v))))
}

func DotV3(a, b Vector3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func CrossV3(x, y Vector3) Vector3 {
	v0X := x.X * y.Y
	v0Y := x.Y * y.Z
	v0Z := x.Z * y.X

	v1X := x.Y * y.X
	v1Y := x.Z * y.Y
	v1Z := x.X * y.Z

	return Vector3{X: v0Y - v1Y, Y: v0Z - v1Z, Z: v0X - v1X}

	//v0 := x.MulV(y.YZX())
	//v1 := x.YZX().MulV(y)
	//return v0.Sub(v1).YZX()
}

func Rsqrt(x float32) float32 {
	return 1.0 / float32(math.Sqrt(float64(x)))
}

func NormalizeV3(x Vector3) Vector3 {
	return x.Mul(Rsqrt(DotV3(x, x)))
}

func NormalizeSafeV2(x Vector2) Vector2 {
	dVal := Vector2{X: 0, Y: 0}
	l := x.Dot(x)

	return Select(dVal, x.Mul(Rsqrt(l)), l > FltMinNormal)
}

func NormalizeSafeV3(x Vector3) Vector3 {
	dVal := Vector3{X: 0, Y: 0, Z: 0}
	l := DotV3(x, x)

	return SelectV3(dVal, x.Mul(Rsqrt(l)), l > FltMinNormal)
}

func Select(falseValue, trueValue Vector2, test bool) Vector2 {
	if test {
		return trueValue
	}

	return falseValue
}

func SelectV3(falseValue, trueValue Vector3, test bool) Vector3 {
	if test {
		return trueValue
	}

	return falseValue
}

func Clamp(valueToClamp, lowerBound, upperBound float32) float32 {
	return Max(lowerBound, Min(upperBound, valueToClamp))
}

func Min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
