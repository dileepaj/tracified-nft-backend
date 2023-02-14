package otpService

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/sirupsen/logrus"
)

var baseUrl = configs.GetBackeBaseUrl() + "/traceabilityProfiles/ecommerce/nft/"

func GetOtpForBatchURL(productId string, batchId string, otpType string, token string) (string, int, error) {
	if otpType == "Batch" && productId != "" && batchId != "" {
		url := baseUrl + productId + "/" + batchId
		logrus.Info("OTP generate url",url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", req.Response.StatusCode, err
		}
		req.Header.Add("authorization", token)
		req.Header.Add("cache-control", "no-cache")
		res, err := http.DefaultClient.Do(req)
		logrus.Info("OTP generate response",res)
		if err != nil {
			logrus.Error("OTP generate ",err)
			return "", 500, err
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode == 200 || res.StatusCode == 204 {
			return string(body), res.StatusCode, nil
		}
		return "", 500, errors.New("Backend server connection issue")
	} else {
		return "", 400, errors.New("Invalied OTP Type or batch")
	}
}

func GetOtpForArtifactURL(artifactId string, otpType string, token string) (string, int, error) {
	if otpType == "Artifact" && artifactId != "" {
		url := baseUrl + "artifact/id/" + artifactId
		logrus.Info("OTP generate url",url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", req.Response.StatusCode, err
		}
		req.Header.Add("authorization", token)
		req.Header.Add("cache-control", "no-cache")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", 500, err
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode == 200 || res.StatusCode == 204 {
			return string(body), res.StatusCode, nil
		}
		return "", res.StatusCode, errors.New("Backend server connection issue")
	} else {
		return "", 400, errors.New("Invalied OTP Type or artifact")
	}
}

func GetOtpJSON(otpUrl string, bearerToken string) (string, error) {
	var otp responseDtos.OTP
	json.Unmarshal([]byte(otpUrl), &otp)
	req, err := http.NewRequest("GET", otp.Url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	req.Header.Add("authorization", bearerToken)
	req.Header.Add("cache-control", "no-cache")
	body, _ := ioutil.ReadAll(res.Body)
	if strings.HasPrefix(string(body), "[{") {
		return string(body), nil
	} else {
		return "", errors.New("Invalid Json Formate")
	}
}
