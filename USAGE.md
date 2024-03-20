<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"log"
)

func main() {
	s := riskgo.New(
		riskgo.WithSecurity(shared.Security{
			Basic: &shared.SchemeBasic{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Authentication.GetAPIToken(ctx, operations.GetAPITokenRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.LegacyAPITokenOut != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->