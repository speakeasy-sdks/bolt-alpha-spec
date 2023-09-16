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
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.GuestPaymentsInitializeSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Payments.GuestPaymentsInitialize(ctx, operations.GuestPaymentsInitializeRequest{
        RequestBody: operations.GuestPaymentsInitializeRequestBody{
            Cart: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCart{
                Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartAmounts{
                    Currency: "USD",
                    Tax: boltalpha.Int64(900),
                    Total: 900,
                },
                Discounts: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartDiscounts{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartDiscounts{
                        Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalpha.Int64(900),
                            Total: 900,
                        },
                        Code: boltalpha.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalpha.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalpha.String("215614191"),
                Items: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartItems{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartItems{
                        Description: boltalpha.String("Large tote with Bolt logo."),
                        ImageURL: boltalpha.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalpha.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipments{
                    shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipments{
                        Address: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipmentsAddressInput{},
                        Carrier: boltalpha.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalpha.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.GuestPaymentsInitializeRequestBodyPaymentMethod{
                DotTag: operations.GuestPaymentsInitializeRequestBodyPaymentMethodTagPaypal,
                Cancel: "www.example.com/handle_paypal_cancel",
                Success: "www.example.com/handle_paypal_success",
            },
        },
        XPublishableKey: "error",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.OnepaymentsPostResponses200ContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GuestPaymentsInitializeRequest](../../models/operations/guestpaymentsinitializerequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GuestPaymentsInitializeSecurity](../../models/operations/guestpaymentsinitializesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
)

func main() {
    s := boltalpha.New()
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
                    Tax: boltalpha.Int64(900),
                    Total: 900,
                },
                Discounts: []PaymentsInitializeRequestBodyCartDiscounts{
                    operations.PaymentsInitializeRequestBodyCartDiscounts{
                        Amounts: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalpha.Int64(900),
                            Total: 900,
                        },
                        Code: boltalpha.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalpha.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalpha.String("215614191"),
                Items: []PaymentsInitializeRequestBodyCartItems{
                    operations.PaymentsInitializeRequestBodyCartItems{
                        Description: boltalpha.String("Large tote with Bolt logo."),
                        ImageURL: boltalpha.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalpha.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []PaymentsInitializeRequestBodyCartShipments{
                    operations.PaymentsInitializeRequestBodyCartShipments{
                        Address: &operations.PaymentsInitializeRequestBodyCartShipmentsAddressInput{},
                        Carrier: boltalpha.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalpha.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.PaymentsInitializeRequestBodyPaymentMethod{
                DotTag: operations.PaymentsInitializeRequestBodyPaymentMethodTagSavedPaymentMethod,
                ID: "id",
            },
        },
        XPublishableKey: "deserunt",
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

