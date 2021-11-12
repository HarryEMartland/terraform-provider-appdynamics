package appdynamics

import (
	"errors"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"appdynamics_health_rule":                resourceHealthRule(),
			"appdynamics_action":                     resourceAction(),
			"appdynamics_collector":                  resourceCollector(),
			"appdynamics_policy":                     resourcePolicy(),
			"appdynamics_transaction_detection_rule": resourceTransactionDetectionRule(),
			"appdynamics_dashboard":                  resourceDashboard(),
			"appdynamics_import_export_dashboard":    resourceImportExportDashboard(),
			"appdynamics_tier_template_association":  resourceTierTemplateAssociation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"appdynamics_dashboard_widget": dataSourceDashboardWidget(),
			"appdynamics_appd_service":     dataSourceAppdService(),
		},
		Schema: map[string]*schema.Schema{
			"secret":              {Type: schema.TypeString, Sensitive: true, Optional: true},
			"controller_base_url": {Type: schema.TypeString, Required: true},
			"client_name":         {Type: schema.TypeString, Optional: true},
			"client_secret":       {Type: schema.TypeString, Sensitive: true, Optional: true},
		},
		ConfigureFunc: func(data *schema.ResourceData) (interface{}, error) {
			controllerBaseUrl := data.Get("controller_base_url").(string)
			clientName := data.Get("client_name").(string)
			clientSecret := data.Get("client_secret").(string)
			token := data.Get("secret").(string)

			if clientName == "" && clientSecret == "" && token == "" {
				return nil, errors.New("please provide token or client_name/client_secret pair")
			}

			if clientName != "" && clientSecret != "" {
				accessToken, err := client.CreateAccessToken(controllerBaseUrl, clientName, clientSecret)
				if err != nil {
					return nil, err
				}
				token = *accessToken
			}

			return &client.AppDClient{
				BaseUrl: controllerBaseUrl,
				Secret:  token,
			}, nil
		},
	}
}
