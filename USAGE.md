<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	boltalphaspec "github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"log"
)

func main() {
	s := boltalphaspec.New()

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		RequestBody: operations.AccountAddPaymentMethodRequestBodyInput{
			BillingAddress: shared.CreateOnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesShipmentsItemsPropertiesAddressInputOnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesShipmentsItemsPropertiesAddressAddressID(
				shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesShipmentsItemsPropertiesAddressAddressID{
					DotTag: "id",
					ID:     "D4g3h5tBuVYK9",
				},
			),
			Bin:        "411111",
			Expiration: "2025-03",
			Last4:      "1004",
			Network:    operations.AccountAddPaymentMethodRequestBodyNetworkVisa,
			Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
			Type:       operations.AccountAddPaymentMethodRequestBodyTypeCredit,
		},
		XPublishableKey: "string",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.Oneaccount1paymentMethodsPostRequestBodyContentApplication1jsonSchema != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->