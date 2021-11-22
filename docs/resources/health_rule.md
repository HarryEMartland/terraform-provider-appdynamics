# Health Rule Resource

Creates a health rule which defines what normal looks like for an application.

[AppDynamics Health Rule Documentation](https://docs.appdynamics.com/display/PRO45/Health+Rules)  
[AppDynamics Health Rule API](https://docs.appdynamics.com/display/PRO45/Health+Rule+API)

## Example Usage

#### Specific tiers in BT

```hcl
resource "appdynamics_health_rule" "my_value_health_rule" {
  name                                = "My Value Health Rule"
  application_id                      = var.application_id
  affected_entity_type                = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope          = "BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS"
  specific_tiers                      = ["my-tier-name"]
  critical_condition_aggregation_type = "ANY"
  warning_condition_aggregation_type  = "ALL"

  critical_criteria {
    name                        = "Errors per minute"
    shortname                   = "EPM"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "Errors per minute"
    metric_eval_detail_type     = "SPECIFIC_TYPE"
    compare_condition           = "GREATER_THAN_SPECIFIC_VALUE"
    compare_value               = 7
  }
  critical_criteria {
    name                        = "Errors per minute"
    shortname                   = "EPM"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "Errors per minute"
    metric_eval_detail_type     = "BASELINE_TYPE"
    baseline_condition          = "WITHIN_BASELINE"
    baseline_name               = "All data - Last 15 days"
    baseline_unit               = "STANDARD_DEVIATIONS"
    metric_path                 = "95th Percentile Response Time (ms)"
    compare_value               = 1
  }
  warning_criteria {
    name                        = "Errors per minute"
    shortname                   = "EPM"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "Errors per minute"
    metric_eval_detail_type     = "SPECIFIC_TYPE"
    compare_condition           = "GREATER_THAN_SPECIFIC_VALUE"
    compare_value               = 2.5
  }
}
```

#### All BTs Baseline

```hcl
resource "appdynamics_health_rule" "my_baseline_health_rule" {
  name                                = "My baseline health rule"
  application_id                      = var.application_id
  affected_entity_type                = "BUSINESS_TRANSACTION_PERFORMANCE"
  business_transaction_scope          = "ALL_BUSINESS_TRANSACTIONS"
  critical_condition_aggregation_type = "ANY"
  warning_condition_aggregation_type  = "ALL"

  critical_criteria {
    name                        = "average cpu"
    shortname                   = "AC"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "Average CPU Used (ms)"
    metric_eval_detail_type     = "BASELINE_TYPE"
    baseline_condition          = "WITHIN_BASELINE"
    baseline_name               = "All data - Last 15 days"
    baseline_unit               = "PERCENTAGE"
    compare_value               = 33
  }
}
```

#### Tier / Node Health with specific tiers

```hcl
resource "appdynamics_health_rule" "my_specific_tiers_health_rule" {
  name                 = "My specific tiers health rule"
  application_id       = var.application_id
  affected_entity_type = "TIER_NODE_HARDWARE"
  tier_or_node         = "TIER_AFFECTED_ENTITIES"
  affected_tier_scope  = "SPECIFIC_TIERS"
  tiers                = ["my-tier-1", "my-tier-2"]

  critical_condition_aggregation_type = "ALL"
  warning_condition_aggregation_type  = "ALL"

  critical_criteria {
    name                        = "jvm gc allocated objects"
    shortname                   = "EPM"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "JVM|Garbage Collection|Allocated-Objects (MB)"
    metric_eval_detail_type     = "SPECIFIC_TYPE"
    compare_condition           = "GREATER_THAN_SPECIFIC_VALUE"
    compare_value               = 11
  }
}
```

#### Tier / Node Health with nodes matching

```hcl
resource "appdynamics_health_rule" "my_nodes_matching_rule" {
  name                         = "My nodes matching rule"
  application_id               = var.application_id
  use_data_from_last_n_minutes = 70
  wait_time_after_violation    = 13
  affected_entity_type         = "TIER_NODE_TRANSACTION_PERFORMANCE"
  tier_or_node                 = "NODE_AFFECTED_ENTITIES"
  type_of_node                 = "JAVA_NODES"
  affected_node_scope          = "NODES_MATCHING_PATTERN"
  nodes_match                  = "STARTS_WITH"
  nodes_match_value            = "my_prefix"
  nodes_match_negation         = true

  critical_condition_aggregation_type = "ALL"
  warning_condition_aggregation_type  = "ALL"

  critical_criteria {
    name                        = "Errors per minute"
    shortname                   = "EPM"
    evaluate_to_true_on_no_data = false
    eval_detail_type            = "SINGLE_METRIC"
    metric_aggregation_function = "VALUE"
    metric_path                 = "Errors per minute"
    metric_eval_detail_type     = "SPECIFIC_TYPE"
    compare_condition           = "GREATER_THAN_SPECIFIC_VALUE"
    compare_value               = 58
  }
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the health rule to|`32423`|
|`name`|yes|string|Name of health rule to add|`My health rule`|
|`schedule_name`|no|string|Schedule to be associated with the health rule|`Always`|
|`use_data_from_last_n_minutes`|no|number|The time interval during which the data collected is considered for health rule evaluation|`30`|
|`wait_time_after_violation`|no|number|Time to wait after violation|`10`|
|`affected_entity_type`|yes|string|The entity type for the health rule|`"OVERALL_APPLICATION_PERFORMANCE"`|
|`tier_or_node`|no|string|Affected entity type|`TIER_AFFECTED_ENTITIES`|
|`type_of_node`|no|string|Affected nodes type|`JAVA_NODES`|
|`affected_tier_scope`|no|string|Scope of tiers|`SPECIFIC_TIERS`|
|`tiers`|no|set|Tiers to match|`["my-tier1", "my-tier2"]`|
|`affected_node_scope`|no|string|Scope of nodes|`NODES_MATCHING_PATTERN`|
|`nodes_specific_tiers`|no|set|Tiers specific of node|`["my-tier1", "my-tier2"]`|
|`nodes`|no|set|Affected nodes|`["my-node1", "my-node2"]`|
|`nodes_match`|no|string|Type of match|`CONTAINS`|
|`nodes_match_value`|no|string|Value to match|`some value`|
|`nodes_match_negation`|no|bool|Negate match|`false`|
|`business_transaction_scope`|yes|string|Which business transaction are applicable for the health rule|`"ALL_BUSINESS_TRANSACTIONS"`|
|`business_transactions`|no|set|Set containing transactions|`["trans1", "trans2"]`|
|`business_transaction_specific_tiers`|no|set|Set containing tier names|`["my-tier1","my-tier2"]`|
|`business_transaction_match`|no|string|Type of match for transaction matching|`"ENDS_WITH"|
|`business_transaction_match_value`|no|string|Value to match|`some value`|
|`business_transaction_match_negation`|no|bool|If match should be negated|`true`|
|`warning_condition_aggregation_type`|no|string|How to aggregate warning conditions|`ANY`|
|`warning_criteria`|no|list|List of structures defining warning criteria|`{}`|
|`critical_condition_aggregation_type`|no|string|How to aggregate critical conditions|`ANY`|
|`critical_criteria`|no|list|List of structures defining critical criteria|`{}`|

Note:\
One of `warning_criteria` and `critical_criteria` has to be defined.

### Critical and warning criteria arguments

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`name`|yes|string|Name of condition to add|`My condition`|
|`shortname`|yes|string|Name of condition to add `[A-Z]{3}`|`ABC`|
|`evaluate_to_true_on_no_data`|no|bool|Self explanatory|`false`|
|`eval_detail_type`|yes|string|What to evaluate the metric against|`"SINGLE_METRIC"`|
|`metric_aggregation_function`|yes|string|How to aggregate multiple sources of the metric|`"VALUE"`|
|`metric_path`|yes|string|Which metric to use|`"95th Percentile Response Time (ms)"`|
|`metric_eval_detail_type`|yes|string|The type of comparison|`"BASELINE_TYPE"`|
|`baseline_name`|no|string|Which baseline to use|`"All data - Last 15 days"`|
|`baseline_condition`|no|string|How to compare to the baseline|`"WITHIN_BASELINE"`|
|`baseline_unit`|no|string|What unit to compare the baseline with|`"PERCENTAGE"`|
|`compare_condition`|no|string|How to compare the values to the metric|`"GREATER_THAN_SPECIFIC_VALUE"`|
|`compare_value`|yes|number|The value at which the health rule should trigger an error|`2`|

#### affected_entity_type

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

#### tier_or_node

- TIER_AFFECTED_ENTITIES
- NODE_AFFECTED_ENTITIES

#### type_of_node

- ALL_NODES
- JAVA_NODES
- DOT_NET_NODES
- PHP_NODES

#### affected_tier_scope

- ALL_TIERS
- SPECIFIC_TIERS

#### affected_node_scope

- ALL_NODES
- SPECIFIC_NODES
- NODES_OF_SPECIFIC_TIERS
- NODES_MATCHING_PATTERN
- NODE_PROPERTY_VARIABLE_MATCHER

#### nodes_match

- STARTS_WITH
- ENDS_WITH
- CONTAINS
- EQUALS
- MATCH_REG_EX

#### business_transaction_scope

- ALL_BUSINESS_TRANSACTIONS
- SPECIFIC_BUSINESS_TRANSACTIONS
- BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS
- BUSINESS_TRANSACTIONS_MATCHING_PATTERN

#### business_transaction_match

- STARTS_WITH
- ENDS_WITH
- CONTAINS
- EQUALS
- MATCH_REG_EX

#### baseline_unit

- STANDARD_DEVIATIONS
- PERCENTAGE

#### baseline_condition

- WITHIN_BASELINE
- NOT_WITHIN_BASELINE
- GREATER_THAN_BASELINE
- LESS_THAN_BASELINE

#### critical_condition_aggregation_type

- ALL
- ANY

#### warning_condition_aggregation_type

- ALL
- ANY

#### metric_eval_detail_type

- BASELINE_TYPE
- SPECIFIC_TYPE

#### eval_detail_type

- SINGLE_METRIC
- METRIC_EXPRESSION

#### compare_condition

- GREATER_THAN_SPECIFIC_VALUE
- LESS_THAN_SPECIFIC_VALUE

#### metric_aggregation_function

- MINIMUM
- MAXIMUM
- VALUE
- SUM
- COUNT
- CURRENT
- GROUP_COUNT

#### match_type

- AVERAGE
- ANY_NODE
- PERCENTAGE_NODES
- NUMBER_OF_NODES