# Authentication
(*Authentication*)

## Overview

Getting Started: How to create an [API Access Token](https://www.logicgate.com/developer/risk-cloud-api-authentication/) to begin integrating with the Risk Cloud API

### Available Operations

* [GetAPIToken](#getapitoken) - Create an API Access Token

## GetAPIToken

**Permissions:** Authenticated User

Generates a new, expiring access token from the provided Client and Secret keys.

### Example Usage

```go
package main

import(
	riskgo "github.com/speakeasy-sdks/risk-go"
	"github.com/speakeasy-sdks/risk-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := riskgo.New()


    operationSecurity := operations.GetAPITokenSecurity{
            Password: "<YOUR_PASSWORD_HERE>",
            Username: "<YOUR_USERNAME_HERE>",
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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetAPITokenRequest](../../pkg/models/operations/getapitokenrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetAPITokenSecurity](../../pkg/models/operations/getapitokensecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.GetAPITokenResponse](../../pkg/models/operations/getapitokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
