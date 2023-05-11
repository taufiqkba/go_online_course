package dto

import (
	"database/sql"
	entity2 "go_online_course/internal/admin/entity"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/product_category/entity"
	"gorm.io/gorm"
)

type ProductResponseBody struct {
	ID              int64                   `json:"id"`
	ProductCategory *entity.ProductCategory `json:"product_category"`
	Title           string                  `json:"title"`
	Image           string                  `json:"image"`
	Video           string                  `json:"video"`
	Description     string                  `json:"description"`
	Price           int64                   `json:"price"`
	CratedBy        *entity2.Admin          `json:"crated_by"`
	UpdatedBy       *entity2.Admin          `json:"updated_by"`
	CreatedAt       sql.NullTime            `json:"created_at"`
	UpdatedAt       sql.NullTime            `json:"updated_at"`
	DeletedAt       gorm.DeletedAt          `json:"deleted_at"`
}

func CreateProductResponse(product entity3.Product) ProductResponseBody {
	return ProductResponseBody{
		ProductCategory: product.ProductCategory,
		Title:           product.Title,
		Image:           product.Image,
		Video:           product.Video,
		Description:     product.Description,
		Price:           product.Price,
		CratedBy:        product.CreatedBy,
		UpdatedBy:       product.UpdatedBy,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		DeletedAt:       product.DeletedAt,
	}
}

type ProductListResponse []ProductResponseBody

func CreateProductListResponse(products []entity3.Product) ProductListResponse {
	productResp := ProductListResponse{}

	for _, p := range products {
		product := CreateProductResponse(p)
		productResp = append(productResp, product)
	}
	return productResp
}
