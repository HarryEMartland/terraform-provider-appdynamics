package main

import (
"github.com/hashicorp/terraform-plugin-sdk/plugin"
"github.com/hashicorp/terraform-plugin-sdk/terraform"
	. "github.com/HarryEMartland/appd-terraform-provider/appd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}