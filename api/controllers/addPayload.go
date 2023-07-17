package controllers

import (
	"net/http"
	"payload-app/api/entities"
	"payload-app/api/middleware"
	"payload-app/api/models"
)

func (s *Server) AddPayload(w http.ResponseWriter, r *http.Request) {
	var input entities.IncomingPayload
	

	err := middleware.ReadJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	provider := &entities.CreateProvider{
		DepositsId:  input.DepositsId,
		ProviderRef: input.ProviderRef,
	}

	id, err := models.Insert(provider, s.Db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = middleware.WriteJSON(w, http.StatusCreated, id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
