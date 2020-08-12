# Transaction Detection Rule Resource

Creates a rule which defines when to creat a transaction within an application.

[AppDynamics Transaction Detection Rule Documentation](https://docs.appdynamics.com/display/PRO45/Transaction+Detection+Rules)  

## Example Usage

#### NodeJs

```hcl
resource "appdynamics_transaction_detection_rule" "test_rule" {
  application_id      = var.application_id
  account_id          = var.account_id
  scope_id            = var.scope_id
  name                = "My test rule"
  agent_type          = "NODE_JS_SERVER"
  description         = "My test health rule created with Terraform"
  entry_point_type    = "NODEJS_WEB"
  http_uri_match_type = "EQUALS"
  http_uris           = ["/user/account"]
  http_method         = "GET"
}
```

#### NodeJs Regex

```hcl
resource "appdynamics_transaction_detection_rule" "regex_test_rule" {
  application_id      = var.application_id
  account_id          = var.account_id
  scope_id            = var.scope_id
  name                = "My regex test rule"
  agent_type          = "NODE_JS_SERVER"
  description         = "My test regex health rule created with Terraform"
  entry_point_type    = "NODEJS_WEB"
  http_uri_match_type = "MATCHES_REGEX"
  http_uris           = ["/user/.*"]
  http_method         = "GET"
}
```

#### NodeJs List

```hcl
resource "appdynamics_transaction_detection_rule" "list_test_rule" {
  application_id      = var.application_id
  account_id          = var.account_id
  scope_id            = var.scope_id
  name                = "My list test rule"
  agent_type          = "NODE_JS_SERVER"
  description         = "My test regex health rule created with Terraform"
  entry_point_type    = "NODEJS_WEB"
  http_uri_match_type = "IS_IN_LIST"
  http_uris           = ["/user/.*", "/user/.*/"]
  http_method         = "GET"
}
```

#### Java

```hcl
resource "appdynamics_transaction_detection_rule" "java_test_rule" {
  application_id      = var.application_id
  account_id          = var.account_id
  scope_id            = var.scope_id
  name                = "My java test rule"
  agent_type          = "APPLICATION_SERVER"
  description         = "My test java health rule created with Terraform"
  entry_point_type    = "SERVLET"
  http_uri_match_type = "EQUALS"
  http_uris           = ["/user/account"]
  http_method         = "GET"
}
```

## Argument Reference

|Name|Required|Type|Description|Example|
|----|--------|----|-----------|-------|
|`application_id`|yes|number|The application to add the action to|`32423`|
|`account_id`|yes|string|The account id for the rule|`"dffa443e-3634-415c-9755-317ee5ddbbbc"`|
|`name`|yes|string|The name for the transaction detection rule|`"My Rule"`|
|`description`|yes|string|A description of the rule|`"a description"`|
|`enabled`|no|string|If the rule is enabled|`true`|
|`priority`|no|string|The priority of the rule compared to other rules|`10`|
|`agent_type`|yes|string|Which agent types the rule should run for|`"NODE_JS_SERVER"`|
|`entry_point_type`|yes|string|The type of transaction entry point|`"NODEJS_WEB"`|
|`http_uri_match_type`|no|string|How to match uris|`"EQUALS"`|
|`http_uris`|no|string|A list of uris to match for the transaction|`["/test"]`|
|`http_method`|no|string|The method of the transactions|`"GET"`|


#### agent_type
- DOT_NET_APPLICATION_SERVER
- NODE_JS_SERVER
- APPLICATION_SERVER
- PHP_APPLICATION_SERVER
- NATIVE_WEB_SERVER
- PYTHON_SERVER

#### http_uri_match_type
- EQUALS
- CONTAINS
- STARTS_WITH
- ENDS_WITH
- MATCHES_REGEX
- IS_IN_LIST
- IS_NOT_EMPTY

#### entry_point_type
- SERVLET
- ASP_DOTNET
- PHP_WORDPRESS
- PHP_WEB
- PHP_DRUPAL
- PHP_MVC
- PHP_WEB_SERVICE
- NODEJS_WEB
- WEB
- SPRING_BEAN
- STRUTS_ACTION
- WEB_SERVICE
- POJO
- JMS
- EJB
- POCO
- DOTNET_JMS
- DOTNET_REMOTING
- ASP_DOTNET_WEB_SERVICE
- WCF
- PYTHON_WEB

#### http_method
- GET
- POST
- PUT
- DELETE