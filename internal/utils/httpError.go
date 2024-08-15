package utils

import "net/http"

type Error struct {
	Status  uint
	Message string
}

func HttpStatus500(w http.ResponseWriter) {
	AsJson(w, Error{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
}
