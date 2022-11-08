package v1

import (
	"time"
)

type Extend map[string]interface{}

type ListMeta struct {
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	TotalCount  int64 `json:"total_count"`
}

func (meta *ListMeta) GetCurrentPage() int64 {
	return meta.CurrentPage
}

func (meta *ListMeta) SetCurrentPage(offset int64) {
	meta.CurrentPage = offset
}

func (meta *ListMeta) GetTotalPage() int64 {
	return meta.TotalCount
}

func (meta *ListMeta) SetTotalPage(limit int64) {
	totalCount := meta.TotalCount
	totalPages := int64(0)
	if meta.TotalCount > 0 {
		if totalCount%limit != 0 {
			totalPages = (totalCount / limit) + 1
		} else {
			totalPages = totalCount / limit
		}
	}
	meta.TotalPage = totalPages
}

type Model struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time `json:"createdAt" gorm:"index;comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index;comment:最后更新时间"`
}

type ObjectMeta struct {
	ID     uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Extend Extend `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ModelTime
}

type ListOptions struct {
	// LabelSelector is used to find matching REST resources.
	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`

	// FieldSelector restricts the list of returned objects by their fields. Defaults to everything.
	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`

	// TimeoutSeconds specifies the seconds of ClientIP type session sticky time.
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`

	// Offset specify the number of records to skip before starting to return the records.
	Offset *int64 `json:"offset,omitempty" form:"offset"`

	// Limit specify the number of records to be retrieved.
	Limit *int64 `json:"limit,omitempty" form:"limit"`
}
