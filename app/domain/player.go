package domain

import "context"

// GetPlayerByName gets a Player by the name.
func GetPlayerByName(c context.Context, name string) (*Player, error) {
	return repos.PlayerRepository.GetPlayerByName(c, name)
}

// CreatePlayer creates a Player in DB.
func CreatePlayer(c context.Context, name string) (*Player, error) {
	return repos.PlayerRepository.CreatePlayer(c, name)
}

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
