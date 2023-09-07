# Account

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [AccountAddPaymentMethod](#accountaddpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [AccountAddressCreate](#accountaddresscreate) - Add an address
* [AccountAddressDelete](#accountaddressdelete) - Delete an existing address
* [AccountAddressEdit](#accountaddressedit) - Edit an existing address
* [AccountExists](#accountexists) - Determine the existence of a Bolt account
* [AccountGet](#accountget) - Retrieve account details

## AccountAddPaymentMethod

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
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.AccountAddPaymentMethodSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
        RequestBody: operations.AccountAddPaymentMethodRequestBodyInput{
            BillingAddress: shared.OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesShipmentsItemsPropertiesAddressInput{},
            Bin: "411111",
            Expiration: "2025-03",
            Last4: "1004",
            Network: operations.AccountAddPaymentMethodRequestBodyNetworkVisa,
            Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
            Type: operations.AccountAddPaymentMethodRequestBodyTypeCredit,
        },
        XPublishableKey: "provident",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Oneaccount1paymentMethodsPostRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountAddPaymentMethodRequest](../../models/operations/accountaddpaymentmethodrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.AccountAddPaymentMethodSecurity](../../models/operations/accountaddpaymentmethodsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.AccountAddPaymentMethodResponse](../../models/operations/accountaddpaymentmethodresponse.md), error**


## AccountAddressCreate

Add an address to the shopper's account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.AccountAddressCreateSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddressCreate(ctx, operations.AccountAddressCreateRequest{
        RequestBody: operations.AccountAddressCreateRequestBodyInput{
            Company: boltalpha.String("ACME Corporation"),
            CountryCode: "US",
            Email: boltalpha.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: boltalpha.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: boltalpha.String("+14155550199"),
            PostalCode: "94105",
            Region: boltalpha.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: boltalpha.String("c/o Shipping Department"),
        },
        XPublishableKey: "distinctio",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountAddressCreateRequest](../../models/operations/accountaddresscreaterequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.AccountAddressCreateSecurity](../../models/operations/accountaddresscreatesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.AccountAddressCreateResponse](../../models/operations/accountaddresscreateresponse.md), error**


## AccountAddressDelete

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.AccountAddressDeleteSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddressDelete(ctx, operations.AccountAddressDeleteRequest{
        XPublishableKey: "quibusdam",
        ID: "D4g3h5tBuVYK9",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountAddressDeleteRequest](../../models/operations/accountaddressdeleterequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.AccountAddressDeleteSecurity](../../models/operations/accountaddressdeletesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.AccountAddressDeleteResponse](../../models/operations/accountaddressdeleteresponse.md), error**


## AccountAddressEdit

Edit an existing address on the shopper's account. This does not edit addresses
that are already associated with other resources, such as transactions or
shipments.


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
    operationSecurity := operations.AccountAddressEditSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddressEdit(ctx, operations.AccountAddressEditRequest{
        XPublishableKey: "unde",
        ID: "D4g3h5tBuVYK9",
        Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput: shared.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchemaInput{
            Company: boltalpha.String("ACME Corporation"),
            CountryCode: "US",
            Email: boltalpha.String("alice@example.com"),
            FirstName: "Alice",
            IsDefault: boltalpha.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: boltalpha.String("+14155550199"),
            PostalCode: "94105",
            Region: boltalpha.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: boltalpha.String("c/o Shipping Department"),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Oneaccount1addressesPostRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountAddressEditRequest](../../models/operations/accountaddresseditrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.AccountAddressEditSecurity](../../models/operations/accountaddresseditsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.AccountAddressEditResponse](../../models/operations/accountaddresseditresponse.md), error**


## AccountExists

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
)

func main() {
    s := boltalpha.New()

    ctx := context.Background()
    res, err := s.Account.AccountExists(ctx, operations.AccountExistsRequest{
        XPublishableKey: "nulla",
        Identifier: operations.AccountExistsIdentifier{
            IdentifierType: operations.AccountExistsIdentifierIdentifierTypeEmail,
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


## AccountGet

Retrieve a shopper's account details, such as addresses and payment information

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.AccountGetSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountGet(ctx, operations.AccountGetRequest{
        XPublishableKey: "corrupti",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountGet200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.AccountGetRequest](../../models/operations/accountgetrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.AccountGetSecurity](../../models/operations/accountgetsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.AccountGetResponse](../../models/operations/accountgetresponse.md), error**
