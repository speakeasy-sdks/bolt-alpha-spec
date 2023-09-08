// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState string

const (
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailStateMissing    Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState = "missing"
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailStateUnverified Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState = "unverified"
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailStateVerified   Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState = "verified"
)

func (e Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState) ToPointer() *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState {
	return &e
}

func (e *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState) UnmarshalJSON(data []byte) error {
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
		*e = Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState: %v", v)
	}
}

type Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState string

const (
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneStateMissing    Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState = "missing"
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneStateUnverified Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState = "unverified"
	Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneStateVerified   Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState = "verified"
)

func (e Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState) ToPointer() *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState {
	return &e
}

func (e *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState) UnmarshalJSON(data []byte) error {
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
		*e = Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState: %v", v)
	}
}

type Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput struct {
	DeactivateAt time.Time                                                                 `json:"deactivate_at"`
	Email        string                                                                    `json:"email"`
	EmailState   Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState `json:"email_state"`
	OauthCode    string                                                                    `json:"oauth_code"`
	OtpCode      string                                                                    `json:"otp_code"`
	Phone        string                                                                    `json:"phone"`
	PhoneState   Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState `json:"phone_state"`
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetDeactivateAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.DeactivateAt
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetEmailState() Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState {
	if o == nil {
		return Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaEmailState("")
	}
	return o.EmailState
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetOauthCode() string {
	if o == nil {
		return ""
	}
	return o.OauthCode
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetOtpCode() string {
	if o == nil {
		return ""
	}
	return o.OtpCode
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

func (o *Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaOutput) GetPhoneState() Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState {
	if o == nil {
		return Onetesting1accountsPostRequestBodyContentApplication1jsonSchemaPhoneState("")
	}
	return o.PhoneState
}
