package domain_test

import (
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
)

func TestHalfRoundGameScores_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		halfRoundGameScores domain.HalfRoundGameScores
		want                bool
	}{
		{
			name: "success",
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: true,
		},
		{
			name: "over 4",
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
					domain.NewPlayerScore(5, 0, 5),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "invalid ranking",
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 3),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 2),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "sum is not 0",
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 11, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(3, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
		{
			name: "same player ID",
			halfRoundGameScores: map[uint32][]*domain.PlayerScore{
				1: {
					domain.NewPlayerScore(1, 20, 1),
					domain.NewPlayerScore(2, 10, 2),
					domain.NewPlayerScore(3, -30, 3),
				},
				2: {
					domain.NewPlayerScore(1, -20, 3),
					domain.NewPlayerScore(2, 50, 1),
					domain.NewPlayerScore(3, 10, 2),
					domain.NewPlayerScore(4, -40, 4),
				},
				3: {
					domain.NewPlayerScore(1, 12, 2),
					domain.NewPlayerScore(2, -36, 4),
					domain.NewPlayerScore(1, -21, 3),
					domain.NewPlayerScore(4, 45, 1),
				},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.halfRoundGameScores.Validate()
			if tc.want != got {
				t.Fatalf("unexpected result (want : %v, got : %v)", tc.want, got)
			}
		})
	}
}
