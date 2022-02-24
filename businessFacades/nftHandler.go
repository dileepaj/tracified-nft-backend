package businessFacades

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/dao"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/dileepaj/tracified-nft-backend/wrappers/requestWrappers"
)

func SaveNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createNFTObject requestWrappers.CreateNFTRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createNFTObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateCreateNFTObject(createNFTObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	}else{
		_, err1 := dao.SaveNFT(createNFTObject.NFT)
		_, err2 := dao.SaveOwnership(createNFTObject.Ownership)
		if (err1 != nil || err2 !=nil ) {
			ErrorMessage := err1.Error()+err2.Error()
			errors.BadRequest(w,ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "SAVED NFT"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}
