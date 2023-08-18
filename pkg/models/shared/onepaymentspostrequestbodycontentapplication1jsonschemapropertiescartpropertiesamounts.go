// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts struct {
	Currency string `json:"currency"`
	// The total tax amount, in cents for all of the items associated with the cart.
	Tax *int64 `json:"tax,omitempty"`
	// The total amount, in cents, including its items and taxes (if applicable).
	Total int64 `json:"total"`
}

func (o *OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts) GetTax() *int64 {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}
