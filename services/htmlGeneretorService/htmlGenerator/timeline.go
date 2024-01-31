package htmlGenerator

import (
	//"fmt"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/services/mapGenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/mitchellh/mapstructure"
)

type JMACNFT struct {
	BatchID   string
	ProductID string
	ItemName  string
}

var (
	svgStart        = services.ReadFromFile("services/htmlGeneretorService/templates/timelinetemplate/svgNFTHeader.txt")
	svgEnd          = services.ReadFromFile("services/htmlGeneretorService/templates/timelinetemplate/svgNFTFooter.txt")
	styling         = services.ReadFromFile("services/htmlGeneretorService/templates/timelinetemplate/svgNFTStyles.css")
	styleStart1     = `<style>`
	styleEnd        = `</style></head>`
	htmlBody        = "</div></div>"
	mapRepository   customizedNFTrepository.MapRepository
	backendUrl      = configs.GetNftBackendBaseUrl()
	proofModalCount = 0
	txnMap          = make(map[string][]string)
	svgRepository   customizedNFTrepository.SvgRepository
	// filebaseBucket  = os.Getenv("FILEBASE_BUCKET")
)

func (r *JMACNFT) GenerateNFT() (string, error) {
	tdpData, err := customizedNFTFacade.GetDigitalTwinData(r.BatchID, r.ProductID)
	if err != nil {
		return "", fmt.Errorf("Failed to get digital twin data: %v", err)
	}
	if len(tdpData) == 0 {
		return "", fmt.Errorf("no digital twin data found for BatchID: %s, ProductID: %s", r.BatchID, r.ProductID)
	}
	r.ItemName = tdpData[0].Item
	// Svg will be generated using the template
	svgrst, _, svgGenErr := r.GenerateSVGTemplateforNFT(tdpData)
	if svgGenErr != nil {
		logs.InfoLogger.Println("Failed to generate SVG : ", svgGenErr.Error())
		return "", svgGenErr
	}

	return svgrst, nil
}

func (r *JMACNFT) GenerateSVGTemplateforNFT(data []models.Component) (string, string, error) {
	/* batchID := r.BatchID
	productID := r.ProductID  */
	//shopID := r.ShopID

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

	htmlStart := `<div class="base-layer"><body><div class="nft-header default-font layer-1">
						<div class="nft-header-content cont-wrapper d-flex justify-content-between">
							<div class="c1">
								<label id="nftName" class="nftName">` + r.ItemName + `</label>
							</div>
							<div class="c2">
							<label id="nftName" class="nftName">` + r.BatchID + `</label>
							</div>
						</div>
					</div>
					<div class="d-flex justify-content-center align-content-center flex-wrap layer-2" id="container"><div class="proof-toggle-wrapper cont-wrapper">`

	r.GenerateContent(data)

	template := svgStart + styleStart1 + styling + styleEnd + htmlStart + htmlBody + svgEnd
	htmlBody = ""
	txnMap = make(map[string][]string)
	/* template = strings.(template)
	fmt.Println(template) */
	/* template = strings.Replace(template, "\r", " ", -1)
	template = strings.Replace(template, "\t", " ", -1)
	template = strings.Replace(template, "\n", " ", -1) */
	return template, "", nil
}

func (r *JMACNFT) GenerateContent(data []models.Component) {
	for _, data := range data {
		if data.Component == "expandableTab" && len(data.VerticalTab) > 0 {
			r.GenerateVerticalTabs(data)
		}
	}
}

