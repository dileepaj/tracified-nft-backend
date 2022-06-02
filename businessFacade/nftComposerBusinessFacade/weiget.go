package nftComposerBusinessFacade

import (
	"errors"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/otpService"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveWidgetList(widgets []models.Widget) (string, error) {
	return widgetRepository.SaveWidgetList(widgets)
}

/**
*@fun SaveWidget() this method basically  do a insert widget to widget collecon
*! Before insert
*? 1.check widget type  if WidgetType eqila only (BarChart,PieChart,Table.ProofBot) it will check the OTPTtpe
* if OTPType equal to "Batch" it will fetch the BatchOTP from backend using itemId and batchId ,after assign it to otpString and send to InsertWidgetMethod with OTP
*!  if OTPType equal to "Artifact" we do not keep th OTP in DB
*
*? 2.For Other widget type does not have the OTP
*
**/
func SaveWidget(widget models.Widget,token string) (responseDtos.WidgetIdResponse,int, error) {
	var response responseDtos.WidgetIdResponse
	var otpString string = ""
	var codeStatus int
	var err error
	if widget.WidgetType == "BarChart" || widget.WidgetType == "PieChart" || widget.WidgetType == "BubbleChart" || widget.WidgetType == "Table"{
		if widget.OTPType == "Batch" {
			otpString,codeStatus, err = otpService.GetOtpForBatchURL(widget.ProductId, widget.BatchId, widget.OTPType,token)
		} else if widget.OTPType == "Artifact" {
			otpString,codeStatus, err = otpService.GetOtpForArtifactURL(widget.ArtifactId, widget.OTPType,token)
		} else {
			otpString = ""
			err = errors.New("Invalied OTP Type")
		}

		if err != nil {
			return response,codeStatus, err
		} else if strings.HasPrefix(otpString, `{"err"`) {
			return response,500, errors.New(otpString)
		} else {
			widget.OTP = otpString
			rst2, err := widgetRepository.SaveWidget(widget)
			if rst2 == "" {
				return response,500, errors.New("The file can not insert into DB please check the network")
			}
			if err != nil {
				return response,500, err
			}
			response.WidgetId = rst2
			return response,200, nil
		}
	} else {
		return response,400, errors.New("Invalid Widget Type")
	}
}

func ChangeWidget(widget requestDtos.UpdateWidgetRequest,token string) (models.Widget,int, error) {
	var response models.Widget
	updateWidget := bson.M{}
	var otpString string
	var err error
	var codeStatus int
	rst, err := FindWidgetByWidgetId(widget.WidgetId)
	if err != nil {
		return response,500, err
	}
	if rst.WidgetType == "BarChart" || rst.WidgetType == "PieChart" || rst.WidgetType == "BubbleChart" || rst.WidgetType == "Table"{
		if widget.OTPType == "Batch" {
			otpString,codeStatus, err = otpService.GetOtpForBatchURL(widget.ProductId, widget.BatchId, widget.OTPType,token)
			if err != nil {
				return response,codeStatus, err
			} else if strings.HasPrefix(otpString, `{"err"`) {
				return response,500, errors.New(otpString)
			} else {
				updateWidget = bson.M{
					"$set": bson.M{"timestamp": widget.Timestamp, "batchid": widget.BatchId, "productid": widget.ProductId, "productname": widget.ProductName, "otp": otpString},
				}
				rst2, err := widgetRepository.FindWidgetAndUpdate("widgetid", widget.WidgetId, updateWidget)
				if err != nil {
					return response,400, err
				}
				response = rst2
				return response,200, nil
			}
		} else if widget.OTPType == "Artifact" {
			otpString,codeStatus, err = otpService.GetOtpForArtifactURL(widget.ArtifactId, widget.OTPType,token)
			if err != nil {
				return response,codeStatus, err
			} else if strings.HasPrefix(otpString, `{"err"`) {
				return response,500, errors.New(otpString)
			} else {
				updateWidget = bson.M{
					"$set": bson.M{"timestamp": widget.Timestamp, "otptype": widget.OTPType, "artifactid": widget.ArtifactId, "otp": otpString},
				}
				rst2, err := widgetRepository.FindWidgetAndUpdate("widgetid", widget.WidgetId, updateWidget)
				if err != nil {
					return response,400, err
				}
				response = rst2
				return response,200, nil
			}
		} else {
			return response,400, errors.New("Invalied OTP Type")
		}
	} else {
		return response,400, errors.New("Invalied Widget Type")
	}
}

func FindWidgetAndUpdateQuery(widget requestDtos.RequestWidget) (models.Widget, error) {
	update := bson.M{
		"$set": bson.M{"query": widget.Query},
	}
	return widgetRepository.FindWidgetAndUpdate("widgetid", widget.WidgetId, update)
}

func FindWidgetByWidgetId(id string) (models.Widget, error) {
	return widgetRepository.FindWidgetOneById("widgetid", id)
}

func FindWidgetByWidgetIdWithOTP(id string) (models.Widget, error) {
	return widgetRepository.FindWidgetOneByIdWithOtp("widgetid", id)
}
