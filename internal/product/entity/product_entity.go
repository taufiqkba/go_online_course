package entity

import (
	"database/sql"
	entity2 "go_online_course/internal/admin/entity"
	"go_online_course/internal/product_category/entity"
	"gorm.io/gorm"
)

type Product struct {
	ID                int64                   `json:"id"`
	ProductCategory   *entity.ProductCategory `json:"product_category" gorm:"foreignKey:ProductCategoryID;references:ID"`
	ProductCategoryID int64                   `json:"product_category_id"`
	Title             string                  `json:"title"`
	Image             *string                 `json:"image"`
	Video             *string                 `json:"video"`
	Description       string                  `json:"description"`
	Price             int64                   `json:"price"`
	CreatedByID       int64                   `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *entity2.Admin          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64                  `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *entity2.Admin          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         sql.NullTime            `json:"created_at"`
	UpdatedAt         sql.NullTime            `json:"updated_at"`
	DeletedAt         gorm.DeletedAt          `json:"deleted_at"`
}
