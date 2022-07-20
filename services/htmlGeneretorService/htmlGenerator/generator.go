package htmlGenerator

import (
	"encoding/json"
	"fmt"
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

    fmt.Println("sss ss s              ",htmlData)

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
	body := ` <body>
				 <h1 class="text-center">Tracified NFT</h1>
				 <p class="text-center fw-bold text-muted">`+ htmlData.NFTName +`<p>
				</div>
				 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
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
						if element.WidgetId == bar.WidgetId {
							jsScripts += `
		displayBarchart(data.NftContent.BarCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "PieChart" {
				if len(piecharts) != 0 {
					for i, pie := range piecharts {
						if element.WidgetId == pie.WidgetId {
							jsScripts += `
		displayPiechart(data.NftContent.PieCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "BubbleChart" {
				if len(bubbleCharts) != 0 {
					for i, bubble := range bubbleCharts {
						if element.WidgetId == bubble.WidgetId {
							jsScripts += `
		displayBubblechart(data.NftContent.BubbleCharts[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Table" {
				if len(tables) != 0 {
					for i, table := range tables {
						if element.WidgetId == table.WidgetId {
							jsScripts += `
		displayTable(data.NftContent.Tables[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Image" {
				if len(images) != 0 {
					for i, image := range images {
						if element.WidgetId == image.WidgetId {
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
						if element.WidgetId == botData.WidgetId {
							jsScripts += `
			displayProofBot(data.NftContent.ProofBot[
		` + strconv.Itoa(i) + `])`
						}
					}
				}
			} else if element.Type == "Timeline" {
				if len(Timelines) != 0 {
					for i, timelineData := range Timelines {
						if len(timelineData.TimelineData)!=0 {
						if element.WidgetId == timelineData.WidgetId {
							jsScripts += `
			displayTimeline(data.NftContent.Timeline[
		` + strconv.Itoa(i) + `])`
						}
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
