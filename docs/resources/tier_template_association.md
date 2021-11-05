## Import Export Dashboard Resource

Bound already existed tier template with the tier.

## Example Usage

```hcl
resource "appdynamics_tier_template_association" "sample_association" {
  application_id = 123
  tier_id = 123
  template_ids= []
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|int|The application id|`123`|
|`tier_id`|yes|int|The tier id|`324`|
|`template_ids`|yes|int[]|Ids of templates to bind with tier|`[1,2,3]`|
