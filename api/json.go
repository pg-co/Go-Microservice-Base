package api

import (
	"net/http"
	"encoding/json"
)

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		return err
	} else {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.Write(b)
		return nil
	}
}