package configs

import (
	"os"
)

var (
	endorsmentSenderEmailAddress = ""
	endorsmentSenderEmailKey     = ""
)

func GetEndrosmentSenderEmailAddres() string {
	LoadEnv()
	endorsmentSenderEmailAddress = os.Getenv("MK_SENDER_EMAILADDRESS")
	return endorsmentSenderEmailAddress
}

func GetEndorsmentSenderEmailKey() string {
	LoadEnv()
	endorsmentSenderEmailKey = os.Getenv("MK_SENDER_EMAIL_KEY")
	return endorsmentSenderEmailKey
}

func GetAcceptedEndorsmentEmail(name string, rating string, review string) string {
	var startRate string
	if rating == "1" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "1.5" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-half.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "2" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "2.5" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-half.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "3" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "3.5" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-half.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "4" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-empty.png">
        `
	} else if rating == "4.5" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-half.png">
        `
	} else if rating == "5" {
		startRate = `
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/star-full.png">
        `
	}

	var emailTemplate = `<table border="0" cellpadding="0" cellspacing="0" width="100%">
    <!-- LOGO -->
    <tr>
        <td style="background: radial-gradient(107.46% 107.46% at 92.81% 151.49%, #44795A 19.79%, rgba(0, 0, 0, 0) 100%),radial-gradient(74.55% 188.74% at -19.17% 100%, #44795A 19.79%, rgba(0, 0, 0, 0) 100%), #000000; height: 120px;"
            align="center">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                    <td align="center" valign="top" style="padding: 40px 10px 40px 10px;"> </td>
                </tr>
            </table>
        </td>
    </tr>

    <tr>
        <td bgcolor="#f4f4f4" align="center" style="padding: 0px 10px 0px 10px;">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                    <td style="padding: 20px 30px 15px 20px; color: #666666; font-family: Lato, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px; width: 100%;"
                        align="left" bgcolor="#ffffff">
                        <center> <img
                                src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/Tracified-NFT-v2.png"
                                style="width:290px;height: 62.37px;"></center>
                        <p
                            style="text-align: center;font-family: 'Inter';font-weight: 300;font-size: 29.162px;line-height: 0px;letter-spacing: 0.225em;">
                            MARKETPLACE</p>
                        <p style="margin: 0;"><strong>HI name, </strong><br />Tracified Marketplace team has
                            completed evaluating your endorsement and has decided to <strong>accept your
                                endorsement</strong> request. Congratulations!</P>
                    </td>
                <tr>
                    <td bgcolor="#ffffff" align="left">
                        <table width="100%" border="0" cellspacing="0" cellpadding="0">
                            <tr>
                                <td bgcolor="#ffffff" align="left" style="padding: 2px 30px 10px 30px;">
                                    <table border="0" cellspacing="0" cellpadding="0">
                                        <tr>
                                            <td bgcolor="#F1F1F1" align="center" style="border-radius: 3px;">
                                                <div
                                                    style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:center">
                                                    Rating</div>
                                            </td>
                                            <td bgcolor="#F1F1F1" align="left"
                                                style="border-radius: 3px;width:100%;">
                                                <div
                                                    style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:left">
                                                    ` + startRate + `
                                                </div>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td bgcolor="#F1F1F1" align="left" style="border-radius: 3px;">
                                                <div
                                                    style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:center">
                                                    Review</div>
                                            </td>
                                            <td bgcolor="#F1F1F1" align="left" style="border-radius: 3px;">
                                                <div
                                                    style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:left">
                                                    {{Review description}}</div>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                <tr>

                </tr>
                <tr>

                    <td bgcolor="#ffffff" align="left"
                        style="padding: 0px 30px 40px 30px; border-radius: 0px 0px 4px 4px; color: #666666; font-family: 'Lato', Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px;">
                        <span style="margin: 0;">Cheers,<br>Tracified Marketplace Team</span>
                    </td>
                </tr>
            </table>
            <table border="0px" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                    <td>
                        <table border="0" style="width: 100%;">
                            <tr>
                                <td style="text-align: center;"><a href="#">Visit marketplace<img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/external-link.png"></a></td>
                                <td style="text-align: left;"><a href="#">privacy policy<img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/external-link.png"></a></td>
                            </tr>
                            <tr>
                                <td style="text-align: center;" colspan="2">Manage your email preference <a href="#">here</a> or <a href="#">Unsubscribe</a> </td>
                            </tr>
                        </table>
                        <hr style="border:1px solid #CCCCCC;">
                        <p style="color: #878787;">
                            <center>Powered by</center>
                        </p>
                        <center> <img
                                src="https://s3.ap-south-1.amazonaws.com/tracified-image-storage/logos/TracifiedLogo_dark.png"
                                style="width:210px;height: 44px;"></center>
                    </td>
                </tr>
                <tr>
                    <td colspan="2" style="text-align: center;padding-top: 46px;">
                        <ul class="social-icons">                         
                            <li style="display: inline;padding:50px 10px 4px 10px;"><a href="https://www.facebook.com/tracified/?ref=page_internal" target="_blank"><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Facebook.png"></a>
                            <li style="display: inline;padding:20px 10px 4px 10px;"><a href="https://twitter.com/Tracified1" target="_blank"><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Twitter.png"></a>
                            <li style="display: inline;padding:20px 10px 4px 10px;"><a href="https://www.instagram.com/tracified_official/" target="_blank"><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Instagram.png"></a>
                            <li style="display: inline;padding:20px 10px 4px 10px;"><a href="https://www.instagram.com/tracified_official/" target="_blank"><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/LinkedIn.png"></a>
                            <li style="display: inline;padding:20px 10px 4px 10px;"><a href="#"><img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/YouTube.png"></a>
                        </ul>
                    </td>
                </tr>
                <tr>
                    <th colspan="2" style="text-align: center; text-decoration: none;padding-top: 15px;">Stay connected!</th>
                </tr>
            </table>
        </td>
    </tr>
</table>`
	return emailTemplate
}

