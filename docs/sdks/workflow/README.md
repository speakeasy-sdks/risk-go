# Workflow
(*Workflow*)

## Overview

A [Workflow](https://help.logicgate.com/hc/en-us/articles/4402683108756-Create-a-new-Workflow) is a combination of Steps, Paths, Fields, and routing logic that combine to form a system in an Application

### Available Operations

* [Create](#create) - Create a workflow
* [Delete](#delete) - Delete a workflow
* [Read](#read) - Retrieve a workflow
* [ReadAll](#readall) - Retrieve workflows
* [Update](#update) - Update a workflow

## Create

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create a workflow from a JSON request body. The workflow will contain a Default Origin step and a Default End step.

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
    res, err := s.Workflow.Create(ctx, operations.CreateWorkflowRequest{
        WorkflowAPICreateIn: shared.WorkflowAPICreateIn{
            ApplicationID: "a1b2c3d4",
            Name: "Risk Assessments",
            RecordPrefix: "Assessment",
            Xpos: riskgo.Int(20),
            Ypos: riskgo.Int(20),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateWorkflowRequest](../../pkg/models/operations/createworkflowrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateWorkflowResponse](../../pkg/models/operations/createworkflowresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete a workflow specified by the ID in the URL path.

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
    res, err := s.Workflow.Delete(ctx, operations.DeleteWorkflowRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteWorkflowRequest](../../pkg/models/operations/deleteworkflowrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteWorkflowResponse](../../pkg/models/operations/deleteworkflowresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Read

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a workflow specified by the ID in the URL path.

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
    res, err := s.Workflow.Read(ctx, operations.ReadWorkflowRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ReadWorkflowRequest](../../pkg/models/operations/readworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ReadWorkflowResponse](../../pkg/models/operations/readworkflowresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all workflows that the current user has [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

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
    res, err := s.Workflow.ReadAll(ctx, operations.ReadAllWorkflowsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutWorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ReadAllWorkflowsRequest](../../pkg/models/operations/readallworkflowsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ReadAllWorkflowsResponse](../../pkg/models/operations/readallworkflowsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update a workflow specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

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
    res, err := s.Workflow.Update(ctx, operations.UpdateWorkflowRequest{
        WorkflowAPIUpdateIn: shared.WorkflowAPIUpdateIn{
            Name: riskgo.String("Risk Assessments"),
            RecordPrefix: riskgo.String("Assessment"),
            Xpos: riskgo.Int(20),
            Ypos: riskgo.Int(20),
        },
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateWorkflowRequest](../../pkg/models/operations/updateworkflowrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateWorkflowResponse](../../pkg/models/operations/updateworkflowresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
