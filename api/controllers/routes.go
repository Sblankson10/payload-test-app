package controllers

import (
	"github.com/gorilla/mux"
)

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", s.Home)
}
