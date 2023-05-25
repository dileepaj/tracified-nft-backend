package validations

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
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

func ValidateUpdateProject(e requestDtos.UpdateProjectRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateTable(e requestDtos.UpdateTableRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateChart(e requestDtos.UpdateChartRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateProofBot(e requestDtos.UpdateProofBotRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateImage(e requestDtos.UpdateImageRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateTimelineRequest(e requestDtos.UpdateTimelineRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateStats(e requestDtos.UpdateStatsRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil

}

func ValidateContentOrderData(e []models.ContentOrderData) error {
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

func ValidateTimeline(e models.Timeline) error {
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
