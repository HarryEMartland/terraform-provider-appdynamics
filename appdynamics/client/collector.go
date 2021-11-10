package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/imroc/req"
)

type Collector struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AgentName string `json:"agentName"`
	Enabled   bool   `json:"enabled"`
}

func (c *AppDClient) createCollectorBaseUrl() string {
	return fmt.Sprintf("%s/controller/rest/databases/collectors", c.BaseUrl)
}

func (c *AppDClient) createCollectorUrl() string {
	return fmt.Sprintf("%s/create", c.createCollectorBaseUrl())
}

func (c *AppDClient) updateCollectorUrl() string {
	return fmt.Sprintf("%s/update", c.createCollectorBaseUrl())
}

func (c *AppDClient) deleteCollectorUrl(collectorId int) string {
	return fmt.Sprintf("%s/%d", c.createCollectorBaseUrl(), collectorId)
}

func (c *AppDClient) getCollectorUrl(collectorId int) string {
	return fmt.Sprintf("%s/%d", c.createCollectorBaseUrl(), collectorId)
}

func (c *AppDClient) CreateCollector(collector *Collector) (string, error) {
	url := c.createCollectorUrl()
	auth := c.createAuthHeader()
	body := req.BodyJSON(collector)
	resp, err := req.Post(url, auth, body)
	if err != nil {
		return "", err
	}
	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		return "", errors.New(fmt.Sprintf("Error creating Collector: %d, %s\nurl=[%s]", resp.Response().StatusCode, respString, c.createCollectorUrl()))
	}
	//eg.  http://worldremit-test.saas.appdynamics.com/controller/rest/databases/collectors/create/1540
	locationHeader := resp.Response().Header.Get("Location")
	id := locationHeader[strings.LastIndex(locationHeader, "/")+1:]
	fmt.Printf("id=%v\n", id)

	return id, nil
}

func (c *AppDClient) DeleteCollector(collectorId int) error {
	resp, err := req.Delete(c.deleteCollectorUrl(collectorId), c.createAuthHeader())
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Collector: %d, %s", resp.Response().StatusCode, respString))
	}
	return nil
}

func (c *AppDClient) UpdateCollector(collector Collector) (*Collector, error) {
	req.Debug = true
	resp, err := req.Post(c.updateCollectorUrl(), c.createAuthHeader(), req.BodyJSON(collector))
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Collector: %d, %s", resp.Response().StatusCode, respString))
	}
	updated := Collector{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, err
}

func (c *AppDClient) GetCollector(collectorId int) (*Collector, error) {
	resp, err := req.Get(c.getCollectorUrl(collectorId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting collector: %d, %s", resp.Response().StatusCode, respString))
	}
	collector := Collector{}
	err = resp.ToJSON(&collector)
	if err != nil {
		return nil, err
	}
	return &collector, nil
}
