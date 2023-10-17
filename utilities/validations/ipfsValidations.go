package validations

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/go-playground/validator/v10"
)

func ValidateUploadIpfsFile(e models.IpfsObjectForTDP) error {
	validate = validator.New()
	errWhenValiating := validate.Struct(e)
	if errWhenValiating != nil {
		return errWhenValiating
	}
	return nil
}
