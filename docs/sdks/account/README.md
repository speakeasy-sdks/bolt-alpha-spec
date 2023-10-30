# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [AddAddress](#addaddress) - Add an address
* [AddPaymentMethod](#addpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [DeleteAddress](#deleteaddress) - Delete an existing address
* [DeletePaymentMethod](#deletepaymentmethod) - Delete an existing payment method
* [Detect](#detect) - Determine the existence of a Bolt account
* [GetDetails](#getdetails) - Retrieve account details
* [UpdateAddress](#updateaddress) - Edit an existing address

## AddAddress

Add an address to the shopper's account

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
    res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
        XPublishableKey: "string",
        AddressListingInput: shared.AddressListingInput{
            Company: boltalphaspec.String("ACME Corporation"),
            CountryCode: shared.AddressListingCountryCodeUs,
            Email: boltalphaspec.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: boltalphaspec.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: boltalphaspec.String("+14155550199"),
            PostalCode: "94105",
            Region: boltalphaspec.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: boltalphaspec.String("c/o Shipping Department"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountAddressCreateRequest](../../models/operations/accountaddresscreaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.AccountAddressCreateResponse](../../models/operations/accountaddresscreateresponse.md), error**


## AddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from
your backend because authentication requires the use of your private key.<br />
**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function,
which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


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
    res, err := s.Account.AddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
        XPublishableKey: "string",
        PaymentMethodInput: shared.CreatePaymentMethodInputPaymentMethodCreditCardInput(
                shared.PaymentMethodCreditCardInput{
                    DotTag: shared.PaymentMethodCreditCardTagCreditCard,
                    BillingAddress: shared.CreateAddressReferenceInputAddressReferenceAddressReferenceID(
                            shared.AddressReferenceAddressReferenceID{
                                DotTag: shared.AddressReferenceAddressReferenceIDTagID,
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Bin: "411111",
                    Expiration: "2029-03",
                    Last4: "1004",
                    Network: shared.PaymentMethodCreditCardNetworkVisa,
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AccountAddPaymentMethodRequest](../../models/operations/accountaddpaymentmethodrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.AccountAddPaymentMethodResponse](../../models/operations/accountaddpaymentmethodresponse.md), error**


## DeleteAddress

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


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
    res, err := s.Account.DeleteAddress(ctx, operations.AccountAddressDeleteRequest{
        XPublishableKey: "string",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountAddressDeleteRequest](../../models/operations/accountaddressdeleterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.AccountAddressDeleteResponse](../../models/operations/accountaddressdeleteresponse.md), error**


## DeletePaymentMethod

Delete an existing payment method. Deleting a payment method does not invalidate transactions or
orders that are associated with it.


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
    res, err := s.Account.DeletePaymentMethod(ctx, operations.AccountPaymentMethodDeleteRequest{
        XPublishableKey: "string",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AccountPaymentMethodDeleteRequest](../../models/operations/accountpaymentmethoddeleterequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.AccountPaymentMethodDeleteResponse](../../models/operations/accountpaymentmethoddeleteresponse.md), error**


## Detect

Determine whether or not an identifier is associated with an existing Bolt account.

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
    res, err := s.Account.Detect(ctx, operations.AccountExistsRequest{
        XPublishableKey: "string",
        Identifier: shared.Identifier{
            IdentifierType: shared.IdentifierIdentifierTypeEmail,
            IdentifierValue: "alice@example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AccountExistsRequest](../../models/operations/accountexistsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.AccountExistsResponse](../../models/operations/accountexistsresponse.md), error**


## GetDetails

Retrieve a shopper's account details, such as addresses and payment information

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
    res, err := s.Account.GetDetails(ctx, operations.AccountGetRequest{
        XPublishableKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.AccountGetRequest](../../models/operations/accountgetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.AccountGetResponse](../../models/operations/accountgetresponse.md), error**


## UpdateAddress

Edit an existing address on the shopper's account. This does not edit addresses
that are already associated with other resources, such as transactions or
shipments.


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
    res, err := s.Account.UpdateAddress(ctx, operations.AccountAddressEditRequest{
        XPublishableKey: "string",
        AddressListingInput: shared.AddressListingInput{
            Company: boltalphaspec.String("ACME Corporation"),
            CountryCode: shared.AddressListingCountryCodeUs,
            Email: boltalphaspec.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: boltalphaspec.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: boltalphaspec.String("+14155550199"),
            PostalCode: "94105",
            Region: boltalphaspec.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: boltalphaspec.String("c/o Shipping Department"),
        },
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AccountAddressEditRequest](../../models/operations/accountaddresseditrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.AccountAddressEditResponse](../../models/operations/accountaddresseditresponse.md), error**

