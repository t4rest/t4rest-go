package meta

// Meta .
type Meta interface {
	GetMetaData() map[string]interface{}
}

// Pagination .
type Pagination struct {
	ResultCount uint64 `json:"result_count"`
	Limit       uint64 `json:"limit"`
	Offset      uint64 `json:"offset"`
}

// GetMetaData return pagination meta
func (p Pagination) GetMetaData() map[string]interface{} {
	m := map[string]interface{}{}

	m["pagination"] = map[string]uint64{
		"total":  p.ResultCount,
		"limit":  p.Limit,
		"offset": p.Offset,
	}

	return m
}
