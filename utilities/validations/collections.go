package validations

import (
	"log"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/go-playground/validator/v10"
)

func ValidateInsertCollection(e models.NFTCollection) error {
	log.Println("------------------------------------testing 3 ---------------------------------------------------")
	validate = validator.New()
	err := validate.Struct(e)
	log.Println("------------------------------------testing 4 ---------------------------------------------------")
	if err != nil {
		log.Println("------------------------------------err ---------------------------------------------------", err)
		return err
	}
	log.Println("------------------------------------testing 5 ---------------------------------------------------")
	return nil
}

func ValidateUpdateRequestCollection(e requestDtos.UpdateCollection) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
