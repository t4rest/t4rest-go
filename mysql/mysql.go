package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/t4rest/t4rest-go/logger"
)

// Mysql .
type Mysql struct {
	DB  *sqlx.DB
	cfg Conf
	log *logger.Logger
}

// New .
func New(cfg Conf) (*Mysql, error) {
	db, err := sqlx.Connect("mysql", cfg.ConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Connect")
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return &Mysql{DB: db}, nil
}

// SetLogger .
func (ms *Mysql) SetLogger(log *logger.Logger) {
	ms.log = log
}

// Close .
func (ms *Mysql) Close() error {
	return ms.DB.Close()
}
