package entity

import (
	"database/sql"
	"go_online_course/internal/admin/entity"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	CreatedByID int64          `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *entity.Admin  `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *entity.Admin  `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
