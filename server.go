package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"encoding/json"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type info struct {
	ID         int    `json:"id"`
	Deviation  string `json:"deviation"`
	SchoolName string `json:"school_name"`
	Course     string `json:"course"`
	URL        string `json:"url"`
	Prefecture string `json:"prefecture"`
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

func favionHandler(w http.ResponseWriter, r *http.Request) {
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
	db.Where("deviation between ? and ? AND prefecture = ?", deviation-5, deviation+5, prefecture).Find(&info)

	json, err := json.Marshal(info)
	w.Write(json)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.HandleFunc("/favicon.ico", favionHandler)
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
