package main

import (
	"html/template"
	"net/http"

	"siteAlushta/controller"
	"siteAlushta/model"
)

func main() {
	service := model.NewPlaceService()
	if err := service.LoadPlaces(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("/siteAlushta/index.html"))
	controller := controller.NewPlaceController(service, tmpl)

	http.HandleFunc("/siteAlushta/", controller.HomeHandler)
	http.Handle("/siteAlushta/static/", http.StripPrefix("/siteAlushta/static/", http.FileServer(http.Dir("siteAlushta/static"))))

	http.ListenAndServe(":8080", nil)
}
