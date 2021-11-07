package domain

import "context"

// CreatePlayerHandPairs creates data of players_hands table columns.
func CreatePlayerHandPairs(c context.Context, pairs []*CreatePlayerHandArgs) error {
	if len(pairs) == 0 {
		return nil
	}

	return repos.playerHandRepo.CreatePlayerHandPairs(c, pairs)
}
