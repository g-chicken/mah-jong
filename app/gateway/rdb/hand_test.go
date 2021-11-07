package rdb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/google/go-cmp/cmp"
)

func TestHandRepository_CreateHand(t *testing.T) {
	testCases := []struct {
		name      string
		timestamp time.Time
		want      *domain.Hand
		errFunc   func(error) bool
	}{
		{
			name:      "success",
			timestamp: time.Date(2021, time.November, 6, 0, 0, 0, 0, time.UTC),
			want:      domain.NewHand(4, time.Date(2021, time.November, 6, 0, 0, 0, 0, time.UTC)),
			errFunc:   notErrFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			defer initializeHands()

			c := context.Background()

			repo := rdb.NewHandRepository(rdbDetectorRepo)
			got, err := repo.CreateHand(c, tc.timestamp)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			if err == nil {
				query := "SELECT id, game_date FROM hands WHERE id = ?"
				args := []interface{}{tc.want.GetID()}
				ope := rdbDetectorRepo.GetRDBOperator(c)

				var (
					id        uint64
					timestamp time.Time
				)

				_ = ope.Get(c, query, args, &id, &timestamp)
				hand := domain.NewHand(id, timestamp)

				if diff := cmp.Diff(tc.want, hand, allowUnexported); diff != "" {
					t.Fatalf("unexpected result (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestHandRepository_CreateHand_transaction(t *testing.T) {
	testCases := []struct {
		name      string
		timestamp time.Time
		want      *domain.Hand
		errFunc   func(error) bool
	}{
		{
			name:      "success",
			timestamp: time.Date(2021, time.November, 6, 0, 0, 0, 0, time.UTC),
			want:      domain.NewHand(5, time.Date(2021, time.November, 6, 0, 0, 0, 0, time.UTC)),
			errFunc:   notErrFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			defer initializeHands()

			c := context.Background()

			err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewHandRepository(rdbDetectorRepo)
				got, err := repo.CreateHand(c, tc.timestamp)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			})
			if err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if err == nil {
				query := "SELECT id, game_date FROM hands WHERE id = ?"
				args := []interface{}{tc.want.GetID()}
				ope := rdbDetectorRepo.GetRDBOperator(c)

				var (
					id        uint64
					timestamp time.Time
				)

				_ = ope.Get(c, query, args, &id, &timestamp)
				hand := domain.NewHand(id, timestamp)

				if diff := cmp.Diff(tc.want, hand, allowUnexported); diff != "" {
					t.Fatalf("unexpected result (-want +got):\n%s", diff)
				}
			}
		})
	}
}
