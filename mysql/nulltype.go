package mysql

import (
	"database/sql"
	"time"
)

// ToNullTime .
func ToNullTime(value time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  value,
		Valid: !value.IsZero(),
	}
}

// NullTimeToInt .
func NullTimeToInt(value sql.NullTime) int64 {
	if value.Valid {
		return value.Time.Unix()
	}

	return 0
}

// ToNullString .
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// ToNullInt64 .
func ToNullInt64(n int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: n,
		Valid: n != 0,
	}
}
