package songs

import (
	"encoding/json"
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

func GetSongs(log *logrus.Logger, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := 10
		offset := 0

		if !str.IsEmpty(limitStr) {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			} else {
				log.Warn(err)
				renderError(http.StatusBadRequest, w, r, err.Error())

				return
			}
		}
		if !str.IsEmpty(offsetStr) {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			} else {
				log.Warn(err)
				renderError(http.StatusBadRequest, w, r, err.Error())

				return
			}
		}

		filters := make(map[string]string)
		for _, f := range FilterFields {
			fValue := r.URL.Query().Get(f)
			if fValue != "" {
				filters[ConvertFField(f)] = fValue
			}
		}

		songs, err := service.GetSongs(limit, offset, &filters)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		render.JSON(w, r, Response{
			Status: "ok",
			Songs:  songs,
		})
	}
}

func UpdateSong(log *logrus.Logger, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updated models.Song
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		updatedId, err := service.UpdateSong(id, updated)
		if err != nil {
			log.Error(err)
			renderError(http.StatusInternalServerError, w, r, err.Error())

			return
		}

		render.JSON(w, r, Response{
			Status:  "ok",
			Message: fmt.Sprintf("updated with id %d", updatedId),
		})
	}
}

func DeleteSong(log *logrus.Logger, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		deleted, err := service.DeleteSong(id)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		render.JSON(w, r, Response{
			Status:  "ok",
			Message: fmt.Sprintf("delete with id %d", deleted),
		})
	}
}

func GetText(log *logrus.Logger, service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := MaxVersesLimit
		offset := 0

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}
		if !str.IsEmpty(limitStr) {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			} else {
				log.Warn(err)
				renderError(http.StatusBadRequest, w, r, err.Error())

				return
			}
		}
		if !str.IsEmpty(offsetStr) {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			} else {
				log.Warn(err)
				renderError(http.StatusBadRequest, w, r, err.Error())

				return
			}
		}

		verses, err := service.GetSongText(id, limit, offset)
		if err != nil {
			log.Warn(err)
			renderError(http.StatusBadRequest, w, r, err.Error())

			return
		}

		render.JSON(w, r, Response{
			Status: "ok",
			Text:   verses,
		})
	}
}
