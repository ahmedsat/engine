package vectors

type Vec2f32 struct {
	X, Y float32
}

func NewVec2f32(x, y float32) Vec2f32        { return Vec2f32{X: x, Y: y} }
func (v Vec2f32) Spread() (float32, float32) { return v.X, v.Y }

type Vec2f64 struct {
	X, Y float64
}
