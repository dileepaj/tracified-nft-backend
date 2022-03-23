package htmlGenerator

import (
	"encoding/json"
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
)

var (
	documentStart  = services.ReadFromFile("services/htmlGeneretorService/templates/htmlHeader.txt")
	documentEnd = `</html>`
	styleStart = services.ReadFromFile("services/htmlGeneretorService/templates/htmlStyles.css")
	body = services.ReadFromFile("services/htmlGeneretorService/templates/htmlBody.html")
	mainHandler = services.ReadFromFile("services/htmlGeneretorService/templates/htmlScript.html")
	stratScript = `<script type="text/javascript">`
	endScript   = `</script>`
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

	//take json data convert it to string
	dataString, err := json.Marshal(htmlData)
	if err != nil {
		fmt.Println(err)
	}
	if len(barcharts) != 0 {
		jsScripts += `
		displayBarchart(data.NftContent.BarCharts)`
	}
	if len(piecharts) != 0 {
		jsScripts += `
		displayPiechart(data.NftContent.PieCharts)
		`
	}
	if len(tables) != 0 {
		jsScripts += `
		displayTable(data.NftContent.Tables)
		`
	}
	if len(stats) != 0 {
		jsScripts += `
		displayStat(data.NftContent.Stats)
		`
	}
	if len(bubbleCharts) != 0 {
		jsScripts += `
		displayBubblechart(data.NftContent.BubbleCharts)
		`
	}
	if len(images) != 0 {
		jsScripts += `
		displayImages(data.NftContent.Images)
		`
	}
	if len(proofbot) != 0 {
		jsScripts += `
		displayProofBot(data.NftContent.ProofBot)
			`
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
	return template, nil
}