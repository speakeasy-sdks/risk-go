# Step
(*Step*)

## Overview

A [Step](https://help.logicgate.com/hc/en-us/articles/4402674059668-Create-a-Step) lives in a Workflow and is configured with a set of Sections, Subsections and Fields to create a form

### Available Operations

* [Create](#create) - Create a step
* [Delete](#delete) - Delete a step
* [Read](#read) - Retrieve a step
* [ReadAll](#readall) - Retrieve steps
* [Update](#update) - Update a step

## Create

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create a step from a JSON request body.

### Example Usage

```go
package main

import(
	"context"
	"log"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Step.Create(ctx, operations.CreateStepRequest{
        StepAPICreateIn: shared.StepAPICreateIn{
            AssignableUserType: shared.AssignableUserTypeAppUsers.ToPointer(),
            EnableComments: riskgo.Bool(false),
            ExternalUserMfaRequired: riskgo.Bool(false),
            Name: "Identify Risk",
            WorkflowID: "a1b2c3d4",
            Xpos: riskgo.Int(20),
            Ypos: riskgo.Int(20),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateStepRequest](../../pkg/models/operations/createsteprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateStepResponse](../../pkg/models/operations/createstepresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Delete

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete a step specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"context"
	"log"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Step.Delete(ctx, operations.DeleteStepRequest{
        ID: "<ID>",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteStepRequest](../../pkg/models/operations/deletesteprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteStepResponse](../../pkg/models/operations/deletestepresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Read

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a step specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"context"
	"log"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Step.Read(ctx, operations.ReadStepRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ReadStepRequest](../../pkg/models/operations/readsteprequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ReadStepResponse](../../pkg/models/operations/readstepresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all steps that the current user has [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

### Example Usage

```go
package main

import(
	"context"
	"log"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Step.ReadAll(ctx, operations.ReadAllStepsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutStepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ReadAllStepsRequest](../../pkg/models/operations/readallstepsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ReadAllStepsResponse](../../pkg/models/operations/readallstepsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Update

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update a step specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

### Example Usage

```go
package main

import(
	"context"
	"log"
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
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
    res, err := s.Step.Update(ctx, operations.UpdateRequest{
        StepAPIUpdateIn: shared.StepAPIUpdateIn{
            AssignableUserType: shared.StepAPIUpdateInAssignableUserTypeAppUsers.ToPointer(),
            EnableComments: riskgo.Bool(false),
            ExternalUserMfaRequired: riskgo.Bool(false),
            Name: riskgo.String("Identify Risk"),
            Type: shared.StepAPIUpdateInTypeOrigin.ToPointer(),
            Xpos: riskgo.Int(20),
            Ypos: riskgo.Int(20),
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.UpdateRequest](../../pkg/models/operations/updaterequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.UpdateResponse](../../pkg/models/operations/updateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
