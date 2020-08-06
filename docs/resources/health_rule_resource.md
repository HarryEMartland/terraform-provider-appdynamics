## Health Rule Resource

#### Properties

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the action to|`32423`|
|`metric_aggregation_function`|yes|string|How to aggregate multiple sources of the metric|`"VALUE"`|
|`eval_detail_type`|yes|string|What to evaluate the metric against|`"SINGLE_METRIC"`|
|`affected_entity_type`|yes|string|The entity type for the health rule|`"OVERALL_APPLICATION_PERFORMANCE"`|
|`business_transaction_scope`|yes|string|Which business transaction are applicable for the health rule|`"ALL_BUSINESS_TRANSACTIONS"`|
|`baseline_condition`|no|string|How to compare to the baseline|`"WITHIN_BASELINE"`|
|`metric_eval_detail_type`|yes|string|The type of comparison|`"BASELINE_TYPE"`|
|`baseline_name`|no|string|Which baseline to use|`"All data - Last 15 days"`|
|`baseline_unit`|no|string|What unit to compare the baseline with|`"PERCENTAGE"`|
|`metric_path`|yes|string|Which metric to use|`"95th Percentile Response Time (ms)"`|
|`compare_condition`|no|string|How to compare the values to the metric|`"GREATER_THAN_SPECIFIC_VALUE"`|
|`business_transactions`|no|number|A list of transactions to trigger the health rule for|`["/endpoint"]`|
|`warn_compare_value`|yes|number|The value at which the health rule should trigger a warning|`1`|
|`critical_compare_value`|yes|number|The value at which the health rule should trigger an error|`2`|

###### affected_entity_type
- OVERALL_APPLICATION_PERFORMANCE
- BUSINESS_TRANSACTION_PERFORMANCE
- TIER_NODE_TRANSACTION_PERFORMANCE
- TIER_NODE_HARDWARE
- SERVERS_IN_APPLICATION
- BACKENDS
- ERRORS
- SERVICE_ENDPOINTS
- INFORMATION_POINTS
- CUSTOM
- DATABASES
- SERVERS

###### business_transaction_scope
- ALL_BUSINESS_TRANSACTIONS
- SPECIFIC_BUSINESS_TRANSACTIONS
- BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS
- BUSINESS_TRANSACTIONS_MATCHING_PATTERN

###### baseline_unit
- STANDARD_DEVIATIONS
- PERCENTAGE

###### baseline_condition
- WITHIN_BASELINE
- NOT_WITHIN_BASELINE
- GREATER_THAN_BASELINE
- LESS_THAN_BASELINE

###### metric_eval_detail_type
- BASELINE_TYPE
- SPECIFIC_TYPE

###### compare_condition
- GREATER_THAN_SPECIFIC_VALUE
- LESS_THAN_SPECIFIC_VALUE

###### metric_aggregation_function
- MINIMUM
- MAXIMUM
- VALUE
- SUM
- COUNT
- CURRENT
- GROUP_COUNT

###### match_type
- AVERAGE
- ANY_NODE
- PERCENTAGE_NODES
- NUMBER_OF_NODES

#### Examples

###### All BTs Baseline 
```hcl
resource "appd_health_rule" "my_baseline_rule" {
  name = "My Baseline Health Rule"
  application_id = var.application_id
  metric_aggregation_function = "VALUE"
  eval_detail_type = "SINGLE_METRIC"
  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
  baseline_condition = "WITHIN_BASELINE"
  metric_eval_detail_type = "BASELINE_TYPE"
  baseline_name = "All data - Last 15 days"
  baseline_unit = "STANDARD_DEVIATIONS"
  metric_path = "95th Percentile Response Time (ms)"
  warn_compare_value = 1
  critical_compare_value = 2
}
```

###### All BTs Value
```hcl
resource "appd_health_rule" "my_single_metric_rule" {
  name = "My Single Metring Health Rule"
  application_id = var.application_id
  metric_aggregation_function = "VALUE"
  eval_detail_type = "SINGLE_METRIC"
  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
  metric_eval_detail_type = "SPECIFIC_TYPE"
  metric_path = "95th Percentile Response Time (ms)"
  compare_condition="GREATER_THAN_SPECIFIC_VALUE"
  warn_compare_value = 100
  critical_compare_value = 200
}
```

###### Specific BTs Value
```hcl
resource "appd_health_rule" "specific_bts_rule" {
  name = "My Specific BTs Rule"
  application_id = var.application_id
  metric_aggregation_function = "VALUE"
  eval_detail_type = "SINGLE_METRIC"
  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope = "SPECIFIC_BUSINESS_TRANSACTIONS"
  business_transactions = [
    "/route/one",
    "/route/two"
  ]
  metric_eval_detail_type = "SPECIFIC_TYPE"
  metric_path = "95th Percentile Response Time (ms)"
  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
  warn_compare_value = 100
  critical_compare_value = 200
}
```

###### Specific Tiers Value
```hcl
resource "appd_health_rule" "specific_tiers_rule" {
  name = "My Specific Tiers Rule"
  application_id = var.application_id
  metric_aggregation_function = "VALUE"
  eval_detail_type = "SINGLE_METRIC"
  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope = "BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS"
  specific_tiers = [
    "tier1",
    "tier2"
  ]
  metric_eval_detail_type = "SPECIFIC_TYPE"
  metric_path = "95th Percentile Response Time (ms)"
  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
  warn_compare_value = 100
  critical_compare_value = 200
}
```