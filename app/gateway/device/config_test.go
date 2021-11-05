package device_test

import (
	"context"
	"testing"
	"time"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/device"
	"github.com/google/go-cmp/cmp"
)

func TestConfigRepository_GetConfig(t *testing.T) {
	testCases := []struct {
		name string
		envs map[string]string
		want *domain.Config
		err  bool
	}{
		{
			name: "default",
			want: domain.NewConfig(8080, "localhost:3306", "mah_jong", "app", "hoge", 5*time.Second),
		},
		{
			name: "full",
			envs: map[string]string{
				"MAH_JONG_GRPC_PORT":              "7000",
				"MAH_JONG_RDB_URL":                "mysql:3306",
				"MAH_JONG_RDB_NAME":               "test",
				"MAH_JONG_RDB_USER":               "test_user",
				"MAH_JONG_RDB_PASS":               "test_pass",
				"MAH_JONG_RDB_CONNECTION_TIMEOUT": "10m",
			},
			want: domain.NewConfig(7000, "mysql:3306", "test", "test_user", "test_pass", 10*time.Minute),
		},
		{
			name: "not int",
			envs: map[string]string{
				"MAH_JONG_GRPC_PORT": "hoge",
			},
			err: true,
		},
		{
			name: "invalid duration format",
			envs: map[string]string{
				"MAH_JONG_RDB_CONNECTION_TIMEOUT": "hoge",
			},
			err: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := operateEnv(tc.envs)
			defer reset()

			repo := device.NewConfigRepository()
			got, err := repo.GetConfig(context.Background())

			if tc.err && err == nil {
				t.Fatal("should be error but not")
			}
			if !tc.err && err != nil {
				t.Fatalf("should not be error but %v", err)
			}

			if diff := cmp.Diff(tc.want, got, allowUnexported); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
