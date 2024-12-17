package main

import (
	"html/template"
	"net/http"

	"site/controller"
	"site/model"
)

func main() {
	service := model.NewPlaceService()
	if err := service.LoadPlaces(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	controller := controller.NewPlaceController(service, tmpl)

	http.HandleFunc("/", controller.HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
