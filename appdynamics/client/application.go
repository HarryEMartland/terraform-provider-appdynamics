package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

type Tier struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	AgentType     string `json:"agentType"`
	Description   string `json:"description"`
	NumberOfNodes int    `json:"numberOfNodes"`
}

type Tiers []Tier

type Application struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccountGuid int    `json:"accountGuid"`
}

type Applications []Application

func (c *AppDClient) createApplicationsBaseUrl() string {
	return fmt.Sprintf("%s/controller/rest/applications", c.BaseUrl)
}

func (c *AppDClient) createGetApplicationUrl(applicationName string) string {
	return fmt.Sprintf("%s/%s?output=JSON", c.createApplicationsBaseUrl(), applicationName)
}
func (c *AppDClient) createGetApplicationTierUrl(applicationName string, tierName string) string {
	return fmt.Sprintf("%s/%s/tiers/%s?output=JSON", c.createApplicationsBaseUrl(), applicationName, tierName)
}

func (c *AppDClient) GetApplicationByName(applicationName string) (*Application, error) {
	resp, err := req.Get(c.createGetApplicationUrl(applicationName), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting application: %d, %s", resp.Response().StatusCode, respString))
	}
	applications := Applications{}
	err = resp.ToJSON(&applications)

	if len(applications) == 0 {
		return nil, errors.New(fmt.Sprintf("Application not exist"))
	}
	return &applications[0], nil
}

func (c *AppDClient) GetApplicationTiers(applicationName string, tierName string) (*Tier, error) {
	resp, err := req.Get(c.createGetApplicationTierUrl(applicationName, tierName), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting tiers: %d, %s", resp.Response().StatusCode, respString))
	}
	tiers := Tiers{}
	err = resp.ToJSON(&tiers)

	if len(tiers) == 0 {
		return nil, errors.New(fmt.Sprintf("Tier not exist"))
	}
	return &tiers[0], nil
}
