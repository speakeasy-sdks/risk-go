# ApplicationAPICreateIn


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Color`                                                                           | **string*                                                                         | :heavy_minus_sign:                                                                | The hex representation of the icon color of the application (defaults to #00a3de) | #00a3de                                                                           |
| `Icon`                                                                            | [*shared.Icon](../../../pkg/models/shared/icon.md)                                | :heavy_minus_sign:                                                                | The icon type of the application (defaults to CUBES)                              | CUBES                                                                             |
| `Name`                                                                            | *string*                                                                          | :heavy_check_mark:                                                                | The name of the application                                                       | Cyber Risk Management Application                                                 |
| `Type`                                                                            | [*shared.Type](../../../pkg/models/shared/type.md)                                | :heavy_minus_sign:                                                                | The type of Risk Cloud application (defaults to NONE)                             | CONTROLS_COMPLIANCE                                                               |