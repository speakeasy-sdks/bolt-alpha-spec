# PaymentsLoggedIn
(*Payments.LoggedIn*)

### Available Operations

* [Initialize](#initialize) - Initialize a Bolt payment for logged in shoppers
* [PerformAction](#performaction) - Perform an irreversible action (e.g. finalize) on a pending payment
* [Update](#update) - Update an existing payment

## Initialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for logged in shoppers.


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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Initialize(ctx, operations.PaymentsInitializeRequest{
        XPublishableKey: "string",
        PaymentInitializeRequest: shared.PaymentInitializeRequest{
            Cart: shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        Code: boltalphaspec.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalphaspec.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalphaspec.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: boltalphaspec.String("Large tote with Bolt logo."),
                        ImageURL: boltalphaspec.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalphaspec.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceExplicitInput(
                                shared.AddressReferenceAddressReferenceExplicitInput{
                                    DotTag: shared.AddressReferenceAddressReferenceExplicitTagExplicit,
                                    Company: boltalphaspec.String("ACME Corporation"),
                                    CountryCode: shared.AddressReferenceAddressReferenceExplicitCountryCodeUs,
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
                        Cost: &shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
            },
            PaymentMethod: shared.PaymentMethodReference{
                DotTag: shared.PaymentMethodReferenceTagID,
                ID: "id",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PaymentsInitializeRequest](../../models/operations/paymentsinitializerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PaymentsInitializeResponse](../../models/operations/paymentsinitializeresponse.md), error**


## PerformAction

Perform an irreversible action on a pending payment, such as finalizing it.


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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.PerformAction(ctx, operations.PaymentsActionRequest{
        XPublishableKey: "string",
        ID: "iKv7t5bgt1gg",
        PaymentActionRequest: shared.PaymentActionRequest{
            DotTag: shared.PaymentActionRequestTagFinalize,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PaymentsActionRequest](../../models/operations/paymentsactionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PaymentsActionResponse](../../models/operations/paymentsactionresponse.md), error**


## Update

Update a pending payment


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
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.LoggedIn.Update(ctx, operations.PaymentsUpdateRequest{
        XPublishableKey: "string",
        ID: "iKv7t5bgt1gg",
        PaymentUpdateRequest: shared.PaymentUpdateRequest{
            Cart: &shared.Cart{
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        Code: boltalphaspec.String("SUMMER10DISCOUNT"),
                        DetailsURL: boltalphaspec.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: boltalphaspec.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: boltalphaspec.String("Large tote with Bolt logo."),
                        ImageURL: boltalphaspec.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: boltalphaspec.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceID(
                                shared.AddressReferenceAddressReferenceID{
                                    DotTag: shared.AddressReferenceAddressReferenceIDTagID,
                                    ID: "D4g3h5tBuVYK9",
                                },
                        ),
                        Carrier: boltalphaspec.String("FedEx"),
                        Cost: &shared.Amount{
                            Currency: shared.AmountCurrencyUsd,
                            Units: 900,
                        },
                    },
                },
                Tax: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
                Total: shared.Amount{
                    Currency: shared.AmountCurrencyUsd,
                    Units: 900,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PaymentsUpdateRequest](../../models/operations/paymentsupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PaymentsUpdateResponse](../../models/operations/paymentsupdateresponse.md), error**

