package force

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

// SObject interface all standard and custom objects must implement. Needed for uri generation.
type SObject interface {
	APIName() string
	ExternalIDAPIName() string
}

// SObjectResponse is the response received from force.com API after insert of an sobject.
type SObjectResponse struct {
	ID      string    `force:"id,omitempty"`
	Errors  APIErrors `force:"error,omitempty"` //TODO: Not sure if ApiErrors is the right object
	Success bool      `force:"success,omitempty"`
}

// DescribeSObjects returns the SObjects
func (forceAPI *API) DescribeSObjects() (map[string]*SObjectMetaData, error) {
	if err := forceAPI.getAPISObjects(); err != nil {
		return nil, err
	}

	return forceAPI.apiSObjects, nil
}

// DescribeSObject returns the Sobject
func (forceAPI *API) DescribeSObject(in SObject) (resp *SObjectDescription, err error) {
	// Check cache
	resp, ok := forceAPI.apiSObjectDescriptions[in.APIName()]
	if !ok {
		// Attempt retrieval from api
		sObjectMetaData, ok := forceAPI.apiSObjects[in.APIName()]
		if !ok {
			err = fmt.Errorf("unable to find metadata for object: %v", in.APIName())
			return
		}

		uri := sObjectMetaData.URLs[sObjectDescribeKey]

		resp = &SObjectDescription{}
		err = forceAPI.Get(uri, nil, resp, nil)
		if err != nil {
			return
		}

		// Create Comma Separated String of All Field Names.
		// Used for SELECT * Queries.
		length := len(resp.Fields)
		if length > 0 {
			var allFields bytes.Buffer
			for index, field := range resp.Fields {
				// Field type location cannot be directly retrieved from SQL Query.
				if field.Type != "location" {
					if index > 0 && index < length {
						allFields.WriteString(", ")
					}
					allFields.WriteString(field.Name)
				}
			}

			resp.AllFields = allFields.String()
		}

		forceAPI.apiSObjectDescriptions[in.APIName()] = resp
	}

	return
}

// GetSObject fetches the sobject
func (forceAPI *API) GetSObject(id string, fields []string, out SObject) (err error) {
	uri := strings.Replace(forceAPI.apiSObjects[out.APIName()].URLs[rowTemplateKey], idKey, id, 1)

	params := url.Values{}
	if len(fields) > 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	err = forceAPI.Get(uri, params, out.(interface{}), nil)

	return
}

// InsertSObject insert a SObject
func (forceAPI *API) InsertSObject(in SObject) (resp *SObjectResponse, err error) {
	uri := forceAPI.apiSObjects[in.APIName()].URLs[sObjectKey]

	resp = &SObjectResponse{}
	err = forceAPI.Post(uri, nil, in.(interface{}), resp, nil)

	return
}

// UpdateSObject update a SObject
func (forceAPI *API) UpdateSObject(id string, in SObject) (err error) {
	uri := strings.Replace(forceAPI.apiSObjects[in.APIName()].URLs[rowTemplateKey], idKey, id, 1)

	err = forceAPI.Patch(uri, nil, in.(interface{}), nil, nil)

	return
}

// DeleteSObject delete a SObject
func (forceAPI *API) DeleteSObject(id string, in SObject) (err error) {
	uri := strings.Replace(forceAPI.apiSObjects[in.APIName()].URLs[rowTemplateKey], idKey, id, 1)

	err = forceAPI.Delete(uri, nil, nil)

	return
}

// GetSObjectByExternalID get a SObject external ID
func (forceAPI *API) GetSObjectByExternalID(id string, fields []string, out SObject) (err error) {
	uri := fmt.Sprintf("%v/%v/%v", forceAPI.apiSObjects[out.APIName()].URLs[sObjectKey],
		out.ExternalIDAPIName(), id)

	params := url.Values{}
	if len(fields) > 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	err = forceAPI.Get(uri, params, out.(interface{}), nil)

	return
}

// UpsertSObjectByExternalID update a SObject external ID
func (forceAPI *API) UpsertSObjectByExternalID(id string, in SObject) (resp *SObjectResponse, err error) {
	uri := fmt.Sprintf("%v/%v/%v", forceAPI.apiSObjects[in.APIName()].URLs[sObjectKey],
		in.ExternalIDAPIName(), id)

	resp = &SObjectResponse{}
	err = forceAPI.Patch(uri, nil, in.(interface{}), resp, nil)

	return
}

// DeleteSObjectByExternalID delete a SObject external ID
func (forceAPI *API) DeleteSObjectByExternalID(id string, in SObject) (err error) {
	uri := fmt.Sprintf("%v/%v/%v", forceAPI.apiSObjects[in.APIName()].URLs[sObjectKey],
		in.ExternalIDAPIName(), id)

	err = forceAPI.Delete(uri, nil, nil)

	return
}
