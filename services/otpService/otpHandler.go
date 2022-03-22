package otpService

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/configs"
)

var (
	baseUrl = configs.GetBackeBaseUrl() + "/traceabilityProfiles/ecommerce/nft/"
	token   = configs.GetBackenToken()
)

func GetOtpForBatch(productId string, batchId string, otpType string) (string, error) {
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

func GetOtpForArtifact(artifactId string, otpType string) (string, error) {
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
