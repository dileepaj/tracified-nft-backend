package validations

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/go-playground/validator/v10"
)

func ValidateInsertWeiget(e models.Weiget) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateQueryExecuter(e requestDtos.RequestWeiget) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}