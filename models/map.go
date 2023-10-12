package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MapInfo struct {
	Title     string  `json:"title" bson:"title,omitempty"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `jsonL:"latitude"`
}
type GeneratedMap struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MapTemplate string             `json:"template" bson:"template"`
}
type Location struct {
	PlaceID     int     `json:"place_id,omitempty"`
	Licence     string  `json:"licence,omitempty"`
	OSMType     string  `json:"osm_type,omitempty"`
	OSMID       int     `json:"osm_id,omitempty"`
	Latitude    string  `json:"lat,omitempty"`
	Longitude   string  `json:"lon,omitempty"`
	PlaceRank   int     `json:"place_rank,omitempty"`
	Category    string  `json:"category,omitempty"`
	Type        string  `json:"type,omitempty"`
	Importance  float64 `json:"importance,omitempty"`
	AddressType string  `json:"addresstype,omitempty"`
	Name        string  `json:"name,omitempty"`
	DisplayName string  `json:"display_name,omitempty"`
	Address     struct {
		Leisure       string `json:"leisure,omitempty"`
		Road          string `json:"road,omitempty"`
		Neighbourhood string `json:"neighbourhood,omitempty"`
		Town          string `json:"town,omitempty"`
		Postcode      string `json:"postcode,omitempty"`
		Country       string `json:"country,omitempty"`
		CountryCode   string `json:"country_code,omitempty"`
	} `json:"address,omitempty"`
	BoundingBox []string `json:"boundingbox,omitempty"`
}
