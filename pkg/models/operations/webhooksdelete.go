// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type WebhooksDeleteSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *WebhooksDeleteSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type WebhooksDeleteRequest struct {
	// The ID of the webhook to delete
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *WebhooksDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// WebhooksDelete422ApplicationJSON - An error has occured, e.g. the identifier is not associated with an existing Bolt account
type WebhooksDelete422ApplicationJSON struct {
	// The type of error returned
	DotTag string `json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (o *WebhooksDelete422ApplicationJSON) GetDotTag() string {
	if o == nil {
		return ""
	}
	return o.DotTag
}

func (o *WebhooksDelete422ApplicationJSON) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type WebhooksDeleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// An error has occured, e.g. the identifier is not associated with an existing Bolt account
	WebhooksDelete422ApplicationJSONObject *WebhooksDelete422ApplicationJSON
}

func (o *WebhooksDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebhooksDeleteResponse) GetWebhooksDelete422ApplicationJSONObject() *WebhooksDelete422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.WebhooksDelete422ApplicationJSONObject
}
