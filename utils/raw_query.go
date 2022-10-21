package utils

const (
	REGISTER_ACCOUNT = "insert into m_account (id,username,password,usia,gaji)values(:id,:username,:password,:usia,:gaji)"
	GET_ACCOUNT      = "select id,username,usia,gaji,file_ktp from m_account where id = $1"
	UPLOAD_KTP       = "update m_account set file_ktp=$1, where id = $2"
)
