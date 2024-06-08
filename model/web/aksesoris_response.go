package web

type AksesorisResponse struct {
	Id    string `validate:"required" json:"kode"`
	Name  string `validate:"required" json:"nama_barang"`
	Harga int    `validate:"required" json:"harga"`
}
