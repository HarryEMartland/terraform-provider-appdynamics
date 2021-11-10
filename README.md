## AppDynamics Terraform Provider

A Terraform Provider for AppDynamics.
Provides resources for creating Health Rules, Actions and configuring Transaction detection through Terraform.

#### Download

Download the latest version of the provider and place it where terraform will find it e.g. `~/.terraform.d/plugins/`.
See the [terraform documentation](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery) for more information.
The latest downloads can be found in the [Terraform Registry](https://registry.terraform.io/providers/HarryEMartland/appdynamics/latest) and attached to the [latest release](https://github.com/HarryEMartland/terraform-provider-appdynamics/releases/latest).
Make sure to download the correct version for your OS.

#### Configuration

To use the AppDynamics Terraform provider you must configure it with the controller base url and a secret.
A secret can be generated in the AppDynamics UI as documented [here](https://docs.appdynamics.com/display/PRO45/API+Clients).

###### Token access example
```terraform
provider "appdynamics" {
  secret = "<your temporary token>"
  controller_base_url = "https://example.saas.appdynamics.com"
}
```

###### Client name/secret access example
```terraform
provider "appdynamics" {
  controller_base_url = "https://example.saas.appdynamics.com"
  client_name = "<your client name>"
  client_secret = "<your client secret>"
}
```

#### Resources

- [appdynamics_action](docs/resources/action.md)
- [appdynamics_policy](docs/resources/policy.md)
- [appdynamics_health_rule](docs/resources/health_rule.md)
- [appdynamics_transaction_detection_rule](docs/resources/transaction_detection_rule.md)

#### Building

```shell script
make install build
```

#### Testing

###### Unit Tests
```shell script
make test
```
