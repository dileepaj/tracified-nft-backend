package svgNFTGenerator

import (
	"fmt"
	"strings"

	customizedNFTrepository "github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/mitchellh/mapstructure"
)

var (
	svgStart       = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTHeader.txt")
	svgEnd         = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTFooter.txt")
	styling        = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTStyles.css") //!Need to implement
	styleStart     = `<style>`
	styleEnd       = `</style>`
	ruriRepository customizedNFTrepository.SvgRepository
)

func GenerateSVGTemplateforNFT(tdpData []models.TDP, batchID string) (string, error) {
	var htmlBody string
	var htmlStart = `<h1 class="text-center">NFT name : ` + batchID + `</h1>
							 <p class="text-center fw-bold text-muted">` + batchID + `</p>
							 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
						`

	var iframeImg = `<iframe  src="https://tracified.sirv.com/Spins/RURI%20Gems%20Compressed/120614/120614.spin" style="position: absolute; top: 0; left: 0; width: 100%; height: "300px" frameborder="0" allowfullscreen="true"></iframe>`
	for _, maindata := range tdpData {
		numbersArr := make(map[string]string)
		textArr := make(map[string]string)
		for _, v := range maindata.TraceabilityData {
			if v.Type == 1 {
				s := fmt.Sprintf("%f", v.Val)
				numbersArr[v.Key] = s
				/**
				 **Type 5 : string
				 **Type 3 : Date/Time
				 **Type 7 : List Value
				 */
			} else if v.Type == 5 /*|| v.Type == 3*/ || v.Type == 7 {
				textArr[v.Key] = v.Val.(string)
				//* type 6 = Artifact data
			} else if v.Type == 6 {
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tableContent string = `<p class="common-widget-title">` + "Artifact Data" + `</p>
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
							<td>` + key + `</td>`
					if strings.Contains(itmdata.(string), ".jpg") || strings.Contains(itmdata.(string), ".png") || strings.Contains(itmdata.(string), ".jpeg") {
						tableContent += `<td><img src="` + itmdata.(string) + `"/></td>`
					} else {
						tableContent += `<td>` + itmdata.(string) + `</td>`
					}
					tableContent += `</tr>`
				}
				tableContent += "</thead></table>"
				htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
								<p>` + tableContent + `</p>
							</div>`
			} else if v.Type == 4 {
				var geoDataSet []interface{} = v.Val.([]interface{})
				for _, val := range geoDataSet {
					mapdata := val.(map[string]interface{})
					var tempdata models.GeoImageData
					mapstructure.Decode(mapdata, &tempdata)
					lat := fmt.Sprintf("%f", tempdata.GeoCode.Lat)
					long := fmt.Sprintf("%f", tempdata.GeoCode.Long)
					htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
								<p>` + "description" + ` : ` + tempdata.Description + `</p>
								<p><img src="` + tempdata.Image + `"/></p>
								<ul>
									<li>` + "latitude" + ` : ` + lat + `</li>
									<li>` + "Longitude" + ` : ` + long + `</li>
									<li><a href="https://maps.google.com/?q=` + lat + `,` + long + `">View on map</a></li>
								</ul>
								<p>Time Stamp : ` + tempdata.TimeStamp.Time().String() + `</p>
							</div>`

				}
			}

		}
		//? use for loop to print number and text data
		if numbersArr != nil || len(numbersArr) > 0 {
			start := `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">`
			var content string
			for key, val1 := range numbersArr {
				if key != "" && val1 != "" {
					content += `<p>` + key + ` : ` + val1 + `</p>`
				}
			}
			end := "</div>"
			htmlBody += start + content + end
		}
		if textArr != nil || len(textArr) > 0 {
			var content string
			start := `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">`
			for key1, val2 := range textArr {
				if key1 != "" && val2 != "" {
					content += `<p>` + key1 + ` : ` + val2 + `</p>`
				}
			}
			end := "</div>"
			htmlBody += start + content + end
		}
	}
	template := svgStart + styleStart + styling + styleEnd + htmlStart + iframeImg + htmlBody + svgEnd
	return template, nil
}
