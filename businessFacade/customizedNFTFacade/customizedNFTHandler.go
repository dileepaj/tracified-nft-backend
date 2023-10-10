package customizedNFTFacade

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/xlzd/gotp"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
)

// OTP refers to One Time Password
/**
 * Description:Function will be used to generate a new OTP by using a secret of length 2.
 * An otp with 4 alphanumeric characters will be generated and will be sent via email
 * Libary userd : gotp --> github.com/xlzd/gotp
 * *param: email need to pass the user email
 * *returns: OTP that consoist of 4 alphanumeric characters
 */
func GenerateOTP(email string) (string, error) {
	secretLength := 4
	otp := gotp.RandomSecret(secretLength)
	if otp != "" {
		return otp, nil
	}
	return otp, nil
}

/**
 *  Description : Will save the generated OTP along with user email and batchID
 * *paramm : models.OTPData which consist of email,otp and batch ID (all string)
 * *returns : object ID of new DB enetry or error message if saving fails
 */
func SaveOTP(otpDataSet models.UserAuth) (string, error) {
	return otpRepository.SaveOTP(otpDataSet)
}

/**
 * Descprition : checks if a valid OTP exisit
 * *param : email, users email
 * *param : otp, otp entered by user
 * *reutrns : respective batchID if the otp is valid
 */
func ValidateOTP(email string, otp string) (string, string, error) {
	return otpRepository.ValidateOTP(email, otp)
}

/**
 * Descprition : Resends a new OTP and update DB
 * *param : email, users email
 * *param : otp, otp entered by user
 * *reutrns : respective batchID if the otp is valid
 */
func ResendOTP(otpData models.UserAuth) (string, error) {
	return otpRepository.ResendOTP(otpData)
}

/**
 *  Description : this function formats and send the email tha contains the customers OTP to the customer
 *  lib used : gomail ("gopkg.in/gomail.v2")
 * *param : otp, Otp genereted for customer.
 * *paran : email, email of the customer that the mail will be sent to.
 * *returns : an eroor if it fails to sent the mail.
 */
func SendEmail(otp string, email string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", configs.GetSenderEmailAddress())
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "RURI NFT")
	msg.SetBody("text/html", configs.GetEmail(otp))
	n := gomail.NewDialer(configs.GetEmailHost(), configs.GetEmailPort(), configs.GetSenderEmailAddress(), configs.GetSenderEmailAppPWD())
	if err := n.DialAndSend(msg); err != nil {
		logs.ErrorLogger.Println(err)
		return err
	}
	return nil
}

/**
 * Desctiption : Retreives the respective BatchID when the RUIri product ID is passed.
 **param : productID, Ruri product ID
 **returns : models.RuriItemData(contains batchID) and error is there is any
 * TODO:DO a null check for the response
 */
func GetBatchIDDatabyItemID(shopID string) (models.ItemData, error) {
	var itemdata models.ItemData

	shopify := configs.GetRuriShopifyUrl()

	url := shopify + shopID

	rst, err := http.Get(url)

	if err != nil {
		logs.ErrorLogger.Println("APi call err : ", err.Error())
		return itemdata, err
	} else if rst.StatusCode == 404 {
		err1 := errors.New("Item not found")
		logs.ErrorLogger.Println("APi call err : ", err1.Error())
		return itemdata, err1
	}

	body, err := ioutil.ReadAll(rst.Body)
	defer rst.Body.Close()
	var data = string(body)
	itemdata = FormatBatchIDString(data)
	if err != nil {
		logs.ErrorLogger.Println("err: ", err)
	}
	return itemdata, err
}

/**
 * Descrition : formats the string which containts the item data and assigns data into struct RuriItemData
 **Param : text : String to be formated
 **reutrns : models.RuriItemData
 */
func FormatBatchIDString(text string) models.ItemData {
	// TODO: Have some error checking
	var itemdata models.ItemData
	formatedData := strings.Trim(text, "[")
	formatedData2 := strings.Trim(formatedData, "]")
	splitData := strings.Split(formatedData2, ",")
	itemdata.ItemID = strings.Replace(splitData[0], "\"", "", -1)
	itemdata.HasTracability = strings.Replace(splitData[1], "\"", "", -1)
	itemdata.BatchID = strings.Replace(splitData[2], "\"", "", -1)
	return itemdata
}

func GetSVGbyEmailandBatchID(email string, batchID string) (responseDtos.SVGforNFTResponse, error) {
	return svgRepository.GetSVGbyEmailandBatchID(email, batchID)
}

/**
 * Descrition : Retrives the relevent TDP data for a specific batch ID
 **Param : batchID : batch ID of product
 **reutrns : []models.TDP : Contains a list of the TDP data for the provided batchID
 */
func GetTDPDataByBatchID(batchID string) ([][]models.TDPParent, error) {
	var tdpData [][]models.TDPParent
	//url := "https://api.tracified.com/api/v2/traceabilityProfiles/tdparr/" + batchID
	url := "https://api.tracified.com/api/v2/traceabilityProfiles/generic?identifier=" + batchID

	var bearer = configs.GetBearerToken()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.ErrorLogger.Println("unable to get tdp data :", err.Error())
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(string(body)), &tdpData)
	return tdpData, nil
}

/**
 * Descrition : Retrives the relevent digital twin json for a specific batch ID
 **Param : batchID : batch ID of product
 **reutrns : []models.Component : Contains a list of the TDP data for the provided batchID
 */
