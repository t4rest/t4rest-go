package mysql

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

var (
	ErrDuplicate = errors.New("MySQL duplicate field")
	ErrNoRows    = errors.New("MySQL no rows in result set")
	errorsRegexp = regexp.MustCompile(`^Error (?P<code>\d+)`)
)

// Find into dest from query builder
func (ms *Mysql) Find(ctx context.Context, dest interface{}, b squirrel.SelectBuilder, tx ...Tx) error {
	q, args, err := b.ToSql()
	if err != nil {
		return err
	}

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("Find", "query", q, "params", args)
	}

	if len(tx) > 0 && tx[0] != nil {
		err = tx[0].GetTx().SelectContext(ctx, dest, q, args...)
	} else {
		err = ms.DB.SelectContext(ctx, dest, q, args...)
	}

	return err
}

// FindRaw into dest from query
func (ms *Mysql) FindRaw(ctx context.Context, dest interface{}, q string, args ...interface{}) error {
	err := ms.DB.SelectContext(ctx, dest, q, args...)

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("FindRaw", "query", q, "params", args)
	}

	return err
}

// FindFirst row into dest from query builder
func (ms *Mysql) FindFirst(ctx context.Context, dest interface{}, b squirrel.SelectBuilder, tx ...Tx) error {
	q, args, err := b.ToSql()
	if err != nil {
		return err
	}

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("FindFirst", "query", q, "params", args)
	}

	if len(tx) > 0 && tx[0] != nil {
		err = tx[0].GetTx().GetContext(ctx, dest, q, args...)
	} else {
		err = ms.DB.GetContext(ctx, dest, q, args...)
	}

	return ms.parseError(err)
}

// Insert from query builder
func (ms *Mysql) Insert(ctx context.Context, b squirrel.InsertBuilder, tx ...Tx) (uint64, error) {
	q, args, err := b.ToSql()
	if err != nil {
		return 0, err
	}

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("Insert", "query", q, "params", args)
	}

	result, err := ms.exec(ctx, q, args, tx...)
	if err != nil {
		return 0, ms.parseError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), err
}

// Exec query
func (ms *Mysql) Exec(ctx context.Context, q string, args []interface{}, err error, tx ...Tx) (uint64, error) {
	if err != nil {
		return 0, ms.parseError(err)
	}

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("Exec", "query", q, "params", args)
	}

	result, err := ms.exec(ctx, q, args, tx...)
	if err != nil {
		return 0, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(affectedRows), err
}

// CallFunc from db
func (ms *Mysql) CallFunc(ctx context.Context, q string, args []interface{}, tx ...Tx) error {
	var err error

	if ms.cfg.DebugMode && ms.log != nil {
		ms.log.Debug("CallFunc", "query", q, "params", args)
	}

	if len(tx) > 0 && tx[0] != nil {
		err = tx[0].GetTx().QueryRowContext(ctx, q).Scan(args...)
	} else {
		err = ms.DB.QueryRowContext(ctx, q).Scan(args...)
	}

	return err
}

func (ms *Mysql) exec(ctx context.Context, q string, args []interface{}, tx ...Tx) (sql.Result, error) {
	var result sql.Result
	var err error

	if len(tx) > 0 && tx[0] != nil {
		result, err = tx[0].GetTx().ExecContext(ctx, q, args...)
	} else {
		result, err = ms.DB.ExecContext(ctx, q, args...)
	}

	return result, err
}

func (ms *Mysql) parseError(err error) error {
	if err == nil {
		return nil
	}

	// Just a wrapper not to use sql lib directly from code
	if err == sql.ErrNoRows {
		return ErrNoRows
	}

	matches := ms.matchStringGroups(errorsRegexp, err.Error())
	code, ok := matches["code"]
	if !ok {
		return err
	}

	switch code {
	case "1062":
		return ErrDuplicate
	default:
		return err
	}
}

// matchStringGroups matches regexp with capture groups. Returns map string string
func (ms *Mysql) matchStringGroups(re *regexp.Regexp, s string) map[string]string {
	m := re.FindStringSubmatch(s)
	n := re.SubexpNames()

	r := make(map[string]string, len(m))

	if len(m) > 0 {
		m, n = m[1:], n[1:]
		for i := range n {
			r[n[i]] = m[i]
		}
	}

	return r
}
