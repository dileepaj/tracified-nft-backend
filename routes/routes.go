package routes

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/routes/marketPlace"
	"github.com/dileepaj/tracified-nft-backend/routes/nftComposer"
)

var ApplicationRoutes models.Routers

func init() {
	routes := []models.Routers{
		testRoutes,
		marketPlace.NftRoutes,
		marketPlace.UserRoutes,
		marketPlace.WatchListRoutes,
		nftComposer.ComposerRoutes,
		marketPlace.ReviewRoutes,
		marketPlace.NewsLetterRoutes,
		marketPlace.FaqRoutes,
		marketPlace.CollectionRoutes,
		marketPlace.FavouritesRoutes,
		marketPlace.EndorsementRoutes,
		marketPlace.DocsRoutes,
		marketPlace.PartnerRoutes,
		marketPlace.OneTimePassWordRoutes,
		marketPlace.SvgGenerator,
		marketPlace.MapGenerateRoutes,
		marketPlace.WalletNFTRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
