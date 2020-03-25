package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shinnosuke-K/Deviation-value/db"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(d *gorm.DB) *Controller {
	return &Controller{db: d}
}

func (ctr *Controller) FileHandle() http.Handler {
	return http.FileServer(http.Dir("templates"))
}

func (ctr *Controller) AssetsHandle() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
}

func (ctr *Controller) FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/images/favicon.ico")
}

func (ctr *Controller) InfoHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Error parse request", http.StatusBadRequest)
		return
	}

	score := db.Score{
		Prefecture: r.Form["prefecture"][0],
		Deviation:  r.Form["deviation"][0],
	}

	schLists, err := score.GetSchool(ctr.db)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not found school info", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(schLists)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal json", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(res); err != nil {
		log.Println(err)
		http.Error(w, "Not write response", http.StatusInternalServerError)
		return
	}
}
