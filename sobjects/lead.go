package sobjects

// Lead struct
type Lead struct {
	BaseSObject
	Company       string `force:",omitempty"`
	ConvertedDate string `force:",omitempty"`
	FirstName     string `force:",omitempty"`
	IsConverted   bool   `force:",omitempty"`
	IsDeleted     bool   `force:",omitempty"`
	LastName      string `force:",omitempty"`
	OwnerID       string `force:",omitempty"`
	Status        string `force:",omitempty"`
}

// APIName name of the API
func (t *Lead) APIName() string {
	return "Lead"
}

// LeadQueryResponse struct for API response
type LeadQueryResponse struct {
	BaseQuery
	Records []Lead `json:"Records" force:"records"`
}
