package web

type AdminResponse struct {
	Id       int64  `validate:"required" json:"id_admin"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"nama_lengkap"`
	Role     string `validate:"required" json:"role"`
}
