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

type PaymentsInitializeSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
	Oauth  string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *PaymentsInitializeSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *PaymentsInitializeSecurity) GetOauth() string {
	if o == nil {
		return ""
	}
	return o.Oauth
}

type PaymentsInitializeRequestBodyCartAmounts struct {
	Currency string `json:"currency"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax *int64 `json:"tax,omitempty"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total int64 `json:"total"`
}

func (o *PaymentsInitializeRequestBodyCartAmounts) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *PaymentsInitializeRequestBodyCartAmounts) GetTax() *int64 {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *PaymentsInitializeRequestBodyCartAmounts) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

type PaymentsInitializeRequestBodyCartDiscounts struct {
	Amounts shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts `json:"amounts"`
	// Discount code.
	Code *string `json:"code,omitempty"`
	// Used to provide a link to additional details, such as a landing page, associated with the discount offering.
	DetailsURL *string `json:"details_url,omitempty"`
}

func (o *PaymentsInitializeRequestBodyCartDiscounts) GetAmounts() shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts {
	if o == nil {
		return shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{}
	}
	return o.Amounts
}

func (o *PaymentsInitializeRequestBodyCartDiscounts) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *PaymentsInitializeRequestBodyCartDiscounts) GetDetailsURL() *string {
	if o == nil {
		return nil
	}
	return o.DetailsURL
}

type PaymentsInitializeRequestBodyCartItems struct {
	// A human-readable description of this cart item.
	Description *string `json:"description,omitempty"`
	// Used to provide a link to the image associated with the item.
	ImageURL *string `json:"image_url,omitempty"`
	// The name of a given item.
	Name string `json:"name"`
	// The number of units that comprise this cart item.
	Quantity int64 `json:"quantity"`
	// This value is used by Bolt as an external reference to a given item.
	Reference string `json:"reference"`
	// The total amount, in cents, of the item including its taxes if applicable.
	TotalAmount int64 `json:"total_amount"`
	// The price of one unit of the item; for example, the price of one pack of socks.
	UnitPrice int64 `json:"unit_price"`
}

func (o *PaymentsInitializeRequestBodyCartItems) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PaymentsInitializeRequestBodyCartItems) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *PaymentsInitializeRequestBodyCartItems) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PaymentsInitializeRequestBodyCartItems) GetQuantity() int64 {
	if o == nil {
		return 0
	}
	return o.Quantity
}

func (o *PaymentsInitializeRequestBodyCartItems) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *PaymentsInitializeRequestBodyCartItems) GetTotalAmount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalAmount
}

func (o *PaymentsInitializeRequestBodyCartItems) GetUnitPrice() int64 {
	if o == nil {
		return 0
	}
	return o.UnitPrice
}

type PaymentsInitializeRequestBodyCartShipmentsAddressInputType string

const (
	PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeExplicit PaymentsInitializeRequestBodyCartShipmentsAddressInputType = "explicit"
	PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeID       PaymentsInitializeRequestBodyCartShipmentsAddressInputType = "id"
)

type PaymentsInitializeRequestBodyCartShipmentsAddressInput struct {
	AddressID            *shared.AddressID
	AddressExplicitInput *shared.AddressExplicitInput

	Type PaymentsInitializeRequestBodyCartShipmentsAddressInputType
}

func CreatePaymentsInitializeRequestBodyCartShipmentsAddressInputExplicit(explicit shared.AddressExplicitInput) PaymentsInitializeRequestBodyCartShipmentsAddressInput {
	typ := PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeExplicit
	typStr := string(typ)
	explicit.DotTag = typStr

	return PaymentsInitializeRequestBodyCartShipmentsAddressInput{
		AddressExplicitInput: &explicit,
		Type:                 typ,
	}
}

func CreatePaymentsInitializeRequestBodyCartShipmentsAddressInputID(id shared.AddressID) PaymentsInitializeRequestBodyCartShipmentsAddressInput {
	typ := PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeID
	typStr := string(typ)
	id.DotTag = typStr

	return PaymentsInitializeRequestBodyCartShipmentsAddressInput{
		AddressID: &id,
		Type:      typ,
	}
}

func (u *PaymentsInitializeRequestBodyCartShipmentsAddressInput) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "explicit":
		d = json.NewDecoder(bytes.NewReader(data))
		addressExplicitInput := new(shared.AddressExplicitInput)
		if err := d.Decode(&addressExplicitInput); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AddressExplicitInput = addressExplicitInput
		u.Type = PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeExplicit
		return nil
	case "id":
		d = json.NewDecoder(bytes.NewReader(data))
		addressID := new(shared.AddressID)
		if err := d.Decode(&addressID); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AddressID = addressID
		u.Type = PaymentsInitializeRequestBodyCartShipmentsAddressInputTypeID
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentsInitializeRequestBodyCartShipmentsAddressInput) MarshalJSON() ([]byte, error) {
	if u.AddressID != nil {
		return json.Marshal(u.AddressID)
	}

	if u.AddressExplicitInput != nil {
		return json.Marshal(u.AddressExplicitInput)
	}

	return nil, errors.New("could not marshal union type: all fields are null")

}

