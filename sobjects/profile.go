package sobjects

// Profile struct
type Profile struct {
	BaseSObject
	Description               string `force:",omitempty"`
	IsSsoEnabled              bool   `force:",omitempty"`
	LastReferencedDate        string `force:",omitempty"`
	LastViewedDate            string `force:",omitempty"`
	Name                      string `force:",omitempty"`
	PermissionsPermissionName bool   `force:",omitempty"`
	UserLicenseID             string `force:",omitempty"`
	UserType                  string `force:",omitempty"`
}

// APIName name of the API
func (t *Profile) APIName() string {
	return "Profile"
}

// ProfileQueryResponse API response struct
type ProfileQueryResponse struct {
	BaseQuery
	Records []Profile `json:"Records" force:"records"`
}
