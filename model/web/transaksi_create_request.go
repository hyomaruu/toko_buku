package web

type TransaksiCreateRequest struct {
	NomorMember 		string `validate:"required" json:"nomor_member"`
	IdAdmin     		string `validate:"required" json:"id_admin"`
	Tanggal     		string `validate:"required" json:"tanggal"`
	TotalHarga     		int `validate:"required" json:"total_harga"`
	ListKeranjang 		[]int `validate:"required" json:"list_keranjang"`

}
