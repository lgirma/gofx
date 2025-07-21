package data

import "math"

type PageRequest struct {
	Take int64      `json:"take"`
	Page int64      `json:"page"`
	Sort []SortInfo `json:"sort"`
}

type SortInfo struct {
	ColumnName string `json:"columnName"`
	IsDesc     bool   `json:"isDesc"`
}

func (page PageRequest) IsSorted(by string, isDesc bool) bool {
	for i := range page.Sort {
		s := page.Sort[i]
		if s.ColumnName == by {
			return s.IsDesc == isDesc
		}
	}
	return false
}

type PagedList[T any] struct {
	TotalCount  int64 `json:"totalCount"`
	TotalPages  int64 `json:"totalPages"`
	List        []T   `json:"list"`
	Take        int64 `json:"take"`
	Page        int64 `json:"page"`
	Skip        int64 `json:"skip"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
	LastPage    int64 `json:"lastPage"`
}

func NewPagedList[T any](list []T, totalCount int64, take int64, page int64) *PagedList[T] {
	if take < 0 || page < 0 {
		return &PagedList[T]{}
	}
	totalPages := int64(math.Ceil(float64(totalCount) / float64(take)))
	if page > totalPages-1 {
		page = totalPages - 1
	}
	return &PagedList[T]{
		TotalCount:  totalCount,
		List:        list,
		Take:        take,
		Page:        page,
		Skip:        page * take,
		HasNext:     page < totalPages-1,
		HasPrevious: page > 0,
		TotalPages:  totalPages,
	}
}

func NewPagedListWithProjection[T any, TProjected any](list []T, totalCount int64, take int64, page int64, projection func(*T) *TProjected) *PagedList[TProjected] {
	src := NewPagedList(list, totalCount, take, page)
	result := &PagedList[TProjected]{
		TotalCount:  src.TotalCount,
		TotalPages:  src.TotalPages,
		List:        make([]TProjected, 0),
		Take:        src.Take,
		Page:        src.Page,
		Skip:        src.Skip,
		HasNext:     src.HasNext,
		HasPrevious: src.HasPrevious,
		LastPage:    src.TotalPages - 1,
	}
	for i, _ := range src.List {
		result.List = append(result.List, *projection(&src.List[i]))
	}
	return result
}

func EmptyPagedList[T any]() *PagedList[T] {
	return NewPagedList([]T{}, 0, 0, 0)
}

func PrepareFilter(filter *PageRequest) *PageRequest {
	if filter == nil {
		filter = &PageRequest{}
	}
	if filter.Take == 0 {
		filter.Take = 5
	}
	return filter
}
