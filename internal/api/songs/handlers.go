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

// GetSongs get songs
// @Summary Get a list of songs
// @Description Retrieve a list of songs with optional pagination and filtering
// @ID get-songs
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of songs returned" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Param filter query string false "Filter fields for songs" 
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /songs [get]
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

// UpdateSong update song by id
// @Summary Update a song by ID
// @Description Update the details of a song using its ID
// @ID update-song
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Updated song data"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /songs/{id} [put]
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

// DeleteSong godoc
// @Summary Delete a song by ID
// @Description Delete a song from the database using its ID
// @ID delete-song
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /songs/{id} [delete]
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

// GetText get song text by id
// @Summary Get the text of a song by ID
// @Description Retrieve the lyrics of a song using its ID with optional pagination
// @ID get-song-text
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param limit query int false "Limit the number of verses returned" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /songs/{id}/text [get]
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
