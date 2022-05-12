package svgGenerator

import (
	"fmt"
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
	htmlStart     = services.ReadFromFile("services/svgGeneratorService/templates/htmlStart.html")
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

	if len(contentOrderData) > 0 {
		for _, element := range contentOrderData {
			if element.Type == "BarChart" {
				if len(barcharts) > 0 {
					for _, bar := range barcharts {
						if element.WidgetId == bar.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
							<div class="card-header">Bar Chart</div>
							<div class="text-center justify-content-center card-body">
							  <img
								style="max-width: 500px; max-height: 400px"
								src="` + bar.ChartImage + `"
								alt=""
							  />
							</div>
						  </div>`
						}
					}
				}
			} else if element.Type == "PieChart" {
				if len(piecharts) > 0 {
					for _, pie := range piecharts {
						if element.WidgetId == pie.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
								<div class="card-header">Pie Chart</div>
								<div class="card-body">
								  <img
									style="max-width: 500px; max-height: 400px"
									src="` + pie.ChartImage + `"
									alt=""
								  />
								</div>
							  </div>`
						}
					}
				}
			} else if element.Type == "BubbleChart" {
				if len(bubbleCharts) > 0 {
					for _, bubble := range bubbleCharts {
						if element.WidgetId == bubble.WidgetId {
							if element.WidgetId == bubble.WidgetId {
								htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
								<div class="card-header">Bubble Chart</div>
								<div class="card-body">
								  <img
									style="max-width: 500px; max-height: 400px"
									src="` + bubble.ChartImage + `"
									alt=""
								  />
								</div>
							  </div>`
							}
						}
					}
				}
			} else if element.Type == "Table" {
				if len(tables) > 0 {
					for _, table := range tables {
						if element.WidgetId == table.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
											<div class="card-header">Table</div>
												<div class="card-body text-center justify-content-center"  style="min-width: 500px;" >
												<p style="margin-bottom:10px; font-weight:bold;">`+ table.TableTitle +`</p>
												<table class="table table-bordered">` + table.TableContent + `</table>
											</div>
										</div>`
						}
					}
				}
			} else if element.Type == "Image" {
				if len(images) > 0 {
					for _, image := range images {
						if element.WidgetId == image.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
							<div class="card-header">Image</div>
							<div class="card-body">
							  <img
								style="max-width: 500px; max-height: 400px"
								src="` + image.Base64Image + `"
								alt=""
							  />
							</div>
						  </div>`
						}
					}
				}
			} else if element.Type == "ProofBot" {
				if len(proofbot) > 0 {
					for _, botData := range proofbot {
						if element.WidgetId == botData.WidgetId {
							if len(botData.Data) > 0 {
								var htmlBotcard string
								htmlBotHeader := `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
														<div class="card-header">PrrofBot</div>
														<div class="card-body text-center justify-content-center">
														<p style="margin-bottom:10px; font-weight:bold;">`+ botData.Title +`</p>`
								for _, data := range botData.Data {
									htmlBotcard += `<div class="card botCard">
															<p class="text-start">ProductName :` + botData.ProductName + `</p>
															<p class="text-start">Timestamp :` + botData.Timestamp.Time().String() + `</p>
															<p class="text-start"> TxnHash :` + data.TxnHash + `</p>
															<p class="text-start">Availble Proofs :</p>`
									for _, proofUrl := range data.Urls {
										if(proofUrl.Urls!=""){
										var removeAndsymble string=strings.Replace(proofUrl.Urls, "&", "&amp;", -1)
									
										htmlBotcard += `<p class="text-start">
															<a href="` + removeAndsymble + `">
															` + proofUrl.Type + `</a>
															</p>`
														}
									}
									htmlBotcard += `</div>`
								}
								htmlBotFooter := `</div>
													</div>`
								htmlBody += htmlBotHeader + htmlBotcard + htmlBotFooter
							}
						}
					}
				}
			} else if element.Type == "Timeline" {
				if len(Timelines) > 0 {
					for _, timelineData := range Timelines {
						if element.WidgetId == timelineData.WidgetId {
							var htmlTimelineHeader string
							var htmlTimelineBody string
							var htmlTimelineFooter string
							// var htmlTimelineImage string
							htmlTimelineHeader += `<div class="card text-center justify-content-center m-3" style="min-width: 500px; max-height: fit-content;">
														<div class="card-header">Timeline</div>
														<div class="card-body text-center justify-content-center">
															<div class="text-start row" style="width: 500px">
																<ul class="timeline">`
							for _, data := range timelineData.TimelineData {
								htmlTimelineBody += ` <li class="timeline-header">
                                <img class="timeline-icon"
                                    src="`+ data.Icon+`" />
                                <span class="timeline-stage">`+data.Title+`</span></li>
                        		<div class="card p-3 point">`
								for _, timelineChild := range data.Children {
									htmlTimelineBody += `<span class="timeline-key">` + timelineChild.Key + `</span><p><span class="timeline-value">` + timelineChild.Value + `</span></p>`
								}
								for _,image:=range data.Images{
									htmlTimelineBody += `
									<img src="`+ image +`" width="300" height="250" alt=""/>
									`
								}
								htmlTimelineBody += `</div>`
							}
							htmlTimelineFooter = `</ul></div></div></div>`
							htmlBody += htmlTimelineHeader + htmlTimelineBody  + htmlTimelineFooter
						}
					}
				}
			} else {
			}
		}
	}
	template := documentStart + styleStart + style + styleEnd + htmlStart + htmlBody + documentEnd
	fmt.Println("ssssasasa ",template)
	return template, nil
}
