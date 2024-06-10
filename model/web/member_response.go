package web

type MemberResponse struct {
	Id   int64    `validate:"required" json:"id_member"`
	Name string `validate:"required" json:"nama_lengkap"`
	Email string `validate:"required" json:"email"`
	Telepon string `validate:"required" json:"telepon"`
	Alamat string `validate:"required" json:"alamat"`
}
