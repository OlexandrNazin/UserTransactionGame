package logger

import (
	"encoding/json"
	"net/http"
)

func ResWriter(w http.ResponseWriter, data map[string]interface{}) {
	resp, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}
