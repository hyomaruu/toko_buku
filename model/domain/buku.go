package domain

import "time"

type Buku struct {
	Id        	string
	Name      	string
	Penerbit  	string
	Image     	string
	Pengarang 	string
	Genre     	string
	Deskripsi 	string
	Rilis     	time.Time
	Harga     	int
	Stok		int
}