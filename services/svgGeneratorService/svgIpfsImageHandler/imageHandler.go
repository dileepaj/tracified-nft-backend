package svgipfsimagehandler

import (
	"errors"

	"github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var (
	nftComposerRepo = nftComposerRepository.NFTComposerProjectRepository{}
)

func UpdateImageWithNewIpfsHash(updateImage models.ImageData, cid string) error {
	updateObj := requestDtos.SaveUpdatedImage{
		WidgetId:    updateImage.WidgetId,
		Title:       updateImage.Title,
		Type:        updateImage.Type,
		Base64Image: updateImage.Base64Image,
		ProjectId:   updateImage.ProjectId,
		Cid:         cid,
	}

	rst, errWhenUpdatingTheImage := nftComposerRepo.UpdateImage(updateObj)
	if errWhenUpdatingTheImage != nil {
		return errWhenUpdatingTheImage
	} else if rst.ProjectId == "" {
		return errors.New("No widget for the id " + updateImage.WidgetId)
	}

	return nil
}
