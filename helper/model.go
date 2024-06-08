package helper

import (
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

func ToAdminResponse(admin domain.Admin) web.AdminResponse {
	return web.AdminResponse{
		Id : admin.Id,
		Username: admin.Username,
		Password: admin.Password,
		Name: admin.Name,


	}
}

func ToAdminResponses(admins []domain.Admin) []web.AdminResponse  {
	var adminResponses []web.AdminResponse
	for _, admin := range admins{
		adminResponses = append(adminResponses, ToAdminResponse(admin))
	}
	return adminResponses
}

func ToMemberResponse(member domain.Member) web.MemberResponse {
	return web.MemberResponse{
		Id : member.Id,
		Name: member.Name,
<<<<<<< HEAD
		Email: member.Email,
		Telepon: member.Telepon,
		Alamat: member.Alamat,
=======
>>>>>>> 7a1f8e50849a6ee819f6ea835677610a48fcb3f0


	}
}

func ToMemberResponses(members []domain.Member) []web.MemberResponse  {
	var memberResponses []web.MemberResponse
	for _, member := range members{
		memberResponses = append(memberResponses, ToMemberResponse(member))
	}
	return memberResponses
}

func ToBukuResponse(buku domain.Buku) web.BukuResponse{

	Rilises := buku.Rilis.Format("2006-01-02")


	
	return web.BukuResponse{
		Id: buku.Id,
		Name: buku.Name,
		Penerbit: buku.Penerbit,
		Image: buku.Image,
		Pengarang: buku.Pengarang,
		Genre: buku.Genre,
		Rilis: Rilises,
		Harga: buku.Harga,
	}
}

func ToBukuResponses(bukus []domain.Buku) []web.BukuResponse{
	var bukuResponses []web.BukuResponse
	for _, buku := range bukus{
		bukuResponses = append(bukuResponses, ToBukuResponse(buku))
	}
	return bukuResponses
}

func ToAksesorisResponse(aksesoris domain.Aksesoris) web.AksesorisResponse{
	return web.AksesorisResponse{
		Id: aksesoris.Id,
		Name: aksesoris.Name,
		Harga: aksesoris.Harga,
	}
}

func ToAksesorisResponses(aksesoriss []domain.Aksesoris) []web.AksesorisResponse{
	var aksesorisResponses []web.AksesorisResponse
	for _, aksesoris := range aksesoriss{
		aksesorisResponses = append(aksesorisResponses, ToAksesorisResponse(aksesoris))
	}
	return aksesorisResponses
}