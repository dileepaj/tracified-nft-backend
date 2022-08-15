package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var SvgGenerator = models.Routers{

	/**
	 *?Description : Updates the userSvgMapping collection by adding a SVG hash
	 **@params : passed in request body(object ID and hash -> string)
	 **reutrns : returns SVG
	 */
	models.Router{
		Name:    "Update SVG Hash",
		Method:  "PUT",
		Path:    "/svg/",
		Handler: apiHandler.UpdateSVGUserMappingbySha256,
	},

	/**
	 *?Description : when called reutrns the SVG based on the hash providied
	 **@params : {hash} : SVG hash
	 **reutrns : returns SVG
	 */
	models.Router{
		Name:    "Get SVG by Hash",
		Method:  "GET",
		Path:    "/svg/{hash}",
		Handler: apiHandler.GetSVGbySha256,
	},

	//! TEsting methods remove after full impl
	models.Router{
		Name:    "GET TDP Data",
		Method:  "GET",
		Path:    "/svg/{batchID}",
		Handler: apiHandler.SaveTDPDataByBatchID,
	},
	//! TEsting methods remove after full impl
	models.Router{
		Name:    "Generate SVG",
		Method:  "POST",
		Path:    "/svg/getsvg/{batchID}/{email}",
		Handler: apiHandler.GenerateSVG,
	},
}
