package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

type HtmlGeneratorRequest struct {
	NFTComposerProject models.NFTComposerProject
	WidgetDetails []models.Widget
}