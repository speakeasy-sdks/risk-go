# WorkflowMapAPICreateIn

Workflow Map (Create)


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `From`                                                                     | *string*                                                                   | :heavy_check_mark:                                                         | The unique ID of the source workflow of the workflow map relationship      | a1b2c3d4                                                                   |
| `Relationship`                                                             | [shared.Relationship](../../../pkg/models/shared/relationship.md)          | :heavy_check_mark:                                                         | The type of the relationship between workflows                             | ONE_TO_MANY                                                                |
| `To`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | The unique ID of the destination workflow of the workflow map relationship | a1b2c3d4                                                                   |