package sobjects

// Account sobject type
type Account struct {
	BaseSObject
	BillingCity       string `force:",omitempty"`
	BillingCountry    string `force:",omitempty"`
	BillingPostalCode string `force:",omitempty"`
	BillingState      string `force:",omitempty"`
	BillingStreet     string `force:",omitempty"`
}

// APIName returns the name of the API
func (a Account) APIName() string {
	return "Account"
}
