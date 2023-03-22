package svgNFTGenerator

import (
	//"fmt"

	"strconv"
	"strings"

	customizedNFTrepository "github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/mitchellh/mapstructure"
	//"github.com/mitchellh/mapstructure"
)

var (
	svgStart       = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTHeader.txt")
	svgEnd         = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTFooter.txt")
	styling        = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTStyles.css") //!Need to implement
	styleStart     = `<style>`
	styleEnd       = `</style>`
	htmlBody       = ""
	collectionName = ""
	ruriRepository customizedNFTrepository.SvgRepository
)

func GenerateSVGTemplateforNFT(tdpData [][]models.TDPParent, batchID string, productID string, receiverName string, message string) (string, error) {
	//get gem type from tdp data
	var gemVariety string = ""
	var gemDetailsTDP []models.TraceabilityData

	for _, dataArr := range tdpData {
		for _, maindata := range dataArr {
			if maindata.StageID == "104" {
				gemDetailsTDP = maindata.TraceabilityDataPackets[0].TraceabilityData
				gemVariety = GetGemVariety(gemDetailsTDP)
			}
		}
	}

	var htmlStart = `<div class="nft-header default-font">
						<div class="nft-header-content">
							<div class="header-logo-cont">
								<img src="https://ruri-nft.s3.ap-south-1.amazonaws.com/assets/images/RURI%2B1sa+1.png" class="ruri-logo" />
								<img src="https://s3.ap-south-1.amazonaws.com/qa.marketplace.nft.tracified.com/Tracified-RT-Logo-White.svg"
								class="nft-logo" />
							</div>
							<div class="nft-header-title">
								<label id="topTitle">NFT</label>
								<label id="nftName">` + gemVariety + `</label>
							</div>
						</div>
					</div>
					<div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	var iframeImg = `<div class="iframe-wrapper"><iframe  src="https://tracified.sirv.com/Spins/RURI%20Gems%20Compressed/120614/120614.spin" class="iframe-img" frameborder="0" allowfullscreen="true"></iframe><span class="rotate-icon" style="margin-top : 30px;"></span></div>`

	/* var stageStatus map[string]bool = make(map[string]bool)

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

	} */

	//logs.InfoLogger.Println("TDP", tdpData)

	if receiverName != "" && message != "" {
		GenerateOwnership(receiverName, message)
	}

	GenerateGemDetails(gemDetailsTDP, gemVariety)

	//timelineData := GenerateTimeline(tdpData)

	//GenerateDigitalTwin(gemDetailsTDP, timelineData)

	template := svgStart + styleStart + styling + styleEnd + htmlStart + iframeImg + htmlBody + svgEnd
	htmlBody = ""
	/* template = strings.(template)
	fmt.Println(template) */
	/* template = strings.Replace(template, "\r", " ", -1)
	template = strings.Replace(template, "\t", " ", -1)
	template = strings.Replace(template, "\n", " ", -1) */
	return template, nil
}

// generate ownership section
func GenerateOwnership(receiverName string, message string) {
	htmlBody += `<div class="widget-div">
					<div class="wrap-collabsible">
						<input id="collapsible1" class="toggle" type="radio" name="toggle" checked="true"></input>
						<label for="collapsible1" class="lbl-toggle" tabindex="0">
							<span class="profile-icon"></span>
							<label>Ownership</label>
							<span class="arrow-down-icon"></span>
						</label>
						<div class="collapsible-content">
							<div class="content-inner">
								<div class="bdr">
									<table class="table table-bordered rounded-20 overflow-hidden">
										<tbody>
											<tr>
												<td class="tbl-text-normal">Owner's Name </td>
												<td class="tbl-text-bold">` + receiverName + `</td>
											</tr>
											<tr>
												<td class="tbl-text-normal">Message</td>
												<td class="tbl-text-normal">
													<p>` + message + `</p>
												</td>
											</tr>
										</tbody>
									</table>
								</div>
							</div>
						</div>
					</div>
				</div>`
}

// generate gem details section
func GenerateGemDetails(tdp []models.TraceabilityData, gemVariety string) {
	colour := ""
	species := ""
	shape := ""
	carat := ""
	measurements := ""
	treatment := ""
	tableContent := ""

	for _, v := range tdp {
		if v.Key == "colour" {
			colour = v.Val.(string)
		} else if v.Key == "species" {
			species = v.Val.(string)
		} else if v.Key == "shape&cut" {
			shape = v.Val.(string)
		} else if v.Key == "carat" {
			carat = strconv.FormatFloat(v.Val.(float64), 'f', 2, 32)
		} else if v.Key == "measurement" {
			measurements = v.Val.(string)
		} else if v.Key == "treatment" {
			treatment = v.Val.(string)
		}

	}

	if gemVariety != "" {
		tableContent += `<tr><td class="tbl-text-normal">Variety</td><td class="tbl-text-bold">` + gemVariety + `</td></tr>`
	}
	if species != "" {
		tableContent += `<tr><td class="tbl-text-normal">Species</td><td class="tbl-text-bold">` + species + `</td></tr>`
	}
	if colour != "" {
		tableContent += `<tr><td class="tbl-text-normal">Colour</td><td class="tbl-text-bold">` + colour + `</td></tr>`
	}
	if shape != "" {
		tableContent += `<tr><td class="tbl-text-normal">Shape</td><td class="tbl-text-bold">` + shape + `</td></tr>`
	}
	if carat != "" {
		tableContent += `<tr><td class="tbl-text-normal">Caret</td><td class="tbl-text-bold">` + carat + `</td></tr>`
	}
	if measurements != "" {
		tableContent += `<tr><td class="tbl-text-normal">Measurements (mm)</td><td class="tbl-text-bold">` + measurements + `</td></tr>`
	}
	if treatment != "" {
		tableContent += `<tr><td class="tbl-text-normal">Treatment</td><td class="tbl-text-bold">` + treatment + `</td></tr>`
	}

	htmlBody += `<div class="widget-div">
					<div class="wrap-collabsible">
						<input id="collapsible2" class="toggle" type="radio" name="toggle"></input>
						<label for="collapsible2" class="lbl-toggle" tabindex="0">
							<span class="gem-icon"></span>
							<label>Gem Details</label>
							<span class="arrow-down-icon"></span>
						</label>
						<div class="collapsible-content">
							<div class="content-inner">
								<div class="bdr">
									<table class="table table-bordered rounded-20 overflow-hidden">
										<tbody>
											` + tableContent + `
										</tbody>
									</table>
								</div>
							</div>
						</div>
					</div>
				</div>`
}

// get gem variety
func GetGemVariety(tdp []models.TraceabilityData) string {
	variety := ""
	for _, v := range tdp {
		if v.Key == "variety" {
			variety = v.Val.(string)
		}

	}

	return variety
}

func GenerateDigitalTwin(gemDetailsTDP []models.TraceabilityData, timelineData string) {
	certificationImages := ""

	images := GetImages(gemDetailsTDP) //get certification images
	i := 0
	for _, value := range images {
		certificationImages += `<div class="img-wrapper">
									<input type="checkbox" id="cert` + strconv.Itoa(i+1) + `" class="img-zoom-in"></input>
									<div class="img-fullscreen">
										<label for="cert` + strconv.Itoa(i+1) + `">
											<span class="material-symbols-outlined">
												close
											</span>
										</label>
										<div class="img-div"
											style="background-image: url('` + value.Image + `');">
										</div>
									</div>
									<div class="img-div"
										style="background-image: url('` + value.Image + `');">
									</div>
									<label for="cert` + strconv.Itoa(i+1) + `" title="View Image">
										<span class="zoom-icon"></span>
									</label>
								</div>`
		i++

	}

	htmlBody += `<div class="widget-div">
					<div class="wrap-collabsible">
						<input id="collapsible3" class="toggle" type="radio" name="toggle"></input>
						<label for="collapsible3" class="lbl-toggle" tabindex="0">
							<span class="digital-twin-icon"></span>
							<label>Digital Twin</label>
							<span class="arrow-down-icon"></span>
						</label>

						<div class="collapsible-content">
							<div class="toggle-div">
								<input id="sidebar-toggle" type="checkbox"></input>
								<label for="sidebar-toggle"><span class="open-menu-icon"></span></label>
								<div id="sidebar">
									<div id="sidebar-inner">
										<ul class="sidebar-tabs">
											<li class="tab">
												<label for="tab1">
													Certification
													<span class="tab-arrow-icon"></span>
												</label>
											</li>
											<li class="tab"><label for="tab2">Appraisal <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab3">Journey <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab4">Origin <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab5">Quality <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab6">Sustainability <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab7">Compliance <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab8">Timeline <span
														class="tab-arrow-icon"></span> </label></li>
										</ul>
									</div>
								</div>
							</div>
							<div class="content-inner">
								<div class="tabbed">
									<div style="display: flex; flex-direction : row">
										<input type="radio" id="tab1" name="css-tabs" checked="true"></input>
										<input type="radio" id="tab2" name="css-tabs"></input>
										<input type="radio" id="tab3" name="css-tabs"></input>
										<input type="radio" id="tab4" name="css-tabs"></input>
										<input type="radio" id="tab5" name="css-tabs"></input>
										<input type="radio" id="tab6" name="css-tabs"></input>
										<input type="radio" id="tab7" name="css-tabs"></input>
										<input type="radio" id="tab8" name="css-tabs"></input>
										<ul class="tabs">
											<li class="tab">
												<label for="tab1">
													Certification
													<span class="tab-arrow-icon"></span>
												</label>
											</li>
											<li class="tab"><label for="tab2">Appraisal <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab3">Journey <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab4">Origin <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab5">Quality <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab6">Sustainability <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab7">Compliance <span
														class="tab-arrow-icon"></span> </label></li>
											<li class="tab"><label for="tab8">Timeline <span
														class="tab-arrow-icon"></span> </label></li>
										</ul>

										<!--Certification-->
										<div class="tab-content">
											<div class="img-list">
											` + certificationImages + `	
											</div>
										</div>

										<!--Appraisal-->
										<div class="tab-content">
											<div class="img-list">
												<div class="img-wrapper">
													<input type="checkbox" id="appr1" class="img-zoom-in"></input>
													<div class="img-fullscreen">
														<label for="appr1">
															<span class="material-symbols-outlined">
																close
															</span>
														</label>
														<div class="img-div"
															style="background-image : url('https://via.placeholder.com/400x570')">
														</div>
													</div>
													<div class="img-div"
														style="background-image : url('https://via.placeholder.com/400x570')">
													</div>
													<label for="appr1" title="View Image">
														<span class="zoom-icon"></span>
													</label>
												</div>

												<div class="img-wrapper">
													<input type="checkbox" id="appr2" class="img-zoom-in"></input>
													<div class="img-fullscreen">
														<label for="appr2">
															<span class="material-symbols-outlined">
																close
															</span>
														</label>
														<div class="img-div"
															style="background-image : url('https://via.placeholder.com/400x570')">
														</div>
													</div>
													<div class="img-div"
														style="background-image : url('https://via.placeholder.com/400x570')">
													</div>
													<label for="appr2" title="View Image">
														<span class="zoom-icon"></span>
													</label>
												</div>
											</div>
										</div>

										<!--Journey-->
										<div class="tab-content">
											<iframe class="map" frameborder="0" scrolling="no" marginheight="0"
												marginwidth="0"
												src="https://www.openstreetmap.org/export/embed.html?bbox=139.7599768638611%2C35.70130685541025%2C139.78006124496463%2C35.71334671547798&amp;layer=mapnik&amp;marker=35.70732701275398%2C139.77001905441284"></iframe>
										</div>

										<!--Origin-->
										<div class="tab-content">
											<span class="tree-icon"></span>
											<label class="tab-cont-heading">Begin With Mother Earth</label>
											<div class="card-container">
												<div class="tab-cont-card green">
													<div class="card-div-1">
														<span class="hexagon-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Provenance 原産地</label>
														<label>No Records</label>
													</div>
												</div>
												<div class="tab-cont-card green">
													<div class="card-div-1">
														<span class="leaves-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Natural 天然石</label>
														<label>Corundum</label>
													</div>
												</div>
											</div>


										</div>

										<!--Quality-->
										<div class="tab-content">
											<span class="badge-icon"></span>
											<label class="tab-cont-heading">Commitment To Integrity</label>
											<div class="card-container">
												<div class="tab-cont-card blue">
													<div class="card-div-1">
														<span class="certificate-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Certification 鑑別</label>
														<label>No Records</label>
													</div>
												</div>
												<div class="tab-cont-card blue">
													<div class="card-div-1">
														<span class="treatment-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Treatment 処理</label>
														<label>Corundum</label>
													</div>
												</div>
											</div>

										</div>

										<!--Sustainability-->
										<div class="tab-content">
											<span class="sustainability-icon"></span>
											<label class="tab-cont-heading">More Than A Gemstone</label>
											<div class="card-container">
												<div class="tab-cont-card orange">
													<div class="card-div-1">
														<span class="handshake-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Fairtrade 公正取引</label>
														<label>No Records</label>
													</div>
												</div>
												<div class="tab-cont-card orange">
													<div class="card-div-1">
														<span class="social-impact-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Social Impact 社会貢献</label>
														<label>No Records</label>
													</div>
												</div>
											</div>

										</div>

										<!--Compliance-->
										<div class="tab-content">
											<span class="handshake-outline-icon"></span>
											<label class="tab-cont-heading">Ethical Conduct</label>
											<div class="card-container">
												<div class="tab-cont-card brown">
													<div class="card-div-1">
														<span class="mining-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Mining 採鉱</label>
														<label>No Records</label>
													</div>
												</div>
												<div class="tab-cont-card brown">
													<div class="card-div-1">
														<span class="trading-icon"></span>
													</div>
													<div class="card-div-2">
														<label class="bold-text">Trading 貿易</label>
														<label>20216DL3848</label>
													</div>
												</div>
											</div>

										</div>

										<!--Timeline-->
										<div class="tab-content">
											<div class="tl-wrapper">
											` + timelineData + `
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>`
}

// get gem images from stage 103
func GetGemImages(tdpData []models.TDP) []interface{} {
	var imgArr []interface{}
	for _, maindata := range tdpData {
		if maindata.StageID == "103" {
			for _, v := range maindata.TraceabilityData {
				if v.Type == 4 && v.Key == "photoofGem" {

					imgArr = v.Val.([]interface{})
					break

				}
			}
		}
	}

	return imgArr
}

func GenerateTimeline(tdpData [][]models.TDPParent) string {
	var stageData map[string]string = make(map[string]string)
	var stageStatus map[string]bool = make(map[string]bool)
	var timelineData string = ""
	var stageTDP []models.TraceabilityData
	infoStr := ""
	stageName := ""

	for _, dataArr := range tdpData {
		for _, maindata := range dataArr {
			if !stageStatus[maindata.StageID] {
				stageName = ""

				if maindata.StageID == "100" {
					stageName = "Mining"
				} else if maindata.StageID == "101" {
					stageName = "Treatment"

				} else if maindata.StageID == "102" {
					stageName = "Cutting"

				} else if maindata.StageID == "103" {
					stageName = "Collection"
				}

				stageTDP = maindata.TraceabilityDataPackets[0].TraceabilityData
				infoStr = ""

				for _, data := range stageTDP {
					if data.Type == 3 {
						infoStr += `<div class="tl-info-container">
										<label class="grey-text">` + strings.Replace(data.Key, "&", "&amp;", -1) + `</label>
										<label class="tl-bold-text">` + strings.Split(data.Val.(string), "T")[0] + `</label>
									</div>`
					}
				}

				if infoStr == "" {
					infoStr = `<div class="tl-info-container">
									<label class="grey-text">No traceability data available for
										this stage.</label>
								</div>`
				}

				stageData[maindata.StageID] = `<div class="tl-stage">
													<div class="tl-heading">
														<div class="tl-circle">
															<span class="stack-icon"></span>
														</div>
														<label>` + stageName + `</label>
													</div>
													<div class="tl-content">
														` + infoStr + `
													</div>
												</div>`

				stageStatus[maindata.StageID] = true
			}
		}
	}

	if stageData["100"] == "" {
		stageData["100"] = `<div class="tl-stage">
								<div class="tl-heading">
									<div class="tl-circle">
										<span class="stack-icon"></span>
									</div>
									<label>Mining</label>
								</div>
								<div class="tl-content">
									<div class="tl-info-container">
										<label class="grey-text">No traceability data available for
											this stage.</label>
									</div>
								</div>
							</div>`
	}
	if stageData["101"] == "" {
		stageData["101"] = `<div class="tl-stage">
								<div class="tl-heading">
									<div class="tl-circle">
										<span class="stack-icon"></span>
									</div>
									<label>Treatment</label>
								</div>
								<div class="tl-content">
									<div class="tl-info-container">
										<label class="grey-text">No traceability data available for
											this stage.</label>
									</div>
								</div>
							</div>`
	}
	if stageData["102"] == "" {
		stageData["102"] = `<div class="tl-stage">
								<div class="tl-heading">
									<div class="tl-circle">
										<span class="stack-icon"></span>
									</div>
									<label>Cutting</label>
								</div>
								<div class="tl-content">
									<div class="tl-info-container">
										<label class="grey-text">No traceability data available for
											this stage.</label>
									</div>
								</div>
							</div>`
	}
	if stageData["103"] == "" {
		stageData["103"] = `<div class="tl-stage">
								<div class="tl-heading">
									<div class="tl-circle">
										<span class="stack-icon"></span>
									</div>
									<label>Collection</label>
								</div>
								<div class="tl-content">
									<div class="tl-info-container">
										<label class="grey-text">No traceability data available for
											this stage.</label>
									</div>
								</div>
							</div>`
	}

	timelineData += stageData["100"] + stageData["101"] + stageData["102"] + stageData["103"]

	return timelineData
}

// get collection name from stage 103
/* func GetCollectionName(tdpData []models.TDP) {
	for _, maindata := range tdpData {
		if maindata.StageID == "103" {
			for _, v := range maindata.TraceabilityData {
				if v.Type == 6 && v.Key == "collectionname" {
					var artifactData map[string]interface{} = v.Val.(map[string]interface{})
					for key, itmdata := range artifactData {
						if key == "name" {
							collectionName = itmdata.(string)
							break
						}
					}
				}
			}
		}
	}
}

// get and show gem details in nft
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
		} else if i == len-1 {
			next = 1
		}

		lat := fmt.Sprintf("%f", tempdata.GeoCode.Lat)
		long := fmt.Sprintf("%f", tempdata.GeoCode.Long)

		tableContent += `<li id="carousel__slide` + strconv.Itoa(curSlide) + `" tabindex="0" class="carousel__slide"><div class="carousel__snapper">`
		tableContent += `<img src="` + tempdata.Image + `" /><a href="#carousel__slide` + strconv.Itoa(prev) + `" class="carousel__prev">Go to previous slide</a>
						<a href="#carousel__slide` + strconv.Itoa(next) + `" class="carousel__next">Go to next slide</a></div>`
		tableContent += `<div class="map-link-div"><a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span></div>
		<div class="timestamp-div"><label class="timestamp-label">Time Stamp : ` + tempdata.TimeStamp.Time().String() + `</label></div></li>`
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
					break
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

	heading := `<div class="widget-div"><div class="widget-title-div"><span class="` + icon + `-icon"></span><label>` + section + `</label></div></div>`
	tableContent := `<div class="bdr"><table class="table table-bordered table-striped rounded-20 overflow-hidden"><tbody>`

	if stageID == "103" {
		//displaying collector's/ dealer's information
		for _, v := range tdp {
			if v.Type == 6 && v.Key == "collector/dealername" {
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
			if v.Type == 6 && v.Key == "certificationauthorityname" {
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
			} else if i == len-1 {
				next = 1
			}

			lat := fmt.Sprintf("%f", value.GeoCode.Lat)
			long := fmt.Sprintf("%f", value.GeoCode.Long)

			tableContent += `<li id="carousel__slide_cert` + strconv.Itoa(curSlide) + `" tabindex="0" class="carousel__slide"><div class="carousel__snapper">`
			tableContent += `<img src="` + value.Image + `" /><a href="#carousel__slide_cert` + strconv.Itoa(prev) + `" class="carousel__prev">Go to previous slide</a>
							<a href="#carousel__slide_cert` + strconv.Itoa(next) + `" class="carousel__next">Go to next slide</a></div>`
			tableContent += `<div class="map-link-div"><a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span></div>
							<div class="timestamp-div"><label class="timestamp-label">Time Stamp : ` + value.TimeStamp.Time().String() + `</label></div></li>`
			i++

		}

		tableContent += `</ol></section></div></td></tr>`

	} else if stageID == "106" {
		//display export details
		for _, v := range tdp {
			if v.Type == 6 {
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
			if v.Type == 3 {
				tableContent += `<tr><td>Appraisal Date </td><td class="value-label">` + v.Val.(string) + `</td></tr>`
			} else if v.Type == 6 {
				var artifactData map[string]interface{} = v.Val.(map[string]interface{})
				var tempdata models.Appraiser
				mapstructure.Decode(artifactData, &tempdata)
				tableContent += `<tr><td>Appraiser </td><td class="value-label">` + tempdata.Name + `</td></tr>`
				tableContent += `<tr><td>Appraiser Qualification </td><td class="value-label">` + tempdata.Qualification + `</td></tr>`
			} else if v.Type == 4 {
				tableContent += `<tr><td>Appraisal Photos</td>`
				dataArr := v.Val.([]interface{})

				//display appraisal images
				for _, v := range dataArr {
					mapdata := v.(map[string]interface{})
					var tempdata models.GeoImageData
					mapstructure.Decode(mapdata, &tempdata)

					lat := fmt.Sprintf("%f", tempdata.GeoCode.Lat)
					long := fmt.Sprintf("%f", tempdata.GeoCode.Long)

					tableContent += `<td><img src="` + tempdata.Image + `" class="report-img" /><br/><br/>
					<a href="https://maps.google.com/?q=` + lat + `,` + long + `" target="_blank" class="map-link-a">View on map</a><span class="material-symbols-outlined map-link-span">open_in_new</span><br/>
					<label class="timestamp-label" style="margin-bottom: 30px">Time Stamp : ` + tempdata.TimeStamp.Time().String() + `</label></td>`
					break
				}
				tableContent += `</tr>`
			}
		}
	}

	tableContent += `</tbody></table></div>`
	htmlBody += heading + tableContent

} */

// get images in a tdp
func GetImages(tdp []models.TraceabilityData) map[int]models.GeoImageData {
	images := make(map[int]models.GeoImageData)
	var i int = 0
	for _, v := range tdp {

		if v.Type == 4 {
			dataArr := v.Val.([]interface{})
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
