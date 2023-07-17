package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("connection", "close")
	w.WriteHeader(status)
	w.Write(response)

	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576 // set max byte allowed on this body
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}
