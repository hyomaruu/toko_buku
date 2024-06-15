package web

type KeranjangCreateRequest struct {
	TipeBarang 		string `validate:"required" json:"tipe_barang"`
	KodeBarang     	string `validate:"required" json:"kode_barang"`
	TotalBarang     int `validate:"required" json:"total_barang"`
	Harga     		int `validate:"required" json:"harga"`
}
