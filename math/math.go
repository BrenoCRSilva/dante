package logic

// Vector2 is a struct that represents a 2D vector
type Vector struct {
	X float64
	Y float64
}

// Add adds two vectors
func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
