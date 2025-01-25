package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func RegisterRoutes(r *chi.Mux, log *logrus.Logger, s SongsGetter) {
	r.Route("/songs", func(r chi.Router) {
		r.Get("/", getSongs(log, s))
		r.Get("/{id}/text", getText(log, s))
	})
}
