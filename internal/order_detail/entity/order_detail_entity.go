package entity

import (
	"database/sql"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/user/entity"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID          int64            `json:"id"`
	Price       int64            `json:"price"`
	Order       *entity3.Product `json:"order" gorm:"foreignKey:ProductID;references:ID"`
	OrderID     int64            `json:"order_id"`
	Product     *entity3.Product `json:"product" gorm:"foreignKey:ProductID;referenced:ID"`
	ProductID   int64            `json:"product_id"`
	CreatedByID *int64           `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *entity.User     `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64           `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *entity.User     `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime     `json:"created_at"`
	UpdatedAt   sql.NullTime     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at"`
}
