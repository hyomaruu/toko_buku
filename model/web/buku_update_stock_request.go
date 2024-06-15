package web

type BukuUpdateStockRequest struct {
	Stok    	int 	`validate:"required" json:"stok"`
}
