package request

// Create a struct that models the structure of a user, both in the request body, and in the DB
type BeneficiaryEnrollment struct {
	OwnerAccount         string `json:"ownerAccount"`
	BeneficiaryAccount   string `json:"beneficiaryAccount"`
	BeneficiaryFirstName string `json:"beneficiaryFirstName"`
	BeneficiaryLastName  string `json:"beneficiaryLastName"`
	Email                string `json:"email"`
}
