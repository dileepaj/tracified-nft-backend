package otpService

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
)

// baseUrl : url usethe retrive Otp's S3 bucked Url
var baseUrl = configs.GetBackeBaseUrl() + "/traceabilityProfiles/ecommerce/nft/"

/*
	function Get S3 bucke url for Batch Otp, this fuction perform the get request ,It call the tracied backend and retrive otp strored url
	?Parameter
		productId : itemId of the selected ProductName :should be not empty
		batchId : batchId ofthe slected batch :should be not empty
		otpType : selectedOtp Type (id should be "Batch")
		token : user Auth token :should be not empty
	?return
		batch OTP url :string
		error
*/
func GetOtpForBatchURL(productId string, batchId string, otpType string, token string) (string, error) {
	if otpType == "Batch" && productId != "" && batchId != "" {
		url := baseUrl + productId + "/" + batchId
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Add("authorization", token)
		req.Header.Add("cache-control", "no-cache")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return string(body), nil
	} else {
		return "", errors.New("Invalied OTP Type or batch")
	}
}

/*
	function Get S3 bucke url for Stage Otp, this fuction perform the get request ,It call the tracied backend and retrive otp strored url
	?Parameter
		ArtifactId : Artifact Id of the selacted artifact :should be not empty
		otpType : selectedOtp Type (id should be "Batch") :should be not empty
		token : user Auth token :should be not empty
	?return
		stage(master)  OTP url :string
		error
*/
func GetOtpForArtifactURL(artifactId string, otpType string, token string) (string, error) {
	if otpType == "Artifact" && artifactId != "" {
		url := baseUrl + "artifact/id/" + artifactId
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Add("authorization", token)
		req.Header.Add("cache-control", "no-cache")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return string(body), nil
	} else {
		return "", errors.New("Invalied OTP Type or artifact")
	}
}

/*
	function GetOtpJSON() this fuction perform the get request ,It call the s3 bucket stored Url and take the OTP JSON and convert it to string
	?Parameter
		OtpUrl : S3 Bucket Otp Url  : Type string
		bearerToken user Auth token
	?return
		Json OTP :string
		error
*/
func GetOtpJSON(otpUrl string, bearerToken string) (string, error) {
	var otp responseDtos.OTP
	// remove the otpUrl backslashes assignit to OTp struct
	json.Unmarshal([]byte(otpUrl), &otp)

	// http get request
	req, err := http.NewRequest("GET", otp.Url, nil)
	if err != nil {
		return "", err
	}

	// res: response of request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// add headers
	req.Header.Add("authorization", bearerToken)
	req.Header.Add("cache-control", "no-cache")

	// read the body of the response
	body, _ := ioutil.ReadAll(res.Body)

	// checked the Json is valid or not andsend he proper response
	if strings.HasPrefix(string(body), "[{") {
		return string(body), nil
	} else {
		return "", errors.New("Invalid Json Formate")
	}
}
