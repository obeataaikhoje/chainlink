package mathutil

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// Happy path
	assert.Equal(t, 3, Max(3, 2, 1))
	// Single element
	assert.Equal(t, 3, Max(3))
	// Signed
	assert.Equal(t, -1, Max(-2, -1))
	// Uint64
	assert.Equal(t, uint64(2), Max(uint64(0), uint64(2)))
	// String
	assert.Equal(t, "c", Max("a", []string{"b", "c"}...))
}

func TestMin(t *testing.T) {
	// Happy path
	assert.Equal(t, 1, Min(3, 2, 1))
	// Single element
	assert.Equal(t, 3, Min(3))
	// Signed
	assert.Equal(t, -2, Min(-2, -1))
	// Uint64
	assert.Equal(t, uint64(0), Min(uint64(0), uint64(2)))
	// String
	assert.Equal(t, "a", Min("a", []string{"b", "c"}...))
}

func TestAvg(t *testing.T) {
	// happy path
	r, err := Avg(int8(1), -2, 4)
	assert.NoError(t, err)
	assert.Equal(t, int8(1), r)

	// single element
	r, err = Avg(int8(0))
	assert.NoError(t, err)
	assert.Equal(t, int8(0), r)

	// overflow addition
	r, err = Avg(int8(math.MaxInt8), 1)
	assert.ErrorContains(t, err, fmt.Sprintf("overflow: addition"))
	r, err = Avg(int8(math.MinInt8), -1)
	assert.ErrorContains(t, err, fmt.Sprintf("overflow: addition"))

	// overflow length
	a := make([]int8, 256)
	r, err = Avg(a...)
	assert.ErrorContains(t, err, "overflow: array len")
}

func TestMed(t *testing.T) {
	// happy path len = odd
	v, err := Med(2, 1, 5, 4, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, v)

	// happy path len = even
	v, err = Med(10, 11, 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 6, v)

	// zero input
	v, err = Med[int]()
	assert.Error(t, err)
	assert.Equal(t, 0, v)
}
