package domain

import (
	"context"
)

// GetPlayerByID gets a Player by the player ID.
func GetPlayerByID(c context.Context, playerID uint64) (*Player, error) {
	return repos.playerRepo.GetPlayerByID(c, playerID)
}

// GetPlayerByName gets a Player by the name.
func GetPlayerByName(c context.Context, name string) (*Player, error) {
	return repos.playerRepo.GetPlayerByName(c, name)
}

// GetPlayers gets Players.
func GetPlayers(c context.Context) ([]*Player, error) {
	return repos.playerRepo.GetPlayers(c)
}

// CreatePlayer creates a Player in DB.
func CreatePlayer(c context.Context, name string) (*Player, error) {
	return repos.playerRepo.CreatePlayer(c, name)
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

// SetName set the player's name.
func (p *Player) SetName(name string) {
	if p == nil {
		return
	}

	p.name = name
}

// UpdateName updates the player's name.
func (p *Player) UpdateName(c context.Context, name string) error {
	if p == nil {
		return errNilPlayer
	}

	if p.GetName() == name {
		return nil
	}

	if err := repos.playerRepo.UpdatePlayer(c, p.GetID(), name); err != nil {
		return err
	}

	p.SetName(name)

	return nil
}

// Delete delete the player.
func (p *Player) Delete(c context.Context) error {
	if p == nil {
		return errNilPlayer
	}

	return repos.playerRepo.DeletePlayer(c, p.GetID())
}
