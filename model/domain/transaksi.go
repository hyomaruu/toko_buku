package domain

import "time"


type Transaksi struct {
	KodeTransaksi  	int64
	NomorMember  	int
	IdAdmin  		int
	Tanggal     	time.Time
	TotalHarga    	int
	ListKeranjang	interface{}
}