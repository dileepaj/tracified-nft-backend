package apiHandler

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/htmlGeneretorService/htmlGenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// handel the HTML generate POST request(Generatee HTML NFT)
func HTMLFileGenerator(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var generateHTMLRequest models.HtmlGenerator
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&generateHTMLRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateHtmlGenerator(generateHTMLRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
			return
		} else {
			// fmt.Println("generateHTMLRequest",generateHTMLRequest)
			result, err := nftComposerBusinessFacade.GenerateHTMLFile(generateHTMLRequest)
			if err != nil {
				errors.InternalError(w, err.Error())
				return
			}
			commonResponse.SuccessStatus[string](w, result)
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

// handel the svg generate POST request(Generatee HTML NFT)
func SVGFileGenerator(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var generateSVGRequest models.HtmlGenerator
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&generateSVGRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateHtmlGenerator(generateSVGRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
			return
		} else {
			result, err := nftComposerBusinessFacade.GenerateSVGFile(generateSVGRequest)
			if err != nil {
				errors.InternalError(w, err.Error())
				return
			}
			commonResponse.SuccessStatus[string](w, result)
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func TimelineHTMLFileGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Get the path of the requested URL
	path := r.URL.Path

	defer context.Clear(r)

	batchBase64, ok := mux.Vars(r)["batchBase64"]
	if !ok {
		errors.BadRequest(w, "batchBase64 parameter is missing")
		return
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(batchBase64)
	if err != nil {
		logrus.Error("Error decoding base64:", err)
		errors.InternalError(w, "Error decoding base64: "+err.Error())
		return
	}

	batctDecodedString := string(decodedBytes)

	productID, ok := mux.Vars(r)["productId"]
	if !ok {
		errors.InternalError(w, "productId parameter is missing")
		return
	}

	customizedNft := &htmlGenerator.JMACNFT{
		BatchID:   batctDecodedString,
		ProductID: productID,
		ItemName:  "", // Consider getting ItemName from the request if needed
	}

	rst, err := customizedNft.GenerateNFT()
	if err != nil {
		logrus.Error(err)
		errors.BadRequest(w, "Failed to generate HTML template: "+err.Error())
		return
	}
	// Check if the path contains "/nft/timeline/html/hash"
	if strings.Contains(path, "/nft/timeline/html/hash") {
		// Create a new SHA-256 hash object
		hasher := sha256.New()

		// Write the string to the hash object
		hasher.Write([]byte(rst))

		// Get the hashed bytes
		hashedBytes := hasher.Sum(nil)

		// Convert the hashed bytes to a hex-encoded string
		hashedString := hex.EncodeToString(hashedBytes)
		base64String := base64.StdEncoding.EncodeToString([]byte(rst))
		result := models.HTMLTimelineHashGenerationResponse{
			TimelineHtmlBase64: base64String,
			TimelineHtmlHash:   hashedString,
		}
		// Encode the structure as JSON
		responseJSON, err := json.Marshal(result)
		if err != nil {
			errors.InternalError(w, "Failed to marshal JSON response")
			return
		}
		// Example response
		w.Write(responseJSON)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Write the HTML response
	_, err = w.Write([]byte(rst))
	if err != nil {
		errors.InternalError(w, "Failed to write HTML response: "+err.Error())
		return
	}
}
