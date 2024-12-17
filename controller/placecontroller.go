package controller

import (
	"html/template"
	"net/http"

	"site/model"
)

type PlaceController interface {
	HomeHandler(w http.ResponseWriter, r *http.Request)
}

type placeController struct {
	service model.PlaceService
	tmpl    *template.Template
}

func NewPlaceController(service model.PlaceService, tmpl *template.Template) PlaceController {
	return &placeController{
		service: service,
		tmpl:    tmpl,
	}
}

func (pc *placeController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	places := pc.service.GetPlacesByCategory(category)
	pc.tmpl.Execute(w, places)
}
