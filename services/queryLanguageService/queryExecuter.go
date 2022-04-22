package queryLanguageService

import (
	Fcl "github.com/TharinduBalasooriya/fcl-go-lib"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

/*
function QueryExecuter() This method taks 3 parameters quering from the OTPJSON return the result  as sting ,It can be error or query result
	?Paramters
		query : query that user inputs :Type string
		widget : widget struct (this specific widget's query and Otp user for get the result) :Type string
		OtpJson OTP JSON :Type string string
	?return
			responseDtos.QueryResult struct
*/

func QueryExecuter(query string, widget models.Widget, OtpJson string) responseDtos.QueryResult {
	/*
		This method extract the data from OTP using  query , otp and defs file
		!FCL worked as a library in import and use it here
		FCL GetCommonJSON takes 3 parameters arg2 string, arg3 string, arg4 string
		?Parameters
			ars2 is a Defs file it is specific tho quey language :Type File
			arg3 is a FCL query :Type string
			arg4 is a OTP JSON :Type string string
		?return 
			query result :type string
		!! This FCL GetCommonJSON has a memory leaks
	*/
	var result string = Fcl.NewFCLWrapper().GetCommonJSON("services/queryLanguageService/Defs.txt", query, OtpJson)
	// queryResult is result struct this is response for query Both Staging and Batch otp
	queryResult := responseDtos.QueryResult{
		Result:   result,
		OTPType:  widget.OTPType,
		WidgetId: widget.WidgetId,
	}
	return queryResult
}

// func QueryExecuter(query string, widget models.Widget,OTPJSON string) responseDtos.QueryResult {
// 	//var result string = fcl.NewFCLWrapper().GetCommonJSON("services/queryLanguageService/Defs.txt", query, OTPJSON)
// 	queryResult := responseDtos.QueryResult{
// 		Result:   "FCL issue Memory leaks",
// 		OTPType:  widget.OTPType,
// 		WidgetId: widget.WidgetId,
// 	}
// 	return queryResult
// }
