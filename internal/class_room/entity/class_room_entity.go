package entity

import (
	"database/sql"
	entity2 "go_online_course/internal/admin/entity"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/user/entity"
	"gorm.io/gorm"
)

type ClassRoom struct {
	ID          int64            `json:"id"`
	User        *entity.User     `json:"user" gorm:"foreignKey:UserID;References:ID"`
	UserID      int64            `json:"user_id"`
	Product     *entity3.Product `json:"product" gorm:"foreignKey:ProductID;References:ID"`
	ProductID   int64            `json:"product_id"`
	CreatedByID *int64           `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *entity2.Admin   `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64           `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *entity2.Admin   `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime     `json:"created_at"`
	UpdatedAt   sql.NullTime     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at"`
}
