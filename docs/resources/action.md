# Action Resource

Creates an action which can be triggered by a [policy](policy_resource.md).

[AppDynamics Action Information](https://docs.appdynamics.com/display/PRO45/Actions)  
[AppDynamics Action API](https://docs.appdynamics.com/display/PRO45/Actions+API)

## Example Usage

#### Email
```terraform
resource "appd_action" "my_first_email_action" {
  application_id = var.application_id
  action_type = "EMAIL"
  emails = [
    "bob@example.com",
    "sandra@example.com"
  ]
}
```

#### SMS
```hcl
resource "appd_action" "my-first-sms-action" {
  application_id = var.application_id
  action_type = "SMS"
  phone_number = "07421365896"
}
```

#### HTTP Request
```hcl
resource "appd_action" "my-first-http-action" {
  application_id = var.application_id
  name = "My First HTTP Action"
  action_type = "HTTP_REQUEST"
  http_request_template_name = "Slack Alert - Any Channel"
  custom_template_variables = {
    channel: "#alert-channel"
  }
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the action to|`32423`|
|`name`|no|string|The name for the action|`"My Action"`|
|`action_type`|yes|string|The type of the action|`"EMAIL"`|
|`emails`|no|string[]|The email addresses to be notified when he action is performed|`["bob@example.com"]`|
|`http_request_template_name`|no|string|The name of the request template to use|`"Slack Template"`
|`custom_template_variables`|no|map{string:string}|The names and values of variables to pass into the template|`{channel: "#alert-channel"}`

###### action_type
- SMS
- EMAIL
- CUSTOM_EMAIL
- THREAD_DUMP
- HTTP_REQUEST
- CUSTOM
- RUN_SCRIPT_ON_NODES
- DIAGNOSE_BUSINESS_TRANSACTIONS
- CREATE_UPDATE_JIRA