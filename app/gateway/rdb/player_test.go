package rdb_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/google/go-cmp/cmp"
)

func TestPlayerRepository_CreatePlayer_normal(t *testing.T) {
	testCases := []struct {
		testName string
		argName  string
		want     *domain.Player
		errFunc  func(error) bool
	}{
		{
			testName: "success",
			argName:  "testhoge",
			want:     domain.NewPlayer(6, "testhoge"),
			errFunc:  notErrFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			defer initializePlayers()

			c := context.Background()

			repo := rdb.NewPlayerRepository(rdbDetectorRepo)
			got, err := repo.CreatePlayer(c, tc.argName)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}

			query := "SELECT id, name FROM players WHERE id = ?"
			args := []interface{}{tc.want.GetID()}
			ope := rdbDetectorRepo.GetRDBOperator(c)

			var (
				id   uint64
				name string
			)

			_ = ope.Get(c, query, args, &id, &name)
			p := domain.NewPlayer(id, name)

			if diff := cmp.Diff(tc.want, p, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerRepository_CreatePlayer_transaction(t *testing.T) {
	testCases := []struct {
		testName string
		argName  string
		want     *domain.Player
		errFunc  func(error) bool
	}{
		{
			testName: "success",
			argName:  "testhoge",
			want:     domain.NewPlayer(7, "testhoge"),
			errFunc:  notErrFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			defer initializePlayers()

			c := context.Background()

			if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewPlayerRepository(rdbDetectorRepo)
				got, err := repo.CreatePlayer(c, tc.argName)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			query := "SELECT id, name FROM players WHERE id = ?"
			args := []interface{}{tc.want.GetID()}
			ope := rdbDetectorRepo.GetRDBOperator(c)

			var (
				id   uint64
				name string
			)

			_ = ope.Get(c, query, args, &id, &name)
			p := domain.NewPlayer(id, name)

			if diff := cmp.Diff(tc.want, p, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPlayerRepository_CreatePlayer_error(t *testing.T) {
	testCases := []struct {
		testName string
		argName  string
		want     *domain.Player
		errFunc  func(error) bool
	}{
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
			defer initializePlayers()

			c := context.Background()

			repo := rdb.NewPlayerRepository(rdbDetectorRepo)
			got, err := repo.CreatePlayer(c, tc.argName)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		t.Run(tc.testName+"(transaction)", func(t *testing.T) {
			defer initializePlayers()

			c := context.Background()

			if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewPlayerRepository(rdbDetectorRepo)
				got, err := repo.CreatePlayer(c, tc.argName)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}

func TestPlayerRepository_GetPlayerByID(t *testing.T) { // nolint:dupl
	t.Parallel()

	testCases := []struct {
		name    string
		id      uint64
		want    *domain.Player
		errFunc func(err error) bool
	}{
		{
			name:    "success",
			id:      2,
			want:    domain.NewPlayer(2, "hoge"),
			errFunc: notErrFunc,
		},
		{
			name: "not found",
			id:   99,
			errFunc: func(err error) bool {
				return !errors.As(err, &domain.NotFoundError{})
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repo := rdb.NewPlayerRepository(rdbDetectorRepo)
			got, err := repo.GetPlayerByID(context.Background(), tc.id)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		t.Run(tc.name+"(transaction)", func(t *testing.T) {
			t.Parallel()

			c := context.Background()

			if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewPlayerRepository(rdbDetectorRepo)
				got, err := repo.GetPlayerByID(c, tc.id)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}

func TestPlayerRepository_GetPlayerByName(t *testing.T) { // nolint:dupl
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

			repo := rdb.NewPlayerRepository(rdbDetectorRepo)
			got, err := repo.GetPlayerByName(context.Background(), tc.argName)

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		t.Run(tc.testName+"(transaction)", func(t *testing.T) {
			t.Parallel()

			c := context.Background()

			if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewPlayerRepository(rdbDetectorRepo)
				got, err := repo.GetPlayerByName(c, tc.argName)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}

func TestPlayerRepository_GetPlayers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		want    []*domain.Player
		errFunc func(err error) bool
	}{
		{
			name: "success",
			want: []*domain.Player{
				domain.NewPlayer(1, "test"),
				domain.NewPlayer(2, "hoge"),
				domain.NewPlayer(3, "foo"),
				domain.NewPlayer(4, "bar"),
				domain.NewPlayer(5, "fuga"),
			},
			errFunc: notErrFunc,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			repo := rdb.NewPlayerRepository(rdbDetectorRepo)
			got, err := repo.GetPlayers(context.Background())

			if tc.errFunc(err) {
				t.Fatalf("unexpected error (error = %v)", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})

		t.Run(tc.name+"(transaction)", func(t *testing.T) {
			t.Parallel()

			c := context.Background()

			if err := rdbStatementSetRepo.Transaction(c, func(c context.Context) error {
				repo := rdb.NewPlayerRepository(rdbDetectorRepo)
				got, err := repo.GetPlayers(c)

				if tc.errFunc(err) {
					return fmt.Errorf("unexpected error (error = %w)", err)
				}

				if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
					return fmt.Errorf("unexpected result (-want +got):\n%s", diff)
				}

				return nil
			}); err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}
