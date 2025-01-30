package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:admin@localhost:5432/gateway_router?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Panic("unable to open the database connection")
	}
	if err = conn.Ping(); err != nil {
		log.Panic("unable to rach the database")
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
