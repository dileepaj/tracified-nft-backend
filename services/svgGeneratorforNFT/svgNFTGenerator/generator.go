package svgNFTGenerator

import (
	"fmt"
	"strconv"

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
	htmlBody	   = ""
	collectionName = ""
	ruriRepository customizedNFTrepository.SvgRepository
)

func GenerateSVGTemplateforNFT(tdpData []models.TDP, batchID string) (string, error) {
	var htmlStart = `<div class="nft-header default-font">
						<div class="nft-header-content">
							<img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/Tracified-NFT-v5.png" class="nft-logo"/>
							<label>` + batchID + `</label>
						</div>
					</div>
					<div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	var iframeImg = `<div class="iframe-wrapper"><iframe  src="https://tracified.sirv.com/Spins/RURI%20Gems%20Compressed/120614/120614.spin" class="iframe-img" frameborder="0" allowfullscreen="true"></iframe></div>`

	var stageStatus map[string]bool = make( map[string]bool)

	for _, maindata := range tdpData {

		if stageStatus[maindata.StageID] == false {

			if maindata.StageID == "100" {

				imgs := GetGemImages(tdpData)
				GetCollectionName(tdpData)
				GetGemDetails(maindata.TraceabilityData, "Gem Details", imgs)

			} else if maindata.StageID == "103" {

				GenerateContent(maindata.StageID, maindata.TraceabilityData, "Collector's/ Dealer's Information", "profile")

			} else if maindata.StageID == "105" {

				GenerateContent(maindata.StageID, maindata.TraceabilityData, "Certifications", "doc")

			} else if maindata.StageID == "106" {

				GenerateContent(maindata.StageID, maindata.TraceabilityData, "Export", "export")

			} else if maindata.StageID == "109" {

				GenerateContent(maindata.StageID, maindata.TraceabilityData, "Appraisal", "doc")

			}

			stageStatus[maindata.StageID] = true
		}
		
	}	
	template := svgStart + styleStart + styling + styleEnd + htmlStart + iframeImg + htmlBody + svgEnd
	return template, nil
}

//get gem images from stage 103
func GetGemImages (tdpData []models.TDP) []interface{}{
	var imgArr []interface{}
	for _, maindata := range tdpData {
		if maindata.StageID == "103" {
			for _, v := range maindata.TraceabilityData {
				if v.Type == 4 && v.Key == "photoofGem"{
					
					imgArr = v.Val.([]interface{})
					break;
					
				}
			}
		}
	}

	return imgArr
}

//get collection name from stage 103
func GetCollectionName (tdpData []models.TDP) {
	for _, maindata := range tdpData {
		if maindata.StageID == "103" {
			for _, v := range maindata.TraceabilityData {
				if v.Type == 6 && v.Key == "collectionname"{
					var artifactData map[string]interface{} = v.Val.(map[string]interface{})
					for key, itmdata := range artifactData {
						if key == "name" {
							collectionName = itmdata.(string)
							break;
						}
					}
				}
			}
		}
	}
}

//get and show gem details in nft
func GetGemDetails(tdp []models.TraceabilityData, section string, imgArr []interface{}) {
	heading := `<div class="widget-div"><div class="widget-title-div"><span class="gem-icon"></span><label>` + section + `</label></div></div>`
	tableContent := `<div class="bdr"><table class="table table-bordered table-striped rounded-20 overflow-hidden">`

	tableContent += `<tbody><tr><td>Gem Photos</td><td class="td-carousel"><div class="carousel-wrapper"><section class="carousel" aria-label="Gallery">
	<ol class="carousel__viewport">`
	len := len(imgArr)
	i := 0

	//displaying image carousel
	for _, v := range imgArr {
		mapdata := v.(map[string]interface{})
		var tempdata models.GeoImageData
		mapstructure.Decode(mapdata, &tempdata)

		curSlide := i + 1
		prev := curSlide - 1
		next := curSlide + 1

		if i == 0 {
			prev = len
		} else if i == len -1 {
			next = 1
		}

		lat := fmt.Sprintf("%f", tempdata.GeoCode.Lat)
		long := fmt.Sprintf("%f", tempdata.GeoCode.Long)
		
		tableContent += `<li id="carousel__slide` + strconv.Itoa(curSlide) +`" tabindex="0" class="carousel__slide"><div class="carousel__snapper">`
		tableContent += `<img src="` + tempdata.Image + `" /><a href="#carousel__slide`+strconv.Itoa(prev)+`" class="carousel__prev">Go to previous slide</a>
						<a href="#carousel__slide`+strconv.Itoa(next)+`" class="carousel__next">Go to next slide</a></div>`
		tableContent += `<div class="map-link-div"><a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span></div>
		<div class="timestamp-div"><label class="timestamp-label">Time Stamp : `+tempdata.TimeStamp.Time().String()+`</label></div></li>`
		i++

	}

	tableContent += "</ol></section></div></td></tr>"

	//displaying gem details
	for _, v := range tdp {
		if v.Type == 6 {
			var artifactData map[string]interface{} = v.Val.(map[string]interface{})
			for key, itmdata := range artifactData {
				if key == "gemType" {
					tableContent += `<tr><td>Gem Type</td><td class="value-label">` + itmdata.(string) + `</td></tr>`
					break;
				}
			}	
		} else if v.Type == 1 {
			s := fmt.Sprintf("%f", v.Val)
			tableContent += `<tr><td>Rough Stone Weight</td><td class="value-label">` + s + `</td></tr>`
		}
	}

	tableContent += `<tr><td>Collection Name</td><td class="value-label">` + collectionName + `</td></tr>`

	tableContent += `</tbody></table></div>`
	htmlBody += heading + tableContent
	
}

func GenerateContent(stageID string, tdp []models.TraceabilityData, section string, icon string) {

	heading := `<div class="widget-div"><div class="widget-title-div"><span class="`+icon+`-icon"></span><label>` + section + `</label></div></div>`
	tableContent := `<div class="bdr"><table class="table table-bordered table-striped rounded-20 overflow-hidden"><tbody>`

	if stageID == "103" {
		//displaying collector's/ dealer's information
		for _, v := range tdp {
			if  v.Type == 6 && v.Key == "collector/dealername"{
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tempdata models.CollectorInfo
				mapstructure.Decode(artifactData, &tempdata)

				tableContent += `<tr><td>Collector's/Dealer's Photo </td><td><img src="` + tempdata.Photo + `" class="photo" /></td></tr>`
				tableContent += `<tr><td>Collector's/Dealer's Name</td><td class="value-label">` + tempdata.Name + `</td></tr>`
				tableContent += `<tr><td>Collector's/Dealer's Address</td><td class="value-label">` + tempdata.Address + `</td></tr>`
				tableContent += `<tr><td>Collector's/Dealer's Contact Number</td><td class="value-label">` + tempdata.ContactNumber + `</td></tr>`
			}
		}
	} else if stageID == "105" {
		//displaying certification details
		for _, v := range tdp {
			if  v.Type == 6 && v.Key == "certificationauthorityname"{
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tempdata models.CertificationAuthority
				mapstructure.Decode(artifactData, &tempdata)

				tableContent += `<tr><td>Certification Authority Name </td><td class="value-label">` + tempdata.Name + `</td></tr>`
				tableContent += `<tr><td>Certification Authority Address</td><td class="value-label">` + tempdata.Address + `</td></tr>`

				break
			}
			
		}
		for _, v := range tdp {
			if v.Type == 3 {
				tableContent += `<tr><td>Certification Date</td><td class="value-label">` + v.Val.(string) + `</td></tr>`
				break
			}
		}
		for _, v := range tdp {
			if v.Type == 5 {
				tableContent += `<tr><td>` + v.Key + `</td><td class="value-label">` + v.Val.(string) + `</td></tr>`
			}
		}

		images := GetImages(tdp) //get certification images
		tableContent += `<tr><td>Images</td><td class="td-carousel"><div class="carousel-wrapper"><section class="carousel" aria-label="Gallery">
						<ol class="carousel__viewport">`

		//display certification images in a carousel				
		len := len(images)
		i := 0
		for _, value := range images {
			curSlide := i + 1
			prev := curSlide - 1
			next := curSlide + 1

			if i == 0 {
				prev = len
			} else if i == len -1 {
				next = 1
			}

			lat := fmt.Sprintf("%f", value.GeoCode.Lat)
			long := fmt.Sprintf("%f", value.GeoCode.Long)

			tableContent += `<li id="carousel__slide_cert` + strconv.Itoa(curSlide) +`" tabindex="0" class="carousel__slide"><div class="carousel__snapper">`
			tableContent += `<img src="` + value.Image + `" /><a href="#carousel__slide_cert`+strconv.Itoa(prev)+`" class="carousel__prev">Go to previous slide</a>
							<a href="#carousel__slide_cert`+strconv.Itoa(next)+`" class="carousel__next">Go to next slide</a></div>`
							tableContent += `<div class="map-link-div"><a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span></div>
							<div class="timestamp-div"><label class="timestamp-label">Time Stamp : `+value.TimeStamp.Time().String()+`</label></div></li>`
			i++

		}

		tableContent += `</ol></section></div></td></tr>`

	} else if stageID == "106" { 
		//display export details
		for _, v := range tdp {
			if  v.Type == 6 {
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tempdata models.ExporterInfo
				mapstructure.Decode(artifactData, &tempdata)

				tableContent += `<tr><td>Exporter Name </td><td class="value-label">` + tempdata.Name + `</td></tr>`
				tableContent += `<tr><td>Exporter Location</td><td class="value-label">` + tempdata.Address + `</td></tr>`
				tableContent += `<tr><td>EXPORT (Dealer) LICENSE Number</td><td class="value-label">` + tempdata.LicenseNumber + `</td></tr>`
				tableContent += `<tr><td>EXPORT (Dealer) LICENSE Exp.date</td><td class="value-label">` + tempdata.LicenseExpirationDate + `</td></tr>`

				break
			}
		}
	} else if stageID == "109" { 
		//display appraisal details
		for _, v := range tdp {
			if  v.Type == 3 {
				tableContent += `<tr><td>Appraisal Date </td><td class="value-label">` + v.Val.(string) + `</td></tr>`
			} else if  v.Type == 6 {
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tempdata models.Appraiser
				mapstructure.Decode(artifactData, &tempdata)
				tableContent += `<tr><td>Appraiser </td><td class="value-label">` + tempdata.Name + `</td></tr>`
				tableContent += `<tr><td>Appraiser Qualification </td><td class="value-label">` + tempdata.Qualification + `</td></tr>`
			} else if  v.Type == 4 {
				tableContent += `<tr><td>Appraisal Photos</td>`
				dataArr := v.Val.([] interface{})

				//display appraisal images
				for _, v := range dataArr {
					mapdata := v.(map[string]interface{})
					var tempdata models.GeoImageData
					mapstructure.Decode(mapdata, &tempdata)

					lat := fmt.Sprintf("%f", tempdata.GeoCode.Lat)
					long := fmt.Sprintf("%f", tempdata.GeoCode.Long)


					tableContent += `<td><img src="` + tempdata.Image + `" class="report-img" /><br/><br/>
					<a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span><br/>
					<label class="timestamp-label" style="margin-bottom: 30px">Time Stamp : `+tempdata.TimeStamp.Time().String()+`</label></td>`
					break
				}
				tableContent += `</tr>`
			}
		}
	}

	tableContent += `</tbody></table></div>`
	htmlBody += heading + tableContent
	
}

//get images in a tdp
func GetImages (tdp []models.TraceabilityData) map[int]models.GeoImageData{
	images := make(map[int]models.GeoImageData)
	var i int = 0
	for _, v := range tdp {
			
		if v.Type == 4 {
			dataArr := v.Val.([] interface{})
			for _, v := range dataArr {
				mapdata := v.(map[string]interface{})
				var tempdata models.GeoImageData
				mapstructure.Decode(mapdata, &tempdata)
				images[i] = tempdata
				i++
			}
		}
		
	}

	return images
}
