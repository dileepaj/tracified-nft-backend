package ruriBusinessFacade

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/xlzd/gotp"
	"gopkg.in/gomail.v2"
)

/**
 * Description:Function will be used to generate a new OTP by using a secret of length 2.
 * An otp with 4 alphanumeric characters will be generated and will be sent via email
 * Libary userd : gotp --> github.com/xlzd/gotp
 * *param: email need to pass the user email
 * *returns: OTP that consoist of 4 alphanumeric characters
 */
func GenerateOTP(email string) (string, error) {
	secretLength := 2
	otp := gotp.RandomSecret(secretLength)
	if otp != "" {
		fmt.Println("Current OTP is", otp)
		err := SendEmail(otp, email)
		if err != nil {
			return otp, err
		}
		return otp, err
	}
	return otp, nil
}

/**
* 	TODO:implement function
 */
func GetTDPDatabyBatchID(batchID string) error { //? not called
	url := ""
	rst, err := http.Get(url)
	if err != nil {
		return err
	}
	logs.InfoLogger.Println(rst)
	return nil
}

/**
 *  Description : Will save the generated OTP along with user email and batchID
 * *paramm : models.OTPData which consist of email,otp and batch ID (all string)
 * *returns : object ID of new DB enetry or error message if saving fails
 */
func SaveOTP(otpDataSet models.OTPData) (string, error) {
	return ruriRepository.SaveOTP(otpDataSet)
}

/**
 * Descprition : checks if a valid OTP exisit
 * *param : email, users email
 * *param : otp, otp entered by user
 * *reutrns : respective batchID if the otp is valid
 */
func ValidateOTP(email string, otp string) (string, error) {
	return ruriRepository.ValidateOTP(email, otp)
}

func ResendOTP(otpData models.OTPData) (string, error) {
	return ruriRepository.ResendOTP(otpData)
}

/**
 *  TODO need to implement
 */
func SaveTDP(tdp models.TDP) { //! Configure param data type properly
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
	msg.SetHeader("From", "mithilap@tracified.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "RURI NFT")
	msg.SetBody(
		"text/html",
		`"<html>
		<head></head>
		<body>
			<div style="font-family: Helvetica,Arial,sans-serif;min-width:1000px;overflow:auto;line-height:2">
				<div style="margin:50px auto;width:70%;padding:20px 0">
				  <div style="border-bottom:1px solid #eee">
					<a href="" style="font-size:1.4em;color: #00466a;text-decoration:none;font-weight:600">RURI(OTP generation TEST)</a>
				  </div>
				  <p style="font-size:1.1em">Hi,</p>
				  <p>Thank you for choosing RURI. Use the following OTP to complete your Sign Up procedures. OTP is valid for 1 month</p>
				  <h2 style="background: #00466a;margin: 0 auto;width: max-content;padding: 0 10px;color: #fff;border-radius: 4px;">`+otp+`</h2>
				  <p style="font-size:0.9em;">Regards,<br />RURI</p>
				  <hr style="border:none;border-top:1px solid #eee" />
				  <div style="float:right;padding:8px 0;color:#aaa;font-size:0.8em;line-height:1;font-weight:300">
					<p>Tracified Inc</p>
					<p>1600 Amphitheatre Parkway</p>
					<p>California</p>
				  </div>
				</div>
			</div>
		</body>
	</html>"`)
	n := gomail.NewDialer("smtp.gmail.com", 587, "mithilap@tracified.com", "yeevjlnelugrbawf")
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
func GetBatchIDDatabyItemID(productID string) (models.RuriItemData, error) {
	// ?https://qa.ecom.api.tracified.com/shopifymappings/nisaltest.myshopify.com/6779546796091 <-- sample url
	var itemdata models.RuriItemData
	var shopname = "nisaltest.myshopify.com"
	url := "https://qa.ecom.api.tracified.com/shopifymappings/" + shopname + "/" + productID
	logs.InfoLogger.Println("API call url: ", url)
	rst, err := http.Get(url)
	if err != nil {
		return itemdata, err
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
func FormatBatchIDString(text string) models.RuriItemData {
	// TODO: Have some error checking
	var itemdata models.RuriItemData
	formatedData := strings.Trim(text, "[")
	formatedData2 := strings.Trim(formatedData, "]")
	splitData := strings.Split(formatedData2, ",")
	itemdata.ItemID = strings.Replace(splitData[0], "\"", "", -1)
	itemdata.HasTracability = strings.Replace(splitData[1], "\"", "", -1)
	itemdata.BatchID = strings.Replace(splitData[2], "\"", "", -1)
	logs.InfoLogger.Println("data to be returned:", itemdata)
	return itemdata
}

//! NEED FIXING
/*
func FormatBatchIDString(text string) (string, error) {
	logs.ErrorLogger.Println("before convertion: ", text)
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
	formatedText := re.ReplaceAllString(text, ",")
	logs.InfoLogger.Println("formated text: ", formatedText)
	data := strings.Split(formatedText, " ")
	for index, val := range data {
		logs.InfoLogger.Println("index:", index)
		logs.InfoLogger.Println("val: ", val)
	}
	return formatedText, err
}
*/
