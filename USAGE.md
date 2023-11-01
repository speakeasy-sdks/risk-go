<!-- Start SDK Example Usage -->


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
				Password: "",
				Username: "",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Application.Create(ctx, operations.CreateApplicationRequest{
		ApplicationAPICreateIn: shared.ApplicationAPICreateIn{
			Color: riskgo.String("#00a3de"),
			Icon:  shared.ApplicationAPICreateInIconCubes.ToPointer(),
			Name:  "Cyber Risk Management Application",
			Type:  shared.ApplicationAPICreateInTypeControlsCompliance.ToPointer(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.ApplicationAPIOut != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->