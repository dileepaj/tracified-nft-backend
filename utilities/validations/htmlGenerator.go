package validations

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/go-playground/validator/v10"
)

func ValidateHtmlGenerator(e models.HtmlGenerator) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateNFTProject(e models.NFTComposerProject) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateChart(e models.Chart) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateImage(e models.ImageData) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateTable(e models.Table) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateStat(e models.StataArray) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateProofBot(e models.ProofBotData) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
