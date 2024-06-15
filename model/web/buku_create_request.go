package web


type BukuCreateRequest struct {
	Id        	string    	`validate:"required" form:"kode_buku"`
	Name      	string    	`validate:"required" form:"nama_buku"`
	Penerbit  	string    	`validate:"required" form:"penerbit"`
	Image     	string    	`form:"gambar"`
	Pengarang 	string    	`validate:"required" form:"pengarang"`
	Genre     	string    	`validate:"required" form:"genre"`
	Rilis     	string 		`validate:"required" form:"rilis"`
	Deskripsi   string 		`validate:"required" form:"deskripsi"`
	Harga     	int       	`validate:"required" form:"harga"`
}
