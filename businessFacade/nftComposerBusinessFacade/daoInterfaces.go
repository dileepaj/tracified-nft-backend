/*
	This is a Dao repositort intefaces 
	nftProjectRepository include the all crud opration for Project,Chart,proofcbot,images,timeline and tables in nft composer 
	widgetRepository 

*/
package nftComposerBusinessFacade

import "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"

var nftProjectRepository nftComposerRepository.NFTComposerProjectRepository
var widgetRepository nftComposerRepository.WidgetRepository