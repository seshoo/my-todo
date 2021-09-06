package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/seshoo/my-todo/internal/config"
	logger "github.com/sirupsen/logrus"
)

func NewDb(cfg *config.Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.DatabaseName, cfg.Postgres.Password, cfg.Postgres.SslMode))

	if err != nil {
		logger.Error(err)
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
