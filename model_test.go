package SC_727

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNums(t *testing.T) {
	result := AddNums()
	assert.Equal(t, 3, result)
}
