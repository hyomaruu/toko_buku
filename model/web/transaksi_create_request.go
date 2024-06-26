package web

type TransaksiCreateRequest struct {
	NomorMember 		int `validate:"required" json:"nomor_member"`
	IdAdmin     		int `validate:"required" json:"id_admin"`
	Tanggal     		string `validate:"required" json:"tanggal"`
	TotalHarga     		int `validate:"required" json:"total_harga"`
	ListKeranjang 		[]string `validate:"required" json:"list_keranjang"`

}
