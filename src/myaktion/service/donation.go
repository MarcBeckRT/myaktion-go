package service

import (
	"log"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/models"
)

var (
	donationStore map[uint]*models.Donation
	actDonationId uint = 1
)

func init() {
	donationStore = make(map[uint]*models.Donation)
}

func AddDonation(donation *models.Donation) error {
	donation.ID = actDonationId
	donation.Status = "IN_PROCESS"
	donationStore[actDonationId] = donation
	actDonationId += 1
	log.Printf("Successfully stored new donation with ID %v in database.", donation.ID)
	log.Printf("Stored: %v", donation)
	return nil
}