## AppD Terraform Provider

A Terraform Provider for AppDynamics.
Provides resources for creating Health Rules, Actions and configuring Transaction detection through Terraform.

#### Configuration

To use the AppD Teraform provider you must configure it with the controller base url and a secret.
A secret can be generated in the AppD UI as documented [here](https://docs.appdynamics.com/display/PRO45/API+Clients).

###### Example
```terraform
provider "appd" {
  secret = "<your secret>"
  controller_base_url = "https://example.saas.appdynamics.com"
}
```

#### Resources

- [appd_action](docs/action_resource.md)
- [appd_policy](docs/policy_resource.md)
- [appd_health_rule](docs/health_rule_resource.md)

#### Building

```shell script
go get
go build -o terraform-provider-appd && chmod +x ./terraform-provider-appd
```