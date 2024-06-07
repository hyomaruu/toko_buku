package web

type MemberUpdateRequest struct {
	Id   int    `validate:"required" json:"id_member"`
	Name string `validate:"required" json:"nama_lengkap"`
}
