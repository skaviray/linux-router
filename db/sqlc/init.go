package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *Queries

const (
	db_Driver = "postgres"
	db_Source = "postgresql://root:admin@localhost:5453/gateway_router?sslmode=disable"
)

func Init() {
	conn, err := sql.Open(db_Driver, db_Source)
	if err != nil {
		log.Println(err)
		log.Panic("unable to open the database connection")
	}
	if err = conn.Ping(); err != nil {
		log.Panic("unable to rach the database")
	}
	Db = New(conn)
}
