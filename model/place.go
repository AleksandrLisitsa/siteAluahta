package model

type Place struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
