package validations

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/wrappers/requestWrappers"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateInsertNft(e models.NFT) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateInsertOffer(e models.Offer) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateInsertOwnership(e models.Ownership) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateCreateNFTObject(e requestWrappers.CreateNFTRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}