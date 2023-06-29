package svgNFTGenerator

import (
	//"fmt"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"strconv"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/services/mapGenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/mitchellh/mapstructure"
	//"github.com/mitchellh/mapstructure"
)

type RURINFT struct {
	Email        string
	ShopID       string
	BatchID      string
	ProductID    string
	ReceiverName string
	CustomMsg    string
	NFTName      string
	Logo         string
	EmailTitle   string
}

var (
	svgStart        = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTHeader.txt")
	svgEnd          = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTFooter.txt")
	styling         = services.ReadFromFile("services/svgGeneratorforNFT/templates/svgNFTStyles.css") //!Need to implement
	styleStart      = `<style>`
	styleEnd        = `</style>`
	htmlBody        = ""
	mapRepository   customizedNFTrepository.MapRepository
	backendUrl      = configs.GetNftBackendBaseUrl()
	proofModalCount = 0
	txnMap          = make(map[string][]string)
	svgRepository   customizedNFTrepository.SvgRepository
)

func (r *RURINFT) GenerateNFT() (responseDtos.SVGforNFTResponse, error) {
	var userSVGMapRst responseDtos.SVGforNFTResponse
	batchData, err := customizedNFTFacade.GetBatchIDDatabyItemID(r.ShopID)

	if err != nil {
		return userSVGMapRst, err
	}

	r.BatchID = batchData.BatchID
	r.ProductID = batchData.ItemID

	tdpData, _ := customizedNFTFacade.GetDigitalTwinData(r.BatchID, r.ProductID)
	var userNftMapping models.UserNFTMapping
	//Svg will be generated using the template
	svgrst, thumbnail, _ := r.GenerateSVGTemplateforNFT(tdpData)
	userNftMapping.BatchID = r.BatchID
	userNftMapping.SVG = svgrst
	userNftMapping.Email = r.Email
	userNftMapping.NFTName = r.NFTName
	userNftMapping.Thumbnail = thumbnail
	//Generated SVG data will get added to the DB
	rst, err1 := svgRepository.SaveUserMapping(userNftMapping)
	if err1 != nil {
		return userSVGMapRst, err1
	}
	userSVGMapRst = rst
	return userSVGMapRst, nil
}

func (r *RURINFT) GenerateSVGTemplateforNFT(data []models.Component) (string, string, error) {
	/* batchID := r.BatchID
	productID := r.ProductID  */
	//shopID := r.ShopID
	receiverName := r.ReceiverName
	message := r.CustomMsg
	nftname := r.NFTName

	proofModalCount = 0
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
						<div class="nft-header-content cont-wrapper">
							<div class="header-logo-cont">
								<img src="` + r.Logo + `" class="ruri-logo" />
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

	iframeImg, thumb := r.GenerateTopSection(data)

	var proofToggle = `<div class="proof-toggle-wrapper cont-wrapper">
							<label>View available blockchain proofs</label>
							<label class="switch">
							<input id="proveToggle" type="checkbox" onclick="onChangeProofToggle()" />
							<span class="slider round">
							</span>
							</label>
									</div>
						<div class="proof-tip-wrapper cont-wrapper provable-val">
							<label>
							Click on 
							<span class="material-symbols-outlined">
							check_circle
							</span>
							icons to view proofs.
							</label>
						</div>`

	if receiverName != "" && message != "" {
		r.GenerateOwnership(receiverName, message, nftname)
	}

	r.GenerateContent(data)

	template := svgStart + styleStart + styling + styleEnd + htmlStart + iframeImg + proofToggle + htmlBody + svgEnd
	htmlBody = ""
	txnMap = make(map[string][]string)
	/* template = strings.(template)
	fmt.Println(template) */
	/* template = strings.Replace(template, "\r", " ", -1)
	template = strings.Replace(template, "\t", " ", -1)
	template = strings.Replace(template, "\n", " ", -1) */
	return template, thumb, nil
}

