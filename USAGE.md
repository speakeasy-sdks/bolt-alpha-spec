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
	s := boltalphaspec.New(
		boltalphaspec.WithSecurity(shared.Security{
			APIKey: "",
			Oauth:  "",
		}),
	)

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListingInput: shared.AddressListingInput{
			Company:        boltalphaspec.String("ACME Corporation"),
			CountryCode:    shared.AddressListingCountryCodeUs,
			Email:          boltalphaspec.String("alice@example.com"),
			FirstName:      "Alice",
			IsDefault:      boltalphaspec.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          boltalphaspec.String("+14155550199"),
			PostalCode:     "94105",
			Region:         boltalphaspec.String("CA"),
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
<!-- End SDK Example Usage -->