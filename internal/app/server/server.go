package server

import (
	"database/sql"
	"net/http"

	"github.com/vladjong/L0/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DateBaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
