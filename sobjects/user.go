package sobjects

// User struct
type User struct {
	BaseSObject
	Alias             string `force:",omitempty"`
	CommunityNickname string `force:",omitempty"`
	Email             string `force:",omitempty"`
	EmailEncodingKey  string `force:",omitempty"`
	FirstName         string `force:",omitempty"`
	FullPhotoURL      string `force:",omitempty"`
	LanguageLocaleKey string `force:",omitempty"`
	LastName          string `force:",omitempty"`
	LocaleSidKey      string `force:",omitempty"`
	ProfileID         string `force:",omitempty"`
	SmallPhotoURL     string `force:",omitempty"`
	TimeZoneSidKey    string `force:",omitempty"`
	Username          string `force:",omitempty"`
}

// APIName name of the API
func (t *User) APIName() string {
	return "User"
}

// UserQueryResponse API response struct
type UserQueryResponse struct {
	BaseQuery
	Records []User `json:"Records" force:"records"`
}