func GetDigitalTwinData(batchID string, productID string) ([]models.Component, error) {
	var digitalTwinData []models.Component

	bEnc := b64.StdEncoding.EncodeToString([]byte(batchID))

	dtUrl := configs.GetDigitalTwinUrl()

	url := dtUrl + bEnc + `?itemId=` + productID

	var bearer = configs.GetBearerToken()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.ErrorLogger.Println("unable to get data :", err.Error())
		return digitalTwinData, err
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err1 := client.Do(req)

	if err1 != nil {
		logs.ErrorLogger.Println("unable to get data :", err.Error())
		return digitalTwinData, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(string(body)), &digitalTwinData)
	return digitalTwinData, nil
}

/**
 * Descrition : Update the SVGUserMapping colletion with SVG hash
 **Param :  models.UserNFTMapping : Contains object ID and svg hash
 **reutrns : reutrns SVG as a string
 */
func UpdateUserMappingbySha256(request models.UserNFTMapping) (responseDtos.SVGforNFTResponse, error) {
	update := bson.M{
		"$set": bson.M{"hash": request.Hash},
	}
	return svgRepository.UpdateUserMappingbySha256("_id", request.SvgID, update)

}

/**
 * Descrition : Retruns the SVG based on the hash provided
 **Param : hash string : batchID of gem
 **reutrns : reutrns SVG as a string
 */
func GetSVGbySha256(hash string) (string, error) {
	return svgRepository.GetSVGbySha256(hash)
}

func GetNFTStatus(email string, otp string) (string, error) {
	return otpRepository.ValidateNFTStatus(email, otp)
}

func GetNFTStatusbyShopID(shopID string) (string, error) {
	return otpRepository.ValidateNFTStatusbyShopId(shopID)
}

func GenerateOTPExpireDate() time.Time {
	currentDate := time.Now()
	duration := time.Hour * 24 * 30
	expireDate := currentDate.Add(duration)
	return expireDate
}

func ValidateWalletTenant(tenantName string) (models.WalletNFTTenantUser, error) {
	return otpRepository.GetWalletTenant(tenantName)
}

func GetOTPStatus(email string, id string) (responseDtos.OTPStatus, error) {
	return otpRepository.CheckOTPValidatedStatus(email, id)
}

// GetNFTIdentifiersByTenantName fetches a list of NFT identifiers for a given tenantID
// by making an HTTP request to a specified URL and returns the list as a slice of strings.
func GetNFTIdentifersByTenantName(tenantID string) ([]string, error) {
	var identifierList []string

	// Define the URL for the HTTP request
	url := commons.GoDotEnvVariable("BACKEND_BASEURL")

	// Make an HTTP GET request to the URL
	rst, err := http.Get(url + "/tenant/productID/" + tenantID)

	if err != nil {
		// Handle API call error and log it
		logs.ErrorLogger.Println("API call err : ", err.Error())
		return identifierList, err
	} else if rst.StatusCode == 404 {
		// Handle the case where the API returns a 404 status code
		err1 := errors.New("data not found")
		logs.ErrorLogger.Println("API call err : ", err1.Error())
		return identifierList, err1
	}

	// Read and parse the response body
	body, err := io.ReadAll(rst.Body)
	defer rst.Body.Close()
	var data = string(body)
	err1 := json.Unmarshal([]byte(data), &identifierList)
	if err1 != nil {
		// Handle JSON unmarshaling error and log it
		logs.ErrorLogger.Println("Failed to convert api response : ", err1.Error())
		return identifierList, err1
	}

	// Handle any remaining errors and return the identifierList
	if err != nil {
		logs.ErrorLogger.Println("err: ", err)
	}

	return identifierList, nil
}

// GetMintedNFTIdentifierForWallet fetches minted NFT identifiers for a given tenantID
// by calling GetNFTIdentifiersByTenantName to retrieve API data and then matching it
// against database results. It returns a list of common identifiers as a slice of strings.
func GetMintedNFTIdentifierForWallet(tenantID string) ([]string, error) {
	var matchedList []string
	var commonIDs []string

	// Call GetNFTIdentifiersByTenantName to retrieve API data
	apirst, err := GetNFTIdentifersByTenantName(tenantID)
	if err != nil {
		// Handle the case where there is no data from the API and log it
		logs.ErrorLogger.Println("No data from api")
		return matchedList, err
	}

	// Call nftRepository.GetMintedWalletNFTIdentifiers to retrieve database results
	dbrst, err := nftRepository.GetMintedWalletNFTIdentifiers()
	if err != nil {
		// Handle the case where there is a failure to get DB data and log it
		logs.ErrorLogger.Println("failed to get DB data: ", err.Error())
		return matchedList, err
	}

	// Match the identifiers from the API and the database
	commonIDs = _matchStoreIdentifiers(apirst, dbrst)
	return commonIDs, nil
}

// matchStoreIdentifiers finds common values between two string slices (apiResult and databaseResult)
// by using a map to efficiently identify matches. It returns a list of common values as a slice of strings.
func _matchStoreIdentifiers(apiResult []string, databaseResult []string) []string {
	var commonValues []string

	// Create a map to store the unique values from apiResult
	list1map := make(map[string]bool)
	for _, item := range apiResult {
		list1map[item] = true
	}

	// Iterate through databaseResult and check if each item exists in list1map
	for _, item2 := range databaseResult {
		if list1map[item2] {
			commonValues = append(commonValues, item2)
		}
	}

	return commonValues
}
