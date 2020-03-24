package controller

import (
	"net/http"

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
