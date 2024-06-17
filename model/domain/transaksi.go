package domain

import "time"


type Transaksi struct {
	KodeTransaksi  	int64
	NomorMember  	string
	IdAdmin  		string
	Tanggal     	time.Time
	TotalHarga    	int
	ListKeranjang	interface{}
}