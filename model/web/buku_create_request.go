package web


type BukuCreateRequest struct {
	Id        string    `validate:"required" json:"kode_buku"`
	Name      string    `validate:"required" json:"nama_buku"`
	Penerbit  string    `validate:"required" json:"penerbit"`
	Image     string    `json:"gambar"`
	Pengarang string    `validate:"required" json:"pengarang"`
	Genre     string    `validate:"required" json:"genre"`
	Rilis     string `validate:"required" json:"rilis"`
	Harga     int       `validate:"required" json:"harga"`
}
