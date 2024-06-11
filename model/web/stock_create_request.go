package web

type StockCreateRequest struct{
	Id string `validate:"required" form:"kode_barang"`
	Stock int `validate:"required" form:"stok"`
}