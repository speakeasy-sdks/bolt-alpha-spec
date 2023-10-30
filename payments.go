// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltalphaspec

type payments struct {
	Guest    *paymentsGuest
	LoggedIn *paymentsLoggedIn

	sdkConfiguration sdkConfiguration
}

func newPayments(sdkConfig sdkConfiguration) *payments {
	return &payments{
		sdkConfiguration: sdkConfig,
		Guest:            newPaymentsGuest(sdkConfig),
		LoggedIn:         newPaymentsLoggedIn(sdkConfig),
	}
}
