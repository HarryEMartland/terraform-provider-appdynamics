package appdynamics

import (
	"strconv"

	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceCollectorCreate,
		Read:   resourceCollectorRead,
		Update: resourceCollectorUpdate,
		Delete: resourceCollectorDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"COUCHBASE",
					"CASSANDRA",
					"DB2",
					"MONGO",
					"MSSQL",
					"MYSQL",
					"ORACLE",
					"POSTGRESQL",
					"SYBASE",
					"SQLAZURE",
				}),
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"agent_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceCollectorCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)

	collector := createCollector(d)

	id, err := appdClient.CreateCollector(&collector)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func createCollector(d *schema.ResourceData) client.Collector {

	collector := client.Collector{
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
		Hostname:  d.Get("hostname").(string),
		Port:      d.Get("port").(int),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		AgentName: d.Get("agent_name").(string),
		Enabled:   d.Get("enabled").(bool),
	}
	return collector
}

func resourceCollectorUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)

	collector := createCollector(d)
	collectorId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	collector.ID = collectorId

	_, err = appdClient.UpdateCollector(collector)
	if err != nil {
		return err
	}

	return resourceCollectorRead(d, m)
}

func resourceCollectorRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	id := d.Id()

	collectorID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	collector, err := appdClient.GetCollector(collectorID)
	if err != nil {
		return err
	}

	updateCollector(d, *collector)

	return nil
}

func resourceCollectorDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	id := d.Id()

	collectorID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = appdClient.DeleteCollector(collectorID)
	if err != nil {
		return err
	}

	return nil
}

func updateCollector(d *schema.ResourceData, collector client.Collector) {
	d.Set("name", collector.Name)
	d.Set("type", collector.Type)
	d.Set("hostname", collector.Hostname)
	d.Set("port", collector.Port)
	d.Set("username", collector.Username)
	d.Set("enabled", collector.Enabled)
	// Password is always set as `appdynamics_redacted_password` so we need to always overwrite this
	//d.Set("password", collector.Password)
	d.Set("agent_name", collector.AgentName)
}
