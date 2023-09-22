# Configuration

## Overview

Merchant configuration endpoints allow you to retrieve and configure merchant-level
configuration, such as callback URLs, identifiers, secrets, etc...


### Available Operations

* [MerchantCallbacksGet](#merchantcallbacksget) - Retrieve callback URLs for the merchant
* [MerchantCallbacksUpdate](#merchantcallbacksupdate) - Update callback URLs for the merchant
* [MerchantIdentifiersGet](#merchantidentifiersget) - Retrieve identifiers for the merchant

## MerchantCallbacksGet

Return callback URLs configured on the merchant such as OAuth URLs.


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
    res, err := s.Configuration.MerchantCallbacksGet(ctx, operations.MerchantCallbacksGetRequest{
        XPublishableKey: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.MerchantCallbacksGetRequest](../../models/operations/merchantcallbacksgetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.MerchantCallbacksGetResponse](../../models/operations/merchantcallbacksgetresponse.md), error**


## MerchantCallbacksUpdate

Update and configure callback URLs on the merchant such as OAuth URLs.


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
    res, err := s.Configuration.MerchantCallbacksUpdate(ctx, operations.MerchantCallbacksUpdateRequest{
        RequestBody: operations.MerchantCallbacksUpdateRequestBody{
            AccountPage: boltalphaspec.String("https://www.example.com/account"),
            BaseDomain: boltalphaspec.String("https://www.example.com/"),
            ConfirmationRedirect: boltalphaspec.String("https://www.example.com/bolt/redirect"),
            CreateOrder: boltalphaspec.String("https://www.example.com/bolt/order"),
            Debug: boltalphaspec.String("https://www.example.com/bolt/debug"),
            GetAccount: boltalphaspec.String("https://www.example.com/bolt/account"),
            MobileAppDomain: boltalphaspec.String("https://m.example.com/"),
            OauthLogout: boltalphaspec.String("https://www.example.com/bolt/logout"),
            OauthRedirect: boltalphaspec.String("https://www.example.com/bolt/oauth"),
            PrivacyPolicy: boltalphaspec.String("https://www.example.com/privacy-policy"),
            ProductInfo: boltalphaspec.String("https://www.example.com/bolt/product"),
            RemoteAPI: boltalphaspec.String("https://www.example.com/bolt/remote-api"),
            Shipping: boltalphaspec.String("https://www.example.com/bolt/shipping"),
            SupportPage: boltalphaspec.String("https://www.example.com/help"),
            Tax: boltalphaspec.String("https://www.example.com/bolt/tax"),
            TermsOfService: boltalphaspec.String("https://www.example.com/terms-of-service"),
            UniversalMerchantAPI: boltalphaspec.String("https://www.example.com/bolt/merchant-api"),
            UpdateCart: boltalphaspec.String("https://www.example.com/bolt/cart"),
            ValidateAdditionalAccountData: boltalphaspec.String("https://www.example.com/bolt/validate-account"),
        },
        XPublishableKey: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Onemerchant1callbacksPatchRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.MerchantCallbacksUpdateRequest](../../models/operations/merchantcallbacksupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.MerchantCallbacksUpdateResponse](../../models/operations/merchantcallbacksupdateresponse.md), error**


## MerchantIdentifiersGet

Return several identifiers for the merchant, such as public IDs, publishable keys, signing secrets, etc...

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltalphaspec "github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/shared"
)

func main() {
    s := boltalphaspec.New(
        boltalphaspec.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Configuration.MerchantIdentifiersGet(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.MerchantIdentifiersGet200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.MerchantIdentifiersGetResponse](../../models/operations/merchantidentifiersgetresponse.md), error**

