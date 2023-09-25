package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
)

var u, p, name, host, port, scheme string

func InitDB() {
	scheme = "mysql"
	u = os.Getenv("DB_USER")
	p = os.Getenv("DB_PASSWORD")
	name = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")

	dsn := url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%s", host, port),
		User:   url.UserPassword(u, p),
		Path:   name,
	}

	q := dsn.Query()
	q.Add("sslmode", "disabled")
	dsn.RawQuery = q.Encode()

	db, err := sql.Open(scheme, dsn.String())
	if err != nil {
		fmt.Println("failed establishing db connection", err)
	}
	defer func() {
		_ = db.Close()
		fmt.Println("db connection closed")
	}()
	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("failed to ping db", err)
		return
	}

}
