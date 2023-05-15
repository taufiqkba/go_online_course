package dto

type DiscountRequestBody struct {
	Name              string `json:"name" binding:"required"`
	Code              string `json:"code" binding:"required"`
	Quantity          int64  `json:"quantity" binding:"required"`
	Type              string `json:"type" binding:"required"`
	RemainingQuantity int64  `json:"remaining_quantity" binding:"number"`
	Value             int64  `json:"value" binding:"required,number"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	CreatedBy         int64  `json:"created_by"`
	UpdatedBy         int64  `json:"updated_by"`
}
