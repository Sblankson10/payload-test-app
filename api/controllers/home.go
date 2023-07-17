package controllers

import (
	"fmt"
	"net/http"
	"payload-app/api/middleware"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("%s , app is up and running", r.Method)

	if err := middleware.WriteJSON(w, 200, response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
