package force

import (
	"fmt"
	"strings"
)

// APIErrors Custom Error to handle salesforce api responses.
type APIErrors []*APIError

// APIError is a salesforce API error
type APIError struct {
	Fields           []string `json:"fields,omitempty" force:"fields,omitempty"`
	Message          string   `json:"message,omitempty" force:"message,omitempty"`
	ErrorCode        string   `json:"errorCode,omitempty" force:"errorCode,omitempty"`
	ErrorName        string   `json:"error,omitempty" force:"error,omitempty"`
	ErrorDescription string   `json:"error_description,omitempty" force:"error_description,omitempty"`
}

func (e APIErrors) Error() string {
	return e.String()
}

func (e APIErrors) String() string {
	s := make([]string, len(e))
	for i, err := range e {
		s[i] = err.String()
	}

	return strings.Join(s, "\n")
}

// Validate the error
func (e APIErrors) Validate() bool {
	return len(e) != 0
}

func (e APIError) Error() string {
	return e.String()
}

func (e APIError) String() string {
	return fmt.Sprintf("%#v", e)
}

// Validate the error
func (e APIError) Validate() bool {
	if len(e.Fields) != 0 || len(e.Message) != 0 || len(e.ErrorCode) != 0 || len(e.ErrorName) != 0 || len(e.ErrorDescription) != 0 {
		return true
	}

	return false
}
