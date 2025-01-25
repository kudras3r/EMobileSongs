package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	Error = "error"
)

func ErrorBadRequest(w http.ResponseWriter, r *http.Request, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, Response{
		Status: Error,
		Error:  msg,
	})
}
