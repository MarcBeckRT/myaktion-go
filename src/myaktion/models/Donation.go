package models

type Status string

const (
	TRANSFERRED Status = "TRANSFERRED"
	IN_PROCESS  Status = "IN_PROCESS"
)

type Donation struct {
	Amount           float64
	ReceiptRequested bool
	DonorName        string
	Status           Status
	Account          Account
}
