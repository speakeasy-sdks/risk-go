# Application
(*Application*)

## Overview

An [Application](https://help.logicgate.com/hc/en-us/articles/4402674055572-Create-a-new-Application) is a collection of Workflows, Steps, and logic that collectively solve a business use case

### Available Operations

* [Create](#create) - Create an application
* [Delete](#delete) - Delete an application
* [Read](#read) - Retrieve an application
* [ReadAll](#readall) - Retrieve applications
* [Update](#update) - Update an application

## Create

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create an application from a JSON request body.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"context"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Application.Create(ctx, operations.CreateApplicationRequest{
        ApplicationAPICreateIn: shared.ApplicationAPICreateIn{
            Color: riskgo.String("#00a3de"),
            Icon: shared.IconCubes.ToPointer(),
            Name: "Cyber Risk Management Application",
            Type: shared.TypeControlsCompliance.ToPointer(),
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

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateApplicationRequest](../../pkg/models/operations/createapplicationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateApplicationResponse](../../pkg/models/operations/createapplicationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete an application specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"context"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Application.Delete(ctx, operations.DeleteApplicationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AppActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteApplicationRequest](../../pkg/models/operations/deleteapplicationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteApplicationResponse](../../pkg/models/operations/deleteapplicationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Read

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve an application specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"context"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Application.Read(ctx, operations.ReadApplicationRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ReadApplicationRequest](../../pkg/models/operations/readapplicationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ReadApplicationResponse](../../pkg/models/operations/readapplicationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all applications that the current user has [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"context"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Application.ReadAll(ctx, operations.ReadAllApplicationsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PageModelOutApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ReadAllApplicationsRequest](../../pkg/models/operations/readallapplicationsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ReadAllApplicationsResponse](../../pkg/models/operations/readallapplicationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update an application specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"context"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Application.Update(ctx, operations.Update1Request{
        ApplicationAPIUpdateIn: shared.ApplicationAPIUpdateIn{
            Color: riskgo.String("#00a3de"),
            Icon: shared.ApplicationAPIUpdateInIconCubes.ToPointer(),
            Live: riskgo.Bool(false),
            Name: riskgo.String("Cyber Risk Management Application"),
            RestrictBuildAccess: riskgo.Bool(false),
            Type: shared.ApplicationAPIUpdateInTypeControlsCompliance.ToPointer(),
        },
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.Update1Request](../../pkg/models/operations/update1request.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.Update1Response](../../pkg/models/operations/update1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
