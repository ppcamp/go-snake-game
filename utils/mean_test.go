package utils_test

import (
	"go-spyder-game/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

// Median gets the median number in a slice of numbers
func TestMedian(t *testing.T) {
	assert := require.New(t)

	assert.Equal(utils.Median(7), 3)
}
