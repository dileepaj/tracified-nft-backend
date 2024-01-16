package htmlGenerator

import (
	"encoding/json"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	documentStartTimeline = services.ReadFromFile("services/htmlGeneretorService/templates/htmlHeader.txt")
	documentEndTimeline   = `</html>`
	headeEndTimeline      = `</header>`
	styleStartTimeline    = services.ReadFromFile("services/htmlGeneretorService/templates/timeline.css")
	mainHandlerTimeline   = services.ReadFromFile("services/htmlGeneretorService/templates/timeline.html")
	stratScriptTimeline   = `<script type="text/javascript">`
	endScriptTimeline     = `</script>`
)

func GenerateHTMLTemplateForTimeline(timelineData models.Component) (string, error) {
	// Parse the Data
	var jsScripts string
	body := ` <body style="font-family: 'Inter'; color: #021D27">
				 <div class="d-flex justify-content-center align-content-center flex-wrap" id="container">
				</div>
			  </body>`
	// take json data convert it to string
	dataString, err := json.Marshal(timelineData)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}

	jsScripts += `
			displayTimeline(data)
			`

	template := documentStartTimeline + `
	` + `<style>` + `
	` + styleStartTimeline + `
	` + `</style>` + `
	` + headeEndTimeline + `
	` + body + `
	` + mainHandlerTimeline + `
	` + stratScriptTimeline + `
	` + `let data = ` + string(dataString) + `
	` + jsScripts + `
	` + endScriptTimeline + `
	` + documentEndTimeline

	return template, nil
}
