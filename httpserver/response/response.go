package response

import (
	"encoding/json"
	"net/http"

	error2 "github.com/t4rest/t4rest-go/httpserver/error"

	"github.com/t4rest/t4rest-go/meta"
)

// JSON writes to ResponseWriter JSON-data
func JSON(w http.ResponseWriter, data interface{}, md ...meta.Meta) {
	resp := map[string]interface{}{"data": data}

	metaData := map[string]interface{}{}
	for _, mt := range md {
		for key, m := range mt.GetMetaData() {
			metaData[key] = m
		}
	}

	if len(metaData) > 0 {
		resp["meta"] = metaData
	}

	js, err := json.Marshal(resp)
	if err != nil {
		ERROR(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

// ERROR writes to ResponseWriter error
func ERROR(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	resp := map[string]interface{}{
		"code":    error2.ErrService,
		"message": "Internet Server Error",
	}

	switch apiErr := err.(type) {
	case error2.APIError:
		resp["code"] = apiErr.Code
		resp["message"] = apiErr.Message
		code = apiErr.Code.GetHTTPCode()
	case error2.ValidationError:
		resp["code"] = apiErr.Code
		resp["message"] = apiErr.Message
		resp["validation_errors"] = apiErr.Errors
		code = http.StatusBadRequest
	}

	js, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(js)
}
