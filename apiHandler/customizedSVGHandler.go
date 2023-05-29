package apiHandler

import (
	"encoding/json"
	"net/http"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/svgGeneratorforNFT/svgNFTGenerator"
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
	customizedNFTFacade.GetTDPDataByBatchID(vars["batchID"])
}

func GenerateSVG(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateSVG requestDtos.GenerateSVGReqeust
	var customizedNft models.CustomizedNFT

	tenantName, paramerr := r.URL.Query()["tenant"]
	if !paramerr {
		errors.BadRequest(W, "invalid Query param")
		return
	}
	tenantRst, getTenantErr := customizedNFTFacade.ValidateWalletTenant(tenantName[0])

	if getTenantErr != nil || tenantRst.Name == "" {
		errors.BadRequest(W, "Invalid tenant")
		return
	}

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&requestCreateSVG)
	if err != nil {
		errors.BadRequest(W, err.Error())
		return
	}

	if tenantName[0] == "RURI" {
		customizedNft = &svgNFTGenerator.RURINFT{
			Email:        requestCreateSVG.Email,
			ShopID:       requestCreateSVG.ShopID,
			ReceiverName: requestCreateSVG.ReciverName,
			CustomMsg:    requestCreateSVG.CustomMessage,
			NFTName:      requestCreateSVG.NFTName,
			Logo:         tenantRst.Logo,
			EmailTitle:   tenantRst.EmailTopic,
		}

	}

	rst, err1 := customizedNft.GenerateNFT()

	if err1 != nil {
		errors.BadRequest(W, err1.Error())
		return
	}
	commonResponse.SuccessStatus[responseDtos.SVGforNFTResponse](W, rst)
}
