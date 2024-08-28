package utils

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Data         interface{}                `json:"data"`
	PreloadFunc  func(db *gorm.DB) *gorm.DB `json:"-"`
	Page         int32                      `json:"page"`
	PageSize     int32                      `json:"limit"`
	TotalEntries int32                      `json:"total"`
	TotalPages   int32                      `json:"totalPages"`
	Next         int32                      `json:"next"`
	Prev         int32                      `json:"prev"`
}

func HandlePaginate(statementResult *gorm.DB, model interface{}, page, pageSize int32) (Pagination, error) {
	var totalEntries int64
	//Necessary define model because gorm will failed if there is no model or table set. Moreover he will consider deleted items in case of model is not defined.
	r := statementResult.Model(model).Count(&totalEntries)
	if r.Error != nil {
		return Pagination{}, r.Error
	}

	totalPages := math.Ceil((float64(totalEntries) / float64(pageSize)))
	if totalPages < 1 {
		totalPages = 1
	}
	pagination := Pagination{
		Page:         page,
		PageSize:     pageSize,
		TotalEntries: int32(totalEntries),
		TotalPages:   int32(totalPages),
	}
	pagination.setPreloadFunction(page, pageSize)
	pagination.setNextAndPrevious()
	return pagination, nil
}

func SetPaginate(page, pageSize int32) (int32, int32) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	return page, pageSize
}

// GetNext get the next page of pagination
func (p Pagination) GetNext() int32 {
	var maxPage int32 = p.TotalPages
	var expectedPage int32 = p.Page + 1
	nextPage := math.Min(float64(maxPage), float64(expectedPage))
	return int32(nextPage)
}

// GetPrevious get the previous page of pagination
func (p Pagination) GetPrevious() int32 {
	var minPage int32 = 1
	var expectedPage int32 = p.Page - 1
	previousPage := math.Max(float64(minPage), float64(expectedPage))
	return int32(previousPage)
}

// setPreloadFunction set PreloadFunc in Pagination
func (p *Pagination) setPreloadFunction(page, pageSize int32) {
	startRow := (page - 1) * pageSize
	preloadFunc := func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(startRow)).Limit(int(pageSize))
	}
	p.PreloadFunc = preloadFunc
}

// setNextAndPrevious set Next and Prev in Pagination
func (p *Pagination) setNextAndPrevious() {
	p.Next = p.GetNext()
	p.Prev = p.GetPrevious()
}
