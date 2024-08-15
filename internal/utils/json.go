package utils

import (
	"encoding/json"
	"net/http"
)

func AsJson(w http.ResponseWriter, o any) error {
	j, err := json.Marshal(o)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
	return nil
}
