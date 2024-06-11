package web

type AdminUpdateRequest struct {
	Username string `validate:"required,max=255,min=1" json:"username"`
	Password string `validate:"required,max=255,min=1" json:"password"`
	Name     string `validate:"required,max=255,min=1" json:"nama_lengkap"`
	Role     string `validate:"required,max=255,min=1" json:"role"`

}
