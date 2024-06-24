package main

import (
	"log"
	"net/http"

	handler "github.com/EricFrancis12/evoclick/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func main() {
	godotenv.Load(".env.local", ".env")

	server := NewAPIServer(":3001")
	server.Run()
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/click", handler.Click)
	router.HandleFunc("/postback", handler.Postback)
	router.HandleFunc("/t", handler.T)
	router.HandleFunc("/test", handler.Test)

	log.Println("Dev API running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
