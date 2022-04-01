package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifeTime = 5 * time.Minute

// Creates DB pool for postgres
func ConnectSQL(dsn string) (*DB,error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = db
	err = TestDBConn(db)
	if err != nil {return nil,err}
	return dbConn,err
}

// Tries to ping the database 
func TestDBConn(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Creates a new database connection pool
func NewDatabase(dsn string) (*sql.DB,error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	return db,err
}