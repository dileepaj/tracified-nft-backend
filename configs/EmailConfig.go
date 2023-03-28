package configs

import (
	"os"
	"strconv"
)

var (
	senderEmailAddress = ""
	senderEmailAppPwd  = ""
	emailHost          = ""
	emailPort          = 0
	bearerToken        = ""
)

func GetSenderEmailAddress() string {
	LoadEnv()
	senderEmailAddress = os.Getenv("SENDER_EMAILADRESS")
	return senderEmailAddress
}
func GetSenderEmailAppPWD() string {
	LoadEnv()
	senderEmailAppPwd = os.Getenv("SENDER_EMAILADRESS_APPPWD")
	return senderEmailAppPwd
}
func GetEmailHost() string {
	LoadEnv()
	emailHost = os.Getenv("HOST_EMAIL")
	return emailHost
}
func GetEmailPort() int {
	LoadEnv()
	emailPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		return 587
	}
	return emailPort
}

func GetBearerToken() string {
	LoadEnv()
	bearerToken = os.Getenv("BACKEND_TOKEN")
	return bearerToken
}

func GetShopify() string {
	LoadEnv()
	bearerToken = os.Getenv("RURI_SHOPIFY")
	return bearerToken
}

func GetDigitalTwin() string {
	LoadEnv()
	bearerToken = os.Getenv("DIGITALTWIN")
	return bearerToken
}

func GetEmail(otp string) string {
	var emailTemplate = `<table border="0" cellpadding="0" cellspacing="0" width="100%">
    <!-- LOGO -->
    <tr>
        <td bgcolor="#021d28" align="center">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                    <td align="center" valign="top" style="padding: 40px 10px 40px 10px;"> </td>
                </tr>
            </table>
        </td>
    </tr>
    <tr>
        <td bgcolor="#021d28" align="center" style="padding: 0px 10px 0px 10px;">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                   <td style="padding: 40px 20px 20px; border-radius: 4px 4px 0px 0px; color: #111111; font-family: Lato, Helvetica, Arial, sans-serif; font-size: 40px; font-weight: 400; letter-spacing: 4px; line-height: 48px; height: 31px;" align="center" valign="top" bgcolor="#ffffff"><img style="width: 3em;" src="https://tracified-profile-images.s3.ap-south-1.amazonaws.com/RURI+1.png"></td>
                </tr>
            </table>
        </td>
    </tr>
    <tr>
        <td bgcolor="#f4f4f4" align="center" style="padding: 0px 10px 0px 10px;">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
              <tr>
      <td style="padding: 20px 30px 15px 20px; color: #666666; font-family: Lato, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px; width: 100%;" align="left" bgcolor="#ffffff">
      <p style="margin: 0;">Hi<br />Thank you for choosing RURI. Please use the following One Time Password (OTP) to complete the sign up process.</p>
      </td>
      <tr>
                    <td bgcolor="#ffffff" align="left">
                        <table width="100%" border="0" cellspacing="0" cellpadding="0">
                            <tr>
                                <td bgcolor="#ffffff" align="center" style="padding: 2px 30px 10px 30px;">
                                    <table border="0" cellspacing="0" cellpadding="0">
                                        <tr>
                                            <td align="center" style="border-radius: 3px;" bgcolor="#00466a"><div style="font-size: 20px; font-family: Helvetica, Arial, sans-serif; color: #62FFA3; text-decoration: none; color: #62FFA3; text-decoration: none; padding: 15px 25px; border-radius: 2px; border: 1px solid #00466a; display: inline-block; text-align:center">` + otp + `</a></td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                <tr>
                    
                </tr>
        <td bgcolor="#ffffff" align="left" style="padding: 0px 30px 10px 30px; color: #666666; font-family: 'Lato', Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px;">
          <p style="margin: 0;"> <strong>Note - </strong>Please note that the One Time Password is valid for a period of one month only.
                    </td>						
      </tr>
                <tr>
        <td bgcolor="#ffffff" align="left" style="padding: 0px 30px 10px 30px; color: #666666; font-family: 'Lato', Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px;">
          <p style="margin: 0;">Enjoy your NFT !
                    </td>						
      </tr>
                <tr>
                    <td bgcolor="#ffffff" align="left" style="padding: 0px 30px 40px 30px; border-radius: 0px 0px 4px 4px; color: #666666; font-family: 'Lato', Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px;">
                        <p style="margin: 0;">Cheers,<br>Team RURI</p>
                        <hr style="background-color: #D9D9D9; ">
                        <p style="color: #878787;"><center>Powered by</center></p>
                        <center><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/Tracified-NFT-v2.png" style="width:20em"></center>
                    </td>
                </tr>
            </table>
        </td>
    </tr>        
  </table>`
	return emailTemplate
}
