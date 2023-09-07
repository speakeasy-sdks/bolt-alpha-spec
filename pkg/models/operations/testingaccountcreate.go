// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"net/http"
	"time"
)

type TestingAccountCreateSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *TestingAccountCreateSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type TestingAccountCreateRequestBodyEmailState string

const (
	TestingAccountCreateRequestBodyEmailStateMissing    TestingAccountCreateRequestBodyEmailState = "missing"
	TestingAccountCreateRequestBodyEmailStateUnverified TestingAccountCreateRequestBodyEmailState = "unverified"
	TestingAccountCreateRequestBodyEmailStateVerified   TestingAccountCreateRequestBodyEmailState = "verified"
)

func (e TestingAccountCreateRequestBodyEmailState) ToPointer() *TestingAccountCreateRequestBodyEmailState {
	return &e
}

func (e *TestingAccountCreateRequestBodyEmailState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "missing":
		fallthrough
	case "unverified":
		fallthrough
	case "verified":
		*e = TestingAccountCreateRequestBodyEmailState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TestingAccountCreateRequestBodyEmailState: %v", v)
	}
}

type TestingAccountCreateRequestBodyPhoneState string

const (
	TestingAccountCreateRequestBodyPhoneStateMissing    TestingAccountCreateRequestBodyPhoneState = "missing"
	TestingAccountCreateRequestBodyPhoneStateUnverified TestingAccountCreateRequestBodyPhoneState = "unverified"
	TestingAccountCreateRequestBodyPhoneStateVerified   TestingAccountCreateRequestBodyPhoneState = "verified"
)

func (e TestingAccountCreateRequestBodyPhoneState) ToPointer() *TestingAccountCreateRequestBodyPhoneState {
	return &e
}

func (e *TestingAccountCreateRequestBodyPhoneState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "missing":
		fallthrough
	case "unverified":
		fallthrough
	case "verified":
		*e = TestingAccountCreateRequestBodyPhoneState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TestingAccountCreateRequestBodyPhoneState: %v", v)
	}
}

type TestingAccountCreateRequestBodyInput struct {
	DeactivateAt time.Time                                 `json:"deactivate_at"`
	EmailState   TestingAccountCreateRequestBodyEmailState `json:"email_state"`
	HasAddress   *bool                                     `json:"has_address,omitempty"`
	IsMigrated   *bool                                     `json:"is_migrated,omitempty"`
	PhoneState   TestingAccountCreateRequestBodyPhoneState `json:"phone_state"`
}

func (o *TestingAccountCreateRequestBodyInput) GetDeactivateAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.DeactivateAt
}

func (o *TestingAccountCreateRequestBodyInput) GetEmailState() TestingAccountCreateRequestBodyEmailState {
	if o == nil {
		return TestingAccountCreateRequestBodyEmailState("")
	}
	return o.EmailState
}

func (o *TestingAccountCreateRequestBodyInput) GetHasAddress() *bool {
	if o == nil {
		return nil
	}
	return o.HasAddress
}

func (o *TestingAccountCreateRequestBodyInput) GetIsMigrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsMigrated
}

func (o *TestingAccountCreateRequestBodyInput) GetPhoneState() TestingAccountCreateRequestBodyPhoneState {
	if o == nil {
		return TestingAccountCreateRequestBodyPhoneState("")
	}
	return o.PhoneState
}

type TestingAccountCreateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The account was successfully created
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchema *shared.Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput
}

func (o *TestingAccountCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestingAccountCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestingAccountCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestingAccountCreateResponse) GetOnetesting1accountsPostRequestBodyContentApplication1jsonSchema() *shared.Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput {
	if o == nil {
		return nil
	}
	return o.Onetesting1accountsPostRequestBodyContentApplication1jsonSchema
}