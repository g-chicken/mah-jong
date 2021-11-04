package device_test

import (
	"os"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/google/go-cmp/cmp"
)

var allowUnexported = cmp.AllowUnexported(domain.Config{})

func MainTest(m *testing.M) {
	os.Exit(m.Run())
}

func operateEnv(envs map[string]string) func() {
	setEnv := map[string]string{}

	for k, v := range envs {
		e, exist := os.LookupEnv(k)
		if exist {
			setEnv[k] = e
		}

		_ = os.Setenv(k, v)
	}

	return func() {
		for k := range envs {
			e, ok := setEnv[k]
			if ok {
				_ = os.Setenv(k, e)

				continue
			}

			_ = os.Unsetenv(k)
		}
	}
}
