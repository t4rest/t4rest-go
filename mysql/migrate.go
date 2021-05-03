package mysql

import (
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

// Migrate .
func (ms *Mysql) Migrate(dir string) error {

	err := goose.SetDialect("mysql")
	if err != nil {
		return errors.Wrap(err, "goose.SetDialect")
	}

	//goose.SetLogger(logger.Log.With("goose", "migrate"))

	err = goose.Run("up", ms.DB.DB, dir)
	if err != nil {
		return errors.Wrap(err, "goose.Run")
	}

	return nil
}
