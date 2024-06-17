package web

type TransaksiUpdateStockRequest struct {
	TotalBarang    	int 	`validate:"required" json:"total_barang"`
}
