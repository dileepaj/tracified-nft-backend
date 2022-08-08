package htmlGenerator

import (
	"encoding/json"
	"strconv"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	documentStart = services.ReadFromFile("services/htmlGeneretorService/templates/htmlHeader.txt")
	documentEnd   = `</html>`
	headeEnd      = `</header>`
	styleStart    = services.ReadFromFile("services/htmlGeneretorService/templates/htmlStyles.css")
	mainHandler   = services.ReadFromFile("services/htmlGeneretorService/templates/htmlScript.html")
	stratScript   = `<script type="text/javascript">`
	endScript     = `</script>`
)

/**
*	Generate complete  HTML template for NFT (css,javascript)
* @function GenerateNFTTemplate
*
**/
func GenerateHTMLTemplate(htmlData models.HtmlGenerator) (string, error) {
	// Parse the Data
	var jsScripts string
	var barcharts []models.Chart = htmlData.NftContent.BarCharts
	var piecharts []models.Chart = htmlData.NftContent.PieCharts
	var tables []models.Table = htmlData.NftContent.Tables
	var stats []models.StataArray = htmlData.NftContent.Stats
	var proofbot []models.ProofBotData = htmlData.NftContent.ProofBot
	var bubbleCharts []models.Chart = htmlData.NftContent.BubbleCharts
	var images []models.ImageData = htmlData.NftContent.Images
	var Timelines []models.Timeline = htmlData.NftContent.TimeLine
	var contentOrderData []models.ContentOrderData = htmlData.ContentOrderData
	body := ` <body style="font-family: 'Inter'; color: #021D27">
				 <div class="nft-header">
				 <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/Tracified-NFT-v5.png" class="nft-logo"/>
				 <label>` + htmlData.NFTName + `</label>
				 </div>
				 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
				</div>
				<div class="text-center nft-footer ">
					<div class="nft-footer-content">
						<label>Powered by </label> <img src="https://tracified-platform-images.s3.ap-south-1.amazonaws.com/tracified-logo+(1).png" class="logo"/>
					</div>
				</div>
			  </body>`
	// take json data convert it to string
	dataString, err := json.Marshal(htmlData)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
	if len(contentOrderData) != 0 {
		for _, element := range contentOrderData {
			if element.Type == "BarChart" {
				if len(barcharts) != 0 {
					for i, bar := range barcharts {
						if len(bar.ChartData) != 0 && element.WidgetId == bar.WidgetId {
							jsScripts += `
		displayBarchart(data.NftContent.BarCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "PieChart" {
				if len(piecharts) != 0 {
					for i, pie := range piecharts {
						if len(pie.ChartData) != 0 && element.WidgetId == pie.WidgetId {
							jsScripts += `
		displayPiechart(data.NftContent.PieCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "BubbleChart" {
				if len(bubbleCharts) != 0 {
					for i, bubble := range bubbleCharts {
						if len(bubble.ChartData) != 0 && element.WidgetId == bubble.WidgetId {
							jsScripts += `
		displayBubblechart(data.NftContent.BubbleCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Table" {
				if len(tables) != 0 {
					for i, table := range tables {
						if table.TableContent != "" && table.TableContent != "EMPTY" && element.WidgetId == table.WidgetId {
							jsScripts += `
		displayTable(data.NftContent.Tables[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Image" {
				if len(images) != 0 {
					for i, image := range images {
						if image.Base64Image != "" && element.WidgetId == image.WidgetId {
							jsScripts += `
		displayImages(data.NftContent.Images[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Stat" {
				if len(stats) != 0 {
					for i, stat := range stats {
						if element.WidgetId == stat.WidgetId {
							jsScripts += `
		displayStat(data.NftContent.Stats[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "ProofBot" {
				if len(proofbot) != 0 {
					for i, botData := range proofbot {
						if len(botData.Data) != 0 && element.WidgetId == botData.WidgetId {
							jsScripts += `
			displayProofBot(data.NftContent.ProofBot[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Timeline" {
				if len(Timelines) != 0 {
					for i, timelineData := range Timelines {
						if len(timelineData.TimelineData) != 0 && element.WidgetId == timelineData.WidgetId {
							jsScripts += `
			displayTimeline(data.NftContent.Timeline[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else {
			}
		}
	}

	template := documentStart + `
	` + `<style>` + `
	` + styleStart + `
	` + `</style>` + `
	` + headeEnd + `
	` + body + `
	` + mainHandler + `
	` + stratScript + `
	` + `let data = ` + string(dataString) + `
	` + jsScripts + `
	` + endScript + `
	` + documentEnd

	return template, nil
}
