package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func Init(config *DBConfig) {
	connString := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	db.SetMaxIdleConns(config.MaxConn)
	db.SetMaxOpenConns(config.MaxConn)
	conn = db
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(db *sqlx.DB) {
	conn = db
}
