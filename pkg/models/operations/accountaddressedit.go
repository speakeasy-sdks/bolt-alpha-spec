// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"net/http"
)

type AccountAddressEditSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
	Oauth  string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *AccountAddressEditSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AccountAddressEditSecurity) GetOauth() string {
	if o == nil {
		return ""
	}
	return o.Oauth
}

type AccountAddressEditRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// The ID of the address to edit
	ID                                                                    string                                                                       `pathParam:"style=simple,explode=false,name=id"`
	Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput `request:"mediaType=application/json"`
}

func (o *AccountAddressEditRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressEditRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountAddressEditRequest) GetOneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput() shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput {
	if o == nil {
		return shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput{}
	}
	return o.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput
}

// AccountAddressEdit400ApplicationJSON2Tag - The type of error returned
type AccountAddressEdit400ApplicationJSON2Tag string

const (
	AccountAddressEdit400ApplicationJSON2TagInvalidRegion AccountAddressEdit400ApplicationJSON2Tag = "invalid_region"
)

func (e AccountAddressEdit400ApplicationJSON2Tag) ToPointer() *AccountAddressEdit400ApplicationJSON2Tag {
	return &e
}

func (e *AccountAddressEdit400ApplicationJSON2Tag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_region":
		*e = AccountAddressEdit400ApplicationJSON2Tag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountAddressEdit400ApplicationJSON2Tag: %v", v)
	}
}

type AccountAddressEdit400ApplicationJSON2 struct {
	// The type of error returned
	DotTag AccountAddressEdit400ApplicationJSON2Tag `json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (o *AccountAddressEdit400ApplicationJSON2) GetDotTag() AccountAddressEdit400ApplicationJSON2Tag {
	if o == nil {
		return AccountAddressEdit400ApplicationJSON2Tag("")
	}
	return o.DotTag
}

func (o *AccountAddressEdit400ApplicationJSON2) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// AccountAddressEdit400ApplicationJSON1Tag - The type of error returned
type AccountAddressEdit400ApplicationJSON1Tag string

const (
	AccountAddressEdit400ApplicationJSON1TagInvalidPostalCode AccountAddressEdit400ApplicationJSON1Tag = "invalid_postal_code"
)

func (e AccountAddressEdit400ApplicationJSON1Tag) ToPointer() *AccountAddressEdit400ApplicationJSON1Tag {
	return &e
}

func (e *AccountAddressEdit400ApplicationJSON1Tag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_postal_code":
		*e = AccountAddressEdit400ApplicationJSON1Tag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountAddressEdit400ApplicationJSON1Tag: %v", v)
	}
}

type AccountAddressEdit400ApplicationJSON1 struct {
	// The type of error returned
	DotTag AccountAddressEdit400ApplicationJSON1Tag `json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (o *AccountAddressEdit400ApplicationJSON1) GetDotTag() AccountAddressEdit400ApplicationJSON1Tag {
	if o == nil {
		return AccountAddressEdit400ApplicationJSON1Tag("")
	}
	return o.DotTag
}

func (o *AccountAddressEdit400ApplicationJSON1) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type AccountAddressEdit400ApplicationJSONType string

const (
	AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON1 AccountAddressEdit400ApplicationJSONType = "accountAddressEdit_400ApplicationJSON_1"
	AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON2 AccountAddressEdit400ApplicationJSONType = "accountAddressEdit_400ApplicationJSON_2"
)

type AccountAddressEdit400ApplicationJSON struct {
	AccountAddressEdit400ApplicationJSON1 *AccountAddressEdit400ApplicationJSON1
	AccountAddressEdit400ApplicationJSON2 *AccountAddressEdit400ApplicationJSON2

	Type AccountAddressEdit400ApplicationJSONType
}

func CreateAccountAddressEdit400ApplicationJSONAccountAddressEdit400ApplicationJSON1(accountAddressEdit400ApplicationJSON1 AccountAddressEdit400ApplicationJSON1) AccountAddressEdit400ApplicationJSON {
	typ := AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON1

	return AccountAddressEdit400ApplicationJSON{
		AccountAddressEdit400ApplicationJSON1: &accountAddressEdit400ApplicationJSON1,
		Type:                                  typ,
	}
}

func CreateAccountAddressEdit400ApplicationJSONAccountAddressEdit400ApplicationJSON2(accountAddressEdit400ApplicationJSON2 AccountAddressEdit400ApplicationJSON2) AccountAddressEdit400ApplicationJSON {
	typ := AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON2

	return AccountAddressEdit400ApplicationJSON{
		AccountAddressEdit400ApplicationJSON2: &accountAddressEdit400ApplicationJSON2,
		Type:                                  typ,
	}
}

func (u *AccountAddressEdit400ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	accountAddressEdit400ApplicationJSON1 := new(AccountAddressEdit400ApplicationJSON1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&accountAddressEdit400ApplicationJSON1); err == nil {
		u.AccountAddressEdit400ApplicationJSON1 = accountAddressEdit400ApplicationJSON1
		u.Type = AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON1
		return nil
	}

	accountAddressEdit400ApplicationJSON2 := new(AccountAddressEdit400ApplicationJSON2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&accountAddressEdit400ApplicationJSON2); err == nil {
		u.AccountAddressEdit400ApplicationJSON2 = accountAddressEdit400ApplicationJSON2
		u.Type = AccountAddressEdit400ApplicationJSONTypeAccountAddressEdit400ApplicationJSON2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AccountAddressEdit400ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.AccountAddressEdit400ApplicationJSON1 != nil {
		return json.Marshal(u.AccountAddressEdit400ApplicationJSON1)
	}

	if u.AccountAddressEdit400ApplicationJSON2 != nil {
		return json.Marshal(u.AccountAddressEdit400ApplicationJSON2)
	}

	return nil, errors.New("could not marshal union type: all fields are null")

}

type AccountAddressEditResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request is missing required fields, or its fields have invalid values
	AccountAddressEdit400ApplicationJSONOneOf *AccountAddressEdit400ApplicationJSON
	// The address was successfully edited
	Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema *shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema
}

func (o *AccountAddressEditResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressEditResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressEditResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddressEditResponse) GetAccountAddressEdit400ApplicationJSONOneOf() *AccountAddressEdit400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.AccountAddressEdit400ApplicationJSONOneOf
}

func (o *AccountAddressEditResponse) GetOneaccount1addressesPostRequestBodyContentApplication1jsonSchema() *shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema {
	if o == nil {
		return nil
	}
	return o.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema
}
