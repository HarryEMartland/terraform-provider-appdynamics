package appdynamics

import (
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTierTemplateAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTierTemplateAssociationSet,
		Read:   resourceTierTemplateAssociationRead,
		Update: resourceTierTemplateAssociationSet,
		Delete: resourceTierTemplateAssociationDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tier_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"template_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceTierTemplateAssociationSet(d *schema.ResourceData, m interface{}) error {
	tierId := d.Get("tier_id").(int)
	applicationId := d.Get("application_id").(int)
	d.SetId(fmt.Sprintf("%d%d", tierId, applicationId))
	templateIds := d.Get("template_ids").(*schema.Set).List()
	appdClient := m.(*client.AppDClient)
	err := appdClient.SetTemplateDashboardAssociations(tierId, templateIds)
	return err
}

func resourceTierTemplateAssociationRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	tierId := d.Get("tier_id").(int)
	templates, err := appdClient.GetAllDashboardTemplatesByTier(tierId)
	if err != nil {
		return err
	}
	templateIds := schema.NewSet(schema.HashInt, []interface{}{})
	for _, template := range templates {
		templateIds.Add(template.ID)
	}
	d.Set("template_ids", templateIds)
	return nil
}

func resourceTierTemplateAssociationDelete(d *schema.ResourceData, m interface{}) error {
	tierId := d.Get("tier_id").(int)
	templateIds := make([]int, 0)
	appdClient := m.(*client.AppDClient)
	err := appdClient.SetTemplateDashboardAssociations(tierId, templateIds)
	return err
}
