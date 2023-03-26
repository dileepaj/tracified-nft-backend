package validations

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
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

func ValidateInsertNftStory(e models.NFTStory) error {
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

func ValidateMakeSale(e requestDtos.UpdateNFTSALERequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateRequestNFTObject(e requestDtos.CreateNFTRequest) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateRequestTXNObject(e models.TXN) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateWalletNft(e models.WalletNFT) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
