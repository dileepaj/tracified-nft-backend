package queryLanguageService

import (
	fcl "github.com/TharinduBalasooriya/fcl-go-lib"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

/**
*This method extract the data from OTP using  query , otp and defs file
*
**/
func QueryExecuter(query string, widget models.Widget,OTPJSON string) responseDtos.QueryResult {
	var result string = fcl.NewFCLWrapper().GetCommonJSON("services/queryLanguageService/Defs.txt", query, OTPJSON)
	queryResult := responseDtos.QueryResult{
		Result:   result,
		OTPType:  widget.OTPType,
		WidgetId: widget.WidgetId,
	}
	return queryResult
}