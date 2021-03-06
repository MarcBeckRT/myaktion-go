package models

type Campaign struct {
	ID              uint
	Name            string
	OrganizerName   string
	DonationMinimum float64
	TargetAmount    float64
	Donations       []Donation
	Account         Account
}
