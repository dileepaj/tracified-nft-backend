package svgGenerator

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
)

var (
	Start      = services.ReadFromFile("services/svgGeneratorService/templates/svgHeader.txt")
	End        = services.ReadFromFile("services/svgGeneratorService/templates/svgFooter.txt")
	svgStyle   = services.ReadFromFile("services/svgGeneratorService/templates/svgStyle.css")
	StartStyle = `<style>`
	EndStyle   = `</style>`
	svgTemp    = services.ReadFromFile("services/svgGeneratorService/templates/temp.svg")
)

func GenerateSVGRuriTemplate(svgData models.SvgCreator) (string, error) {
	var htmlBody string
	var htmlStart = `<h1 class="text-center">Tracified NFT</h1>
							 <p class="text-center fw-bold text-muted">` + svgData.ProductName + `</p>
							 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	template := Start + StartStyle + svgStyle + EndStyle + htmlStart + htmlBody + End
	return template, nil
}
