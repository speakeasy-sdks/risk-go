# github.com/speakeasy-sdks/risk-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/risk-go.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/risk-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
# SDK Installation

```bash
go get github.com/speakeasy-sdks/risk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
# Available Resources and Operations


## [Application](docs/sdks/application/README.md)

* [Create](docs/sdks/application/README.md#create) - Create an application
* [Delete](docs/sdks/application/README.md#delete) - Delete an application
* [Read](docs/sdks/application/README.md#read) - Retrieve an application
* [ReadAll](docs/sdks/application/README.md#readall) - Retrieve applications
* [Update](docs/sdks/application/README.md#update) - Update an application

## [Authentication](docs/sdks/authentication/README.md)

* [GetAPIToken](docs/sdks/authentication/README.md#getapitoken) - Create an API Access Token

## [Field](docs/sdks/field/README.md)

* [ReadAll](docs/sdks/field/README.md#readall) - Retrieve fields

## [Record](docs/sdks/record/README.md)

* [ReadAll](docs/sdks/record/README.md#readall) - Retrieve records

## [Step](docs/sdks/step/README.md)

* [Create](docs/sdks/step/README.md#create) - Create a step
* [Delete](docs/sdks/step/README.md#delete) - Delete a step
* [Read](docs/sdks/step/README.md#read) - Retrieve a step
* [ReadAll](docs/sdks/step/README.md#readall) - Retrieve steps
* [Update](docs/sdks/step/README.md#update) - Update a step

## [Workflow](docs/sdks/workflow/README.md)

* [Create](docs/sdks/workflow/README.md#create) - Create a workflow
* [Delete](docs/sdks/workflow/README.md#delete) - Delete a workflow
* [Read](docs/sdks/workflow/README.md#read) - Retrieve a workflow
* [ReadAll](docs/sdks/workflow/README.md#readall) - Retrieve workflows
* [Update](docs/sdks/workflow/README.md#update) - Update a workflow

## [WorkflowMap](docs/sdks/workflowmap/README.md)

* [Create](docs/sdks/workflowmap/README.md#create) - Create a workflow map
* [Delete](docs/sdks/workflowmap/README.md#delete) - Delete a workflow map
* [Read](docs/sdks/workflowmap/README.md#read) - Retrieve a workflow map
* [ReadAll](docs/sdks/workflowmap/README.md#readall) - Retrieve workflow maps
* [Update](docs/sdks/workflowmap/README.md#update) - Update a workflow map
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->

<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->

<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `http://localhost` | None |

For example:


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
		riskgo.WithServerIndex(0),
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


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


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
		riskgo.WithServerURL("http://localhost"),
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
<!-- End Server Selection -->

<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