// Generate section with video and physical tag
func (r *RURINFT) GenerateTopSection(data []models.Component) (string, string) {

	content := ""
	video := ""
	physicalTag := ""
	thumbnail := ""

	for _, data := range data {
		if data.Component == "expandableTab" && len(data.Tabs) > 0 && data.Title == "NFT Content" {
			arr := data.Tabs[0].Children

			for _, val := range arr {
				if val.Component == "key-value" && val.Key == "Video Link" {

					var valueWithProof models.ValueWithProof

					decodeErr := mapstructure.Decode(val.Value, &valueWithProof)
					if decodeErr != nil {
						logs.ErrorLogger.Println("Failed to decode map data : ", decodeErr.Error())
					}

					video += `<video autoplay="true" controls="true" width="500px" height="300px" allow="autoplay" loop="true" muted="muted">
									<source src="` + valueWithProof.Value.(string) + `"  type="video/mp4" />
							</video>`
				} else if val.Component == "key-value" && val.Key == "Thumbnail" {
					var valueWithProof models.ValueWithProof

					decodeErr := mapstructure.Decode(val.Value, &valueWithProof)
					if decodeErr != nil {
						logs.ErrorLogger.Println("Failed to decode map data : ", decodeErr.Error())
					}

					thumbnail = valueWithProof.Value.(string)

				} else if val.Component == "image-slider" {

					var imgs []models.ImageValue

					decodeErr := mapstructure.Decode(val.Slides.Value, &imgs)
					if decodeErr != nil {
						logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
					}

					imgstr := ""

					for i, img := range imgs {
						imgstr += `<div id="gemimg` + strconv.Itoa(i) + `" style="position: relative; min-width: 150px; height: 125px; background-position: center center; background-repeat: no-repeat; background-size: contain; background-image:url('` + img.Img + `');">
										<span class="material-symbols-outlined provable-tick-wrapper provable-val" style="position: absolute; top: 5px; right: 5px; cursor: pointer;" onclick="openModal('PhysicalTag-modal')">
											check_circle
										</span>
										<span class="material-symbols-outlined tl-view-image" style="position: absolute; bottom: 5px; right: 5px; cursor: pointer; font-size: 18px" onclick="openFullScreenImg('gemimg` + strconv.Itoa(i) + `')">
											web_asset
										</span>
									</div>`
					}

					fmt.Println(len(val.Slides.TdpId))

					proofStr := ""

					if val.Slides.Provable && len(val.Slides.TdpId) > 0 {
						proofStr, _ = r.GenerateProofContentStr("Physical Tag", val.Slides)
					}

					physicalTag += `<div class="gemimages" style="margin-top: 10px; max-width: 100%; display: flex; flex-direction: row; column-gap: 10px; row-gap: 10px; overflow-x: auto;">
											` + imgstr + proofStr + `
									</div>
									`

				}
			}

		}
	}

	if video != "" || physicalTag != "" {
		content = `<div class="iframe-wrapper cont-wrapper">` + video + physicalTag + `</div>`
	}

	return content, thumbnail
}

