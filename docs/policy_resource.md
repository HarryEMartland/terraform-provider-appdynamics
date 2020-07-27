## Policy Resource

#### Properties

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the action to|`32423`|
|`name`|yes|string|The name of the policy|`"My Policy"`|
|`action_name`|yes|string|The name of the policy to execute (may be the phone number or emails)|`"bob@example.com"`|
|`acion_type`|yes|string|The type of the action to execute|`"EMAIL"`|
|`health_rule_event_types`|yes|string[]|A list of event types to execute the action on|`["HEALTH_RULE_OPEN_CRITICAL"]`|
|`health_rule_scope_type`|yes|string|The type of health rules to execute the action on|`"SPECIFIC_HEALTH_RULES"`|
|`health_rules`|no (default `[]`)|string[]|The health rules to trigger the action on|`["My Health Rule"]`|
|`other_events`|no (default `[]`)|string[]|Other events to trigger the action on|`["SPECIFIC_HEALTH_RULES"]`|
|`execute_actions_in_batch`|no (default `true`)|boolean|Other events to trigger the action on|`true`|
|`enabled`|no (default `true`)|boolean|Triggers the action when enabled|`true`|

###### Health Rule Scope
- ALL_HEALTH_RULES
- SPECIFIC_HEALTH_RULES

###### Health Rule Event Types
- HEALTH_RULE_CONTINUES_CRITICAL
- HEALTH_RULE_OPEN_CRITICAL
- HEALTH_RULE_OPEN_WARNING
- HEALTH_RULE_UPGRADED
- HEALTH_RULE_DOWNGRADED
- HEALTH_RULE_CONTINUES_WARNING
- HEALTH_RULE_CLOSE_WARNING
- HEALTH_RULE_CLOSE_CRITICAL
- HEALTH_RULE_CANCELED_WARNING
- HEALTH_RULE_CANCELED_CRITICAL

###### Other Events
- CLR_CRASH
- APPLICATION_CRASH
- DEADLOCK
- RESOURCE_POOL_LIMIT
- APPLICATION_DEPLOYMENT
- APP_SERVER_RESTART
- APPLICATION_CONFIG_CHANGE
- AGENT_CONFIGURATION_ERROR
- APPLICATION_DISCOVERED
- TIER_DISCOVERED
- NODE_DISCOVERED
- MACHINE_DISCOVERED
- BT_DISCOVERED
- SERVICE_ENDPOINT_DISCOVERED
- BACKEND_DISCOVERED
- EUM_CLOUD_SYNTHETIC_HEALTHY_EVENT
- EUM_CLOUD_SYNTHETIC_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_CONFIRMED_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_ONGOING_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_ERROR_EVENT
- EUM_CLOUD_SYNTHETIC_CONFIRMED_ERROR_EVENT
- EUM_CLOUD_SYNTHETIC_ONGOING_ERROR_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_HEALTHY_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_CONFIRMED_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_ONGOING_WARNING_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_CRITICAL_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_CONFIRMED_CRITICAL_EVENT
- EUM_CLOUD_SYNTHETIC_PERF_ONGOING_CRITICAL_EVENT
- MOBILE_NEW_CRASH_EVENT, SLOW, VERY_SLOW, STALL
- ERROR

#### Examples

###### Specific Health Rule

```terraform
resource "appd_policy" "my_policy" {
  name = "My Policy"
  application_id = var.application_id
  action_name = "my action"
  action_type = appd_action.my-first-action.action_type
  health_rule_event_types = [
    "HEALTH_RULE_OPEN_WARNING",
    "HEALTH_RULE_OPEN_CRITICAL"]
  health_rule_scope_type = "SPECIFIC_HEALTH_RULES"
  health_rules = ["my health rule"]
}
```

###### All Health Rules

```hcl
resource "appd_policy" "all_health_rules_email_on_call" {
  name = "All Health Rules Email On call"
  application_id = var.application_id
  action_name = join(", ",appd_action.on-call-email-action.emails)
  action_type = appd_action.my-first-action.action_type
  health_rule_event_types = ["HEALTH_RULE_OPEN_CRITICAL"]
  health_rule_scope_type = "ALL_HEALTH_RULES"
}
```