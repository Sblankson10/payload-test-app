package controllers

import (
	"github.com/gorilla/mux"
)

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", s.Home)
	s.Router.HandleFunc("/send-payload", s.AddPayload).Methods("POST")
}
