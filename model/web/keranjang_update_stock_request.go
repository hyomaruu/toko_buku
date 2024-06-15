package web

type KeranjangUpdateStockRequest struct {
	TotalBarang    	int 	`validate:"required" json:"total_barang"`
}
