// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"net/http"
)

type WebhooksGetAllRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *WebhooksGetAllRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

// WebhooksGetAll200ApplicationJSON - All existing webhook information has been retrieved
type WebhooksGetAll200ApplicationJSON struct {
	Webhooks []shared.OnewebhooksPutRequestBodyContentApplication1jsonSchema `json:"webhooks,omitempty"`
}

func (o *WebhooksGetAll200ApplicationJSON) GetWebhooks() []shared.OnewebhooksPutRequestBodyContentApplication1jsonSchema {
	if o == nil {
		return nil
	}
	return o.Webhooks
}

type WebhooksGetAllResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// All existing webhook information has been retrieved
	WebhooksGetAll200ApplicationJSONObject *WebhooksGetAll200ApplicationJSON
}

func (o *WebhooksGetAllResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksGetAllResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksGetAllResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebhooksGetAllResponse) GetWebhooksGetAll200ApplicationJSONObject() *WebhooksGetAll200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.WebhooksGetAll200ApplicationJSONObject
}
