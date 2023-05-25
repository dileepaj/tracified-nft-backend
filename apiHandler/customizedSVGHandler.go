package apiHandler

import (
	"encoding/json"
	"net/http"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/mux"
)

/**
 ** Description : Updaes the DB with the image hash provided. Seareches by mongo DB object ID
 ** Returns : returns the svg if successfully updated.
 */
func UpdateSVGUserMappingbySha256(W http.ResponseWriter, r *http.Request) {
	var UpdateSVG models.UserNFTMapping
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode((&UpdateSVG))
	if err != nil {
		logs.ErrorLogger.Println("Error Decoding request body data : ", err.Error())
		errors.BadRequest(W, err.Error())
		return
	} else {
		_, err := customizedNFTFacade.UpdateUserMappingbySha256(UpdateSVG)
		if err != nil {
			logs.ErrorLogger.Println("Error While updating hash : ", err.Error())
			errors.BadRequest(W, err.Error())
			return
		}
		commonResponse.SuccessStatus[string](W, "SVG updated Scuccessfully")

	}
}

/**
 ** Description : Will be used to get the the svg by the hash
 ** Reutrns : will reutrn the  SVG is a valide hash is provided
 */
func GetSVGbySha256(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	if vars["hash"] != "" {
		rst, err := customizedNFTFacade.GetSVGbySha256(vars["hash"])
		if err != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, rst)
		}
	}
}

// ! Testing methods remove after full impl
func SaveTDPDataByBatchID(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	logs.InfoLogger.Println("URL Param:", vars["batchID"])
	_, err := customizedNFTFacade.GetTDPDataByBatchID(vars["batchID"])
	if err != nil {
		logs.ErrorLogger.Println("failed to get batch data: ", err.Error())
	}
}

func GenerateSVG(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateSVG requestDtos.GenerateSVGReqeust
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&requestCreateSVG)
	if err != nil {
		errors.BadRequest(W, err.Error())
		return
	}
	batchData, err := customizedNFTFacade.GetBatchIDDatabyItemID(requestCreateSVG.ShopID)
	if err != nil {
		errors.BadRequest(W, err.Error())
		return
	}
	rst, err1 := SVGGen(batchData.BatchID, requestCreateSVG.Email, requestCreateSVG.ReciverName, requestCreateSVG.CustomMessage, batchData.ItemID, requestCreateSVG.ShopID, requestCreateSVG.NFTName)
	if err1 != nil {
		errors.BadRequest(W, err1.Error())
		return
	}
	commonResponse.SuccessStatus[responseDtos.SVGforNFTResponse](W, rst)
}

func SVGGen(batchID, email, reciverName, msg, productID, shopID string, nftname string) (responseDtos.SVGforNFTResponse, error) {
	// var tempBatchID = "RURI_VSAPPH_013" //? Templary hardcoded
	// tempBatchID := base64.StdEncoding.EncodeToString([]byte(`{"id":"` + "VSAPPH_013" + `","type":"barcode"}`))
	// var tempBatchID = "eyJpZCI6IlJVUklfVlNBUFBIXzAxMyIsInR5cGUiOiJiYXJjb2RlIn0=" //identifier is base64 encoded {"id":"RURI_VSAPPH_013","type":"barcode"}
	svg, err := customizedNFTFacade.GenerateandSaveSVG(batchID, email, reciverName, msg, productID, shopID, nftname)
	if err != nil {
		return svg, err
	}
	return svg, nil
}
