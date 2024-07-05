package vectors

type Vec4f32 struct {
	X, Y, Z, W float32
}

func (v Vec4f32) Spread() (float32, float32, float32, float32) { return v.X, v.Y, v.Z, v.W }

type Vec4f64 struct {
	X, Y, Z, W float64
}
