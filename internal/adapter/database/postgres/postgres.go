package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	appconfig "github.com/pawannn/famly/internal/pkg/appConfig"
)

func InitDatabase(c appconfig.Config) (*sql.DB, error) {
	var conn_url = fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		c.DB_host, c.DB_port, c.DB_user, c.DB_name, c.DB_SSL, c.DB_pass,
	)
	fmt.Println(conn_url)

	db, err := sql.Open("postgres", conn_url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
