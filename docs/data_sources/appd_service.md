## AppD Service Data source

Small helper to get tier and application id from names.

## Example Usage

```hcl
data "appdynamics_appd_service" "test" {
  application_name = "APPLICATION_NAME"
  tier_name = "TIER_NAME"
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_name`|string|int|Existing application name|`APPLICATION_NAME`|
|`tier_name`|yes|string|Existing tier name|`TIER_NAME`|


## Output References
|`application_id`|yes|int|The application id|`123`|
|`tier_id`|yes|int|The tier id|`324`|


