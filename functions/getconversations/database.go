package function

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var database *sql.DB

func getDB() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}

	db, err := sql.Open("postgres", getDBConnectionString(getDBCredentials()))
	if err != nil {
		return nil, err
	}

	database = db
	return database, err
}

type dbCredentials struct {
	host    string
	port    string
	user    string
	pass    string
	name    string
	sslmode string
}

func getDBCredentials() dbCredentials {
	return dbCredentials{
		host:    os.Getenv("DB_HOST"),
		port:    os.Getenv("DB_PORT"),
		user:    os.Getenv("DB_USER"),
		pass:    os.Getenv("DB_PASS"),
		name:    os.Getenv("DB_NAME"),
		sslmode: os.Getenv("DB_SSLMODE"),
	}
}

func getDBConnectionString(c dbCredentials) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.host, c.port, c.user, c.pass, c.name, c.sslmode)
}
