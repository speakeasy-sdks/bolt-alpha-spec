<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	boltalphaspec "github.com/speakeasy-sdks/bolt-alpha-spec"
	"log"
	"net/http"
)

func main() {
	s := boltalphaspec.New()

	ctx := context.Background()
	res, err := s.Pets.CreatePets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->