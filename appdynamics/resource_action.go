package appdynamics

import (
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func resourceAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceActionCreate,
		Read:   resourceActionRead,
		Update: resourceActionUpdate,
		Delete: resourceActionDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"action_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"SMS",
					"EMAIL",
					"CUSTOM_EMAIL",
					"THREAD_DUMP",
					"HTTP_REQUEST",
					"CUSTOM",
					"RUN_SCRIPT_ON_NODES",
					"DIAGNOSE_BUSINESS_TRANSACTIONS",
					"CREATE_UPDATE_JIRA",
				}),
			},
			"emails": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"phone_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_request_template_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_template_variables": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
func resourceActionCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	action := createAction(d)

	updatedHealthRule, err := appdClient.CreateAction(&action, applicationId)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(updatedHealthRule.ID))

	return resourceActionRead(d, m)
}

func createAction(d *schema.ResourceData) client.Action {

	name := d.Get("name").(string)
	actionType := d.Get("action_type").(string)
	emails := d.Get("emails").(*schema.Set).List()

	varialbesMap := d.Get("custom_template_variables").(map[string]interface{})
	var varialbesList []*client.ActionVariable

	for k, v := range varialbesMap {
		varialbesList = append(varialbesList, &client.ActionVariable{
			Key:   k,
			Value: v.(string),
		})
	}

	healthRule := client.Action{
		Name:                    name,
		ActionType:              actionType,
		Emails:                  emails,
		PhoneNumber:             d.Get("phone_number").(string),
		HttpRequestTemplateName: d.Get("http_request_template_name").(string),
		CustomTemplateVariables: varialbesList,
	}
	return healthRule
}

func updateAction(d *schema.ResourceData, action client.Action) {
	d.Set("name", action.Name)
	d.Set("action_type", action.ActionType)
	d.Set("emails", trimSpaceA(action.Emails))
	d.Set("phone_number", action.PhoneNumber)
	d.Set("http_request_template_name", action.HttpRequestTemplateName)
	d.Set("custom_template_variables", action.CustomTemplateVariables)
}

func trimSpaceA(array []interface{}) []interface{} {
	for i, str := range array {
		array[i] = strings.TrimSpace(str.(string))
	}
	return array
}

func resourceActionRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	actionId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	action, err := appdClient.GetAction(actionId, applicationId)
	if err != nil {
		return err
	}

	updateAction(d, *action)

	return nil
}

func resourceActionUpdate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)

	healthRule := createAction(d)

	healthRuleId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	healthRule.ID = healthRuleId

	_, err = appdClient.UpdateAction(&healthRule, applicationId)
	if err != nil {
		return err
	}

	return resourceActionRead(d, m)
}

func resourceActionDelete(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	actionId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = appdClient.DeleteAction(applicationId, actionId)
	if err != nil {
		return err
	}

	return nil
}
