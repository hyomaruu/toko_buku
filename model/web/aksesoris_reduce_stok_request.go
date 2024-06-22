package web

type AksesorisReduceStockRequest struct {
	Qty    	int 	`validate:"required" json:"qty"`
}
