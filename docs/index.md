# AppDynamics Provider

The AppDynamics Provider allows for the configuration of various settings within AppDynamics.
Both SaaS and self hosted instances are supported as long as terraform is run from a place which can reach the controller.
A secret is required to make changes to the controller, documentation to create a secret can be found [here](https://docs.appdynamics.com/display/PRO45/API+Clients).
You must also provide terraform with the base url of the controller, an example is provided below.


## Example Usage

```hcl
provider "appdynamics" {
  secret = "<your secret>"
  controller_base_url = "https://example.saas.appdynamics.com"
}
```

## Argument Reference

|       Name          |Required| Type |            Description             |
|---------------------|--------|------|------------------------------------|
|`secret`             |yes     |number|The application to add the action to|
|`controller_base_url`|yes     |string|The name for the action             |
