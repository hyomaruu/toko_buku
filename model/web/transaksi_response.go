package web

type TransaksiResponse struct {
	KodeTransaksi 		int64 
	NomorMember 		int 
	IdAdmin     		int 
	Tanggal     		string 
	TotalHarga     		int 
	ListKeranjang	 	interface{}
}
