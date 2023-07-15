package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLog(t *testing.T) {
	t.Logf("~> mock InitLog")
	// mock InitLog

	t.Logf("~> do InitLog")
	// do InitLog
	err := InitLog(true, true)
	assert.Nil(t, err)
}
