package web

type KeranjangResponse struct {
	KodeKeranjang 	int64 
	KodeTransaksi 	string 
	TipeBarang 		string 
	KodeBarang     	string 
	TotalBarang     int 
	Harga     		int 
	Data	 		interface{}
}
