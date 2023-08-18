// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MerchantIdentifiersGetSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *MerchantIdentifiersGetSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type MerchantIdentifiersGet200ApplicationJSONMerchantDivisions struct {
	PublishableKey string `json:"publishable_key"`
}

func (o *MerchantIdentifiersGet200ApplicationJSONMerchantDivisions) GetPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.PublishableKey
}

// MerchantIdentifiersGet200ApplicationJSON - Identifiers were successfully retrieved
type MerchantIdentifiersGet200ApplicationJSON struct {
	MerchantDivisions []MerchantIdentifiersGet200ApplicationJSONMerchantDivisions `json:"merchant_divisions"`
	MerchantID        string                                                      `json:"merchant_id"`
	SigningSecret     string                                                      `json:"signing_secret"`
}

func (o *MerchantIdentifiersGet200ApplicationJSON) GetMerchantDivisions() []MerchantIdentifiersGet200ApplicationJSONMerchantDivisions {
	if o == nil {
		return []MerchantIdentifiersGet200ApplicationJSONMerchantDivisions{}
	}
	return o.MerchantDivisions
}

func (o *MerchantIdentifiersGet200ApplicationJSON) GetMerchantID() string {
	if o == nil {
		return ""
	}
	return o.MerchantID
}

func (o *MerchantIdentifiersGet200ApplicationJSON) GetSigningSecret() string {
	if o == nil {
		return ""
	}
	return o.SigningSecret
}

type MerchantIdentifiersGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Identifiers were successfully retrieved
	MerchantIdentifiersGet200ApplicationJSONObject *MerchantIdentifiersGet200ApplicationJSON
}

func (o *MerchantIdentifiersGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MerchantIdentifiersGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MerchantIdentifiersGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MerchantIdentifiersGetResponse) GetMerchantIdentifiersGet200ApplicationJSONObject() *MerchantIdentifiersGet200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MerchantIdentifiersGet200ApplicationJSONObject
}
