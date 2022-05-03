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

func GetCampaigns() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	for _, campaign := range campaignStore {
		campaigns = append(campaigns, *campaign)
	}
	log.Printf("Retrieved: %v", campaigns)
	return campaigns, nil
}
