package data

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewPagedList(t *testing.T) {
	pagedList := NewPagedList([]int{0, 1, 2}, 10, 3, 2)
	assert.Len(t, pagedList.List, 3)
	assert.EqualValues(t, 3, pagedList.Take)
	assert.EqualValues(t, 2, pagedList.Page)
	assert.EqualValues(t, 10, pagedList.TotalCount)

	assert.EqualValues(t, 4, pagedList.TotalPages)
	assert.EqualValues(t, 6, pagedList.Skip)
	assert.EqualValues(t, true, pagedList.HasPrevious)
	assert.EqualValues(t, true, pagedList.HasNext)
}

func TestNewPagedListWithProjection(t *testing.T) {
	pagedList := NewPagedListWithProjection([]int{0, 1, 2}, 10, 3, 2, func(t *int) *string {
		res := "Num-" + strconv.Itoa(*t)
		return &res
	})
	assert.Len(t, pagedList.List, 3)
	assert.EqualValues(t, 3, pagedList.Take)
	assert.EqualValues(t, 2, pagedList.Page)
	assert.EqualValues(t, 10, pagedList.TotalCount)

	assert.EqualValues(t, 4, pagedList.TotalPages)
	assert.EqualValues(t, 6, pagedList.Skip)
	assert.EqualValues(t, true, pagedList.HasPrevious)
	assert.EqualValues(t, true, pagedList.HasNext)

	assert.Equal(t, pagedList.List[0], "Num-0")
	assert.Equal(t, pagedList.List[1], "Num-1")
	assert.Equal(t, pagedList.List[2], "Num-2")
}
