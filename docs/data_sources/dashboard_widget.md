## AppD Service Data source

Helper to add all needed fields to widget before using it with [dashboard](../resources/dashboard.md)`resource.

## Example Usage

```hcl
data "appdynamics_dashboard_widget" "test" {
  json = file("./dashboards/jenkins-cicd.json")
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`json`|yes|string|Sample widget json|`{}`|


## Output References
|`widget_json`|yes|string|Properly formatted widget|`{}`|

## Note
Unfortunately AppDynamics do not share API about widgets. For samples look [here](../../appdynamics/widgets).  


