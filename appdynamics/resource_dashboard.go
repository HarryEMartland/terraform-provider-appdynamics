package appdynamics

import (
	"encoding/json"
	"fmt"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

//"name":"abcde",
//"width":1024,
//"height":768,
//"canvasType":
//"CANVAS_TYPE_GRID",
//"templateEntityType":"APPLICATION_COMPONENT_NODE",
//"refreshInterval":120000,
//"backgroundColor":15856629,
//"warRoom":false,
//"template":false,
//"widgets":[],
//"version":0,
//"minutesBeforeAnchorTime":-1,
//"startTime":-1,
//"endTime":-1

func resourceDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceDashboardCreate,
		Read:   resourceDashboardRead,
		Update: resourceDashboardUpdate,
		Delete: resourceDashboardDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"width": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"height": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  768,
			},
			"canvas_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CANVAS_TYPE_GRID",
			},
			"template_entity_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "APPLICATION_COMPONENT_NODE",
			},
			"refresh_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  120000,
			},
			"background_color": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15856629,
			},
			"war_room": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"template": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"widgets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"minutes_before_anchor_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"end_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
		},
	}
}

func resourceDashboardCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	dashboard := createDashboard(d)
	dash, _ := appdClient.CreateDashboard(dashboard)
	d.SetId(strconv.Itoa(dash.ID))
	return resourceDashboardRead(d, m)
}

func createDashboard(d *schema.ResourceData) client.Dashboard {
	widgetList := d.Get("widgets").([]interface{})
	var dashboardWidgetList []client.DashboardWidget
	for _, widget := range widgetList {
		test := client.DashboardWidget{}
		json.Unmarshal([]byte(widget.(string)), &test)
		dashboardWidgetList = append(dashboardWidgetList, test)
	}

	dashboard := client.Dashboard{
		Name:                    d.Get("name").(string),
		Width:                   d.Get("width").(int),
		Height:                  d.Get("height").(int),
		CanvasType:              d.Get("canvas_type").(string),
		TemplateEntityType:      d.Get("template_entity_type").(string),
		RefreshInterval:         d.Get("refresh_interval").(int),
		BackgroundColor:         d.Get("background_color").(int),
		WarRoom:                 d.Get("war_room").(bool),
		Template:                d.Get("template").(bool),
		Widgets:                 dashboardWidgetList,
		Version:                 d.Get("version").(int),
		MinutesBeforeAnchorTime: d.Get("minutes_before_anchor_time").(int),
		StartTime:               d.Get("start_time").(int),
		EndTime:                 d.Get("end_time").(int),
	}
	return dashboard
}

func updateDashboard(d *schema.ResourceData, dashboard client.Dashboard) {
	d.Set("ID", dashboard.ID)
	d.Set("Name", dashboard.Name)
	//d.Set("Template", dashboard.Template)

}

func resourceDashboardRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)

	fmt.Println("1")
	id := d.Id()
	dashboardId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	dashboard, err := appdClient.GetDashboard(dashboardId)
	if err != nil {
		return err
	}
	updateDashboard(d, *dashboard)
	return nil
}

func resourceDashboardUpdate(d *schema.ResourceData, m interface{}) error {
	//appdClient := m.(*client.AppDClient)
	//applicationId := d.Get("application_id").(int)
	//
	//healthRule := createDashboard(d)
	//
	//healthRuleId, err := strconv.Atoi(d.Id())
	//if err != nil {
	//	return err
	//}
	//healthRule.ID = healthRuleId
	//
	//_, err = appdClient.UpdateDashboard(&healthRule, applicationId)
	//if err != nil {
	//	return err
	//}
	//
	//return resourceDashboardRead(d, m)
	return nil
}

func resourceDashboardDelete(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Trying to delete resource")
	appdClient := m.(*client.AppDClient)

	id := d.Id()
	dashboardId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	err = appdClient.DeleteDashboard(dashboardId)
	if err != nil {
		return err
	}

	return nil
}

//func contains(s []string, e string) bool {
//	for _, a := range s {
//		if a == e {
//			return true
//		}
//	}
//	return false
//}
