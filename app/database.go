package app

import (
	"database/sql"
	"time"
	"toko-buku/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql","root@tcp(localhost:3306)/toko_buku?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}