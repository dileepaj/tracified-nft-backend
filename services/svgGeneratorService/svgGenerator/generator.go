package svgGenerator

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/sirupsen/logrus"
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
	const stageIcon = "https://s3.ap-south-1.amazonaws.com/tracified-image-storage/mobile/stage-icons/Harvesting+stage.png"

	htmlStart := `	<div class="cont-div">
					<div class="nft-header default-font">
					<img src="https://s3.ap-south-1.amazonaws.com/qa.marketplace.nft.tracified.com/Tracified-RT-Logo-White.svg"
					class="nft-logo"/>
					<label>` + svgData.NFTName + `</label>
					</div>
				<div class="d-flex justify-content-center align-content-center flex-wrap" id="container">`

	if len(contentOrderData) > 0 {
		for _, element := range contentOrderData {
			if element.Type == "BarChart" {
				if len(barcharts) > 0 {
					for _, bar := range barcharts {
						if len(bar.ChartData) != 0 && element.WidgetId == bar.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3  default-font round-card" style="max-height: fit-content;">
											<div class="card-header round-card-header">` + bar.ChartTitle + `</div>
											<div class="text-center card-body justify-content-center scroll">
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
							htmlBody += `<div class="card text-center justify-content-center m-3 default-font round-card" style=" max-height: fit-content;">
											<div class="card-header round-card-header">` + pie.ChartTitle + `</div>
											<div class="card-body justify-content-center scroll">
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
							htmlBody += `<div class="card text-center justify-content-center m-3 default-font round-card" style="max-height: fit-content;">
												<div class="card-header round-card-header">` + bubble.ChartTitle + `</div>
												<div class="card-body justify-content-center scroll">
												<div class="img-widget-image" style="background-image: url(` + bubble.ChartImage + `);"></div>
												</div>
											</div>`
						}
					}
				}
			} else if element.Type == "Table" {
				if len(tables) > 0 {
					for _, table := range tables {
						if table.TableContent != "" && table.TableContent != "EMPTY" && element.WidgetId == table.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3 default-font round-card" style="max-height: fit-content;">
											<div class="card-header round-card-header scroll">` + table.TableTitle + `</div>
											<div class="card-body-2 text-center scroll"  style="width: 100%;" >
											<div>
											<div class="table-responsive scroll m-3" style="height:370px;">
											<table class="table text-wrap table-hover table-bordered" style="word-wrap: break-word;">` + table.TableContent + `</table>
											</div>
											</div>
											</div>
										</div>`
						}
					}
				}
			} else if element.Type == "Image" {
				if len(images) > 0 {
					for _, image := range images {
						if image.Base64Image != "" && element.WidgetId == image.WidgetId {
							htmlBody += `<div class="card text-center justify-content-center m-3 default-font round-card" style="max-height: fit-content;">
											<div class="card-header round-card-header">` + image.Title + `</div>
											<div class="card-body justify-content-center scroll">
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
							htmlBotHeader := `<div class="card text-center justify-content-center m-3 default-font round-card" style="max-height: fit-content;">
													<div class="card-header round-card-header">` + botData.Title + `</div>
													<div class="card-body text-center scroll">`
							for _, data := range botData.Data {
								htmlBotcard += `<div class="botCard">
														<div class="proof-section"><label class="proofbot-data-field">Product Name : </label><label class="proofbot-value-field">` + botData.ProductName + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Batch ID : </label><label class="proofbot-value-field">` + data.BatchId + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Timestamp : </label><label class="proofbot-value-field">` + data.Timestamp + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Transaction ID : </label><label class="proofbot-value-field">` + data.TxnHash + `</label></div>
														<div class="proof-section"><label class="proofbot-data-field">Available Proofs : </label>
														`
								for _, proofUrl := range data.Urls {
									if (proofUrl.Urls != "") && (strings.ToLower(proofUrl.Type) != "poc") {
										var removeAndsymble string = strings.Replace(proofUrl.Urls, "&", "&amp;", -1)

										htmlBotcard += `<a class="proof-anchor1" href="` + removeAndsymble + `" target="_blank" rel="noopener noreferrer">
														<div class="proof-url"><a class="proof-anchor">
														` + GetProofName(proofUrl.Type) + `</a><span class="material-symbols-outlined open-icon">open_in_new</span>
														</div></a>`
									}
								}
								htmlBotcard += `</div></div>`
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
							htmlTimelineHeader += `<div class="card text-center justify-content-center m-3 default-font round-card" style="max-height: fit-content;">
														<div class="card-header round-card-header">` + timelineData.Title + `</div>
														<div class="card-body text-center scroll">
														<div class="text-start row" style="width: 100%">
														<ul class="timeline"><div class="timeline-heading">
														<label class="timeline-product"><span class="bold-text">Product Name : </span>` + timelineData.ProductName + `</label>
														<label class="timeline-batch"><span class="bold-text">Batch ID : </span>` + timelineData.BatchId + `</label>
													  </div>`
							for _, dataPackets := range timelineData.TimeLineTDPData {
								htmlTimelineBody += ` <li class="timeline-header">
                                						<img class="timeline-icon" src="` + stageIcon + `" /><span class="timeline-stage">` + replaceAndSymbol(dataPackets.Title) + `</span>
													  </li>
													  <div class="card p-3 point">`
								if len(dataPackets.TraceabilityDataPackets) > 0 {
									for _, traceabilityData := range dataPackets.TraceabilityDataPackets {
										for _, tdp := range traceabilityData.TraceabilityData {
											if tdp.Type == 1 {
												floatValue := 0.0 // Replace SomeStruct with the expected data structure
												if err := json.Unmarshal(tdp.Val, &floatValue); err != nil {
													logrus.Error("Error unmarshaling TraceabilityData: %v\n", err)
												} else {
													htmlTimelineBody += `<p><span class="timeline-key">` + replaceAndSymbol(tdp.Key) + ` : </span><span class="timeline-value">` + strconv.FormatFloat(floatValue, 'f', -1, 64) + `</span></p>`
												}
											} else if tdp.Type == 5 || tdp.Type == 3 {
												strData := ""
												if err := json.Unmarshal(tdp.Val, &strData); err != nil {
													logrus.Error("Error unmarshaling TraceabilityData: %v\n", err)
												} else {
													htmlTimelineBody += `<p><span class="timeline-key">` + replaceAndSymbol(tdp.Key) + ` : </span><span class="timeline-value">` + strData + `</span></p>`
												}
											} else if tdp.Type == 4 {
												var images []models.ImageValuesData
												if err := json.Unmarshal(tdp.Val, &images); err != nil {
													logrus.Error("Error unmarshaling TraceabilityData: %v\n", err)
												} else {
													htmlTimelineBody += `<p><span class="timeline-value">` + replaceAndSymbol(tdp.Key) + `</span></p>`
													for _, img := range images {
														htmlTimelineBody += `<span class="timeline-key">` + img.Timestamp + `</span>`
														htmlTimelineBody += `<span class="timeline-key">` + replaceAndSymbol(img.Description) + `</span>`
														htmlTimelineBody += `
																	<div class="img-timeline-image" style="background-image: url(` + img.Image + `);">
																	 </div>
															`
													}
												}
											}
										}
									}
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
		return proofType
	}
}

func replaceAndSymbol(value string) string {
	var newValue string = strings.Replace(value, "&", "&amp;", -1)
	return newValue
}
