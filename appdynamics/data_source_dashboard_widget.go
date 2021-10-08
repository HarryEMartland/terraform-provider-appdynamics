package appdynamics

import (
	"encoding/json"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDashboardWidget() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDashboardWidgetRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TIMESERIES_GRAPH",
			},
			"height": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  "4",
			},
			"width": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  "4",
			},
			"x": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  "0",
			},
			"y": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  "0",
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//ID int `json:"id"`
//Version int `json:"version"`
//GUID int `json:"guid"`
//Type string `json:"type"`
//DashboardId int `json:"dashboardId"`
//Height int `json:"height"`
//Width int `json:"width"`
//MinHeight int `json:"minHeight"`
//MinWidth int `json:"minWidth"`
//X int `json:"x"`
//Y int `json:"y"`

func dataSourceDashboardWidgetRead(d *schema.ResourceData, meta interface{}) error {
	//fmt.Println(d.Get("json"))
	widgetType := d.Get("type").(string)
	height := d.Get("height").(int)
	width := d.Get("width").(int)
	x := d.Get("x").(int)
	y := d.Get("y").(int)
	//d.SetId(fmt.Sprintf("%s-%d-%d%d%d", widgetType, height, width, x, y))
	d.SetId(widgetType)

	widget := client.DashboardWidget{
		Type:   widgetType,
		Height: height,
		Width:  width,
		//MinHeight: d.Get("minHeight").(int),
		//MinWidth: d.Get("minWidth").(int),
		X: x,
		Y: y,
	}
	jsonDoc, err := json.MarshalIndent(widget, "", "  ")
	if err != nil {
		return err
	}
	jsonString := string(jsonDoc)
	d.Set("json", jsonString)
	return nil
}
