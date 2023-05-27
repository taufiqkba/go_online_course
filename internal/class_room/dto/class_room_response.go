package dto

import (
	"database/sql"
	entity2 "go_online_course/internal/admin/entity"
	entity4 "go_online_course/internal/class_room/entity"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/user/entity"
	"gorm.io/gorm"
)

type ClassRoomResponseBody struct {
	ID        int64            `json:"id"`
	User      *entity.User     `json:"user"`
	Product   *entity3.Product `json:"product"`
	CreatedBy *entity2.Admin   `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedBy *entity2.Admin   `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt sql.NullTime     `json:"created_at"`
	UpdatedAt sql.NullTime     `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

func CreateClassRoomResponse(entity entity4.ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        entity.ID,
		User:      entity.User,
		Product:   entity.Product,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []entity4.ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoURL = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}
	return classRoomResp
}
