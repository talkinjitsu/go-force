package force

// Limits ia a limit map
type Limits map[string]Limit

// Limit with remaining and max
type Limit struct {
	Remaining float64
	Max       float64
}

// GetLimits returns all limits
func (forceAPI *API) GetLimits() (limits *Limits, err error) {
	uri := forceAPI.apiResources[limitsKey]

	limits = &Limits{}
	err = forceAPI.Get(uri, nil, limits)

	return
}
