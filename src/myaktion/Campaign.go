package models

type Campaign struct {
	Name            string
	DonationMinimum float64
	TargetAmount    float64
	Donations       []Donation
	Account         Account
}