func GetDeclinedEndorsmentEmail(name string, rating string, review string) string {
	var emailTemplate = `<table border="0" cellpadding="0" cellspacing="0" width="100%">
    <!-- LOGO -->
    <tr>
        <td bgcolor="#00C820 " align="center">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
                <tr>
                    <td align="center" valign="top" style="padding: 40px 10px 40px 10px;"> </td>
                </tr>
            </table>
        </td>
    </tr>

    <tr>
        <td bgcolor="#f4f4f4" align="center" style="padding: 0px 10px 0px 10px;">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
              <tr>
      <td style="padding: 20px 30px 15px 20px; color: #666666; font-family: Lato, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px; width: 100%;" align="left" bgcolor="#ffffff">
      <p style="margin: 0;">Hi ` + name + `<br />Tracified Marketplace team has completed Evaluating your endorsment and has decided to decline your endrosment request as it has not met with tracified marketplace requirments.Please feeel free to contact Tracified team for further clarifications.</p>
      <P>Please Refer below for the feedback given on your endorsment.<a href="#">Visit marketplace</a></P>
      </td>
      <tr>
                    <td bgcolor="#ffffff" align="left">
                        <table width="100%" border="0" cellspacing="0" cellpadding="0">
                            <tr>
                                <td bgcolor="#ffffff" align="center" style="padding: 2px 30px 10px 30px;">
                                    <table border="0" cellspacing="0" cellpadding="0">
                                        <tr>
                                            <td align="center" style="border-radius: 3px;"><div style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:center">Rating</div></td>
                                            <td align="center" style="border-radius: 3px;"><div style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:left">` + rating + `</div></td>
                                        </tr>
                                        <tr>
                                            <td align="center" style="border-radius: 3px;"><div style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:center">Review</div></td>
                                            <td align="center" style="border-radius: 3px;"><div style="font-size: 12px; font-family: Helvetica, Arial, sans-serif;  text-decoration: none; text-decoration: none; padding: 15px 25px; border-radius: 2px;  display: inline-block; text-align:left">` + review + `</div></td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </td>
                <tr>
                    
                </tr>
                <tr>
                    
                    <td bgcolor="#ffffff" align="left" style="padding: 0px 30px 40px 30px; border-radius: 0px 0px 4px 4px; color: #666666; font-family: 'Lato', Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 25px;">
                        <p style="margin: 0;">Cheers,<br>Tracified Marketplace Team</p>
                        <hr style="background-color: #D9D9D9; ">
                        <p style="color: #878787;"><center>Powered by</center></p>
                        <center> <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/Tracified-NFT-v2.png" style="width:20em"></center>
                    </td>
                </tr>
            </table>
        </td>
    </tr>        
  </table>`
	return emailTemplate
}
