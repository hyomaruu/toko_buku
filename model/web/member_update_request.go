package web

type MemberUpdateRequest struct {
	Id   int    `validate:"required" json:"id_member"`
	Name string `validate:"required" json:"nama_lengkap"`
<<<<<<< HEAD
	Email string `validate:"required" json:"email"`
	Telepon string `validate:"required" json:"telepon"`
	Alamat string `validate:"required" json:"alamat"`
=======
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0
}
