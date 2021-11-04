package appdynamics

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAppdService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAppdServiceRead,
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tier_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tier_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceAppdServiceRead(d *schema.ResourceData, m interface{}) error {
	tierName := d.Get("tier_name").(string)
	applicationName := d.Get("application_name").(string)
	hash := sha256.Sum224([]byte(tierName))
	hashString := hex.EncodeToString(hash[:])
	appdClient := m.(*client.AppDClient)

	d.SetId(hashString)
	application, err := appdClient.GetApplicationByName(applicationName)
	if err != nil {
		return err
	}

	tier, err := appdClient.GetApplicationTiers(applicationName, tierName)

	if err != nil {
		return err
	}

	d.Set("application_id", application.ID)
	d.Set("tier_id", tier.ID)

	return nil
}
