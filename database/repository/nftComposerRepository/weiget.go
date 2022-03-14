package nftcomposerrepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type WeigetRepository struct{}

func (r *WeigetRepository) SaveWeigetList(weigetList []models.Weight) (string, error) {
	var docs []interface{}
	for _, t := range weigetList{
		docs = append(docs, t)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := connections.Connect().Collection("weiget").InsertMany(ctx, docs)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	} else {
		return "SAVED", nil
	}

}
