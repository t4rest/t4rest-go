package response

import (
	"encoding/json"
	"net/http"

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
	HTTPStatus := http.StatusInternalServerError
	resp := map[string]interface{}{
		"code":    ErrService,
		"message": "Internet Server Error",
	}

	switch apiErr := err.(type) {
	case APIError:
		resp["code"] = apiErr.Code
		resp["message"] = apiErr.Message
		HTTPStatus = apiErr.Code.GetHTTPCode()
	case ValidationError:
		resp["code"] = apiErr.Code
		resp["message"] = apiErr.Message
		resp["validation_errors"] = apiErr.Errors
		HTTPStatus = http.StatusBadRequest
	}

	js, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatus)
	_, _ = w.Write(js)
}
