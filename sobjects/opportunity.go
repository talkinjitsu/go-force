package sobjects

// Opportunity struct
type Opportunity struct {
	BaseSObject
	AccountID       string  `force:",omitempty"`
	Amount          float64 `force:",omitempty"`
	CloseDate       string  `force:",omitempty"`
	CurrencyIsoCode string  `force:",omitempty"`
	Description     string  `force:",omitempty"`
	ExpectedRevenue string  `force:",omitempty"`
	IsClosed        bool    `force:",omitempty"`
	IsDeleted       bool    `force:",omitempty"`
	IsSplit         bool    `force:",omitempty"`
	IsWon           bool    `force:",omitempty"`
	Name            string  `force:",omitempty"`
	OwnerID         string  `force:",omitempty"`
	StageName       string  `force:",omitempty"`
}

// APIName name of the API
func (t *Opportunity) APIName() string {
	return "Opportunity"
}

// OpportunityQueryResponse query reponse struct
type OpportunityQueryResponse struct {
	BaseQuery
	Records []Opportunity `json:"Records" force:"records"`
}
