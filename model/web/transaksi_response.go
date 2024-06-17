package web

type TransaksiResponse struct {
	KodeTransaksi 		int64 
	NomorMember 		string 
	IdAdmin     		string 
	Tanggal     		string 
	TotalHarga     		int 
	ListKeranjang	 	interface{}
}
