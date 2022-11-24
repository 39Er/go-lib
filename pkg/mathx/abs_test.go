package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, int64(10), Abs(int64(-10)))
	assert.Equal(t, int64(10), Abs(int64(10)))
	assert.Equal(t, int64(0), Abs(int64(0)))
}
