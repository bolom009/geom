package geom

type Matrix3x3 struct {
	c0 Vector3
	c1 Vector3
	c2 Vector3
}

func NewFloat3x3(c0, c1, c2 Vector3) Matrix3x3 {
	return Matrix3x3{
		c0: c0,
		c1: c1,
		c2: c2,
	}
}
