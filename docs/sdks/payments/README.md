# Payments

## Overview

Use the Payments API to tokenize and process alternative payment methods including Paypal with Bolt. This API is for the Bolt
Accounts package.


### Available Operations

* [GuestPaymentsInitialize](#guestpaymentsinitialize) - Initialize a Bolt payment for guest shoppers
* [PaymentsInitialize](#paymentsinitialize) - Initialize a Bolt payment for logged in shoppers

## GuestPaymentsInitialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for guest shoppers.


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltalphaspec "github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
)

func main() {
    s := boltalphaspec.New(
        boltalphaspec.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.GuestPaymentsInitialize(ctx, operations.GuestPaymentsInitializeRequest{
        RequestBody: operations.GuestPaymentsInitializeRequestBody{
            Cart: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCart{
                Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartAmounts{
                    Currency: "USD",
                    Tax: boltalphaspec.Int64(900),
                    Total: 900,
                },
                Discounts: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartDiscounts{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartDiscounts{
                        Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                        Code: boltalphaspec.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalphaspec.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalphaspec.String("215614191"),
                Items: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartItems{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartItems{
                        Description: boltalphaspec.String("Large tote with Bolt logo."),
                        ImageURL: boltalphaspec.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalphaspec.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipments{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipments{
                        Address: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipmentsAddressInput{},
                        Carrier: boltalphaspec.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.GuestPaymentsInitializeRequestBodyPaymentMethod{
                DotTag: "paypal",
                Cancel: "www.example.com/handle_paypal_cancel",
                Success: "www.example.com/handle_paypal_success",
            },
        },
        XPublishableKey: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OnepaymentsPostResponses200ContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GuestPaymentsInitializeRequest](../../models/operations/guestpaymentsinitializerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GuestPaymentsInitializeResponse](../../models/operations/guestpaymentsinitializeresponse.md), error**


## PaymentsInitialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for logged in shoppers.


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltalphaspec "github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
)

func main() {
    s := boltalphaspec.New()
    operationSecurity := operations.PaymentsInitializeSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Payments.PaymentsInitialize(ctx, operations.PaymentsInitializeRequest{
        RequestBody: operations.PaymentsInitializeRequestBody{
            Cart: operations.PaymentsInitializeRequestBodyCart{
                Amounts: operations.PaymentsInitializeRequestBodyCartAmounts{
                    Currency: "USD",
                    Tax: boltalphaspec.Int64(900),
                    Total: 900,
                },
                Discounts: []PaymentsInitializeRequestBodyCartDiscounts{
                    operations.PaymentsInitializeRequestBodyCartDiscounts{
                        Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                        Code: boltalphaspec.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalphaspec.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalphaspec.String("215614191"),
                Items: []PaymentsInitializeRequestBodyCartItems{
                    operations.PaymentsInitializeRequestBodyCartItems{
                        Description: boltalphaspec.String("Large tote with Bolt logo."),
                        ImageURL: boltalphaspec.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalphaspec.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []PaymentsInitializeRequestBodyCartShipments{
                    operations.PaymentsInitializeRequestBodyCartShipments{
                        Address: &operations.PaymentsInitializeRequestBodyCartShipmentsAddressInput{},
                        Carrier: boltalphaspec.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.PaymentsInitializeRequestBodyPaymentMethod{
                DotTag: "saved_payment_method",
                ID: "id",
            },
        },
        XPublishableKey: "suscipit",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentsInitialize200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PaymentsInitializeRequest](../../models/operations/paymentsinitializerequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.PaymentsInitializeSecurity](../../models/operations/paymentsinitializesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.PaymentsInitializeResponse](../../models/operations/paymentsinitializeresponse.md), error**

