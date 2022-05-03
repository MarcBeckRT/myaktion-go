package service

import (
	"log"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/models"
)

var (
	campaignStore map[uint]*models.Campaign
	actCampaignId uint = 1
)

func init() {
	campaignStore = make(map[uint]*models.Campaign)
}

func CreateCampaign(campaign *models.Campaign) error {
	campaign.ID = actCampaignId
	campaignStore[actCampaignId] = campaign
	actCampaignId += 1
	log.Printf("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Printf("Stored: %v", campaign)
	return nil

}
