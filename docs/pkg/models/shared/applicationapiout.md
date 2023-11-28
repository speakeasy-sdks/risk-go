# ApplicationAPIOut


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Color`                                                                                                 | **string*                                                                                               | :heavy_minus_sign:                                                                                      | The hex representation of the icon color of the application                                             | #00a3de                                                                                                 |
| `Icon`                                                                                                  | [*shared.ApplicationAPIOutIcon](../../../pkg/models/shared/applicationapiouticon.md)                    | :heavy_minus_sign:                                                                                      | The icon type of the application                                                                        | CUBES                                                                                                   |
| `ID`                                                                                                    | **string*                                                                                               | :heavy_minus_sign:                                                                                      | The unique ID of this Risk Cloud resource                                                               | a1b2c3d4                                                                                                |
| `Live`                                                                                                  | **bool*                                                                                                 | :heavy_minus_sign:                                                                                      | Whether the application is live                                                                         | false                                                                                                   |
| `Name`                                                                                                  | **string*                                                                                               | :heavy_minus_sign:                                                                                      | The name of the application                                                                             | Cyber Risk Management Application                                                                       |
| `Object`                                                                                                | **string*                                                                                               | :heavy_minus_sign:                                                                                      | Identifies the type of object this data represents                                                      | application                                                                                             |
| `RestrictBuildAccess`                                                                                   | **bool*                                                                                                 | :heavy_minus_sign:                                                                                      | Whether users with the Build entitlement must be explicitly granted permission to edit this application | false                                                                                                   |
| `Type`                                                                                                  | [*shared.ApplicationAPIOutType](../../../pkg/models/shared/applicationapiouttype.md)                    | :heavy_minus_sign:                                                                                      | The type of Risk Cloud application                                                                      | CONTROLS_COMPLIANCE                                                                                     |