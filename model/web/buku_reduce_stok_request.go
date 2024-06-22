package web

type BukuReduceStockRequest struct {
	Qty    	int 	`validate:"required" json:"qty"`
}
