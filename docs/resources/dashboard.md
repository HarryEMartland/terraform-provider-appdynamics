## Dashboard Resource

Creates a dashboard using widgets from [dashboard_widgets](../data_sources/dashboard_widget.md) data source.
This resource is still in experimental status, have some limitations and can raise unexpected errors.

## Example Usage

```hcl
resource "appdynamics_dashboard" "test_dashboard" {
  name = "test dashboard name"
  template_entity_type = "APPLICATION_COMPONENT_NODE"
  minutes_before_anchor_time = -1
  refresh_interval = 120000
  background_color = 15856629
  height = 768
  width = 1024
  canvas_type = "CANVAS_TYPE_GRID"
  widgets = []
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`name`|yes|number|The name of the dashboard|`My dashboard`|
|`template_entity_type`|no|string|Dashboard template entity type|`"APPLICATION_COMPONENT_NODE"`|
|`minutes_before_anchor_time`|no|int|n/a|`"bob@example.com"`|
|`refresh_interval`|no|int|n/a|`120000`|
|`background_color`|no|int|n/a|`15856629`|
|`height`|no|int|Height of the dashboard|`768`|
|`width`|no|int|Width of the dashboard|`1024`|
|`canvas_type`|no|string|n/a|`"CANVAS_TYPE_GRID"`|
|`widgets`|no (default `[]`)|string[]|Must come from data_source_dashboard_widget type|`true`|
