// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/utils"
	"net/http"
)

type WebhooksCreateRequestBodyEventType string

const (
	WebhooksCreateRequestBodyEventTypeGroup WebhooksCreateRequestBodyEventType = "group"
	WebhooksCreateRequestBodyEventTypeList  WebhooksCreateRequestBodyEventType = "list"
)

type WebhooksCreateRequestBodyEvent struct {
	EventGroup *shared.EventGroup
	EventList  *shared.EventList

	Type WebhooksCreateRequestBodyEventType
}

func CreateWebhooksCreateRequestBodyEventGroup(group shared.EventGroup) WebhooksCreateRequestBodyEvent {
	typ := WebhooksCreateRequestBodyEventTypeGroup
	typStr := shared.EventGroupTag(typ)
	group.DotTag = typStr

	return WebhooksCreateRequestBodyEvent{
		EventGroup: &group,
		Type:       typ,
	}
}

func CreateWebhooksCreateRequestBodyEventList(list shared.EventList) WebhooksCreateRequestBodyEvent {
	typ := WebhooksCreateRequestBodyEventTypeList
	typStr := shared.EventListTag(typ)
	list.DotTag = typStr

	return WebhooksCreateRequestBodyEvent{
		EventList: &list,
		Type:      typ,
	}
}

func (u *WebhooksCreateRequestBodyEvent) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "group":
		eventGroup := new(shared.EventGroup)
		if err := utils.UnmarshalJSON(data, &eventGroup, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventGroup = eventGroup
		u.Type = WebhooksCreateRequestBodyEventTypeGroup
		return nil
	case "list":
		eventList := new(shared.EventList)
		if err := utils.UnmarshalJSON(data, &eventList, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventList = eventList
		u.Type = WebhooksCreateRequestBodyEventTypeList
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u WebhooksCreateRequestBodyEvent) MarshalJSON() ([]byte, error) {
	if u.EventGroup != nil {
		return utils.MarshalJSON(u.EventGroup, "", true)
	}

	if u.EventList != nil {
		return utils.MarshalJSON(u.EventList, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type WebhooksCreateRequestBodyInput struct {
	Event WebhooksCreateRequestBodyEvent `json:"event"`
	// The webhook's URL
	URL string `json:"url"`
}

func (o *WebhooksCreateRequestBodyInput) GetEvent() WebhooksCreateRequestBodyEvent {
	if o == nil {
		return WebhooksCreateRequestBodyEvent{}
	}
	return o.Event
}

func (o *WebhooksCreateRequestBodyInput) GetEventGroup() *shared.EventGroup {
	return o.GetEvent().EventGroup
}

func (o *WebhooksCreateRequestBodyInput) GetEventList() *shared.EventList {
	return o.GetEvent().EventList
}

func (o *WebhooksCreateRequestBodyInput) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type WebhooksCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The webhook was successfully created
	OnewebhooksPutRequestBodyContentApplication1jsonSchema *shared.OnewebhooksPutRequestBodyContentApplication1jsonSchema
}

func (o *WebhooksCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebhooksCreateResponse) GetOnewebhooksPutRequestBodyContentApplication1jsonSchema() *shared.OnewebhooksPutRequestBodyContentApplication1jsonSchema {
	if o == nil {
		return nil
	}
	return o.OnewebhooksPutRequestBodyContentApplication1jsonSchema
}
