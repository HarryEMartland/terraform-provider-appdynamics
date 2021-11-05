## Import Export Dashboard Resource

Creates a dashboard from json template. All modification result in dashboard recreation.

## Example Usage

```hcl
resource "appdynamics_import_export_dashboard" "jenkins_cicd" {
  json = file("./dashboards/jenkins-cicd.json")
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`json`|yes|number|Json type appd dashboard template|`My dashboard`|

For more details about template schema
check [docs](https://docs.appdynamics.com/21.4/en/appdynamics-essentials/dashboards-and-reports/custom-dashboards/import-and-export-custom-dashboards-and-templates-using-the-ui#ImportandExportCustomDashboardsandTemplatesUsingtheUI-ExportCustomDashboardTemp