// generate vertical tab component
func (r *JMACNFT) GenerateVerticalTabs(data models.Component) {
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

	htmlBody += `<div class="widget-div cont-wrapper">
					<div class="wrap-collabsible">
						<div class="collapsible-content-none">
							<div class="content-inner">
								<div class="tabbed">
									<div style="display: flex; flex-direction : column">
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
func (r *JMACNFT) GenerateOverview(data models.Component) (string, string, string, string) {
	mainTabs := ""
	sidebarTabs := ""
	radioButtons := ""
	content := ""
	isHasMapAndTimeline := hasMapAndTimeline(data.Children)

	for index, tab := range data.Children {
		// if tab.Component == "vertical-card-container" {
		// 	cont, mainTbs, sidebarTbs, radioBtns := r.GenerateVerticalCardContainer(tab, index)
		// 	content += cont
		// 	mainTabs += mainTbs
		// 	sidebarTabs += sidebarTbs
		// 	radioButtons += radioBtns
		// // } else
		if tab.Component == "map" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateJourneyMap(tab, index)
			content += cont
			if isHasMapAndTimeline {
				mainTabs += mainTbs
			}
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		} else if tab.Component == "timeline" {
			cont, mainTbs, sidebarTbs, radioBtns := r.GenerateTimeline(tab, index)
			content += cont
			if isHasMapAndTimeline {
				mainTabs += mainTbs
			}
			sidebarTabs += sidebarTbs
			radioButtons += radioBtns
		}
	}

	proofToggle := `<div class="proof-toggle-btn"><div class="proof-toggle-wrapper cont-wrapper">
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
</div></div>`
	content += proofToggle
	return content, mainTabs, sidebarTabs, radioButtons
}

// Generate vertical card container
func (r *JMACNFT) GenerateVerticalCardContainer(data models.Component, index int) (string, string, string, string) {
	content := ""
	mainTab, sidebarTab, radioButton := r.GenerateTabLabels(data.Title, index)

	for i, childComponent := range data.Children {
		if childComponent.Component == "image-slider" {
			res := r.GenerateImageSlider(childComponent, index, data.Title)
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

func (r *JMACNFT) GenerateDecoratedKeyValuesHeading(data models.Component, cards string) string {
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
func (r *JMACNFT) GenerateDecoratedKeyValues(data models.Component, index int) string {
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
			// img := ""
			val := "No Records"
			var decoratedVal models.ValueWithProof

			decodeErr := mapstructure.Decode(child.Value, &decoratedVal)
			if decodeErr != nil {
				logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
			}
			// keyValIcon := GetDecoratedKeyValueIcon(child.Key)

			/* if child.Icon != "" {
				img = `<img class="dt-icon-img" src="` + child.Icon + `" />`
			} */
			if decoratedVal.Value != "" {
				val = decoratedVal.Value.(string)
			}

			proofContentStr := ""
			proofTick := ""
			if decoratedVal.Provable && len(decoratedVal.TdpId) > 0 {
				proofContentStr, proofTick = r.GenerateProofContentStr(child.Key, decoratedVal.TdpId[0])
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
func (r *JMACNFT) GenerateImageSlider(imageSlider models.Component, parentIndex int, title string) string {
	content := ""

	for i, image := range imageSlider.Images.Value {
		if image.Img == "" {
			content += `<p>No Records</p>`
		} else {
			// imgKey := title + "_" + strconv.Itoa(i) // create image key

			imgUrl := r.fetchImgURL(image.Img) // upload image to ipfs and get the url

			content += `<div class="img-wrapper">
			<div id="cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `" class="img-div"
				style="background-image: url('` + imgUrl + `');">
			</div>
			<label onclick="openFullScreenImg('cert` + strconv.Itoa(parentIndex) + strconv.Itoa(i+1) + `')" title="View Image">
				<span class="zoom-icon"></span>
			</label>
		</div>`
		}
	}

	return content
}

// Generate proof modal for journey map
// func (r *JMACNFT) GenerateMapProofContentStr(key string, proofInfo models.ValueWithProof) (string, string) {
// 	id := strings.ReplaceAll(key, " ", "") + "-map-modal"

// 	var valueWithProof models.ValueWithProof

// 	valueWithProof.Proofs = proofInfo.Proofs
// 	valueWithProof.Provable = proofInfo.Provable

// 	var coordinates []models.CoordinateValue
// 	decodeErr := mapstructure.Decode(proofInfo.Value, &coordinates)
// 	if decodeErr != nil {
// 		logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
// 	}

// 	valueWithProof.TdpId = append(valueWithProof.TdpId, coordinates[0].TdpId[0])
// 	valueWithProof.UserId = append(valueWithProof.UserId, coordinates[0].UserId[0])

// 	txnHash, url, err := r.GetTxnHash(valueWithProof.TdpId[0])

// 	tab1 := proofModalCount
// 	tab2 := proofModalCount + 1
// 	tab3 := proofModalCount + 2

// 	proofModalCount += 3

// 	if err != nil {
// 		return "", ""
// 	}

// 	table := r.GenerateProofTable(txnHash, url, valueWithProof)

// 	users, err1 := r.GetUsers(valueWithProof.UserId)

// 	if err1 != nil {
// 		return "", ""
// 	}

// 	table2, cards := r.GenerateUsersContent(users, tab1, tab2, tab3)

// 	proofTick := `<span class="material-symbols-outlined provable-tick-wrapper provable-val" onclick="openModal('` + id + `')">
// 						check_circle
// 					</span>`

// 	proofContentStr := `<!--modal for proof-->
// 					<div id="` + id + `" class="modal-window">
// 						<div>
// 							<div class="modal-header">
// 								<h4 class="modal-heading">Proof Details Of: ` + key + `</h4>
// 								<span class="material-symbols-outlined modal-close" onclick="closeModal('` + id + `')">
// 									close
// 								</span>
// 							</div>
// 							<div class="modal-cont">
// 								<div class="modal-tab">
// 									<label id="modal-tab-lbl-` + strconv.Itoa(tab1) + `" class="modal-tab-label modal-tab-label-active" onclick="openTab('modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Overview</label>
// 									<label id="modal-tab-lbl-` + strconv.Itoa(tab2) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `')">Blockchain Proofs</label>
// 									<label  id="modal-tab-lbl-` + strconv.Itoa(tab3) + `" class="modal-tab-label" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">People &amp; Technologies</label>
// 								</div>
// 								<div id="modal-tab-` + strconv.Itoa(tab1) + `" class="modal-tab-cont modal-tab-overview" style="display: flex">
// 									<div class="overview-card">
// 										<div class="header">
// 											<span class="material-symbols-outlined">
// 												visibility
// 											</span>
// 											<label>Blockchain Proofs</label>
// 										</div>
// 										<div class="body">
// 											<p>Visit TilliT Explorer to view transaction details and blockchain proofs.</p>
// 											<a href="` + configs.GetTillitUrl() + `/txn/` + txnHash + `" target="_blank" >
// 											Tillit Explorer <span class="material-symbols-outlined">
// 											open_in_new
// 											</span>
// 											</a>
// 										</div>
// 									</div>
// 									<div class="overview-card">
// 										<div class="header">
// 											<span class="material-symbols-outlined">
// 												account_circle
// 											</span>
// 											<label>People</label>
// 										</div>
// 										<div class="body people-content">
// 											` + cards + `
// 										</div>
// 									</div>
// 									<div class="overview-card">
// 										<div class="header">
// 											<span class="material-symbols-outlined">
// 												phone_android
// 											</span>
// 											<label>Technology</label>
// 										</div>
// 										<div class="body">
// 											<button class="overview-btn tech-btn" onclick="openTab('modal-tab-` + strconv.Itoa(tab3) + `', 'modal-tab-` + strconv.Itoa(tab2) + `', 'modal-tab-` + strconv.Itoa(tab1) + `', 'modal-tab-lbl-` + strconv.Itoa(tab3) + `', 'modal-tab-lbl-` + strconv.Itoa(tab2) + `', 'modal-tab-lbl-` + strconv.Itoa(tab1) + `')">
// 												<img class="btn-img" src="data:image/png;base64,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" ></img>
// 											</button>
// 										</div>
// 									</div>
// 									<div class="overview-card">
// 										<div class="header">
// 											<span class="blockchain-icon">
// 											</span>
// 											<label>Blockchain Network</label>
// 										</div>
// 										<div class="body">
// 											<div class="bc-net-card">
// 												<span class="stellar-icon"></span>
// 												<label>Stellar</label>
// 											</div>

// 										</div>
// 									</div>
// 								</div>
// 								<div id="modal-tab-` + strconv.Itoa(tab2) + `" class="modal-tab-cont">
// 									<table class="table proof-table">
// 										<thead>
// 											<tr>
// 												<th scope="col">Proof Type</th>
// 												<th scope="col">Transaction ID</th>
// 												<th scope="col">Description</th>
// 												<th scope="col">Proofs</th>
// 											</tr>
// 										</thead>
// 										<tbody>
// 											` + table + `
// 										</tbody>
// 									</table>
// 								</div>
// 								<div id="modal-tab-` + strconv.Itoa(tab3) + `" class="modal-tab-cont">
// 									<table class="table people-and-tech-table">
// 										<thead>
// 											<tr>
// 												<th scope="col">Data Added By</th>
// 												<th scope="col">Data Source</th>
// 												<th scope="col">Trace Power</th>
// 												<th scope="col">Endorsement</th>
// 												<th scope="col">Badges</th>
// 											</tr>
// 										</thead>
// 										<tbody>
// 											` + table2 + `
// 										</tbody>
// 									</table>
// 								</div>
// 							</div>
// 						</div>
// 					</div>`

// 	return proofContentStr, proofTick
// }

// Generate Journey Map
func (r *JMACNFT) GenerateJourneyMap(tab models.Component, index int) (string, string, string, string) {
	var mapInfo []models.MapInfo
	proofCards := `<div class="map-proof-cont">`

	for index, c := range tab.Coordinates {
		var coordinates []models.CoordinateValue
		decodeErr := mapstructure.Decode(c.Values.Value, &coordinates)
		if decodeErr != nil {
			logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
		}

		proofContentStr := ""
		proofTick := ""

		if c.Values.Provable && len(coordinates) > 0 {
			// key := c.Title

			// proofContentStr, proofTick = r.GenerateMapProofContentStr(key, c.Values)
		}

		lat := "0"
		long := "0"

		if len(coordinates) > 0 {
			lat = strconv.FormatFloat(coordinates[0].Lat, 'g', 7, 64)
			long = strconv.FormatFloat(coordinates[0].Lng, 'g', 7, 64)
		}

		cityName, err := mapGenerator.GetCityName(lat, long)
		if err != nil {
			logs.ErrorLogger.Println("failed to get city name : ", err.Error())
		}

		proofCards += `	<div class="map-proof-card">
									<div class="map-proof-title-cont">
										<span class="marker-icon"></span>
										<label class="map-proof-title">` + strconv.Itoa(index+1) + ` - ` + c.Title + `</label>
									</div>
									<label class="map-proof-text">` + cityName + `</label>
									<div class="map-proof-title-cont">
										<label class="map-proof-text">` + lat + `, ` + long + `</label>
										` + proofTick + `
									</div>
									` + proofContentStr + `
								</div>`

		var cmap models.MapInfo
		cmap.Title = ""
		cmap.Latitude = 0
		cmap.Longitude = 0

		if len(coordinates) > 0 {
			cmap.Title = coordinates[0].Name
			cmap.Latitude = coordinates[0].Lat
			cmap.Longitude = coordinates[0].Lng
		}
		mapInfo = append(mapInfo, cmap)

		/* for _, c1 := range coordinates {
			coordinate := c1
			var cmap models.MapInfo
			cmap.Title = coordinate.Name
			cmap.Latitude = coordinate.Lat
			cmap.Longitude = coordinate.Long
			mapInfo = append(mapInfo, cmap)
		} */

	}

	/* mapInfo = append(mapInfo, models.MapInfo{6.927079, 79.861244})
	mapInfo = append(mapInfo, models.MapInfo{35.652832, 139.839478}) */

	generatedMap := mapGenerator.GenerateMap(mapInfo)
	var newMap models.GeneratedMap
	newMap.MapTemplate = generatedMap
	rst, mapSaveErr := mapRepository.SaveMap(newMap)

	proofCards += `</div>`

	if mapSaveErr != nil {
		logs.ErrorLogger.Println("Failed to save map : ", mapSaveErr.Error())
	}

	content := `<div class="tab-content">
					<embed class="map" frameborder="0" scrolling="no" marginheight="0"
						marginwidth="0"
						src="` + backendUrl + `/GetMap/` + rst + `"></embed>
						` + proofCards + `

				</div>`

	mainTab, sidebarTab, radioButton := r.GenerateTabLabels("Journey", index)

	return content, mainTab, sidebarTab, radioButton
}

// Generate Timeline
func (r *JMACNFT) GenerateTimeline(data models.Component, index int) (string, string, string, string) {
	content := ``
	tlCont := ""
	imgSliderCount := 0

	mainTab, sidebarTab, radioButton := r.GenerateTabLabels(data.Title, index)

	for i, stage := range data.Children {
		infoStr := ""
		if stage.Icon == "" {
			stage.Icon = "https://s3.ap-south-1.amazonaws.com/nft.tracified.com/assets/icons/common-stage.png"
		}
		for j, info := range stage.Children {
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
					proofContentStr, proofTick = r.GenerateProofContentStr(info.Key, decoratedVal.TdpId[0])
				}

				infoStr += `<div class="tl-info-container tl-key-value">
								<label class="grey-text">` + strings.Replace(info.Key, "&", "&amp;", -1) + `</label>
								<label class="tl-bold-text">` + val + " " + proofTick + `</label>
							</div>` + proofContentStr

			} else if info.Component == "image-slider" {
				imgSliderId := "slider_" + strconv.Itoa(i) + "_" + strconv.Itoa(j)
				imgCont := ""
				var imgs []models.ImageValue
				decodeErr := mapstructure.Decode(info.Slides.Value, &imgs)
				if decodeErr != nil {
					logs.ErrorLogger.Println("failed to decode map : ", decodeErr.Error())
				}

				proofModalStr := ""

				if len(imgs) > 0 {

					for j, image := range imgs {
						proofTickIcon := ""
						if info.Slides.Provable && len(info.Slides.TdpId) > 0 {
							id := "slider" + strconv.Itoa(imgSliderCount) + "-modal"
							proofTickIcon = `<span class="material-symbols-outlined provable-tick-wrapper provable-val" onclick="openModal('` + id + `')">
																check_circle
															</span>
											`
							proofModalStr += r.GenerateImgProofModalStr(info.Slides.TdpId[j], id)
							imgSliderCount++
						}

						dateStr := strings.ReplaceAll(strings.Split(image.Time, "T")[0], "-", "/")

						imgUrl := image.Img

						imgCont += `<li id="carousel__slide` + strconv.Itoa(i) + strconv.Itoa(j) + `"
							tabindex="0" class="carousel__slide">										

							<div>
								<div class="row pb-2"> <label class="image-text-field">` + image.FieldName + `</label> </div>
							
								<div class="row flex justify-content-center">
									<img class="carosal-img" id="img` + strconv.Itoa(i) + strconv.Itoa(j) + `"
									onclick="openFullScreenImg('img` + strconv.Itoa(i) + strconv.Itoa(j) + `')"
									src="` + imgUrl + `"> 
								</div>
						
								<div class="row d-grid justify-content-center pt-1">` + proofTickIcon + `</div>																									
								<div class="row"> <label class="date-text-field">` + dateStr + `</label> </div>
								<div class="row"> <label class="comment-text-field">` + image.Comment + `</label> </div>
							</div>
						</li>`

					}

					disabledClass := ""

					if len(imgs) == 1 {
						disabledClass = "disabled-carousel"
					}

					if imgCont != "" {
						if len(imgs) == 1 {
							infoStr += `<div class="tl-info-container">
							<section class="carousel ` + disabledClass + `" aria-label="Gallery">
		
							<div class=" arrow-right" id="right ` + imgSliderId + `" onclick="moveRight('` + imgSliderId + `','left ` + imgSliderId + `','right ` + imgSliderId + `')">
							</div>
							<div class=" arrow-left" id="left ` + imgSliderId + `" onclick="moveLeft('` + imgSliderId + `','left ` + imgSliderId + `','right ` + imgSliderId + `')">
							</div>
	
								<ol id="` + imgSliderId + `" class="carousel__viewport">
								` + imgCont + `
								</ol>
							</section>
							` + proofModalStr + `
						</div>`
						} else {
							infoStr += `<div class="tl-info-container">
							<section class="carousel ` + disabledClass + `" aria-label="Gallery">
		
							<div class=" arrow-right" id="right ` + imgSliderId + `" onclick="moveRight('` + imgSliderId + `','left ` + imgSliderId + `','right ` + imgSliderId + `')">
								<span class="material-symbols-outlined icon-container"> chevron_right</span>
							</div>
							<div class=" arrow-left" id="left ` + imgSliderId + `" onclick="moveLeft('` + imgSliderId + `','left ` + imgSliderId + `','right ` + imgSliderId + `')">
							   <span class="material-symbols-outlined icon-container">chevron_left</span>
							</div>
	
								<ol id="` + imgSliderId + `" class="carousel__viewport">
								` + imgCont + `
								</ol>
							</section>
							` + proofModalStr + `
						</div>`
						}
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

func (r *JMACNFT) GetWordWrap(value string) string {
	if len(strings.Split(value, " ")) == 1 {
		return `style="word-break: break-all;"`
	} else {
		return ""
	}
}

func (r *JMACNFT) GetDecoratedKeyValueIcon(key string) string {
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
func (r *JMACNFT) GenerateTabLabels(title string, index int) (string, string, string) {
	sidebarTab := `<li class="tab">
						<label for="tab` + strconv.Itoa(index+1) + `" onclick="closeSidebar('` + title + `')">
							` + title + `	
						</label>
					</li>`
	checked := ""

	if index == 0 {
		checked = `checked="true"`
	}

	radioButton := `<input type="radio" id="tab` + strconv.Itoa(index+1) + `" name="css-tabs" ` + checked + ` checked></input>`

	mainTab := `<li class="tab">
					<label for="tab` + strconv.Itoa(index+1) + `" onclick="setTabName('` + title + `')">
					` + title + `
					</label>
				</li>`

	return mainTab, sidebarTab, radioButton
}

// Generate proof modal for key value pairs
func (r *JMACNFT) GenerateProofContentStr(key, tdpId string) (string, string) {
	id := strings.ReplaceAll(key, " ", "") + "-modal"
	// txnHash, url, err := r.GetTxnHash(proofInfo.TdpId[0])
	// tab1 := proofModalCount
	// tab2 := proofModalCount + 1
	// tab3 := proofModalCount + 2
	proofModalCount += 3

	table := r.GenerateProofTable(tdpId, "")

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
							<table class="table proof-table">
										<thead>
											<tr>
												<th scope="col">Type</th>
												<th scope="col">Description</th>
												<th scope="col">Proofs</th>
											</tr>
										</thead>
										<tbody>
											` + table + `
										</tbody>
									</table>
							</div>
						</div>
					</div>`

	return proofContentStr, proofTick
}

// Generate proof modal for image sliders
func (r *JMACNFT) GenerateImgProofModalStr(tdpId string, id string) string {
	// tab1 := proofModalCount
	// tab2 := proofModalCount + 1
	// tab3 := proofModalCount + 2

	proofModalCount += 3

	table := r.GenerateProofTable(tdpId, "")

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
												<table class="table proof-table">
															<thead>
																<tr>
																	<th scope="col">Type</th>
																	<th scope="col">Description</th>
																	<th scope="col">Proofs</th>
																</tr>
															</thead>
															<tbody>
																` + table + `
															</tbody>
														</table>
												</div>
											</div>
										</div>`

	return proofContentStr
}

// Generate proof table displayed in the modal
func (r *JMACNFT) GenerateProofTable(tdpid string, url string) string {
	table := ""
	descStyle := ""
	proofName := r.GetProofName("POE")
	table += `<tr>
				<td style="width : 30%">` + proofName + `</td>
				<td style="width : 40%` + descStyle + `">` + "Proof that a given data packet existed in the blockchain" + `</td>
				<td><a class="proof-link" href="` + configs.GetTillitUrl() + `/search/` + tdpid + `" target="_blank">Proof <span class="material-symbols-outlined">
					open_in_new
				</span></a>
				</td>
			</tr>`
	return table
}

// Get proof name
func (r *JMACNFT) GetProofName(proof string) string {
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
func (r *JMACNFT) GetTxnHash(tdpID string) (string, string, error) {
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

func (r *JMACNFT) fetchImgURL(img string) string {
	streamingAPIUrl := configs.GetStreamingAPIUrl()
	b64 := base64.StdEncoding.EncodeToString([]byte(img))
	url := streamingAPIUrl + b64
	return url
}

func (r *JMACNFT) fetchImg(img string) ([]byte, string) {
	streamingAPIUrl := configs.GetStreamingAPIUrl()

	b64 := base64.StdEncoding.EncodeToString([]byte(img))
	url := streamingAPIUrl + b64

	res, err := http.Get(url)
	if err != nil {
		logs.ErrorLogger.Println("unable to get data :", err.Error())
	}

	fType := strings.Split(res.Header.Get("Content-Type"), "/")[1]

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.ErrorLogger.Println("ioutil.ReadAll :", err)
	}

	res.Body.Close()

	return data, fType
}

// Check for the presence of both "map" and "timeline" components
func hasMapAndTimeline(components []models.Component) bool {
	hasMap := false
	hasTimeline := false

	for _, comp := range components {
		if comp.Component == "map" {
			hasMap = true
		} else if comp.Component == "timeline" {
			hasTimeline = true
		}

		// Recursively check in the VerticalTab and Tabs fields
		hasMap = hasMap || hasMapAndTimeline(comp.VerticalTab)
		hasTimeline = hasTimeline || hasMapAndTimeline(comp.Tabs)
	}

	return hasMap && hasTimeline
}
