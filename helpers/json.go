package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJson(w http.ResponseWriter, content map[string]interface{}) {
	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, fmt.Sprintf("cannot encode json: %s\n", err), http.StatusBadRequest)
		return
	}
}
