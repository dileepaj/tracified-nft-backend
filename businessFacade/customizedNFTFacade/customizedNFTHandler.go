package customizedNFTFacade

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/svgGeneratorforNFT/svgNFTGenerator"
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
		fmt.Println("Current OTP is", otp)
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
	logs.InfoLogger.Println("result:", string(body))
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
	logs.InfoLogger.Println("data to be returned:", itemdata)
	return itemdata
}

/**
 * Descrition : Function is responsible for generating the SVG and sending it to the DB for it tp get saved
 **Param : batchID : batch ID of product
 **Param : email : email address of user
 **reutrns : models.UserNFTMapping : Contains the generated SVG
 */
func GenerateandSaveSVG(batchID, email, reciverName, msg, productID, shopID string, nftname string) (responseDtos.SVGforNFTResponse, error) {
	var userSVGMapRst responseDtos.SVGforNFTResponse
	tdpData, _ := GetDigitalTwinData(batchID, productID)
	var userNftMapping models.UserNFTMapping
	//Svg will be generated using the template
	svgrst, _ := GenerateSVG(tdpData, batchID, productID, shopID, reciverName, msg, nftname)
	userNftMapping.BatchID = batchID
	userNftMapping.SVG = svgrst
	userNftMapping.Email = email
	userNftMapping.NFTName = nftname
	//Generated SVG data will get added to the DB
	rst, err := svgRepository.SaveUserMapping(userNftMapping)
	if err != nil {
		return userSVGMapRst, err
	}
	userSVGMapRst = rst
	return userSVGMapRst, nil
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
 * Descrition : Generates and returns the SVG
 **Param : []models.TDP : Contains a list of the TDP data for the provided batchID
 **Param : batchID string : batchID
 **reutrns : reutrns the generated SVG as a string
 */
func GenerateSVG(data []models.Component, batchID string, productID string, shopID string, receiverName string, message string, nftname string) (string, error) {
	return svgNFTGenerator.GenerateSVGTemplateforNFT(data, batchID, productID, shopID, receiverName, message, nftname)
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
