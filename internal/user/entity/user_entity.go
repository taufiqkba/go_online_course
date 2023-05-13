package entity

import (
	"database/sql"
	entity2 "go_online_course/internal/admin/entity"
	"gorm.io/gorm"
)

type User struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	Password        string         `json:"-"`
	CodeVerified    string         `json:"-"`
	EmailVerifiedAt sql.NullTime   `json:"email_verified_at"`
	CreatedByID     int64          `json:"created_by" gorm:"column:created_by"`
	CreatedBy       *entity2.Admin `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID     *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy       *entity2.Admin `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt       sql.NullTime   `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
