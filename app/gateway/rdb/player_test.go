package rdb_test

import (
	"context"
	"errors"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/google/go-cmp/cmp"
)

func TestPlayerRepository_GetPlayerByName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		argName  string
		want     *domain.Player
		errFunc  func(err error) bool
	}{
		{
			testName: "success",
			argName:  "hoge",
			want:     domain.NewPlayer(2, "hoge"),
			errFunc:  notErrFunc,
		},
		{
			testName: "not found",
			argName:  "testtesttest",
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.NotFoundError{})
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			repo := rdb.NewPlayerRepository(rdbGetterRepo)
			got, err := repo.GetPlayerByName(context.Background(), tc.argName)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerRepository_CreatePlayer(t *testing.T) {
	testCases := []struct {
		testName string
		argName  string
		want     *domain.Player
		errFunc  func(err error) bool
	}{
		{
			testName: "success",
			argName:  "testhoge",
			want:     domain.NewPlayer(6, "testhoge"),
			errFunc:  notErrFunc,
		},
		{
			testName: "conflict",
			argName:  "hoge",
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.ConflictError{})
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			repo := rdb.NewPlayerRepository(rdbGetterRepo)
			got, err := repo.CreatePlayer(context.Background(), tc.argName)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		initializePlayers()
	}
}
