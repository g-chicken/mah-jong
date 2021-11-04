package usecase_test

import (
	"os"
	"testing"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/google/go-cmp/cmp"
)

var allowUnexported = cmp.AllowUnexported(domain.Config{})

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
