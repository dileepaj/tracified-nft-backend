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
        font-size: 15px;"><strong>Hi ` + name + `, </strong><br /><br/>Tracified Marketplace team has
                          completed evaluating your endorsement and has  <strong> decided to accept your
                              endorsement request.</strong> Congratulations!
                              <br/><br/>
                              Please refer below for the feedback given on your endorsement.
                          </p>
                          
                  </td>
              <tr>
                  <td  style="border-radius: 20px 20px 20px 20px;" bgcolor="#ffffff" align="left">
                      <table width="100%" border="0" cellspacing="0" cellpadding="0">
                          <tr>
                              <td bgcolor="#ffffff" align="left" style="border-radius: 0px 0px 20px 20px;">
                                  <table style="text-align: left;
                                  margin: 0px 20px 0px 20px;
                                  border: 1px solid white;
                                  border-collapse: collapse;" border="0" cellspacing="0" cellpadding="0">
                                      <tr>
                                          <td bgcolor="#F1F1F1" align="center" style="border-radius: 3px; @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                          font-family: 'Inter', sans-serif;
                                          border: 10px solid white;
                                          border-collapse: collapse;
                                          padding: 20px 30px 20px 15px;
                                          background: #f1f1f1;
                                          font-style: normal;
                                          font-weight: 700;
                                          font-size: 15px;
                                          line-height: 24px; ">
                                              <div>
                                                  Rating</div>
                                          </td>
                                          <td bgcolor="#F1F1F1" align="left" style=" @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                          font-family: 'Inter', sans-serif;
                                          border: 10px solid white;
                                          border-collapse: collapse;
                                          padding: 20px 30px 20px 15px;
                                          background: #f1f1f1;
                                          font-style: normal;
                                          font-weight: 200;
                                          font-size: 15px;
                                          line-height: 24px; border-radius: 3px;width:100%;">
                                              <div>
                                                  ` + startRate + `
                                              </div>
                                          </td>
                                      </tr>
                                      <tr>
                                          <td bgcolor="#F1F1F1" align="left" style="
                                               @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                          font-family: 'Inter', sans-serif;
                                          border: 10px solid white;
                                          border-collapse: collapse;
                                          padding: 20px 30px 20px 15px;
                                          background: #f1f1f1;
                                          font-style: normal;
                                          font-weight: 700;
                                          font-size: 15px;
                                          line-height: 24px; border-radius: 3px;">
                                              <div>Review</div>
                                          </td>
                                          <td bgcolor="#F1F1F1" align="left" style="
                                           @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                          font-family: 'Inter', sans-serif;
                                          border: 10px solid white;
                                          border-collapse: collapse;
                                          padding: 20px 30px 20px 15px;
                                          background: #f1f1f1;
                                          font-style: normal;
                                          font-weight: 200;
                                          font-size: 15px;
                                          line-height: 24px; border-radius: 3px;">
                                              <div>` + review + `</div>
                                          </td>
                                      </tr>
                                  </table>
                              </td>
                          </tr>
                      </table>
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

func GetDeclinedEndorsmentEmail(name string, rating string, review string) string {
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
          font-size: 15px;"><strong>Hi ` + name + `, </strong><br/><br/>
                                Tracified Marketplace team has completed evaluating your endorsment and had <strong> decided to decline your endorsment request as it has not met the tracified marketplace requirments</strong>. Please feel free to <a href="#">contact Tracified team </a> for further clarifications
                                <br><br>
                                Please refer below for the feedback given on your endorsement.
                            </p>
                            
                    </td>
                <tr>
                    <td  style="border-radius: 20px 20px 20px 20px;" bgcolor="#ffffff" align="left">
                        <table width="100%" border="0" cellspacing="0" cellpadding="0">
                            <tr>
                                <td bgcolor="#ffffff" align="left" style="border-radius: 0px 0px 20px 20px;">
                                    <table style="text-align: left;
                                    margin: 0px 20px 0px 20px;
                                    border: 1px solid white;
                                    border-collapse: collapse;" border="0" cellspacing="0" cellpadding="0">
                                        <tr>
                                            <td bgcolor="#F1F1F1" align="center" style="border-radius: 3px; @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                            font-family: 'Inter', sans-serif;
                                            border: 10px solid white;
                                            border-collapse: collapse;
                                            padding: 20px 30px 20px 15px;
                                            background: #f1f1f1;
                                            font-style: normal;
                                            font-weight: 700;
                                            font-size: 15px;
                                            line-height: 24px; ">
                                                <div>
                                                    Rating</div>
                                            </td>
                                            <td bgcolor="#F1F1F1" align="left" style=" @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                            font-family: 'Inter', sans-serif;
                                            border: 10px solid white;
                                            border-collapse: collapse;
                                            padding: 20px 30px 20px 15px;
                                            background: #f1f1f1;
                                            font-style: normal;
                                            font-weight: 200;
                                            font-size: 15px;
                                            line-height: 24px; border-radius: 3px;width:100%;">
                                                <div>
                                                    ` + startRate + `
                                                </div>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td bgcolor="#F1F1F1" align="left" style="
                                                 @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                            font-family: 'Inter', sans-serif;
                                            border: 10px solid white;
                                            border-collapse: collapse;
                                            padding: 20px 30px 20px 15px;
                                            background: #f1f1f1;
                                            font-style: normal;
                                            font-weight: 700;
                                            font-size: 15px;
                                            line-height: 24px; border-radius: 3px;">
                                                <div>Review</div>
                                            </td>
                                            <td bgcolor="#F1F1F1" align="left" style="
                                             @import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
                                            font-family: 'Inter', sans-serif;
                                            border: 10px solid white;
                                            border-collapse: collapse;
                                            padding: 20px 30px 20px 15px;
                                            background: #f1f1f1;
                                            font-style: normal;
                                            font-weight: 200;
                                            font-size: 15px;
                                            line-height: 24px; border-radius: 3px;">
                                                <div>` + review + `</div>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
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
