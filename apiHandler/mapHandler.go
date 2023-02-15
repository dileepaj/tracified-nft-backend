package apiHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/context"
)

func GenerateMap(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "text/html")
	var mapData []models.MapInfo

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&mapData)
	if err != nil {
		logs.ErrorLogger.Println("Error While Decoding JSON in GenerateMap:mapHandler : ", err.Error())
	}

	rst := customizedNFTFacade.GetMap(mapData)

	fmt.Fprint(w, rst)
}
