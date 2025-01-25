package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/kudras3r/EMobile/internal/models"
	"github.com/kudras3r/EMobile/pkg/str"
	"github.com/sirupsen/logrus"
)

const MaxVersesLimit = 10000

type SongsGetter interface {
	GetSongs(limit, offset int, filters *map[string]string) ([]models.Song, error)
	GetSongText(id, limit, offset int) ([]models.Verse, error)
	DeleteSong(id int) (int, error)
}

var filterFields []string = []string{"song", "group", "link", "releaseDate", "versesCount", "text"}

func convertFField(fField string) string {
	var res string

	switch fField {
	case "song":
		res = "title"
	case "group":
		res = "group_performer"
	case "releaseDate":
		res = "release_date"
	case "versesCount":
		res = "verses_count"
	default:
		res = fField
	}

	return res
}

func getSongs(log *logrus.Logger, sGetter SongsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := 10
		offset := 0

		if !str.IsEmpty(limitStr) {
			if l, err := strconv.Atoi(limitStr); err == nil && l >= 0 {
				limit = l
			} else {
				msg := fmt.Sprintf("invalid limit %s", limitStr)
				log.Warn(msg)
				ErrorBadRequest(w, r, msg)

				return
			}
		}
		if !str.IsEmpty(offsetStr) {
			if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
				offset = o
			} else {
				msg := fmt.Sprintf("invalid offset %s", offsetStr)
				log.Warn(msg)
				ErrorBadRequest(w, r, msg)

				return
			}
		}

		filters := make(map[string]string)
		for _, f := range filterFields {
			fValue := r.URL.Query().Get(f)
			if fValue != "" {
				filters[convertFField(f)] = fValue
			}
		}

		songs, err := sGetter.GetSongs(limit, offset, &filters)
		if err != nil {
			msg := "cannot get songs"
			log.Warn(msg)
			ErrorBadRequest(w, r, msg)

			return
		}

		render.JSON(w, r, songs)
	}
}

func getText(log *logrus.Logger, sGetter SongsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		IdStr := chi.URLParam(r, "id")
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := MaxVersesLimit
		offset := 0

		id, err := strconv.Atoi(IdStr)
		if err != nil || id <= 0 {
			msg := fmt.Sprintf("invalid id %s", IdStr)
			log.Warn(msg)
			ErrorBadRequest(w, r, msg)

			return
		}
		if !str.IsEmpty(limitStr) {
			if l, err := strconv.Atoi(limitStr); err == nil && l >= 0 {
				limit = l
			} else {
				msg := fmt.Sprintf("invalid limit %s", limitStr)
				log.Warn(msg)
				ErrorBadRequest(w, r, msg)

				return
			}
		}
		if !str.IsEmpty(offsetStr) {
			if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
				offset = o
			} else {
				msg := fmt.Sprintf("invalid offset %s", offsetStr)
				log.Warn(msg)
				ErrorBadRequest(w, r, msg)

				return
			}
		}

		verses, err := sGetter.GetSongText(id, limit, offset)
		if err != nil {
			msg := "cannot get text"
			log.Warn(msg)
			ErrorBadRequest(w, r, msg)

			return
		}

		render.JSON(w, r, verses)
	}
}
