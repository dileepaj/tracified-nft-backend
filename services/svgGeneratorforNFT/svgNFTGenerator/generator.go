package svgNFTGenerator

import (
	//"fmt"
	"fmt"
	"strconv"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/services/mapGenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
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
	mapRepository  customizedNFTrepository.MapRepository
	backendUrl     = configs.GetNftBackendBaseUrl()
)

func GenerateSVGTemplateforNFT(data []models.Component, batchID string, productID string, shopID string, receiverName string, message string, nftname string) (string, error) {
	//get gem type from tdp data
	/* var gemVariety string = ""
	var gemDetailsTDP []models.TraceabilityData

	for _, dataArr := range tdpData {
		for _, maindata := range dataArr {
			if maindata.StageID == "104" {
				gemDetailsTDP = maindata.TraceabilityDataPackets[0].TraceabilityData
				gemVariety = GetGemVariety(gemDetailsTDP)
			}
		}
	} */

	var htmlStart = `<div class="nft-header default-font">
						<div class="nft-header-content">
							<div class="header-logo-cont">
								<img src="https://ruri-nft.s3.ap-south-1.amazonaws.com/assets/images/RURI%2B1sa+1.png" class="ruri-logo" />
								<img src="https://s3.ap-south-1.amazonaws.com/qa.marketplace.nft.tracified.com/Tracified-RT-Logo-White.svg"
								class="nft-logo" />
							</div>
							<div class="nft-header-title">
								<label id="topTitle">NFT</label>
								<label id="nftName">` + data[0].Item + `</label>
							</div>
						</div>
					</div>
					<div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	var iframeImg = `<div class="iframe-wrapper"><iframe  src="https://tracified.sirv.com/Spins/RURI%20Gems/` + shopID + `/` + shopID + `.spin" class="iframe-img" frameborder="0" allowfullscreen="true"></iframe><span class="rotate-icon" style="margin-top : 30px;"></span></div>`

	if receiverName != "" && message != "" {
		GenerateOwnership(receiverName, message, nftname)
	}

	GenerateContent(data)

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
func GenerateOwnership(receiverName string, message string, nftname string) {
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
											<td class="tbl-text-normal">NFT Name</td>
											<td class="tbl-text-normal">
												<p>` + nftname + `</p>
											</td>
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

// new digital twin related development

func GenerateContent(data []models.Component) {

	for _, data := range data {
		if data.Component == "expandableTab" && len(data.Tabs) > 0 {
			GenerateTable(data)
		} else if data.Component == "expandableTab" && len(data.VerticalTab) > 0 {
			GenerateVerticalTabs(data)
		}
	}
}

func GenerateTable(data models.Component) {
	tableContent := ""
	var icon string

	for _, tab := range data.Tabs {
		for _, component := range tab.Children {
			if component.Component == "key-value" {
				var valueWithProof models.ValueWithProof
				mapstructure.Decode(component.Value, &valueWithProof)
				tableContent += `<tr><td class="tbl-text-normal">` + component.Key + `</td><td class="tbl-text-bold">` + valueWithProof.Value.(string) + `</td></tr>`
			}
		}
	}

	if data.Title == "Gem Details" {
		icon = `<span class="gem-icon"></span>`
	} else {
		icon = `<img src="` + data.Icon + `" />`
	}

	htmlBody += `<div class="widget-div">
					<div class="wrap-collabsible">
						<input id="collapsible2" class="toggle" type="radio" name="toggle"></input>
						<label for="collapsible2" class="lbl-toggle" tabindex="0">
							` + icon + `
							<label>` + data.Title + `</label>
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

// generate vertical tab component
func GenerateVerticalTabs(data models.Component) {
	mainTabs := ""
	sidebarTabs := ""
	radioButtons := ""
	content := ""
	for _, tab := range data.VerticalTab {
		if tab.Component == "overview" {
			cont, mainTbs, sidebarTbs, radioBtns := GenerateOverview(tab)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		}

	}

	icon := ""

	if data.Title == "Digital Twin" {
		icon = `<span class="digital-twin-icon"></span>`
	} else {
		icon = `<img src="` + data.Icon + `" />`
	}

	htmlBody += `<div class="widget-div">
					<div class="wrap-collabsible">
						<input id="collapsible3" class="toggle" type="radio" name="toggle"></input>
						<label for="collapsible3" class="lbl-toggle" tabindex="0">
							` + icon + `
							<label>` + data.Title + `</label>
							<span class="arrow-down-icon"></span>
						</label>

						<div class="collapsible-content">
							<div class="toggle-div">
								<input id="sidebar-toggle" type="checkbox"></input>
								<label for="sidebar-toggle"><span class="open-menu-icon"></span></label>
								<div id="sidebar">
									<div id="sidebar-inner">
										<ul class="sidebar-tabs">
											` + sidebarTabs + `
										</ul>
									</div>
								</div>
							</div>
							<div class="content-inner">
								<div class="tabbed">
									<div style="display: flex; flex-direction : row">
										` + radioButtons + `
										<ul class="tabs">
											` + mainTabs + `
										</ul>

										` + content + `
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>`

}

// Generate overview tabs
func GenerateOverview(data models.Component) (string, string, string, string) {
	mainTabs := ""
	sidebarTabs := ""
	radioButtons := ""
	content := ""

	for index, tab := range data.Children {
		if tab.Component == "vertical-card-container" {
			cont, mainTbs, sidebarTbs, radioBtns := GenerateVerticalCardContainer(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		} else if tab.Component == "map" {
			cont, mainTbs, sidebarTbs, radioBtns := GenerateJourneyMap(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		} else if tab.Component == "timeline" {
			cont, mainTbs, sidebarTbs, radioBtns := GenerateTimeline(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		}

	}

	return content, mainTabs, sidebarTabs, radioButtons

}

// Generate vertical card container
func GenerateVerticalCardContainer(data models.Component, index int) (string, string, string, string) {
	content := ""
	mainTab, sidebarTab, radioButton := GenerateTabLabels(data.Title, index)

	for i, childComponent := range data.Children {
		if childComponent.Component == "image-slider" {
			res := GenerateImageSlider(childComponent, index)
			content += `<div class="tab-content">
							<div class="img-list">
							` + res + `	
							</div>
						</div>`
		} else if childComponent.Component == "decorated-key-value" {
			if i == 0 {
				res1 := GenerateDecoratedKeyValues(data, i)
				res2 := GenerateDecoratedKeyValuesHeading(data, res1)
				content += `<div class="tab-content">
							` + res2 + `
						</div>`
			}

		}
	}

	return content, mainTab, sidebarTab, radioButton
}

func GenerateDecoratedKeyValuesHeading(data models.Component, cards string) string {
	icon := ""

	if data.Title == "Origin" {
		icon = `<span class="tree-icon"></span>`
	} else if data.Title == "Quality" {
		icon = `<span class="badge-icon"></span>`
	} else if data.Title == "Sustainability" {
		icon = `<span class="sustainability-icon"></span>`
	} else if data.Title == "Compliance" {
		icon = `<span class="handshake-outline-icon"></span>`
	} else {
		icon = data.Icon
	}

	content := `` + icon + `
				<label class="tab-cont-heading">` + data.Subtitle + `</label>
				<div class="card-container">
					` + cards + `
				</div>`

	return content
}

// Generate decorated key values
func GenerateDecoratedKeyValues(data models.Component, index int) string {
	cards := ""
	color := ""

	if data.Title == "Origin" {
		color = "green"
	} else if data.Title == "Quality" {
		color = "blue"
	} else if data.Title == "Sustainability" {
		color = "orange"
	} else if data.Title == "Compliance" {
		color = "brown"
	} else {
		color = ""
	}

	for _, child := range data.Children {
		if child.Component == "decorated-key-value" {
			//img := ""
			val := "No Records"
			var decoratedVal models.ValueWithProof
			mapstructure.Decode(child.Value, &decoratedVal)
			keyValIcon := GetDecoratedKeyValueIcon(child.Key)

			/* if child.Icon != "" {
				img = `<img class="dt-icon-img" src="` + child.Icon + `" />`
			} */
			if decoratedVal.Value != nil && decoratedVal.Value.(string) != "" {
				val = decoratedVal.Value.(string)
			}

			cards += `<div class="tab-cont-card ` + color + `">
							<div class="card-div-1">
								<span class="` + keyValIcon + `"></span>
							</div>
							<div class="card-div-2">
								<label class="bold-text">` + child.Key + `</label>
								<label>` + val + `</label>
							</div>
						</div>`
		}
	}

	return cards
}

// Generate image slider
func GenerateImageSlider(imageSlider models.Component, parentIndex int) string {
	content := ""

	for i, image := range imageSlider.Images.Value {
		if image.Img == "" {
			content += `<p>No Records</p>`
		} else {
			content += `<div class="img-wrapper">
			<input type="checkbox" id="cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `" class="img-zoom-in"></input>
			<div class="img-fullscreen">
				<label for="cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `">
					<span class="material-symbols-outlined">
						close
					</span>
				</label>
				<div class="img-div"
					style="background-image: url('` + image.Img + `');">
				</div>
			</div>
			<div class="img-div"
				style="background-image: url('` + image.Img + `');">
			</div>
			<label for="cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `" title="View Image">
				<span class="zoom-icon"></span>
			</label>
		</div>`
		}
	}

	return content
}

// Generate Journey Map
func GenerateJourneyMap(tab models.Component, index int) (string, string, string, string) {

	var mapInfo []models.MapInfo

	/* for _, c := range tab.Coordinates {
		for _, c1 := range c.Values {
			coordinate := c1
			var cmap models.MapInfo
			cmap.Latitude = coordinate.Lat
			cmap.Longitude = coordinate.Long
			mapInfo = append(mapInfo, cmap)
		}

	} */

	mapInfo = append(mapInfo, models.MapInfo{6.927079, 79.861244})
	mapInfo = append(mapInfo, models.MapInfo{35.652832, 139.839478})

	generatedMap := mapGenerator.GenerateMap(mapInfo)
	var newMap models.GeneratedMap
	newMap.MapTemplate = generatedMap
	rst, mapSaveErr := mapRepository.SaveMap(newMap)

	if mapSaveErr != nil {
		logs.ErrorLogger.Println("Failed to save map : ", mapSaveErr.Error())
	}

	content := `<div class="tab-content">
					<embed class="map" frameborder="0" scrolling="no" marginheight="0"
						marginwidth="0"
						src="` + backendUrl + `/GetMap/` + rst + `"></embed>
				</div>`

	mainTab, sidebarTab, radioButton := GenerateTabLabels("Journey", index)

	return content, mainTab, sidebarTab, radioButton
}

// Generate Timeline
func GenerateTimeline(data models.Component, index int) (string, string, string, string) {
	content := ``
	tlCont := ""

	mainTab, sidebarTab, radioButton := GenerateTabLabels(data.Title, index)

	for _, stage := range data.Children {
		infoStr := ""
		for _, info := range stage.Children {
			if info.Component == "key-value" {

				val := "No Data Available"
				var decoratedVal models.ValueWithProof
				mapstructure.Decode(info.Value, &decoratedVal)

				if decoratedVal.Value != nil && decoratedVal.Value.(string) != "" {
					val = decoratedVal.Value.(string)
				}

				infoStr += `<div class="tl-info-container">
								<label class="grey-text">` + strings.Replace(info.Key, "&", "&amp;", -1) + `</label>
								<label class="tl-bold-text">` + val + `</label>
							</div>`
			}
		}

		tlCont += `<div class="tl-stage">
					<div class="tl-heading">
						<div class="tl-circle">
							<span class="stack-icon"></span>
						</div>
						<label>` + stage.Title + `</label>
					</div>
					<div class="tl-content">
						` + infoStr + `
					</div>
					</div>`
	}

	content = `<div class="tab-content">
					<div class="tl-wrapper">
					` + tlCont + `
					</div>
				</div>`

	return content, mainTab, sidebarTab, radioButton
}

func GetWordWrap(value string) string {
	fmt.Println(len(strings.Split(value, " ")))
	if len(strings.Split(value, " ")) == 1 {
		return `style="word-break: break-all;"`
	} else {
		return ""
	}
}

func GetDecoratedKeyValueIcon(key string) string {
	switch key {
	case "Provenance 原産地":
		return "hexagon-icon"
	case "Natural 天然石":
		return "leaves-icon"
	case "Certification 鑑別":
		return "certificate-icon"
	case "Treatment 処理":
		return "treatment-icon"
	case "Fairtrade 公正取引":
		return "handshake-icon"
	case "Social Impact 社会貢献":
		return "social-impact-icon"
	case "Mining 採鉱":
		return "mining-icon"
	case "Trading 貿易":
		return "trading-icon"
	default:
		return ""
	}
}

// Create labels for vertical tabs
func GenerateTabLabels(title string, index int) (string, string, string) {
	sidebarTab := `<li class="tab">
						<label for="tab` + strconv.Itoa(index+1) + `" onclick="closeSidebar()">
							` + title + `
							<span class="tab-arrow-icon"></span>
						</label>
					</li>`
	checked := ""

	if index == 0 {
		checked = `checked="true"`
	}

	radioButton := `<input type="radio" id="tab` + strconv.Itoa(index+1) + `" name="css-tabs" ` + checked + `></input>`

	mainTab := `<li class="tab">
					<label for="tab` + strconv.Itoa(index+1) + `" >
					` + title + `
						<span class="tab-arrow-icon"></span>
					</label>
				</li>`

	return mainTab, sidebarTab, radioButton
}
