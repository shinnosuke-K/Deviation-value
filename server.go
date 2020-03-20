package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"log"
	"net/http"
	"os"
	"strconv"

	"encoding/json"
)

type info struct {
	ID         int    `json:"id"`
	Deviation  string `json:"deviation"`
	SchoolName string `json:"school_name"`
	Course     string `json:"course"`
	URL        string `json:"url"`
	Prefecture string `json:"prefecture"`
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/images/favicon.ico")
}

func getHost() string {
	if host := os.Getenv("DATABASE_URL"); host != "" {
		return host
	}
	return "host=db user=prefectures password=root dbname=prefectures port=5432 sslmode=disable"
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	prefecture := r.Form["prefecture"][0]
	deviation, err := strconv.Atoi(r.Form["deviation"][0])
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", getHost())
	defer db.Close()
	if err != nil {
		log.Println(err)
	}

	info := []info{}
	db.Where("deviation = ? and prefecture = ?", deviation, prefecture).Find(&info)

	json, err := json.Marshal(info)
	w.Write(json)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/info", infoHandler)

	// for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
