package appdynamics

import (
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceImportExportDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceImportExportDashboardCreate,
		Read:   resourceImportExportDashboardRead,
		Update: resourceImportExportDashboardUpdate,
		Delete: resourceDashboardDelete,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceImportExportDashboardCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	dashboardJSON := d.Get("json").(string)
	dashboard, _ := appdClient.ImportDashboard(dashboardJSON)
	dashboardId := strconv.Itoa(dashboard.ID)
	d.SetId(dashboardId)
	return nil
}

func resourceImportExportDashboardUpdate(d *schema.ResourceData, m interface{}) error {
	err := resourceDashboardDelete(d, m)
	if err != nil {
		return err
	}
	err = resourceImportExportDashboardCreate(d, m)
	if err != nil {
		return err
	}
	return nil
}

func resourceImportExportDashboardRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	id := d.Id()
	dashboardId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	dashboard, err := appdClient.GetDashboard(dashboardId)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(dashboard.ID))

	return nil
}
