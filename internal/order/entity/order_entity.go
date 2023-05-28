package entity

import (
	"database/sql"
	entity2 "go_online_course/internal/discount/entity"
	entity4 "go_online_course/internal/order_detail/entity"
	"go_online_course/internal/user/entity"
	"gorm.io/gorm"
)

type Order struct {
	ID           int64                 `json:"id"`
	User         *entity.User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID       int64                 `json:"user_id"`
	OrderDetails []entity4.OrderDetail `json:"order_details"`
	Discount     *entity2.Discount     `json:"discount" gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID   *int64                `json:"discount_id"`
	CheckoutLink string                `json:"checkout_link"`
	ExternalID   string                `json:"external_id"`
	Price        int64                 `json:"price"`
	TotalPrice   int64                 `json:"total_price"`
	Status       string                `json:"status"`
	CreatedByID  *int64                `json:"created_by" gorm:"column:created_by"`
	CreatedBy    *entity.User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID  *int64                `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy    *entity.User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    sql.NullTime          `json:"created_at"`
	UpdatedAt    sql.NullTime          `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `json:"deleted_at"`
}
