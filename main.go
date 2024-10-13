package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"net/http"
	_ "song_library/docs"
	"song_library/model"
	"strconv"
	"time"
)

// Song used for swagger
type Song struct {
	Title       string    `json:"Title" example:"Song title"`
	Author      string    `json:"Author" example:"Song author"`
	SongGroup   string    `json:"SongGroup" example:"Song group"`
	Link        string    `json:"Link" example:"https://www.youtube.com/watch?v=b_h8kh-PEfI9999"`
	Description string    `json:"Description" example:"Song description text"`
	ReleaseDate time.Time `json:"ReleaseDate" example:"2006-02-01T15:04:05Z" gorm:"default:current_timestamp"`
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		pageNum, err := strconv.Atoi(q.Get("page"))
		if err == nil {
			pageSize, _ := strconv.Atoi(q.Get("items_per_page"))
			switch {
			case pageSize > 100:
				pageSize = 100
			case pageSize <= 0:
				pageSize = 5
			}

			offset := (pageNum - 1) * pageSize
			return db.Offset(offset).Limit(pageSize)
		}

		return db
	}
}

// getSongVerses godoc
// @Summary Retrieves song's verses based on given song ID
// @Produce json
// @Param id path integer true "Song ID"
// @Param page query string false "paginating results - ?page=1"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /song/{id}/verses [GET]
func getSongVerses(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving songs data")

	w.Header().Set("Content-Type", "application/json")
	var verses []model.Verses

	// Extract query parameters
	params := mux.Vars(r)

	// Get song ID from parameters
	songId, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Requesting song's verses data by song_id
	model.DB.Where("song_id = ?", songId).Scopes(Paginate(r)).Find(&verses)
	if len(verses) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, err := json.Marshal(verses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// getSongs godoc
// @Summary Retrieves songs
// @Produce json
// @Param page query string false "paginating results - ?page=1"
// @Param title query string  false "song search - ?title=Some title"
// @Param description query string  false "song search - ?description=Some descr"
// @Param author query string  false "song search - ?author=Some author"
// @Param song_group query string  false "song search - ?song_group=Some group"
// @Param release_date query string  false "song release_date - ?release_date=1995-07-16"
// @Success 200 {object} model.Song
// @Failure 500
// @Router /songs [get]
func getSongs(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving songs data")
	w.Header().Set("Content-Type", "application/json")
	var songs []model.Song

	db := model.DB.Preload("Verses")

	// Check get parameters for filtration
	query := r.URL.Query()
	if len(query) != 0 {
		for key, value := range query {
			if key != "page" && key != "items_per_page" && key != "release_date" {
				//db.Where(key+" LIKE ?", "%"+value[0]+"%")
				db.Where(key+"=?", value[0])
				log.Printf("Using filter %s", key)
			}
			if key == "release_date" {
				db.Where("CAST(release_date AS date)=?", value[0])
			}
		}
	}

	// Exclude song with "deleted_at"
	db.Where("deleted_at IS NULL").Find(&songs)

	// Add a pagination
	db.Scopes(Paginate(r)).Find(&songs)

	res, err := json.Marshal(songs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
	}

}

// getSong godoc
// @Summary Retrieves song based on given ID
// @Produce json
// @Param id path integer true "Song ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /song/{id} [GET]
func getSong(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving song data")
	w.Header().Set("Content-Type", "application/json")
	var song model.Song

	db := model.DB.Preload("Verses")

	params := mux.Vars(r)

	// Receiving song's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	db.Where("deleted_at IS NULL").First(&song, id)
	if song.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, err := json.Marshal(song)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(res)
		}
	}
}

// createSong godoc
// @Summary Create a song
// @Produce json
// @Accept json
// @Param body body model.Song true "body"
// @Success 201
// @Failure 500
// @Router /song [POST]
func createSong(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating song...")
	var song model.Song
	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &song)
	if err != nil {
		log.Println("Error on data: ", err.Error())
	}

	model.DB.Create(&song)

	w.Header().Set("Content-Type", "application/json")

	// Send song's data in response
	res, err := json.Marshal(song)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Song was created:")
		log.Println(song)
		_, err = w.Write(res)
	}
}

// updateSong godoc
// @Summary Updates song based on given ID
// @Produce json
// @Accept json
// @Param id path integer true "Song ID"
// @Param body body Song true "body"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /song/{id} [PUT]
func updateSong(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating song...")
	var songData model.Song
	var song model.Song
	params := mux.Vars(r)

	// Receiving song's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &songData)
	if err != nil {
		log.Println(err.Error())
	} else {
		// Get song to update
		model.DB.First(&song, id)
		if song.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			// Update song with received data
			model.DB.Model(&song).Updates(songData)

			w.WriteHeader(http.StatusOK)
			log.Printf("Song with id=%s was updated", id)
		}

	}

}

// deleteSong godoc
// @Summary Delete a song by id
// @Param id path integer true "PAGE"
// @Success 200
// @Failure 400
// @Router /song/{id} [DELETE]
func deleteSong(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting song...")
	var song model.Song
	params := mux.Vars(r)

	// Receiving song's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	model.DB.Select(clause.Associations).Delete(&song, id)
	w.WriteHeader(http.StatusOK)
	log.Printf("Song with id=%s was deleted", id)
}

// @title Songs library API
// @version 1.0
// @description Swagger API for Songs library API.
// @termsOfService http://swagger.io/terms/

// @contact.name Dmitrii
// @contact.email ladovod@gmail.com

// @BasePath /api/v1
func main() {
	fmt.Println("Initializing database")
	model.InitDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/songs", getSongs).Methods("GET")
	router.HandleFunc("/api/v1/song/{id}", getSong).Methods("GET")
	router.HandleFunc("/api/v1/song/{id}/verses", getSongVerses).Methods("GET")
	router.HandleFunc("/api/v1/song", createSong).Methods("POST")
	router.HandleFunc("/api/v1/song/{id}", deleteSong).Methods("DELETE")
	router.HandleFunc("/api/v1/song/{id}", updateSong).Methods("PUT")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
