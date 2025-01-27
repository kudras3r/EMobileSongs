package api

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/go-chi/chi/v5"
	"github.com/kudras3r/EMobile/internal/api/songs"
	"github.com/sirupsen/logrus"

)

func RegisterRoutes(r *chi.Mux, log *logrus.Logger, s songs.Service) {
	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs(log, s))
		r.Get("/{id}/text", songs.GetText(log, s))
		r.Put("/{id}", songs.UpdateSong(log, s))
		r.Delete("/{id}", songs.DeleteSong(log, s))
	})
	r.Get("/swagger/{path:.*}", httpSwagger.WrapHandler)
}
