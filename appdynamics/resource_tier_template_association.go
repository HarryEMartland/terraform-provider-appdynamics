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
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceTierTemplateAssociationSet(d *schema.ResourceData, m interface{}) error {
	tierId := d.Get("tier_id").(int)
	applicationId := d.Get("application_id").(int)
	d.SetId(fmt.Sprintf("%d%d", tierId, applicationId))
	templateIdsRaw := d.Get("template_ids").([]interface{})
	templateIds := make([]int, len(templateIdsRaw))
	for i, rawId := range templateIdsRaw {
		templateIds[i] = rawId.(int)
	}
	appdClient := m.(*client.AppDClient)
	_, err := appdClient.SetTemplateDashboardAssociations(tierId, templateIds)
	if err != nil {
		return err
	}
	return nil
}

func resourceTierTemplateAssociationRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	tierId := d.Get("tier_id").(int)
	templates, err := appdClient.GetAllDashboardTemplatesByTier(tierId)

	if err != nil {
		return err
	}

	templateIds := make([]int, len(templates))
	for i, template := range templates {
		templateIds[i] = template.ID
	}
	d.Set("template_ids", templateIds)
	return nil
}

func resourceTierTemplateAssociationDelete(d *schema.ResourceData, m interface{}) error {
	tierId := d.Get("tier_id").(int)
	templateIdsRaw := d.Get("template_ids").([]interface{})
	templateIds := make([]int, len(templateIdsRaw))
	appdClient := m.(*client.AppDClient)
	_, err := appdClient.SetTemplateDashboardAssociations(tierId, templateIds)

	if err != nil {
		return err
	}
	return nil
}
