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
	PlaceID     int     `json:"place_id"`
	Licence     string  `json:"licence"`
	OSMType     string  `json:"osm_type"`
	OSMID       int     `json:"osm_id"`
	Latitude    string  `json:"lat"`
	Longitude   string  `json:"lon"`
	PlaceRank   int     `json:"place_rank"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	Importance  float64 `json:"importance"`
	AddressType string  `json:"addresstype"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Address     struct {
		Leisure       string `json:"leisure"`
		Road          string `json:"road"`
		Neighbourhood string `json:"neighbourhood"`
		City          string `json:"city"`
		Postcode      string `json:"postcode"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
	} `json:"address"`
	BoundingBox []string `json:"boundingbox"`
}
