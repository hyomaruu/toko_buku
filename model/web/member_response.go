package web

type MemberResponse struct {
	Id   int    `validate:"required" json:"id_member"`
	Name string `validate:"required" json:"nama_lengkap"`
}
