package rdb_test

import (
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
)

func TestNewRDBDetectorRepository(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config *domain.Config
		err    bool
	}{
		{
			name: "success",
			config: domain.NewConfig(
				8080, "localhost:3306", "mah_jong_test", "app", "hoge", 5*time.Second,
			),
		},
		{
			name: "not connected",
			config: domain.NewConfig(
				8080, "hoge:3306", "mah_jong_test", "app", "hoge", 1*time.Second,
			),
			err: true,
		},
		{
			name: "invalid db name",
			config: domain.NewConfig(
				8080, "localhot:3306", "test", "app", "hoge", 1*time.Second,
			),
			err: true,
		},
		{
			name: "invalid db user",
			config: domain.NewConfig(
				8080, "localhost:3306", "mah_jong_test", "test", "hoge", 1*time.Second,
			),
			err: true,
		},
		{
			name: "invalid db pass",
			config: domain.NewConfig(
				8080, "localhost:3306", "mah_jong_test", "app", "test", 1*time.Second,
			),
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, _, closeFunc, err := rdb.NewRDBDetectorRepository(tc.config)
			defer closeFunc()

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}
		})
	}
}
