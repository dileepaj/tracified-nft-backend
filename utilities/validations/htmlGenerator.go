package validations

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/go-playground/validator/v10"
)

func ValidateInsertHtmlNft(e requestDtos.HtmlGeneratorRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
