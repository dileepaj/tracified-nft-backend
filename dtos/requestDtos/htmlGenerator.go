package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

type HtmlGeneratorRequest struct {
	HtmlGenerator models.HtmlGenerator
	WidgetDetails []models.Widget
}

