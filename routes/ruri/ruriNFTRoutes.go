package ruri

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var RuriNFtRoutes = models.Routers{
	/**
	 * ?Description : POST request, when RURI product ID and user email is providied OTP will be generated and sent as a email to customer
	 * ?relevent batchID for the RURI product ID will be recived and stored in the DB along with email and otp
	 * *Param : {email} --> Users email
	 * *reutrns : Object code of new OTP that was saved in DB
	 */
	models.Router{
		Name:    "Generate OTP",
		Method:  "POST",
		Path:    "/ruri/{productID}/{email}",
		Handler: apiHandler.RuriNewOTP,
	},
	/**
	 * ?Description : GET request, that will validate id the OTP is valid. Will check if the entered OTP and email exisit in the DB
	 * *URL Param : {email} --> Users email
	 * *URL Param : {otp} --> One time password that the user will enter after referring to the email that was sent to them
	 * *reutrns : Generated SVG if OTP is valid or an error message if the OTP is invalid.
	 */
	models.Router{
		Name:    "Validate OTP",
		Method:  "GET",
		Path:    "/ruri/{email}/{otp}/",
		Handler: apiHandler.ValidateOTP,
	},

	/**
	 * ?Description :PUT Request that will Re generate an OTP and send an email to the customer with the new OTP code
	 * *URL Param : {productID} --> Shopify RURI product ID
	 * *URL Param : {email} --> Users email
	 * *reutrns : Returns the BatchID if the OTP was geneerated Succesfully.
	 */
	models.Router{
		Name:    "Resend OTP",
		Method:  "PUT",
		Path:    "/ruri/{productID}/{email}/",
		Handler: apiHandler.ResentOTP,
	},
	//! TEsting methods remove after full impl
	models.Router{
		Name:    "GET TDP Data",
		Method:  "GET",
		Path:    "/ruri/{batchID}",
		Handler: apiHandler.SaveTDPDataByBatchID,
	},
	//! TEsting methods remove after full impl
	models.Router{
		Name:    "Generate SVG",
		Method:  "POST",
		Path:    "/ruri/getsvg/{batchID}/{email}",
		Handler: apiHandler.GenerateSVG,
	},
}
