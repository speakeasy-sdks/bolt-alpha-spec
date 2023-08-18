# Testing

## Overview

Endpoints that allow you to generate and retrieve test data to verify certain
flows in non-production environments.


### Available Operations

* [TestingAccountCreate](#testingaccountcreate) - Create a test account
* [TestingCreditCardGet](#testingcreditcardget) - Retrieve a test credit card, including its token
* [TestingShipmentTrackingCreate](#testingshipmenttrackingcreate) - Simulate a shipment tracking update

## TestingAccountCreate

Create a Bolt shopper account for testing purposes.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/types"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.TestingAccountCreateSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Testing.TestingAccountCreate(ctx, operations.TestingAccountCreateRequestBodyInput{
        DeactivateAt: types.MustTimeFromString("2017-07-21T17:32:28Z"),
        EmailState: operations.TestingAccountCreateRequestBodyEmailStateUnverified,
        HasAddress: boltalpha.Bool(true),
        IsMigrated: boltalpha.Bool(true),
        PhoneState: operations.TestingAccountCreateRequestBodyPhoneStateVerified,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Onetesting1accountsPostRequestBodyContentApplication1jsonSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.TestingAccountCreateRequestBodyInput](../../models/operations/testingaccountcreaterequestbodyinput.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.TestingAccountCreateSecurity](../../models/operations/testingaccountcreatesecurity.md)                 | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.TestingAccountCreateResponse](../../models/operations/testingaccountcreateresponse.md), error**


## TestingCreditCardGet

Retrieve test credit card information. This includes its token, which is
generated against the `4111 1111 1111 1004` test card.


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
    operationSecurity := operations.TestingCreditCardGetSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Testing.TestingCreditCardGet(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TestingCreditCardGet200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.TestingCreditCardGetSecurity](../../models/operations/testingcreditcardgetsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.TestingCreditCardGetResponse](../../models/operations/testingcreditcardgetresponse.md), error**


## TestingShipmentTrackingCreate

Simulate a shipment tracking update, such as those that Bolt would ingest from
third-party shipping providers that would allow updating tracking and delivery
information to shipments associated with orders.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-alpha-spec"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-alpha-spec/pkg/types"
)

func main() {
    s := boltalpha.New()
    operationSecurity := operations.TestingShipmentTrackingCreateSecurity{
            APIKey: "",
        }

    ctx := context.Background()
    res, err := s.Testing.TestingShipmentTrackingCreate(ctx, operations.TestingShipmentTrackingCreateRequestBody{
        DeliveryDate: types.MustTimeFromString("2014-08-23:T06:00:00Z"),
        Status: operations.TestingShipmentTrackingCreateRequestBodyStatusInTransit,
        TrackingDetails: []TestingShipmentTrackingCreateRequestBodyTrackingDetails{
            operations.TestingShipmentTrackingCreateRequestBodyTrackingDetails{
                CountryCode: boltalpha.String("US"),
                EventDate: boltalpha.String("2014-08-21:T14:24:00Z"),
                Locality: boltalpha.String("San Francisco"),
                Message: boltalpha.String("Billing information received"),
                PostalCode: boltalpha.String("94105"),
                Region: boltalpha.String("CA"),
                Status: operations.TestingShipmentTrackingCreateRequestBodyTrackingDetailsStatusPreTransit.ToPointer(),
            },
            operations.TestingShipmentTrackingCreateRequestBodyTrackingDetails{
                CountryCode: boltalpha.String("US"),
                EventDate: boltalpha.String("2014-08-21:T14:24:00Z"),
                Locality: boltalpha.String("San Francisco"),
                Message: boltalpha.String("Billing information received"),
                PostalCode: boltalpha.String("94105"),
                Region: boltalpha.String("CA"),
                Status: operations.TestingShipmentTrackingCreateRequestBodyTrackingDetailsStatusPreTransit.ToPointer(),
            },
        },
        TrackingNumber: "MockBolt-143292",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.TestingShipmentTrackingCreateRequestBody](../../models/operations/testingshipmenttrackingcreaterequestbody.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.TestingShipmentTrackingCreateSecurity](../../models/operations/testingshipmenttrackingcreatesecurity.md)       | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.TestingShipmentTrackingCreateResponse](../../models/operations/testingshipmenttrackingcreateresponse.md), error**

