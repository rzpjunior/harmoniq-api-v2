package dto

type CartRequestUpdate struct {
	ProductId int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required,gte=1,lte=1000"`
}

type CartResponse struct {
	Id      int             `json:"id"`
	Product ProductResponse `json:"product"`
	Qty     int             `json:"qty"`
}
