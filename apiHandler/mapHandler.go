package apiHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GenerateMap(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var mapData []models.MapInfo

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&mapData)
	if err != nil {
		logs.ErrorLogger.Println("Error While Decoding JSON in GenerateMap:mapHandler : ", err.Error())
	}

	rst, saveerr := customizedNFTFacade.SaveMap(mapData)
	if saveerr != nil {
		logs.ErrorLogger.Println("failed to save map: ", saveerr.Error())
		errors.BadRequest(w, "Failed to save map")
		return
	}
	commonResponse.SuccessStatus[string](w, rst)
}

func GetMapByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	mapRst, mapErr := customizedNFTFacade.GetMapByID(vars["mapid"])
	if mapErr != nil {
		fmt.Fprint(w, "<h1>Error</h1>")
	}
	fmt.Fprint(w, mapRst)
}
