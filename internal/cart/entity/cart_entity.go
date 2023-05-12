package entity

import (
	"database/sql"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/user/entity"
)

type Cart struct {
	ID        int64            `json:"id"`
	User      *entity.User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID    int64            `json:"user_id"`
	Product   *entity3.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID int64            `json:"product_id"`
	CreatedAt sql.NullTime     `json:"created_at"`
	UpdatedAt sql.NullTime     `json:"updated_at"`
}
