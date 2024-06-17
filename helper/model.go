package helper

import (
	"toko-buku/model/domain"
	"toko-buku/model/web"
)

func ToAdminResponse(admin domain.Admin) web.AdminResponse {
	return web.AdminResponse{
		Id:       admin.Id,
		Username: admin.Username,
		Password: admin.Password,
		Name:     admin.Name,
		Role:     admin.Role,
	}
}

func ToAdminResponses(admins []domain.Admin) []web.AdminResponse {
	var adminResponses []web.AdminResponse
	for _, admin := range admins {
		adminResponses = append(adminResponses, ToAdminResponse(admin))
	}
	return adminResponses
}

func ToMemberResponse(member domain.Member) web.MemberResponse {
	return web.MemberResponse{
		Id:      member.Id,
		Name:    member.Name,
		Email:   member.Email,
		Telepon: member.Telepon,
		Alamat:  member.Alamat,
	}
}

func ToMemberResponses(members []domain.Member) []web.MemberResponse {
	var memberResponses []web.MemberResponse
	for _, member := range members {
		memberResponses = append(memberResponses, ToMemberResponse(member))
	}
	return memberResponses
}

func ToBukuResponse(buku domain.Buku) web.BukuResponse {
	Rilises := buku.Rilis.Format("2006-01-02")
	return web.BukuResponse{
		Id:        buku.Id,
		Name:      buku.Name,
		Penerbit:  buku.Penerbit,
		Image:     buku.Image,
		Pengarang: buku.Pengarang,
		Genre:     buku.Genre,
		Rilis:     Rilises,
		Harga:     buku.Harga,
		Stok:     	buku.Stok,
		Deskripsi:     	buku.Deskripsi,

	}
}

func ToBukuResponses(bukus []domain.Buku) []web.BukuResponse {
	var bukuResponses []web.BukuResponse
	for _, buku := range bukus {
		bukuResponses = append(bukuResponses, ToBukuResponse(buku))
	}
	return bukuResponses
}

func ToAksesorisResponse(aksesoris domain.Aksesoris) web.AksesorisResponse {
	return web.AksesorisResponse{
		Id:    aksesoris.Id,
		Name:  aksesoris.Name,
		Image: aksesoris.Image,
		Harga: aksesoris.Harga,
		Deskripsi: aksesoris.Deskripsi,
		Stok: aksesoris.Stok,
	}
}

func ToAksesorisResponses(aksesoriss []domain.Aksesoris) []web.AksesorisResponse {
	var aksesorisResponses []web.AksesorisResponse
	for _, aksesoris := range aksesoriss {
		aksesorisResponses = append(aksesorisResponses, ToAksesorisResponse(aksesoris))
	}
	return aksesorisResponses
}

func ToKeranjangResponse(keranjang domain.Keranjang) web.KeranjangResponse {
	return web.KeranjangResponse{
		KodeKeranjang 	: keranjang.KodeKeranjang, 
		KodeTransaksi 	: keranjang.KodeTransaksi, 
		TipeBarang 		: keranjang.TipeBarang, 
		KodeBarang     	: keranjang.KodeBarang, 
		TotalBarang     : keranjang.TotalBarang, 
		Harga     		: keranjang.Harga, 
		Barang     		: keranjang.Barang, 
	}
}

func ToKeranjangResponses(keranjangs []domain.Keranjang) []web.KeranjangResponse {
	var keranjangResponses []web.KeranjangResponse
	for _, keranjang := range keranjangs {
		keranjangResponses = append(keranjangResponses, ToKeranjangResponse(keranjang))
	}
	return keranjangResponses
}

func ToTransaksiResponse(keranjang domain.Transaksi) web.TransaksiResponse {
	Tanggal := keranjang.Tanggal.Format("2006-01-02")
	return web.TransaksiResponse{
		KodeTransaksi 	: keranjang.KodeTransaksi, 
		NomorMember 	: keranjang.NomorMember, 
		IdAdmin 		: keranjang.IdAdmin, 
		Tanggal 		: Tanggal, 
		ListKeranjang 	: keranjang.ListKeranjang, 
	}
}

func ToTransaksiResponses(keranjangs []domain.Transaksi) []web.TransaksiResponse {
	var keranjangResponses []web.TransaksiResponse
	for _, keranjang := range keranjangs {
		keranjangResponses = append(keranjangResponses, ToTransaksiResponse(keranjang))
	}
	return keranjangResponses
}

