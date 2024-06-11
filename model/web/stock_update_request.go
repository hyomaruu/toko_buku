package web

type StockUpdateRequest struct{
	Stock int `validate:"required" form:"stok"`
}