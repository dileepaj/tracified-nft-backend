package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

type HtmlGeneratorRequest struct {
	HtmlGenerator models.NFTComposerProject
	WidgetDetails []models.Widget
}

