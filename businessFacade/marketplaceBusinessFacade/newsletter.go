package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

/*
	Function below Calls the respective newsletterRepository methods and parses data accordingly
*/
func CreateNewsLetter(newsletter models.NewsLetter) (string, error) {
	return newsletterRepository.SaveNewsLetter(newsletter)
}
func GetAllNewsLetters() ([]models.NewsLetter, error) {
	return newsletterRepository.GetAllNewsLetters()
}
func GetNewsLetterByAuthor(authorname string) ([]models.NewsLetter, error) {
	return newsletterRepository.GetNewsLetterByAuthor(authorname)
}