type PaymentsInitializeRequestBodyCartShipments struct {
	// The Address object is used for shipping, and physical store address use cases.
	Address *PaymentsInitializeRequestBodyCartShipmentsAddressInput `json:"address,omitempty"`
	// The name of the carrier selected.
	Carrier *string                                                                                        `json:"carrier,omitempty"`
	Cost    *shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts `json:"cost,omitempty"`
}

func (o *PaymentsInitializeRequestBodyCartShipments) GetAddress() *PaymentsInitializeRequestBodyCartShipmentsAddressInput {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *PaymentsInitializeRequestBodyCartShipments) GetAddressExplicit() *shared.AddressExplicitInput {
	if v := o.GetAddress(); v != nil {
		return v.AddressExplicitInput
	}
	return nil
}

func (o *PaymentsInitializeRequestBodyCartShipments) GetAddressID() *shared.AddressID {
	if v := o.GetAddress(); v != nil {
		return v.AddressID
	}
	return nil
}

func (o *PaymentsInitializeRequestBodyCartShipments) GetCarrier() *string {
	if o == nil {
		return nil
	}
	return o.Carrier
}

func (o *PaymentsInitializeRequestBodyCartShipments) GetCost() *shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts {
	if o == nil {
		return nil
	}
	return o.Cost
}

type PaymentsInitializeRequestBodyCart struct {
	Amounts   PaymentsInitializeRequestBodyCartAmounts     `json:"amounts"`
	Discounts []PaymentsInitializeRequestBodyCartDiscounts `json:"discounts,omitempty"`
	// This field corresponds to the merchant's order reference associated with this Bolt transaction.
	DisplayID *string                                  `json:"display_id,omitempty"`
	Items     []PaymentsInitializeRequestBodyCartItems `json:"items,omitempty"`
	// Used optionally to pass additional information like order numbers or other IDs as needed.
	OrderDescription *string `json:"order_description,omitempty"`
	// This value is used by Bolt as an external reference to a given order. This reference must be unique per successful transaction.
	OrderReference string                                       `json:"order_reference"`
	Shipments      []PaymentsInitializeRequestBodyCartShipments `json:"shipments,omitempty"`
}

func (o *PaymentsInitializeRequestBodyCart) GetAmounts() PaymentsInitializeRequestBodyCartAmounts {
	if o == nil {
		return PaymentsInitializeRequestBodyCartAmounts{}
	}
	return o.Amounts
}

func (o *PaymentsInitializeRequestBodyCart) GetDiscounts() []PaymentsInitializeRequestBodyCartDiscounts {
	if o == nil {
		return nil
	}
	return o.Discounts
}

func (o *PaymentsInitializeRequestBodyCart) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *PaymentsInitializeRequestBodyCart) GetItems() []PaymentsInitializeRequestBodyCartItems {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *PaymentsInitializeRequestBodyCart) GetOrderDescription() *string {
	if o == nil {
		return nil
	}
	return o.OrderDescription
}

func (o *PaymentsInitializeRequestBodyCart) GetOrderReference() string {
	if o == nil {
		return ""
	}
	return o.OrderReference
}

func (o *PaymentsInitializeRequestBodyCart) GetShipments() []PaymentsInitializeRequestBodyCartShipments {
	if o == nil {
		return nil
	}
	return o.Shipments
}

type PaymentsInitializeRequestBodyPaymentMethodTag string

const (
	PaymentsInitializeRequestBodyPaymentMethodTagSavedPaymentMethod PaymentsInitializeRequestBodyPaymentMethodTag = "saved_payment_method"
)

func (e PaymentsInitializeRequestBodyPaymentMethodTag) ToPointer() *PaymentsInitializeRequestBodyPaymentMethodTag {
	return &e
}

func (e *PaymentsInitializeRequestBodyPaymentMethodTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saved_payment_method":
		*e = PaymentsInitializeRequestBodyPaymentMethodTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentsInitializeRequestBodyPaymentMethodTag: %v", v)
	}
}

type PaymentsInitializeRequestBodyPaymentMethod struct {
	DotTag PaymentsInitializeRequestBodyPaymentMethodTag `json:".tag"`
	// Payment ID of the saved Bolt Payment method.
	ID string `json:"id"`
}

func (o *PaymentsInitializeRequestBodyPaymentMethod) GetDotTag() PaymentsInitializeRequestBodyPaymentMethodTag {
	if o == nil {
		return PaymentsInitializeRequestBodyPaymentMethodTag("")
	}
	return o.DotTag
}

func (o *PaymentsInitializeRequestBodyPaymentMethod) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PaymentsInitializeRequestBody struct {
	Cart          PaymentsInitializeRequestBodyCart          `json:"cart"`
	PaymentMethod PaymentsInitializeRequestBodyPaymentMethod `json:"payment_method"`
}

func (o *PaymentsInitializeRequestBody) GetCart() PaymentsInitializeRequestBodyCart {
	if o == nil {
		return PaymentsInitializeRequestBodyCart{}
	}
	return o.Cart
}

