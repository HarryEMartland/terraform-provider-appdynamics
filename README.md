## AppD Terraform Provider

A Terraform Provider for AppDynamics.
Provides resources for creating Health Rules, Actions and configuring Transaction detection through Terraform.

#### Download

Download the latest version of the provider and place it where terraform will find it e.g. `~/.terraform.d/plugins/`.
See the [terraform documentation](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery) for more information.
The latest downloads can be found attached to the [latest release](https://github.com/HarryEMartland/appd-terraform-provider/releases/latest) in github or on the below links.
Make sure to download the correct version for your OS.

#### Configuration

To use the AppD Terraform provider you must configure it with the controller base url and a secret.
A secret can be generated in the AppD UI as documented [here](https://docs.appdynamics.com/display/PRO45/API+Clients).

###### Example
```terraform
provider "appd" {
  secret = "<your secret>"
  controller_base_url = "https://example.saas.appdynamics.com"
}
```

#### Resources

- [appd_action](docs/resources/action.md)
- [appd_policy](docs/resources/policy.md)
- [appd_health_rule](docs/resources/health_rule.md)

#### Building

```shell script
make install build
```

#### Testing

###### Unit Tests
```shell script
make test
```