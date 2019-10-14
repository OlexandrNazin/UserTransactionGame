package logger

import (
	"encoding/json"
	"net/http"
)

func ResWriter(w http.ResponseWriter, str string) {
	resp, _ := json.Marshal(str)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
