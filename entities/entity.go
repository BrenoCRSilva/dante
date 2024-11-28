package boxvszombies

// Entity is a struct that represents a generic entity in the game

type Entity struct {
	Position     Vector2
	Velocity     Vector2
	Sprite       *Sprite
	isHidden     bool
	isRotateable bool
}

func NewEntity() {
	return Entity{
		Position:     Vector2{X: 0, Y: 0},
		Velocity:     Vector2{X: 0, Y: 0},
		Sprite:       sprite,
		isHidden:     false,
		isRotateable: false,
	}
}
