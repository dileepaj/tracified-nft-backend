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
	styleStart    = services.ReadFromFile("services/htmlGeneretorService/templates/htmlStyles.css")
	body          = services.ReadFromFile("services/htmlGeneretorService/templates/htmlBody.html")
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
	var contentOrderData []models.ContentOrderData = htmlData.ContentOrderData

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
			}else{

			}
		}
	}

	template := documentStart + `
	` + styleStart + `
	` + body + `
	` + mainHandler + `
	` + stratScript + `
	` + `let data = ` + string(dataString) + `
	` + jsScripts + `
	` + endScript + `
	` + documentEnd

	fmt.Println("--------------", jsScripts)
	return template, nil
}
