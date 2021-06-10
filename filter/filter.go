package filter

import (
	"strings"
)

// MaxLimit .
const MaxLimit = 100
const defaultLimit = 50
const asc = "ASC"
const desc = "DESC"

// Filter .
type Filter struct {
	Offset         uint64
	Limit          uint64
	OrderDirection string
	OrderBy        string
}

// GetLimit .
func (f Filter) GetLimit() uint64 {
	if f.Limit == 0 {
		return defaultLimit
	} else if f.Limit > MaxLimit {
		return MaxLimit
	}
	return f.Limit
}

// GetOrderDirection .
func (f Filter) GetOrderDirection() string {
	if strings.ToUpper(f.OrderDirection) != asc {
		return desc
	}
	return asc
}
