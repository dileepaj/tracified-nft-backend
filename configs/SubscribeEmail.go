package configs

import (
	"os"
)

var (
	subscribeSenderEmailAddress = ""
	subscribeSenderEmailKey     = ""
)

func GetSubscribeSenderEmailAddres() string {
	LoadEnv()
	subscribeSenderEmailAddress = os.Getenv("MK_SENDER_EMAILADDRESS")
	return subscribeSenderEmailAddress
}

func GetSubscribeSenderEmailKey() string {
	LoadEnv()
	subscribeSenderEmailKey = os.Getenv("MK_SENDER_EMAIL_KEY")
	return subscribeSenderEmailKey
}

func GetAcceptedSubscription() string {
	var emailTemplate = `<table border="0" cellpadding="0" cellspacing="0" width="100%">
  <!-- LOGO -->
  <tr>
      <td style="background: black" align="center">
          <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
              <tr>
                  <td align="center" valign="top" style="padding: 40px 10px 40px 10px;"> </td>
              </tr>
          </table>
      </td>
  </tr>
  <tr>
      <td style="background: linear-gradient(black, #44795A)"
          align="center" style="padding: 0px 10px 0px 10px;">
          <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
              <tr>
                  <td style="padding: 40px 20px 20px; border-radius: 20px 20px 0px 0px; color: #111111; font-family: Lato, Helvetica, Arial, sans-serif; font-size: 40px; font-weight: 400; letter-spacing: 4px; line-height: 48px; height: 31px;"
                      align="center" valign="top" bgcolor="#ffffff">
                      <img style="width: 290x; height: 97px; margin-top: 20px"
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Tracified-NFT-Marketplace-D.png">
                  </td>
              </tr>
          </table>
      </td>
  </tr>
  <tr >
      <td bgcolor="#f4f4f4" align="center" style="padding: 0px 10px 0px 10px;">
          <table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
              <tr>
                  <td align="left" bgcolor="#ffffff">
                      <p style="@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
        font-family: 'Inter', sans-serif;
        margin: 0px 30px 20px 30px;
        text-align: left;
        line-height: 30px;
        font-style: normal;
        font-weight: 400;
        font-size: 15px;"><strong>Hi User,</strong><br /><br />
		You have successfully subscribed with our NFT Marketplace News Letter. As a subscriber, now you will get to enjoy our latest updates first hand, early on.
		 <br /><br />
		 Stay tuned!
	   </p>
                          
                  </td>
              <tr>
                  <td  style="border-radius: 20px 20px 20px 20px;" bgcolor="#ffffff" align="left">
                      <div
      style="
        @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
        font-family: 'Inter', sans-serif;
        text-align: left;
        font-style: normal;
        font-weight: 400;
        font-size: 15px;
        line-height: 25px;
        margin-left: 30px;
        margin-bottom: 40px;
      "
    >
      <p>Cheers, <br /><strong>Tracified Marketplace Team</strong></p>
    </div>
  </td>
  </tr>
  <tr>

</table>
<table border="0px" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px; margin-top: 30px;">
  <tr>
      <td>
          <center>
              <table border="0" style="width: 360px">
                  <tr>
                      <td style="text-align: center;">
                          <p
                          style="
                            @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                            font-family: 'Inter', sans-serif;
                            color: #878787 !important;
                            font-style: normal;
                            font-weight: 400;
                            font-size: 13px;
                          "
                        >
                          <a
                            href="#"
                            style="color: #878787 !important; text-decoration: underline"
                            >Visit Marketplace</a
                          >
                          <a href="#" target="_blank"
                            ><img
                              src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/external-link.png"
                              height="10px"
                              width="10px"
                          /></a>
                        </p>
                      </td>
                      <td style="text-align: left;">
                          <p
                          style="
                            @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                            font-family: 'Inter', sans-serif;
                            color: #878787 !important;
                            font-style: normal;
                            font-weight: 400;
                            font-size: 13px;
                          "
                        >
                          <a href="#" style="color: #878787 !important">Privacy Policy</a>
                          <a href="#" target="_blank"
                            ><img
                              src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/external-link.png"
                              height="10px"
                              width="10px"
                          /></a>
                        </p>
                      </td>
                  </tr>
                  <tr>
                      <td style="text-align: center; padding-right: 25px;" colspan="2">
                              <p style="
                              margin-top: 10px;
                              margin-bottom: 10px;
                              @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                              font-family: 'Inter', sans-serif;
                              color: #878787 !important;
                              font-style: normal;
                              font-weight: 400;
                              font-size: 13px;
                              ">
                                                  Manage your email preferences
                                                  <a href="#" style="
                                  @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                  font-family: 'Inter', sans-serif;
                                  color: #878787 !important;
                                  font-style: normal;
                                  font-weight: 400;
                                  font-size: 13px;
                              ">here</a>
                                                  or
                                                  <a href="#" style="
                                  @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                  font-family: 'Inter', sans-serif;
                                  color: #878787 !important;
                                  font-style: normal;
                                  font-weight: 400;
                                  font-size: 13px;
                              ">unsubscribe.</a>
                          </p>
                      </td>
                  </tr>
              </table>
          </center>
          
          <hr style="width: 600px; border-top: 1px solid #cccccc; border-bottom: 0px">
          <p style="color: #878787;">
              <center>
                  <p style="
        @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
        font-family: 'Inter', sans-serif;
        font-style: normal;
        font-weight: 400;
        font-size: 14px;
        line-height: 22px;
        color: #000000;
      ">
                      Powered by
                  </p> <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/tracified-logo+(1).png"
              style="width:210px;height: 44px; margin-top: -14px; margin-right: 6px;">
              </center>
          </p>
      </td>
  </tr>
  <tr>
      <td colspan="2" style="text-align: center;padding-top: 20px; width: 270px; padding-right: 20px;">
          <ul style="display:inline;">
              <li style="display: inline;padding:20px 20px 4px 5px;"><a
                      href="https://www.facebook.com/tracified/?ref=page_internal" target="_blank"><img
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Facebook.png" /></a>
              </li>
              <li style="display: inline;padding:20px 20px 4px 5px;"><a href="https://twitter.com/Tracified1"
                      target="_blank"><img
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Twitter.png" /></a>
              </li>
              <li style="display: inline;padding:20px 20px 4px 5px;"><a
                      href="https://www.instagram.com/tracified_official/" target="_blank"><img
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/Instagram.png" /></a>
              </li>
              <li style="display: inline;padding:20px 20px 4px 5px;"><a
                      href="https://www.instagram.com/tracified_official/" target="_blank"><img
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/LinkedIn.png" /></a>
              </li>
              <li style="display: inline;padding:20px 20px 4px 5px;"><a href="https://www.youtube.com/channel/UCavZhMjJ1jw5DDmKXNyiXmA" target="_blank"><img
                          src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/NFT_Market/YouTube.png" /></a>
              </li>
          </ul>
      </td>
  </tr>
  <tr>
      <th colspan="2" style="
      @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
      font-family: 'Inter', sans-serif;
      text-align: center;
      font-weight: 400;
      text-decoration: none;
      margin-top: 2px;
      margin-right: 2px;
      font-size: 14px;
    ">Stay connected!
      </th>
  </tr>
</table>
</td>
</tr>
</table>`
	return emailTemplate
}
