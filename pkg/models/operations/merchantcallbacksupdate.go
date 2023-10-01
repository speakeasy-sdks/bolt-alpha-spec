// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/utils"
	"net/http"
)

type MerchantCallbacksUpdateRequestBody struct {
	AccountPage                   *string `json:"account_page,omitempty"`
	BaseDomain                    *string `json:"base_domain,omitempty"`
	ConfirmationRedirect          *string `json:"confirmation_redirect,omitempty"`
	CreateOrder                   *string `json:"create_order,omitempty"`
	Debug                         *string `json:"debug,omitempty"`
	GetAccount                    *string `json:"get_account,omitempty"`
	MobileAppDomain               *string `json:"mobile_app_domain,omitempty"`
	OauthLogout                   *string `json:"oauth_logout,omitempty"`
	OauthRedirect                 *string `json:"oauth_redirect,omitempty"`
	PrivacyPolicy                 *string `json:"privacy_policy,omitempty"`
	ProductInfo                   *string `json:"product_info,omitempty"`
	RemoteAPI                     *string `json:"remote_api,omitempty"`
	Shipping                      *string `json:"shipping,omitempty"`
	SupportPage                   *string `json:"support_page,omitempty"`
	Tax                           *string `json:"tax,omitempty"`
	TermsOfService                *string `json:"terms_of_service,omitempty"`
	UniversalMerchantAPI          *string `json:"universal_merchant_api,omitempty"`
	UpdateCart                    *string `json:"update_cart,omitempty"`
	ValidateAdditionalAccountData *string `json:"validate_additional_account_data,omitempty"`
}

func (o *MerchantCallbacksUpdateRequestBody) GetAccountPage() *string {
	if o == nil {
		return nil
	}
	return o.AccountPage
}

func (o *MerchantCallbacksUpdateRequestBody) GetBaseDomain() *string {
	if o == nil {
		return nil
	}
	return o.BaseDomain
}

func (o *MerchantCallbacksUpdateRequestBody) GetConfirmationRedirect() *string {
	if o == nil {
		return nil
	}
	return o.ConfirmationRedirect
}

func (o *MerchantCallbacksUpdateRequestBody) GetCreateOrder() *string {
	if o == nil {
		return nil
	}
	return o.CreateOrder
}

func (o *MerchantCallbacksUpdateRequestBody) GetDebug() *string {
	if o == nil {
		return nil
	}
	return o.Debug
}

func (o *MerchantCallbacksUpdateRequestBody) GetGetAccount() *string {
	if o == nil {
		return nil
	}
	return o.GetAccount
}

func (o *MerchantCallbacksUpdateRequestBody) GetMobileAppDomain() *string {
	if o == nil {
		return nil
	}
	return o.MobileAppDomain
}

func (o *MerchantCallbacksUpdateRequestBody) GetOauthLogout() *string {
	if o == nil {
		return nil
	}
	return o.OauthLogout
}

func (o *MerchantCallbacksUpdateRequestBody) GetOauthRedirect() *string {
	if o == nil {
		return nil
	}
	return o.OauthRedirect
}

func (o *MerchantCallbacksUpdateRequestBody) GetPrivacyPolicy() *string {
	if o == nil {
		return nil
	}
	return o.PrivacyPolicy
}

func (o *MerchantCallbacksUpdateRequestBody) GetProductInfo() *string {
	if o == nil {
		return nil
	}
	return o.ProductInfo
}

func (o *MerchantCallbacksUpdateRequestBody) GetRemoteAPI() *string {
	if o == nil {
		return nil
	}
	return o.RemoteAPI
}

func (o *MerchantCallbacksUpdateRequestBody) GetShipping() *string {
	if o == nil {
		return nil
	}
	return o.Shipping
}

func (o *MerchantCallbacksUpdateRequestBody) GetSupportPage() *string {
	if o == nil {
		return nil
	}
	return o.SupportPage
}

func (o *MerchantCallbacksUpdateRequestBody) GetTax() *string {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *MerchantCallbacksUpdateRequestBody) GetTermsOfService() *string {
	if o == nil {
		return nil
	}
	return o.TermsOfService
}

