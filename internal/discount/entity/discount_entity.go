package entity

import (
	"database/sql"
	"go_online_course/internal/admin/entity"
	"gorm.io/gorm"
)

type Discount struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Quantity          int64          `json:"quantity"`
	RemainingQuantity int64          `json:"remaining_quantity"`
	Type              string         `json:"type"`
	Value             int64          `json:"value"`
	StartDate         sql.NullTime   `json:"start_date"`
	EndDate           sql.NullTime   `json:"end_date"`
	CreatedByID       *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *entity.Admin  `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *entity.Admin  `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}
