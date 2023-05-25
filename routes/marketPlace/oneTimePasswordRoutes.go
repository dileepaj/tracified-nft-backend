package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var OneTimePassWordRoutes = models.Routers{
	/**
	 * ?Description : POST request, when RURI product ID and user email is provided OTP will be generated and sent as a email to customer
	 * ?relevant batchID for the product ID will be received and stored in the DB along with email and otp
	 * *Param : {email} --> Users email
	 * *returns : Object code of new OTP that was saved in DB
	 */
	models.Router{
		Name:    "Initialize OTP", //! Initialize NFT
		Method:  "POST",
		Path:    "/otpgen",
		Handler: apiHandler.InitNFT,
	},
	/**
	 * ?Description : GET request, that will validate id the OTP is valid. Will check if the entered OTP and email exisit in the DB
	 * *URL Param : {email} --> Users email
	 * *URL Param : {otp} --> One time password that the user will enter after referring to the email that was sent to them
	 * *reutrns : Generated SVG if OTP is valid or an error message if the OTP is invalid.
	 */
	models.Router{
		Name:    "Validate OTP",
		Method:  "POST",          //! Convert to post method
		Path:    "/validateOTP/", //? /validateOTP/..
		Handler: apiHandler.ValidateOTP,
	},

	/**
	 * ?Description :PUT Request that will Re generate an OTP and send an email to the customer with the new OTP code
	 * *URL Param : {productID} --> Shopify product ID
	 * *URL Param : {email} --> Users email
	 * *returns : Returns the BatchID if the OTP was generated Successfully.
	 */
	models.Router{
		Name:    "Resend OTP",
		Method:  "PUT",
		Path:    "/resendOTP",
		Handler: apiHandler.ResentOTP,
	},
}
