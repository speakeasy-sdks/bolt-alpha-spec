# OnepaymentsPostRequestBodyContentApplication1jsonSchemaPropertiesCartPropertiesAmounts


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Currency`                                                                    | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           | USD                                                                           |
| `Tax`                                                                         | **int64*                                                                      | :heavy_minus_sign:                                                            | The total tax amount, in cents for all of the items associated with the cart. | 900                                                                           |
| `Total`                                                                       | *int64*                                                                       | :heavy_check_mark:                                                            | The total amount, in cents, including its items and taxes (if applicable).    | 900                                                                           |