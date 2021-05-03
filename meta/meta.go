package meta

// Meta .
type Meta interface {
	GetMetaData() map[string]interface{}
}

// Pagination .
type Pagination struct {
	ResultCount uint `json:"result_count"`
	Limit       uint `json:"limit"`
	Offset      uint `json:"offset"`
}

// GetMetaData return pagination meta
func (p Pagination) GetMetaData() map[string]interface{} {
	m := map[string]interface{}{}

	m["pagination"] = map[string]uint{
		"total":  p.ResultCount,
		"limit":  p.Limit,
		"offset": p.Offset,
	}

	return m
}
