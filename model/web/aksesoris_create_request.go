package web

type AksesorisCreateRequest struct {
	Id    string `validate:"required" json:"kode"`
	Name  string `validate:"required" json:"nama_barang"`
	Harga int    `validate:"required" json:"harga"`
	Stok int    `validate:"required" json:"stok"`
}