func (o *MerchantCallbacksUpdateRequestBody) GetUniversalMerchantAPI() *string {
	if o == nil {
		return nil
	}
	return o.UniversalMerchantAPI
}

func (o *MerchantCallbacksUpdateRequestBody) GetUpdateCart() *string {
	if o == nil {
		return nil
	}
	return o.UpdateCart
}

func (o *MerchantCallbacksUpdateRequestBody) GetValidateAdditionalAccountData() *string {
	if o == nil {
		return nil
	}
	return o.ValidateAdditionalAccountData
}

type MerchantCallbacksUpdateRequest struct {
	RequestBody MerchantCallbacksUpdateRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *MerchantCallbacksUpdateRequest) GetRequestBody() MerchantCallbacksUpdateRequestBody {
	if o == nil {
		return MerchantCallbacksUpdateRequestBody{}
	}
	return o.RequestBody
}

func (o *MerchantCallbacksUpdateRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

type MerchantCallbacksUpdate400ApplicationJSON1 struct {
	// The type of error returned
	dotTag string `const:"invalid_url" json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (m MerchantCallbacksUpdate400ApplicationJSON1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MerchantCallbacksUpdate400ApplicationJSON1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MerchantCallbacksUpdate400ApplicationJSON1) GetDotTag() string {
	return "invalid_url"
}

func (o *MerchantCallbacksUpdate400ApplicationJSON1) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type MerchantCallbacksUpdate400ApplicationJSONType string

const (
	MerchantCallbacksUpdate400ApplicationJSONTypeMerchantCallbacksUpdate400ApplicationJSON1 MerchantCallbacksUpdate400ApplicationJSONType = "merchantCallbacksUpdate_400ApplicationJSON_1"
)

type MerchantCallbacksUpdate400ApplicationJSON struct {
	MerchantCallbacksUpdate400ApplicationJSON1 *MerchantCallbacksUpdate400ApplicationJSON1

	Type MerchantCallbacksUpdate400ApplicationJSONType
}

func CreateMerchantCallbacksUpdate400ApplicationJSONMerchantCallbacksUpdate400ApplicationJSON1(merchantCallbacksUpdate400ApplicationJSON1 MerchantCallbacksUpdate400ApplicationJSON1) MerchantCallbacksUpdate400ApplicationJSON {
	typ := MerchantCallbacksUpdate400ApplicationJSONTypeMerchantCallbacksUpdate400ApplicationJSON1

	return MerchantCallbacksUpdate400ApplicationJSON{
		MerchantCallbacksUpdate400ApplicationJSON1: &merchantCallbacksUpdate400ApplicationJSON1,
		Type: typ,
	}
}

func (u *MerchantCallbacksUpdate400ApplicationJSON) UnmarshalJSON(data []byte) error {

	merchantCallbacksUpdate400ApplicationJSON1 := new(MerchantCallbacksUpdate400ApplicationJSON1)
	if err := utils.UnmarshalJSON(data, &merchantCallbacksUpdate400ApplicationJSON1, "", true, true); err == nil {
		u.MerchantCallbacksUpdate400ApplicationJSON1 = merchantCallbacksUpdate400ApplicationJSON1
		u.Type = MerchantCallbacksUpdate400ApplicationJSONTypeMerchantCallbacksUpdate400ApplicationJSON1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u MerchantCallbacksUpdate400ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.MerchantCallbacksUpdate400ApplicationJSON1 != nil {
		return utils.MarshalJSON(u.MerchantCallbacksUpdate400ApplicationJSON1, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type MerchantCallbacksUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request is missing required fields, or its fields have invalid values
	MerchantCallbacksUpdate400ApplicationJSONOneOf *MerchantCallbacksUpdate400ApplicationJSON
	// Callbacks URLs were successfully updated
	Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema *shared.Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema
}

func (o *MerchantCallbacksUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MerchantCallbacksUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MerchantCallbacksUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MerchantCallbacksUpdateResponse) GetMerchantCallbacksUpdate400ApplicationJSONOneOf() *MerchantCallbacksUpdate400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.MerchantCallbacksUpdate400ApplicationJSONOneOf
}

func (o *MerchantCallbacksUpdateResponse) GetOnemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema() *shared.Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema {
	if o == nil {
		return nil
	}
	return o.Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema
}
