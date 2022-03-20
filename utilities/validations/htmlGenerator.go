package validations

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/go-playground/validator/v10"
)

func ValidateInsertHtmlNft(e models.NFTComposerProject) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}