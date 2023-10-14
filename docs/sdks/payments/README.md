# Payments
(*Payments*)

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
        boltalphaspec.WithSecurity(""),
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
                        Address: shared.CreateOnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipmentsAddressInputOnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipmentsAddressAddressID(
                                shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartShipmentsAddressAddressID{
                                    DotTag: "id",
                                    ID: "D4g3h5tBuVYK9",
                                },
                        ),
                        Carrier: boltalphaspec.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.CreateGuestPaymentsInitializeRequestBodyPaymentMethodGuestPaymentsInitializeRequestBodyPaymentMethod1(
                    operations.GuestPaymentsInitializeRequestBodyPaymentMethod1{
                        DotTag: operations.GuestPaymentsInitializeRequestBodyPaymentMethod1TagPaypal,
                        Cancel: "www.example.com/handle_paypal_cancel",
                        Success: "www.example.com/handle_paypal_success",
                    },
            ),
        },
        XPublishableKey: "Hyundai",
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
                Discounts: []operations.PaymentsInitializeRequestBodyCartDiscounts{
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
                Items: []operations.PaymentsInitializeRequestBodyCartItems{
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
                Shipments: []operations.PaymentsInitializeRequestBodyCartShipments{
                    operations.PaymentsInitializeRequestBodyCartShipments{
                        Address: operations.CreatePaymentsInitializeRequestBodyCartShipmentsAddressInputPaymentsInitializeRequestBodyCartShipmentsAddressAddressExplicitInput(
                                operations.PaymentsInitializeRequestBodyCartShipmentsAddressAddressExplicitInput{
                                    DotTag: "explicit",
                                    Company: boltalphaspec.String("ACME Corporation"),
                                    CountryCode: "US",
                                    Email: boltalphaspec.String("alice@example.com"),
                                    FirstName: "Alice",
                                    LastName: "Baker",
                                    Locality: "San Francisco",
                                    Phone: boltalphaspec.String("+14155550199"),
                                    PostalCode: "94105",
                                    Region: boltalphaspec.String("CA"),
                                    StreetAddress1: "535 Mission St, Ste 1401",
                                    StreetAddress2: boltalphaspec.String("c/o Shipping Department"),
                                },
                        ),
                        Carrier: boltalphaspec.String("FedEx"),
                        Cost: &shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts{
                            Currency: "USD",
                            Tax: boltalphaspec.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: operations.CreatePaymentsInitializeRequestBodyPaymentMethodPaymentsInitializeRequestBodyPaymentMethod1(
                    operations.PaymentsInitializeRequestBodyPaymentMethod1{
                        DotTag: operations.PaymentsInitializeRequestBodyPaymentMethod1TagSavedPaymentMethod,
                        ID: "id",
                    },
            ),
        },
        XPublishableKey: "Hybrid South",
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

