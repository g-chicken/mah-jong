package domain

// Player expresses the player model.
type Player struct {
	id   uint64
	name string
}

// NewPlayer creates *Player.
func NewPlayer(id uint64, name string) *Player {
	return &Player{
		id:   id,
		name: name,
	}
}

// GetID returns the player ID.
func (p *Player) GetID() uint64 {
	if p == nil {
		return 0
	}

	return p.id
}

// GetName returns the player name.
func (p *Player) GetName() string {
	if p == nil {
		return ""
	}

	return p.name
}
