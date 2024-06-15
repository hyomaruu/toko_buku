package web

type AksesorisUpdateStockRequest struct {
	Stok    	int 	`validate:"required" json:"stok"`
}
