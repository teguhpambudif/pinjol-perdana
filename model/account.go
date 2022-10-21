package model

import "database/sql"

type Account struct {
	Id       string         `json:"id"`
	Username string         `json:"username" binding:"required"`
	Password string         `json:"password" binding:"required"`
	Usia     int            `json:"usia" binding:"required"`
	Gaji     int            `json:"gaji" binding:"required"`
	FileKtp  sql.NullString `json:"file_ktp" db:"file_ktp"`
}
