package web

type AdminCreateRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"nama_lengkap"`
	Role     string `validate:"required" json:"role"`
}
