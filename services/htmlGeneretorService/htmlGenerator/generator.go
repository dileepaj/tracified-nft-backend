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
				 <div class="text-center nft-header">` + htmlData.NFTName + `</div>
				 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
				</div>
				<div class="text-center nft-footer ">
					<div class="nft-footer-content">
						<label>Powered by </label> <img src="data:image/png;base64,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" />
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
