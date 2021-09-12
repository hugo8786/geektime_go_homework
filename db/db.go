package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Error        error
	Sql          *sql.DB
	Value        interface{}
	RowsAffected int64
}

var dbConn = &DB{}

func (db *DB) Clone() (cloneDb *DB) {
	cloneDb = &DB{
		Error:        db.Error,
		Sql:          db.Sql,
		Value:        db.Value,
		RowsAffected: db.RowsAffected,
	}
	return
}
func (db *DB) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}

func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.Sql = d
	return dbConn, err
}