// generate ownership section
func (r *RURINFT) GenerateOwnership(receiverName string, message string, nftname string) {

	const wordBreakStyle = `style="word-break: break-word;"`
	receiverStyle := ""
	nftnameStyle := ""
	messageStyle := ""

	if len(strings.Split(receiverName, " ")) == 1 {
		receiverStyle = wordBreakStyle
	}
	if len(strings.Split(nftname, " ")) == 1 {
		nftnameStyle = wordBreakStyle
	}
	if len(strings.Split(message, " ")) == 1 {
		messageStyle = wordBreakStyle
	}

	htmlBody += `<div class="widget-div cont-wrapper">
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
												<td class="tbl-text-bold" ` + receiverStyle + `>` + receiverName + `</td>
											</tr>
											<tr>
											<td class="tbl-text-normal">NFT Name</td>
											<td class="tbl-text-normal" ` + nftnameStyle + `>
												<p>` + nftname + `</p>
											</td>
										</tr>
											<tr>
												<td class="tbl-text-normal">Message</td>
												<td class="tbl-text-normal" ` + messageStyle + `>
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

func (r *RURINFT) GenerateContent(data []models.Component) {

	for _, data := range data {
		if data.Component == "expandableTab" && len(data.Tabs) > 0 && data.Title != "NFT Content" {
			r.GenerateTable(data)
		} else if data.Component == "expandableTab" && len(data.VerticalTab) > 0 {
			r.GenerateVerticalTabs(data)
		}
	}
}

func (r *RURINFT) GenerateTable(data models.Component) {
	tableContent := ""
	var icon string

	for _, tab := range data.Tabs {
		for _, component := range tab.Children {
			if component.Component == "key-value" {
				var valueWithProof models.ValueWithProof

				decodeErr := mapstructure.Decode(component.Value, &valueWithProof)
				if decodeErr != nil {
					logs.ErrorLogger.Println("Failed to decode map data : ", decodeErr.Error())
				}

				proofContentStr := ""
				proofTick := ""
				if valueWithProof.Provable && len(valueWithProof.TdpId) > 0 {
					proofContentStr, proofTick = r.GenerateProofContentStr(component.Key, valueWithProof)
				}

				style := ""

				if len(strings.Split(valueWithProof.Value.(string), " ")) == 1 {
					style = `style="word-break: break-word;"`
				}

				tableContent += `<tr><td class="tbl-text-normal">` + component.Key + `</td><td class="tbl-text-bold" ` + style + `>` + valueWithProof.Value.(string) + " " + proofTick + proofContentStr + `</td></tr>`
			}
		}
	}

	if data.Title == "Gem Details" {
		icon = `<span class="gem-icon"></span>`
	} else {
		icon = `<img src="` + data.Icon + `" />`
	}

	htmlBody += `<div class="widget-div cont-wrapper">
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
func (r *RURINFT) GenerateVerticalTabs(data models.Component) {
	mainTabs := ""
	sidebarTabs := ""
	radioButtons := ""
	content := ""
	for _, tab := range data.VerticalTab {
		if tab.Component == "overview" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateOverview(tab)
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
	htmlBody += `<div class="widget-div cont-wrapper">
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
								<label class="tab-header" for="sidebar-toggle"><span class="open-menu-icon"></span> <label id="tab-name"></label></label>
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
func (r *RURINFT) GenerateOverview(data models.Component) (string, string, string, string) {
	mainTabs := ""
	sidebarTabs := ""
	radioButtons := ""
	content := ""

	for index, tab := range data.Children {
		if tab.Component == "vertical-card-container" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateVerticalCardContainer(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		} else if tab.Component == "map" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateJourneyMap(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		} else if tab.Component == "timeline" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateTimeline(tab, index)
			content += cont
			mainTabs += mainTbs
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		}

	}

	return content, mainTabs, sidebarTabs, radioButtons

}

// Generate vertical card container
func (r *RURINFT) GenerateVerticalCardContainer(data models.Component, index int) (string, string, string, string) {
	content := ""
	mainTab, sidebarTab, radioButton := r.GenerateTabLabels(data.Title, index)

	for i, childComponent := range data.Children {
		if childComponent.Component == "image-slider" {
			res := r.GenerateImageSlider(childComponent, index)
			content += `<div class="tab-content">
							<div class="img-list">
							` + res + `	
							</div>
						</div>`
		} else if childComponent.Component == "decorated-key-value" {
			if i == 0 {
				res1 := r.GenerateDecoratedKeyValues(data, i)
				res2 := r.GenerateDecoratedKeyValuesHeading(data, res1)
				content += `<div class="tab-content">
							` + res2 + `
						</div>`
			}

		}
	}

	return content, mainTab, sidebarTab, radioButton
}

func (r *RURINFT) GenerateDecoratedKeyValuesHeading(data models.Component, cards string) string {
	icon := `<span class="tree-icon" style="background-image:url('` + data.Icon + `')"></span>`

	/* if data.Title == "Origin" {
		icon = `<span class="tree-icon"></span>`
	} else if data.Title == "Quality" {
		icon = `<span class="badge-icon"></span>`
	} else if data.Title == "Sustainability" {
		icon = `<span class="sustainability-icon"></span>`
	} else if data.Title == "Compliance" {
		icon = `<span class="handshake-outline-icon"></span>`
	} else {
		icon = data.Icon
	} */

	content := `` + icon + `
				<label class="tab-cont-heading">` + data.Subtitle + `</label>
				<div class="card-container">
					` + cards + `
				</div>`

	return content
}

// Generate decorated key values
func (r *RURINFT) GenerateDecoratedKeyValues(data models.Component, index int) string {
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

			decodeErr := mapstructure.Decode(child.Value, &decoratedVal)
			if decodeErr != nil {
				logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
			}
			//keyValIcon := GetDecoratedKeyValueIcon(child.Key)

			/* if child.Icon != "" {
				img = `<img class="dt-icon-img" src="` + child.Icon + `" />`
			} */
			if decoratedVal.Value != nil && decoratedVal.Value.(string) != "" {
				val = decoratedVal.Value.(string)
			}

			proofContentStr := ""
			proofTick := ""
			if decoratedVal.Provable && len(decoratedVal.TdpId) > 0 {
				proofContentStr, proofTick = r.GenerateProofContentStr(child.Key, decoratedVal)
			}

			cards += `<div class="tab-cont-card ` + color + `">
							<div class="card-div-1">
								<span class="hexagon-icon" style="background-image:url('` + child.Icon + `')"></span>
							</div>
							<div class="card-div-2">
								<label class="bold-text">` + child.Key + `</label>
								<label>` + val + " " + proofTick + `</label>
							</div>
						</div>` + proofContentStr
		}
	}

	return cards
}

// Generate image slider
func (r *RURINFT) GenerateImageSlider(imageSlider models.Component, parentIndex int) string {
	content := ""

	for i, image := range imageSlider.Images.Value {
		if image.Img == "" {
			content += `<p>No Records</p>`
		} else {
			content += `<div class="img-wrapper">
			<div id="cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `" class="img-div"
				style="background-image: url('` + image.Img + `');">
			</div>
			<label onclick="openFullScreenImg('cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `')" title="View Image">
				<span class="zoom-icon"></span>
			</label>
		</div>`
		}
	}

	return content
}

// Generate Journey Map
func (r *RURINFT) GenerateJourneyMap(tab models.Component, index int) (string, string, string, string) {

	var mapInfo []models.MapInfo

	for _, c := range tab.Coordinates {
		for _, c1 := range c.Values {
			coordinate := c1
			var cmap models.MapInfo
			cmap.Title = c.Title
			cmap.Latitude = coordinate.Lat
			cmap.Longitude = coordinate.Long
			mapInfo = append(mapInfo, cmap)
		}

	}

	/* mapInfo = append(mapInfo, models.MapInfo{6.927079, 79.861244})
	mapInfo = append(mapInfo, models.MapInfo{35.652832, 139.839478}) */

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

	mainTab, sidebarTab, radioButton := r.GenerateTabLabels("Journey", index)

	return content, mainTab, sidebarTab, radioButton
}

// Generate Timeline
func (r *RURINFT) GenerateTimeline(data models.Component, index int) (string, string, string, string) {
	content := ``
	tlCont := ""
	imgSliderCount := 0

	mainTab, sidebarTab, radioButton := r.GenerateTabLabels(data.Title, index)

	for i, stage := range data.Children {
		infoStr := ""
		for _, info := range stage.Children {
			if info.Component == "key-value" {

				val := "No Data Available"
				var decoratedVal models.ValueWithProof
				decodeErr := mapstructure.Decode(info.Value, &decoratedVal)
				if decodeErr != nil {
					logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
				}

				if decoratedVal.Value != nil && decoratedVal.Value.(string) != "" {
					val = decoratedVal.Value.(string)
				}

				proofContentStr := ""
				proofTick := ""
				if decoratedVal.Provable && len(decoratedVal.TdpId) > 0 {
					proofContentStr, proofTick = r.GenerateProofContentStr(info.Key, decoratedVal)
				}

				infoStr += `<div class="tl-info-container tl-key-value">
								<label class="grey-text">` + strings.Replace(info.Key, "&", "&amp;", -1) + `</label>
								<label class="tl-bold-text">` + val + " " + proofTick + `</label>
							</div>` + proofContentStr

			} else if info.Component == "image-slider" {
				imgCont := ""
				var imgs []models.ImageValue
				decodeErr := mapstructure.Decode(info.Slides.Value, &imgs)
				if decodeErr != nil {
					logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
				}

				proofModalStr := ""
				proofTickIcon := ""
				if info.Slides.Provable && len(info.Slides.TdpId) > 0 {
					id := "slider" + strconv.Itoa(imgSliderCount) + "-modal"
					proofTickIcon = `<span class="material-symbols-outlined provable-tick-wrapper provable-val" onclick="openModal('` + id + `')">
														check_circle
													</span>
									`
					proofModalStr = r.GenerateImgProofModalStr(info.Slides, id)
					imgSliderCount++
				}

				if len(imgs) > 0 {

					for j, image := range imgs {
						var prev = 0
						var next = 0

						if j == 0 {
							prev = len(imgs) - 1
						} else {
							prev = j - 1
						}

						if j == len(imgs)-1 {
							next = 0
						} else {
							next = j + 1
						}

						prevStr := strconv.Itoa(i) + strconv.Itoa(prev)
						nextStr := strconv.Itoa(i) + strconv.Itoa(next)
						dateStr := strings.ReplaceAll(strings.Split(image.Time, "T")[0], "-", "/")

						if len(imgs) > 1 {
							imgCont += `<li id="carousel__slide` + strconv.Itoa(i) + strconv.Itoa(j) + `"
											tabindex="0"
											class="carousel__slide" style="background-image: url('` + image.Img + `');">
											<div class="carousel__snapper">
											<a href="#carousel__slide` + prevStr + `"
												class="carousel__prev">Go to last slide</a>
											<a href="#carousel__slide` + nextStr + `"
												class="carousel__next">Go to next slide</a>
											</div>
											<label class="date-text">` + dateStr + `<span class="material-symbols-outlined tl-view-image" onclick="openFullScreenImg('carousel__slide` + strconv.Itoa(i) + strconv.Itoa(j) + `')">
												web_asset
												</span></label>
											` + proofTickIcon + `
										</li>`
						} else {
							imgCont += `<li id="carousel__slide` + strconv.Itoa(i) + strconv.Itoa(j) + `"
											tabindex="0"
											class="carousel__slide" style="background-image: url('` + image.Img + `');">
											<div class="carousel__snapper">
											<a
												class="carousel__prev">Go to last slide</a>
											<a
												class="carousel__next">Go to next slide</a>
											</div>
											<label class="date-text">` + dateStr + `<span class="material-symbols-outlined tl-view-image" onclick="openFullScreenImg('carousel__slide` + strconv.Itoa(i) + strconv.Itoa(j) + `')">
												web_asset
												</span></label>
											` + proofTickIcon + `
										</li>`
						}

					}

					disabledClass := ""

					if len(imgs) == 1 {
						disabledClass = "disabled-carousel"
					}

					if imgCont != "" {
						infoStr += `<div class="tl-info-container">
						<section class="carousel ` + disabledClass + `" aria-label="Gallery">
							<ol class="carousel__viewport">
							` + imgCont + `
							</ol>
						</section>
						` + proofModalStr + `
					</div>`
					}
				}
			}
		}

		tlCont += `<div class="tl-stage">
					<div class="tl-heading">
						<div class="tl-circle">
							<span class="tl-stage-icon" style="background-image: url('` + stage.Icon + `');"></span>
						</div>
						<label>` + stage.Title + `</label>
					</div>
					<div class="tl-content">
						<div class="tl-inner-wrapper">` + infoStr + `</div>
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

func (r *RURINFT) GetWordWrap(value string) string {
	fmt.Println(len(strings.Split(value, " ")))
	if len(strings.Split(value, " ")) == 1 {
		return `style="word-break: break-all;"`
	} else {
		return ""
	}
}

func (r *RURINFT) GetDecoratedKeyValueIcon(key string) string {
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
func (r *RURINFT) GenerateTabLabels(title string, index int) (string, string, string) {
	sidebarTab := `<li class="tab">
						<label for="tab` + strconv.Itoa(index+1) + `" onclick="closeSidebar('` + title + `')">
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
					<label for="tab` + strconv.Itoa(index+1) + `" onclick="setTabName('` + title + `')">
					` + title + `
						<span class="tab-arrow-icon"></span>
					</label>
				</li>`

	return mainTab, sidebarTab, radioButton
}

// Generate proof modal for key value pairs
func (r *RURINFT) GenerateProofContentStr(key string, proofInfo models.ValueWithProof) (string, string) {
	id := strings.ReplaceAll(key, " ", "") + "-modal"

	fmt.Println(proofInfo.TdpId)

	txnHash, url, err := r.GetTxnHash(proofInfo.TdpId[0])

	tab1 := proofModalCount
	tab2 := proofModalCount + 1
	tab3 := proofModalCount + 2

	proofModalCount += 3

	if err != nil {
		return "", ""
	}

	table := r.GenerateProofTable(txnHash, url, proofInfo)

	users, err1 := r.GetUsers(proofInfo.UserId)

	if err1 != nil {
		return "", ""
	}

	table2, cards := r.GenerateUsersContent(users, tab1, tab2, tab3)

	proofTick := `<span class="material-symbols-outlined provable-tick-wrapper provable-val" onclick="openModal('` + id + `')">
						check_circle
					</span>`

	proofContentStr := `<!--modal for proof-->
					<div id="` + id + `" class="modal-window">
						<div>
							<div class="modal-header">
								<h4 class="modal-heading">Proof Details Of: ` + key + `</h4>
								<span class="material-symbols-outlined modal-close" onclick="closeModal('` + id + `')">
									close
								</span>
							</div>
							<div class="modal-cont">
								<div class="modal-tab">	
									<label id="modal-tab-lbl-` + strconv.Itoa(tab1) + `" class="modal-tab-label modal-tab-label-active" onclick="openTab('modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Overview</label>
									<label id="modal-tab-lbl-` + strconv.Itoa(tab2) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Blockchain Proofs</label>
									<label  id="modal-tab-lbl-` + strconv.Itoa(tab3) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">People &amp; Technologies</label>
								</div>
								<div id="modal-tab-` + strconv.Itoa(tab1) + `" class="modal-tab-cont modal-tab-overview" style="display: flex">
									<div class="overview-card">
										<div class="header">
											<span class="material-symbols-outlined">
												visibility
											</span>
											<label>Blockchain Proofs</label>
										</div>
										<div class="body">
											<p>Visit TilliT Explorer to view transaction details and blockchain proofs.</p>
											<a href="https://qa.explorer.tillit.world/txn/` + txnHash + `" target="_blank" >
											Tillit Explorer <span class="material-symbols-outlined">
											open_in_new
											</span>
											</a>
										</div>
									</div>
									<div class="overview-card">
										<div class="header">
											<span class="material-symbols-outlined">
												account_circle
											</span>
											<label>People</label>
										</div>
										<div class="body people-content">
											` + cards + `
										</div>
									</div>
									<div class="overview-card">
										<div class="header">
											<span class="material-symbols-outlined">
												phone_android
											</span>
											<label>Technology</label>
										</div>
										<div class="body">
											<button class="overview-btn tech-btn" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">
												<img class="btn-img" src="data:image/png;base64,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" ></img>
											</button>
										</div>
									</div>
									<div class="overview-card">
										<div class="header">
											<span class="blockchain-icon">
											</span>
											<label>Blockchain Network</label>
										</div>
										<div class="body">
											<div class="bc-net-card">
												<span class="stellar-icon"></span>
												<label>Stellar</label>
											</div>
											
										</div>
									</div>
								</div>
								<div id="modal-tab-` + strconv.Itoa(tab2) + `" class="modal-tab-cont">
									<table class="table proof-table">
										<thead>
											<tr>
												<th scope="col">Proof Type</th>
												<th scope="col">Transaction ID</th>
												<th scope="col">Description</th>
												<th scope="col">Proofs</th>
											</tr>
										</thead>
										<tbody>
											` + table + `
										</tbody>
									</table>
								</div>
								<div id="modal-tab-` + strconv.Itoa(tab3) + `" class="modal-tab-cont">
									<table class="table people-and-tech-table">
										<thead>
											<tr>
												<th scope="col">Data Added By</th>
												<th scope="col">Data Source</th>
												<th scope="col">Trace Power</th>
												<th scope="col">Endorsement</th>
												<th scope="col">Badges</th>
											</tr>
										</thead>
										<tbody>
											` + table2 + `
										</tbody>
									</table>
								</div>
							</div>
						</div>
					</div>`

	return proofContentStr, proofTick
}

// Generate proof modal for image sliders
func (r *RURINFT) GenerateImgProofModalStr(proofInfo models.ValueWithProof, id string) string {
	txnHash, url, err := r.GetTxnHash(proofInfo.TdpId[0])

	if err != nil {
		return ""
	}

	tab1 := proofModalCount
	tab2 := proofModalCount + 1
	tab3 := proofModalCount + 2

	proofModalCount += 3

	table := r.GenerateProofTable(txnHash, url, proofInfo)

	users, err1 := r.GetUsers(proofInfo.UserId)

	if err1 != nil {
		return ""
	}

	table2, cards := r.GenerateUsersContent(users, tab1, tab2, tab3)

	proofContentStr := `<!--modal for proof-->
										<div id="` + id + `" class="modal-window">
											<div>
												<div class="modal-header">
													<h4 class="modal-heading">Proof Details Of: ` + `Images` + `</h4>
													<span class="material-symbols-outlined modal-close" onclick="closeModal('` + id + `')">
														close
													</span>
												</div>
												<div class="modal-cont">
													<div class="modal-tab">	
														<label id="modal-tab-lbl-` + strconv.Itoa(tab1) + `" class="modal-tab-label modal-tab-label-active" onclick="openTab('modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Overview</label>
														<label id="modal-tab-lbl-` + strconv.Itoa(tab2) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Blockchain Proofs</label>
														<label  id="modal-tab-lbl-` + strconv.Itoa(tab3) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">People &amp; Technologies</label>
													</div>
													<div id="modal-tab-` + strconv.Itoa(tab1) + `" class="modal-tab-cont modal-tab-overview" style="display: flex">
														<div class="overview-card">
															<div class="header">
																<span class="material-symbols-outlined">
																	visibility
																</span>
																<label>Blockchain Proofs</label>
															</div>
															<div class="body">
																<p>Visit TilliT Explorer to view transaction details and blockchain proofs.</p>
																<a href="https://explorer.tillit.world/txn/` + txnHash + `" target="_blank" >
																Tillit Explorer <span class="material-symbols-outlined">
																open_in_new
																</span>
																</a>
															</div>
														</div>
														<div class="overview-card">
															<div class="header">
																<span class="material-symbols-outlined">
																	account_circle
																</span>
																<label>People</label>
															</div>
															<div class="body people-content">
																` + cards + `
															</div>
														</div>
														<div class="overview-card">
															<div class="header">
																<span class="material-symbols-outlined">
																	phone_android
																</span>
																<label>Technology</label>
															</div>
															<div class="body">
																<button class="overview-btn tech-btn" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">
																	<img class="btn-img" src="data:image/png;base64,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" ></img>
																</button>
															</div>
														</div>
														<div class="overview-card">
															<div class="header">
																<span class="blockchain-icon">
																</span>
																<label>Blockchain Network</label>
															</div>
															<div class="body">
																<div class="bc-net-card">
																	<span class="stellar-icon"></span>
																	<label>Stellar</label>
																</div>
																
															</div>
														</div>
													</div>
													<div id="modal-tab-` + strconv.Itoa(tab2) + `" class="modal-tab-cont">
														<table class="table proof-table">
															<thead>
																<tr>
																	<th scope="col">Proof Type</th>
																	<th scope="col">Transaction ID</th>
																	<th scope="col">Description</th>
																	<th scope="col">Proofs</th>
																</tr>
															</thead>
															<tbody>
																` + table + `
															</tbody>
														</table>
													</div>
													<div id="modal-tab-` + strconv.Itoa(tab3) + `" class="modal-tab-cont">
														<table class="table people-and-tech-table">
															<thead>
																<tr>
																	<th scope="col">Data Added By</th>
																	<th scope="col">Data Source</th>
																	<th scope="col">Trace Power</th>
																	<th scope="col">Endorsement</th>
																	<th scope="col">Badges</th>
																</tr>
															</thead>
															<tbody>
																` + table2 + `
															</tbody>
														</table>
													</div>
												</div>
											</div>
										</div>`

	return proofContentStr
}

// Generate proof table displayed in the modal
func (r *RURINFT) GenerateProofTable(txnHash string, url string, proofInfo models.ValueWithProof) string {
	table := ""

	for _, proof := range proofInfo.Proofs {
		txnStr := ""

		if proof.Name == "" {
			continue
		}

		if len(txnHash) > 0 {
			txnStr = txnHash[0:10] + "..."
		} else {
			txnStr = "N/A"
		}

		descStyle := ""

		if len(strings.Split(proof.Description, " ")) == 1 {
			descStyle = "; word-break: break-all;"
		}

		proofName := r.GetProofName(proof.Name)

		table += `<tr>
						<td style="width : 15%">` + proofName + `</td>
						<td style="max-width : 25%; ">
							<div class="txn-wrapper">
								<a href="` + url + `" target="_blank" class="txn-hash-link">` + txnStr + `</a>
								<span class="copy-icon" onclick="copyToClipboard('` + txnHash + `')"></span>
							</div>
						</td>
						<td style="width : 40%` + descStyle + `">` + proof.Description + `</td>
						<td><a class="proof-link" href="https://qa.explorer.tillit.world/txn/` + txnHash + `" target="_blank">Proof <span class="material-symbols-outlined">
							open_in_new
							</span></a>
						</td>
					</tr>`
	}

	return table
}

// Generate user table and cards
func (r *RURINFT) GenerateUsersContent(users []models.Users, tab1 int, tab2 int, tab3 int) (string, string) {
	table := ""
	cards := ""

	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	for _, user := range users {

		userType := ""

		splitType := re.FindAllString(user.Type, -1)

		for _, s := range splitType {
			userType += s + " "

		}

		table += `<tr>
						<td style="width : 25%">` + user.FirstName + ` ` + user.LastName + ` <br/> ` + userType + `</td>
						<td style="max-width : 20%; ">
							<img class="tbl-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEEAAABwCAYAAABB0R1NAAAABHNCSVQICAgIfAhkiAAACF1JREFUeJztnVtsFNcZx39nd9bg9dqsMWAIqNjY63IJ2BFJGgiITQmUm5QtOEhVJWKqtmpwHkpbCH2gMW0fSmgEDwmR0oeWqlGqxKFExQZqR6wpt5aYrBNhaAHblAV2ob57d33Z9fTBsmPwjL32zoyBzv/Jey5z/vPTzJnPM9/MEYwg9xpPls1ifUkWwiNknAgKRuoz7pLxyYIWIctHenpjn3iPH2kYrrlQq3hxQ6EbmTcEuLV3aaxk8CLYU3m01KtUrwjhxfWF+wX8WF9rxkuGA5VlpdsfLL8Pgtvtcdrs0slH4pAfq2R8PeHoC17vkZb+Isvg+sceAICgwGaXTg4usvb/sWrdpgMI4THe1ThIMD3HNS+97url430/6ZsEhcxJtT7Tpk5ho2cD+QsXkDMnyyirY9b1ugZqvrzEH9//kHA4rNpOFrxQebTUawWY45r/ewGKe7dq5Qre+s0vmTc3j8npTp1sa6vJ6U7mzc1jw9rVNDU3U1d/Q61pVt3V2kPCvcaTZbNK9UotVq1cwY7tr+nn1iDt2/82FZ9WKdb1xKLZFskqKc4D06ZOeSwAALz6g63Y7XbFOskqeSQh41GKFjZ6NmhqpD3WQUcsFHd7hzWFVKtDk7EdjhQ2vrSeP33w0ZA6IeORgElKHfMXLtDEQHusg4qmkzT2NI+674ykTFZnvECSSErYx/NLnlWEAEyyqMUFWl0FzrVeGBMAgDvdQc62XtDEh+r+CAosyjXa6UbnzcT6R/6jkRN16Q7BYU1JqH+SJfFTYSTpDiHPnptQ/ycd8zVyoi5J7wEWp+aTanXw7/C1Ufd90jGPrIlf08HV/dIdAkCePYc8e44RQ41Jup8Oj4JMCJgQABMCYEIATAiACQEwIQAmBMCEAJgQABMCYEIATAiACQEwIQAmBMCEAJgQAB3vMXZEQ5zwV3E68E9qmmpH3T8zeSq5aVksy3yGb83SN21KFwi+xkvsrt5HKKqeGzCSgpF7BCP3OBO8QGlDOa8v2kZumj65EZqfDr7GS/zkH3sSAvCgrrc1sP18CYHIXc22OViaQuiIhthdvU/LTQ4oFA1z3K+YgZewNIVQWl+m6REwWFtchRS5NuuybU3nhDMBbZ4gP6idi7axRsfJUVMI19tVc4PGpBTJzv7nSnSbEPtlWJyweuaKUbXPSZ1tCAAw4FlkimTnV4t3UJCxAIcthcMN5SP2yUmdzf4lJTikxB7rxytdj4Sc1Nl88M13KMjoS/0pynuZnNTZw/ZZPXMFv1u+TxGAPxKk6MwezX3qBkFpZxxSCq/nF5MiKWeSbXEVsiu/WLHOHwny/LFtnLt3WXOvukEIRu4pluemZfGK6+Uh5TsXbVO9BFYFqnmu/Ee09HRp6rFfukGoaapl7xcHFesKs9ezNPNpoG/OeG/Zm6qXwKP+02w8VUJ7tEcvq/rOCSf8XtUob1d+MUsznx72CvDuvz7klbN7icqynjb1vzq8+cVBctOyhuyoQ0rh14t3qvbbVf027147rrc9wKA4Yfv5Ejqi8Wezbj1TYhgAMAhCKBpm+/k9cYFYWVHMYf9nBrj6SoZFjNfbGnin9pBqfUc0xKK/fpfPmrQNveORobfXTvi9lNaXDSn3R4I8U/Z9boRbjbQzIE0hjBQNAhy8fAhf46WB3xcba3m27Ifc7myPa4wpExXz0ROSphDyM+LLjN9dvY9A5C5VgWo2nPw5oVgs7jGWT104Vnuq0hRCYfa6uNqFomGKz/2Cb596Y1QAki1WXps/NNpMVJpCmJ48jS2uwrjaNnc2MSPJNqrt/3T+JmYlZ47F2rDSfGIscm2O+95BZpKNyZJ15IbAkim57FiwJRFrqtLl6rArvzjuI2LWRBvJFtVXtrEKeDV3DcdXHtDK3hDpFjYXuTazZpab04ELnAkOf+/x65O6+Pt/h76Q587MZ3fB93Q5BQZL1/8dpidPozB7PYXZ6/UcJmGZzyIxIQAmBMCEAJgQABMCYEIATAiACQEwIQAmBMCEAJgQABMCYEIATAiACQEwIQAmBMCEAJgQABMCYEIATAiACQEY5glUWePfjPQxrlKFcLsrYKSPcZV5OmDQZ8fUNC1pCpKQCMcitETHJ2kLxgmCyz6HhSkLyLClD5Td6rpNdXsNwW7lxHA9ZfjpkGfPxe1cdh8AgJkTnmBtxiom24z/LLKhEBxWB8smfUO13iYk3M7lBjrqk6EQcu3ZWMXwOUoZtnQmP3CU6C1DIdgtyXG1S5eMPSUMhdDTG9+LG1E5qrOT+2UohEBPfO86B7qDOjv5SrJMqwUZxY+c3/pc+4jxZuct7owQiVa319DV26352Gr7I8BnQciKi8Lc8t3R3AhAZXMVLdE2xbrrkXouttfoMq7q/gi5wZrjmudUWtyi1d/GzIIZTJw0QVMzUTlGbegKzdFmJliSCMXC3O25x6mWs1wKXdF0rH613Gzj4vtfKtb19rJHuN0ep2SXGoQY+k1356w0VvxsKUn20eUgP0zqDvdQ9duztPiHHn2yTGtleanT2tBwpXOOa/5EIYYuddTZ1kVdVQOp01NJm6HNF/aN1K3PA1S9dZZwU0Styd66q7VeAX2r/qgdDf2yZyTzRMF0Jtj1/6x4ouoKd3PbFyDcqLrzyLJ8IxqOFXi9R1oGMqtXrt3ksVjEXwxx+TBIFk9VlH/kg0Gr/9Rfu3wlxzX3xv/FCkBy79aK8o8H3jm8L5Cvu3rZ9ziDkGVaZVn+TuWxw38eXD4kYqwoP/wHZPGUWhD1yEqmSiDcnx77+MiDVepvWwCr1m0sAksRgtF9BuMhkizzCRYOqC2OByNA6Jfb7XFak61uIcQjs2SaLMu+WCTmHbwGnJr+B36gc8CxSWd6AAAAAElFTkSuQmCC" ></img>
						</td>
						<td style="width : 20%;">N/A</td>
						<td style="width : 20%;">
							N/A
						</td>
						<td style="width : 15%;">
							N/A
						</td>
					</tr>`

		cards += `<button class="overview-btn people-btn" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">
						<img class="btn-img" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR4lq60B8QUO-IJ7Ocv-5Nn77keApow2URygWAZRwdxYltUVb8jlQ" ></img>
						<div class="people-info">
							<label>` + user.FirstName + ` ` + user.LastName + `</label>
							<label>` + userType + `</label>
						</div>
						<span class="material-symbols-outlined">
							chevron_right
						</span>
					</button>`
	}

	return table, cards
}

// Get proof name
func (r *RURINFT) GetProofName(proof string) string {
	switch strings.ToLower(proof) {
	case "poe":
		return "Proof of Existence"
	case "pog":
		return "Proof of Genesis"
	case "poc":
		return "Proof of Continuity"
	case "poac":
		return "Proof of Authorize Change"
	default:
		return proof
	}
}

// Get transaction hash for tdp
func (r *RURINFT) GetTxnHash(tdpID string) (string, string, error) {
	txnHash := ""
	stellarUrl := ""

	val, exists := txnMap[tdpID]

	if exists {
		return val[0], val[1], nil
	}

	gatewayUrl := configs.GetGatewayUrl()

	url := gatewayUrl + `/GetTransactions?txn=` + tdpID + `&page=1&perPage=10`
	var txnResp []models.TxnResp

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.ErrorLogger.Println("unable to get data :", err.Error())
		return txnHash, stellarUrl, err
	}

	client := &http.Client{}
	resp, err1 := client.Do(req)

	if err1 != nil {
		logs.ErrorLogger.Println("unable to get data :", err.Error())
		return txnHash, stellarUrl, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(string(body)), &txnResp)

	if len(txnResp) > 0 {
		txnHash = txnResp[0].TxnHash
		stellarUrl = txnResp[0].URL
	}

	txnMap[tdpID] = append(txnMap[tdpID], txnHash)
	txnMap[tdpID] = append(txnMap[tdpID], stellarUrl)

	return txnHash, stellarUrl, nil
}

// Get user details from admin backend
func (r *RURINFT) GetUsers(userID []string) ([]models.Users, error) {
	adminBackend := configs.GetAdminBackendUrl()
	url := adminBackend + `/sign/getUsersDetails`
	var users []models.Users

	values := map[string][]string{"ids": userID}

	jsonValue, err1 := json.Marshal(values)

	if err1 != nil {
		logs.ErrorLogger.Println("unable to get data :", err1.Error())
		return users, err1
	}

	req, err2 := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	if err2 != nil {
		logs.ErrorLogger.Println("unable to get data :", err2.Error())
		return users, err2
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err3 := client.Do(req)

	if err3 != nil {
		logs.ErrorLogger.Println("unable to get data :", err3.Error())
		return users, err3
	}

	body, err3 := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(string(body)), &users)

	return users, nil
}
