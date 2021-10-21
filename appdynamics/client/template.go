package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func (c *AppDClient) createTemplateBaseUrl() string {
	return fmt.Sprintf("%s/controller/restui/templates", c.BaseUrl)
}

func (c *AppDClient) createSetAssociatedDashboardsUrl(tierId int) string {
	return fmt.Sprintf("%s/setAssociatedDashboards/%d?isTierDashboard=true", c.createTemplateBaseUrl(), tierId)
}

func (c *AppDClient) createGetAllDashboardTemplatesByTierUrl(tierId int) string {
	return fmt.Sprintf("%s/getAllDashboardTemplatesByTier/%d?isTierDashboard=true", c.createTemplateBaseUrl(), tierId)
}

func (c *AppDClient) SetTemplateDashboardAssociations(tierId int, associations interface{}) error {
	resp, err := req.Post(c.createSetAssociatedDashboardsUrl(tierId), c.createAuthHeader(), req.BodyJSON(associations))
	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error creating association: %d, %s", resp.Response().StatusCode, respString))
	}
	return err
}

func (c *AppDClient) GetAllDashboardTemplatesByTier(tierId int) ([]Dashboard, error) {
	resp, err := req.Get(c.createGetAllDashboardTemplatesByTierUrl(tierId), c.createAuthHeader())
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating association: %d, %s", resp.Response().StatusCode, respString))
	}
	dashboardTemplates := make([]Dashboard, 0)
	err = resp.ToJSON(&dashboardTemplates)

	if err != nil {
		return nil, err
	}
	return dashboardTemplates, nil
}
