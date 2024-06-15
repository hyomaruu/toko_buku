package web

type AksesorisUpdateRequest struct {
	Id    		string 	`validate:"required" form:"kode"`
	Name  		string 	`validate:"required" form:"nama_barang"`
	Image 		string  `form:"gambar"`
	Deskripsi  	string 	`validate:"required" form:"deskripsi"`
	Harga 		int    	`validate:"required" form:"harga"`
}
