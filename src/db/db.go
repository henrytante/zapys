package db

import (
	"database/sql"
	

	_ "github.com/mattn/go-sqlite3"
)

func ConnectV1() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "databases/dbV1.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func ConnectV2() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "databases/dbV2.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
func DBToken() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "databases/tokens.sqlite")
	if err != nil{
		return nil, err
	}
	err = db.Ping()
	if err != nil{
		return nil, err
	}
	return db, nil
}
func DBADM() (*sql.DB, error)  {
	db, err := sql.Open("sqlite3", "databases/admin.sqlite")
	if err != nil{
		return nil, err
	}
	err = db.Ping()
	if err != nil{
		return nil, err
	}
	return db, nil
}