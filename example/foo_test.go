package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	t.Logf("~> mock Foo")
	// mock Foo

	t.Logf("~> do Foo")
	// do Foo

	// verify Foo
	assert.Equal(t, 3, Foo(1, 2))
}
