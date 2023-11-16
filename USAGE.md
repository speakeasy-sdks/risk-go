<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
	"log"
)

func main() {
	s := riskgo.New()

	operationSecurity := operations.GetAPITokenSecurity{
		Password: "",
		Username: "",
	}

	ctx := context.Background()
	res, err := s.Authentication.GetAPIToken(ctx, operations.GetAPITokenRequest{}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.LegacyAPITokenOut != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->