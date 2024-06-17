package domain

type Keranjang struct {
	KodeKeranjang   int64
	KodeTransaksi  	string
	TipeBarang  	string
	KodeBarang    	string
	namaBarang    	string
	TotalBarang 	int
	Harga     		int
	Barang			interface{}
}