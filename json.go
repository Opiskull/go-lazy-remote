package main

import "encoding/json"

// JSONError wrapper
type JSONError struct {
	Error string `json:"error"`
}

func writeJSONError(err error) []byte {
	var jsonError = JSONError{Error: err.Error()}
	return writeJSON(jsonError)
}

func writeJSON(v interface{}) []byte {
	bytes, _ := json.Marshal(&v)
	return bytes
}
