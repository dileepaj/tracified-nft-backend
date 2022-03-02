package businessFacades

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/controllers"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/dileepaj/tracified-nft-backend/wrappers/requestWrappers"
	"github.com/gorilla/mux"
)

func CreateNFT(w http.ResponseWriter, r *http.Request) {
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
	} else {
		_, err1 := controllers.CreateNFT(createNFTObject.NFT)
		_, err2 := controllers.SaveOwnership(createNFTObject.Ownership)
		if err1 != nil || err2 != nil {
			ErrorMessage := err1.Error() + err2.Error()
			errors.BadRequest(w, ErrorMessage)
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

func MakeSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var udpateNFTObj requestWrappers.UpdateNFTSALERequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&udpateNFTObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateMakeSale(udpateNFTObj)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err1 := controllers.MakeSaleNFT(udpateNFTObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
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

func GetAllONSaleNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	fmt.Println(vars["status"], vars["userpk"])

	results, err1 := controllers.GetNFTBySellingStatusAndNotUserCreated(vars["status"], vars["userpk"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

func GetNFTbyBlockChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	fmt.Println(vars["status"], vars["userpk"])

	results, err1 := controllers.GetNFTbyBlockChain(vars["blockchain"])
	fmt.Println("results-----------------------", results)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

func GetNFTbyTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	var arr []string
	_ = json.Unmarshal([]byte(vars["tags"]), &arr)
	fmt.Println(vars["tags"], arr)

	results, err1 := controllers.GetNFTbyTags(arr)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

func GetNFTFromWatchList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	results1, err1 := controllers.FindNFTIdentifieryByUserId(vars["userId"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		result2, err2 := controllers.GetNFTbyNFTIdentifier(results1)
		if err2 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(result2)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func GetNFTByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	results1, err1 := controllers.GetBCAccountPKByUserId(vars["userId"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if len(results1) != 0 {
			result2, err2 := controllers.GetNFTbyAccount(results1)
			if err2 != nil {
				ErrorMessage := err1.Error()
				errors.BadRequest(w, ErrorMessage)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(result2)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode("Can not find the accounts for given Tenet Name")
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func GetNFTByTenentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	results1, err1 := controllers.GetBCAccountPKByTenetName(vars["tenentname"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if len(results1) != 0 {
			result2, err2 := controllers.GetNFTbyAccount(results1)
			if err2 != nil {
				ErrorMessage := err1.Error()
				errors.BadRequest(w, ErrorMessage)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(result2)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode("Can not find the accounts for given userId")
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}