func (o *PaymentsInitializeRequestBody) GetPaymentMethod() PaymentsInitializeRequestBodyPaymentMethod {
	if o == nil {
		return PaymentsInitializeRequestBodyPaymentMethod{}
	}
	return o.PaymentMethod
}

type PaymentsInitializeRequest struct {
	RequestBody PaymentsInitializeRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *PaymentsInitializeRequest) GetRequestBody() PaymentsInitializeRequestBody {
	if o == nil {
		return PaymentsInitializeRequestBody{}
	}
	return o.RequestBody
}

func (o *PaymentsInitializeRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

type PaymentsInitialize200ApplicationJSONActionMethod string

const (
	PaymentsInitialize200ApplicationJSONActionMethodGet  PaymentsInitialize200ApplicationJSONActionMethod = "GET"
	PaymentsInitialize200ApplicationJSONActionMethodPost PaymentsInitialize200ApplicationJSONActionMethod = "POST"
)

func (e PaymentsInitialize200ApplicationJSONActionMethod) ToPointer() *PaymentsInitialize200ApplicationJSONActionMethod {
	return &e
}

func (e *PaymentsInitialize200ApplicationJSONActionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		*e = PaymentsInitialize200ApplicationJSONActionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentsInitialize200ApplicationJSONActionMethod: %v", v)
	}
}

type PaymentsInitialize200ApplicationJSONActionType string

const (
	PaymentsInitialize200ApplicationJSONActionTypeRedirect PaymentsInitialize200ApplicationJSONActionType = "redirect"
	PaymentsInitialize200ApplicationJSONActionTypeFinalize PaymentsInitialize200ApplicationJSONActionType = "finalize"
)

func (e PaymentsInitialize200ApplicationJSONActionType) ToPointer() *PaymentsInitialize200ApplicationJSONActionType {
	return &e
}

func (e *PaymentsInitialize200ApplicationJSONActionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redirect":
		fallthrough
	case "finalize":
		*e = PaymentsInitialize200ApplicationJSONActionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentsInitialize200ApplicationJSONActionType: %v", v)
	}
}

type PaymentsInitialize200ApplicationJSONAction struct {
	Method *PaymentsInitialize200ApplicationJSONActionMethod `json:"method,omitempty"`
	Type   *PaymentsInitialize200ApplicationJSONActionType   `json:"type,omitempty"`
	URL    *string                                           `json:"url,omitempty"`
}

func (o *PaymentsInitialize200ApplicationJSONAction) GetMethod() *PaymentsInitialize200ApplicationJSONActionMethod {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *PaymentsInitialize200ApplicationJSONAction) GetType() *PaymentsInitialize200ApplicationJSONActionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PaymentsInitialize200ApplicationJSONAction) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type PaymentsInitialize200ApplicationJSONStatus string

const (
	PaymentsInitialize200ApplicationJSONStatusAwaitingUserConfirmation PaymentsInitialize200ApplicationJSONStatus = "awaiting_user_confirmation"
	PaymentsInitialize200ApplicationJSONStatusPaymentReady             PaymentsInitialize200ApplicationJSONStatus = "payment_ready"
	PaymentsInitialize200ApplicationJSONStatusUpdatePaymentMethod      PaymentsInitialize200ApplicationJSONStatus = "update_payment_method"
	PaymentsInitialize200ApplicationJSONStatusSuccess                  PaymentsInitialize200ApplicationJSONStatus = "success"
)

func (e PaymentsInitialize200ApplicationJSONStatus) ToPointer() *PaymentsInitialize200ApplicationJSONStatus {
	return &e
}

func (e *PaymentsInitialize200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "awaiting_user_confirmation":
		fallthrough
	case "payment_ready":
		fallthrough
	case "update_payment_method":
		fallthrough
	case "success":
		*e = PaymentsInitialize200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentsInitialize200ApplicationJSONStatus: %v", v)
	}
}

// PaymentsInitialize200ApplicationJSON - Payment token retrieved
type PaymentsInitialize200ApplicationJSON struct {
	Action *PaymentsInitialize200ApplicationJSONAction `json:"action,omitempty"`
	ID     *string                                     `json:"id,omitempty"`
	Status *PaymentsInitialize200ApplicationJSONStatus `json:"status,omitempty"`
}

func (o *PaymentsInitialize200ApplicationJSON) GetAction() *PaymentsInitialize200ApplicationJSONAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *PaymentsInitialize200ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymentsInitialize200ApplicationJSON) GetStatus() *PaymentsInitialize200ApplicationJSONStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type PaymentsInitializeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Payment token retrieved
	PaymentsInitialize200ApplicationJSONObject *PaymentsInitialize200ApplicationJSON
}

func (o *PaymentsInitializeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PaymentsInitializeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PaymentsInitializeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PaymentsInitializeResponse) GetPaymentsInitialize200ApplicationJSONObject() *PaymentsInitialize200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PaymentsInitialize200ApplicationJSONObject
}
