package main

import (
	"github.com/shinnosuke-K/Deviation-value/controller"
	"github.com/shinnosuke-K/Deviation-value/db"
	"github.com/shinnosuke-K/Deviation-value/util"

	"log"
	"net/http"
)

type Server struct {
	Engine *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Engine: http.NewServeMux(),
	}
}

func (router *Server) Init() error {
	db, err := db.Open()
	if err != nil {
		return err
	}

	ctr := controller.NewController(db)

	router.Engine.Handle("/", ctr.FileHandle())
	router.Engine.Handle("/assets/", ctr.AssetsHandle())
	router.Engine.HandleFunc("/favicon.ico", ctr.FaviconHandler)
	router.Engine.HandleFunc("/info", ctr.InfoHandler)

	return nil
}

func (router *Server) Run() {
	if err := http.ListenAndServe(":"+util.GetPort(), router.Engine); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func main() {

	router := NewServer()
	err := router.Init()
	if err != nil {
		log.Fatal(err)
	}
	router.Run()
}
