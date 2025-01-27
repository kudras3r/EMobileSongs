package songs

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/kudras3r/EMobile/internal/models"
)

type Response struct {
	Status  string         `json:"status"`
	Error   string         `json:"error,omitempty"`
	Message string         `json:"message,omitempty"`
	Songs   []models.Song  `json:"songs,omitempty"`
	Text    []models.Verse `json:"text,omitempty"`
}

const (
	Error = "error"
)

func renderError(status int, w http.ResponseWriter, r *http.Request, msg string) {
	w.WriteHeader(status)
	render.JSON(w, r, Response{
		Status: Error,
		Error:  msg,
	})
}
