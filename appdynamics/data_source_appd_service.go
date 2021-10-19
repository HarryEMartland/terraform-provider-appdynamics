package appdynamics

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAppdService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAppdServiceRead,
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tier_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func dataSourceAppdServiceRead(d *schema.ResourceData, meta interface{}) error {
	tierName := d.Get("tier_name").(string)
	applicationName := d.Get("application_name").(string)
	hash := sha256.Sum224([]byte(tierName))
	hashString := hex.EncodeToString(hash[:])

	fmt.Println(applicationName)
	fmt.Println(tierName)
	d.SetId(hashString)
	return nil
}
