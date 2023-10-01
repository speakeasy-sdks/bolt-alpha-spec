// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/utils"
	"time"
)

type OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventType string

const (
	OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeGroup OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventType = "group"
	OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeList  OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventType = "list"
)

type OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent struct {
	EventGroup *EventGroup
	EventList  *EventList

	Type OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventType
}

func CreateOnewebhooksPutRequestBodyContentApplication1jsonSchemaEventGroup(group EventGroup) OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent {
	typ := OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeGroup

	return OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent{
		EventGroup: &group,
		Type:       typ,
	}
}

func CreateOnewebhooksPutRequestBodyContentApplication1jsonSchemaEventList(list EventList) OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent {
	typ := OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeList

	return OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent{
		EventList: &list,
		Type:      typ,
	}
}

func (u *OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "group":
		eventGroup := new(EventGroup)
		if err := utils.UnmarshalJSON(data, &eventGroup, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventGroup = eventGroup
		u.Type = OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeGroup
		return nil
	case "list":
		eventList := new(EventList)
		if err := utils.UnmarshalJSON(data, &eventList, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventList = eventList
		u.Type = OnewebhooksPutRequestBodyContentApplication1jsonSchemaEventTypeList
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent) MarshalJSON() ([]byte, error) {
	if u.EventGroup != nil {
		return utils.MarshalJSON(u.EventGroup, "", true)
	}

	if u.EventList != nil {
		return utils.MarshalJSON(u.EventList, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type OnewebhooksPutRequestBodyContentApplication1jsonSchema struct {
	// The time at which the webhook was created
	CreatedAt *time.Time                                                  `json:"created_at,omitempty"`
	Event     OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent `json:"event"`
	// The webhook's ID
	ID *string `json:"id,omitempty"`
	// The webhook's URL
	URL string `json:"url"`
}

func (o OnewebhooksPutRequestBodyContentApplication1jsonSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetEvent() OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent {
	if o == nil {
		return OnewebhooksPutRequestBodyContentApplication1jsonSchemaEvent{}
	}
	return o.Event
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetEventGroup() *EventGroup {
	return o.GetEvent().EventGroup
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetEventList() *EventList {
	return o.GetEvent().EventList
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OnewebhooksPutRequestBodyContentApplication1jsonSchema) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
