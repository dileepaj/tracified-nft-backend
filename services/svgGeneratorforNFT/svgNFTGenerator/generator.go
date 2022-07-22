package svgNFTGenerator

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	svgStart   = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTHeader.txt")
	svgEnd     = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTFooter.txt")
	styling    = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTStyles.css") //!Need to implement
	styleStart = `<style>`
	styleEnd   = `</style>`
)

func GenerateSVGTemplateforNFT(tdpData []models.TDP, batchID string) (string, error) {
	logs.InfoLogger.Println("starting svg gen")
	var htmlBody string
	var htmlStart = `<h1 class="text-center">NFT name : ` + batchID + `</h1>
							 <p class="text-center fw-bold text-muted">` + batchID + `</p>
							 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
						`

	for _, maindata := range tdpData {
		for _, data := range maindata.TraceabilityData {
			//* type 1 = numbers
			if data.Type == 1 {
				// s := fmt.Sprintf("%f", data.Val)
				// htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
				// 				<p>` + data.Key + ` : ` + s + `</p>
				// 			</div>`

				/**
				 **Type 5 : string
				 **Type 3 : Date/Time
				 **Type 7 : List Value
				 */
			} else if data.Type == 5 || data.Type == 3 || data.Type == 7 {
				// htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
				// 				<p>` + data.Key + ` : ` + data.Val.(string) + `</p>
				// 			</div>`
				//* type 6 = Artifavt data
			} else if data.Type == 6 {
				var artifactData map[string]interface{} = data.Val.(map[string]interface{})
				var tableContent string = `<p class="common-widget-title">` + "TenantID:" + maindata.TenantID + `</p>
				<table class="table table-bordered"><thead>
					<tr>
						<th><strong>name</strong></th>
						<th><strong>Description</strong></th>
					</tr>
				`
				for key, itmdata := range artifactData {
					tableContent +=
						`
						<tr>
							<td>` + key + `</td>
							<td>` + itmdata.(string) + `</td>
						</tr>
					`
					//htmlBody += `<p> Artifiact Data: ` + string(itmdata.(string)) + `</p><br>`
				}
				tableContent += "</thead></table>"
				htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
								<p>` + tableContent + `</p>
							</div>`
			}
		}
	}
	template := svgStart + styleStart + styling + styleEnd + htmlStart + htmlBody + svgEnd
	// logs.InfoLogger.Println("--------------generated template start--------------")
	//logs.InfoLogger.Println("\n", template)
	// logs.InfoLogger.Println("--------------generated template end--------------")
	return template, nil
}
