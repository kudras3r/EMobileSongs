package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/kudras3r/EMobile/internal/models"
	"github.com/sirupsen/logrus"
)

type SongsGetter interface {
	GetSongs(limit, offset int, filters *map[string]string) ([]models.Song, error)
	GetSongText(id, limit, offset int) ([]models.Verse, error)
	DeleteSong(id int) (int, error)
}

var filterFields []string = []string{"id", "song", "group", "link", "releaseDate", "versesCount", "text"}

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
		fmt.Println(limitStr, offsetStr)

		limit := 10
		offset := 0

		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l >= 0 {
				limit = l
			} else {
				log.Warn(fmt.Sprintf("invalid limit %s", limitStr))
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, Response{
					Status: "error",
					Error:  "invalid limit",
				})
				return
			}
		}
		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
				offset = o
			} else {
				log.Warn(fmt.Sprintf("invalid offset %s", offsetStr))
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, Response{
					Status: "error",
					Error:  "invalid offset",
				})
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
			log.Error("failed to get songs")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, Response{
				Status: "error",
				Error:  "cannot get songs",
			})
			return
		}

		render.JSON(w, r, songs)
	}
}
