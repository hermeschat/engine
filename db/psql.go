package db

import (
	"database/sql"
	"fmt"
	"github.com/hermeschat/engine/config"
	_ "github.com/lib/pq"
)

type Postgres struct {
	conn *sql.DB
}

func (p *Postgres) DB() (*sql.DB, error) {
	setPostgresDefaultConfiguration()
	host := config.C.GetString("database.host")
	port := config.C.GetString("database.port")
	user := config.C.GetString("database.user")
	password := config.C.GetString("database.password")
	name := config.C.GetString("database.name")
	sslmode := config.C.GetString("database.sslmode")
	if p.conn == nil {
		conn, err := postgresConnect(host, port, user, password, name, sslmode)
		if err != nil {
			return nil, err
		}
		p.conn = conn
		return p.conn, nil
	}
	if err := p.conn.Ping(); err != nil {
		return nil, err
	}
	return p.conn, nil
}

func postgresConnect(host, port, user, password, name, sslmode string) (*sql.DB, error) {
	conString := postgresConnectionString(host, port, user, password, name, sslmode)
	conn, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, fmt.Errorf("Error in openning postgres connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func postgresConnectionString(host, port, user, password, name, sslmode string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, name, sslmode)
}

func setPostgresDefaultConfiguration() {
	config.C.Set("database.host", "localhost")
	config.C.Set("database.port", "5432")
	config.C.Set("database.user", "postgres")
	config.C.Set("database.password", "")
	config.C.Set("database.sslmode", "disable")
}
