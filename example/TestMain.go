//go:build !test

package example

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup
	runRes := m.Run()
	// teardown
	os.Exit(runRes)
}
