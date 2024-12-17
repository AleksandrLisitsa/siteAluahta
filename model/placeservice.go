package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PlaceService interface {
	LoadPlaces() error
	GetPlacesByCategory(category string) []Place
}

type placeService struct {
	places []Place
}

func NewPlaceService() PlaceService {
	return &placeService{}
}

func (ps *placeService) LoadPlaces() error {
	file, err := os.Open("siteAlushta/places.json")
	if err != nil {
		if os.IsNotExist(err) {
			ps.places = []Place{}
			return nil
		}
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &ps.places)
}

func (ps *placeService) GetPlacesByCategory(category string) []Place {
	if category == "" {
		return ps.places
	}

	var filteredPlaces []Place
	for _, place := range ps.places {
		if place.Category == category {
			filteredPlaces = append(filteredPlaces, place)
		}
	}
	return filteredPlaces
}
