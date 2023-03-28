package mapGenerator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func GetCityName(lat, lon string) (string, error) {
	url := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=" + lon + "&lon=" + lat
	logs.InfoLogger.Println("URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var mapTemplate models.Location
	err = json.NewDecoder(resp.Body).Decode(&mapTemplate)
	if err != nil {
		return "", err
	}
	if mapTemplate.Address.City == "" {
		return mapTemplate.Address.Country, nil
	}
	return mapTemplate.Address.City, nil

}

func GeneratePoints(mapdata []models.MapInfo) string {
	var geoData string = ""
	for i, v := range mapdata {

		long := fmt.Sprintf("%f", v.Longitude)
		lat := fmt.Sprintf("%f", v.Latitude)
		cityname, cityerr := GetCityName(lat, long)
		if cityerr != nil {
			logs.ErrorLogger.Println("failed to get city name: ", cityerr.Error())
			return ""
		}
		geoData += `["` + cityname + `",` + long + `,` + lat + `]`
		if i != len(mapdata)-1 {
			geoData += `,`
		}
	}
	logs.InfoLogger.Println("Cordinates list : ", geoData)
	return geoData
}
func GenerateMap(mapdata []models.MapInfo) string {
	mapPoints := GeneratePoints(mapdata)
	mapHTML :=
		configs.GetMapHead() +
			configs.GetMapBody() +
			configs.GetScriptStart() +
			mapPoints +
			configs.GetLowerScript()
	return mapHTML
}
