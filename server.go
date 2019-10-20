package main

import (
	"fmt"

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

func infoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println("prefectures: ", r.Form["prefecture"])
	//fmt.Println("deviation: ", r.Form["deviation"])
	fmt.Println(r.Form)

	prefecture := r.Form["prefecture"][0]
	deviation, err := strconv.Atoi(r.Form["deviation"][0])
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", "user=prefectures password=pre dbname=prefectures sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	info := []info{}
	db.Where("deviation = ? prefecture = ?", deviation, prefecture).Find(&info)

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

	// for Localhost
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("ListenAndServer: ", err)
	// }
}
