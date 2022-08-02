package svgGenerator

import (
	"strings"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
)

var (
	documentStart = services.ReadFromFile("services/svgGeneratorService/templates/svgHeader.txt")
	documentEnd   = services.ReadFromFile("services/svgGeneratorService/templates/svgFooter.txt")
	style         = services.ReadFromFile("services/svgGeneratorService/templates/svgStyle.css")
	styleStart    = `<style>`
	styleEnd      = `</style>`
	svg           = services.ReadFromFile("services/svgGeneratorService/templates/temp.svg")
)

/**
*	Generate complete  SVG template for NFT
* @function GenerateSVGTemplate
*
**/
func GenerateSVGTemplate(svgData models.HtmlGenerator) (string, error) {
	// Parse the Data
	var htmlBody string
	var barcharts []models.Chart = svgData.NftContent.BarCharts
	var piecharts []models.Chart = svgData.NftContent.PieCharts
	var tables []models.Table = svgData.NftContent.Tables
	var proofbot []models.ProofBotData = svgData.NftContent.ProofBot
	var bubbleCharts []models.Chart = svgData.NftContent.BubbleCharts
	var images []models.ImageData = svgData.NftContent.Images
	var Timelines []models.Timeline = svgData.NftContent.TimeLine
	var contentOrderData []models.ContentOrderData = svgData.ContentOrderData
	htmlStart := `<h1 class="text-center">Tracified NFT</h1>
							 <p class="text-center fw-bold text-muted">` + svgData.NFTName + `</p>
							 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	if len(contentOrderData) > 0 {
		for _, element := range contentOrderData {
			if element.Type == "BarChart" {
				if len(barcharts) > 0 {
					for _, bar := range barcharts {
						if len(bar.ChartData) != 0 && element.WidgetId == bar.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
											<div class="card-header">Bar Chart</div>
											<div class="text-center justify-content-center card-body">
											<div class="img-widget-image" style="background-image: url(` + bar.ChartImage + `);"></div>
											</div>
										</div>`
						}
					}
				}
			} else if element.Type == "PieChart" {
				if len(piecharts) > 0 {
					for _, pie := range piecharts {
						if len(pie.ChartData) != 0 && element.WidgetId == pie.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
											<div class="card-header">Pie Chart</div>
											<div class="card-body">
											<div class="img-widget-image" style="background-image: url(` + pie.ChartImage + `);"></div>
											</div>
										</div>`
						}
					}
				}
			} else if element.Type == "BubbleChart" {
				if len(bubbleCharts) > 0 {
					for _, bubble := range bubbleCharts {
						if len(bubble.ChartData) != 0 && element.WidgetId == bubble.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
												<div class="card-header">Bubble Chart</div>
												<div class="card-body">
												<div class="img-widget-image" style="background-image: url(` + bubble.ChartImage + `);"></div>
												</div>
											</div>`
						}
					}
				}
			} else if element.Type == "Table" {
				if len(tables) > 0 {
					for _, table := range tables {
						if table.TableContent != "" && table.TableContent!="EMPTY" && element.WidgetId == table.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
											<div class="card-header">Table</div>
											<div class="card-body text-center justify-content-center"  style="min-width: 500px;" >
											<p class="common-widget-title">` + table.TableTitle + `</p>
											<table class="table table-bordered">` + table.TableContent + `</table>
											</div>
										</div>`
						}
					}
				}
			} else if element.Type == "Image" {
				if len(images) > 0 {
					for _, image := range images {
						if image.Base64Image != "" && element.WidgetId == image.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
											<div class="card-header">Image</div>
											<div class="card-body">
											<p class="common-widget-title">` + image.Title + `</p>
											<a href="` + image.Base64Image + `"><div class="img-widget-image" style="background-image: url(` + image.Base64Image + `);"></div>
											</a></div>
										</div>`
						}
					}
				}
			} else if element.Type == "ProofBot" {
				if len(proofbot) > 0 {
					for _, botData := range proofbot {
						if len(botData.Data) > 0 && element.WidgetId == botData.WidgetId {
							var htmlBotcard string
							htmlBotHeader := `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
													<div class="card-header">Proof Bot</div>
													<div class="card-body text-center justify-content-center">
													<p class="common-widget-title">` + botData.Title + `</p>`
							for _, data := range botData.Data {
								htmlBotcard += `<div class="botCard">
														<div class="proof-section"><label class="proofbot-data-field">Product Name : </label><label class="proofbot-value-field">` + botData.ProductName + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Batch ID : </label><label class="proofbot-value-field">` + data.BatchId + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Timestamp : </label><label class="proofbot-value-field">` + botData.Timestamp.Time().String() + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Transaction ID : </label><label class="proofbot-value-field">` + data.TxnHash + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Available Proofs : </label>
														<div class="proof-url">`
								for _, proofUrl := range data.Urls {
									if proofUrl.Urls != "" {
										var removeAndsymble string = strings.Replace(proofUrl.Urls, "&", "&amp;", -1)

										htmlBotcard += `<a class="proof-anchor" href="` + removeAndsymble + `">
														` + GetProofName(proofUrl.Type) + `</a><span class="material-symbols-outlined open-icon">open_in_new</span>`
									}
								}
								htmlBotcard += `</div></div></div>`
							}
							htmlBotFooter := `</div></div>`
							htmlBody += htmlBotHeader + htmlBotcard + htmlBotFooter
						}
					}
				}
			} else if element.Type == "Timeline" {
				if len(Timelines) > 0 {
					for _, timelineData := range Timelines {
						if len(timelineData.TimelineData) != 0 && element.WidgetId == timelineData.WidgetId {
							var htmlTimelineHeader string
							var htmlTimelineBody string
							var htmlTimelineFooter string
							// var htmlTimelineImage string
							htmlTimelineHeader += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
														<div class="card-header">Timeline</div>
														<div class="card-body text-center justify-content-center">
														<p class="common-widget-title">` + timelineData.Title + `</p>
														<div class="text-start row" style="width: 500px">
														<ul class="timeline"><div class="timeline-heading">
														<label class="timeline-product">` + timelineData.ProductName + `</label>
														<label class="timeline-batch">` + timelineData.BatchId + `</label>
													  </div>`
							for _, data := range timelineData.TimelineData {
								htmlTimelineBody += ` <li class="timeline-header">
                                						<img class="timeline-icon" src="` + data.Icon + `" /><span class="timeline-stage">` + data.Title + `</span>
													  </li>
													  <div class="card p-3 point">`
								for _, timelineChild := range data.Children {
									htmlTimelineBody += `<span class="timeline-key">` + timelineChild.Key + `</span><p><span class="timeline-value">` + timelineChild.Value + `</span></p>`
								}
								for _, image := range data.Images {
									htmlTimelineBody += `
														<div class="img-timeline-image" style="background-image: url(` + image + `);">
								 						</div>
														`
								}
								htmlTimelineBody += `</div>`
							}
							htmlTimelineFooter = `</ul></div></div></div>`
							htmlBody += htmlTimelineHeader + htmlTimelineBody + htmlTimelineFooter
						}
					}
				}
			} else {
			}
		}
	}
	template := documentStart + styleStart + style + styleEnd + htmlStart + htmlBody + documentEnd
	return template, nil
}

func GetProofName(proofType string) string {
	if proofType == "poe" {
		return "Proof of Existence"
	} else if proofType == "poc" {
		return "Proof of Continuity"
	} else if proofType == "pog" {
		return "Proof of Genesis"
	} else if proofType == "pococ" {
		return "Proof of Change of Custody"
	} else {
		return proofType;
	}
}
