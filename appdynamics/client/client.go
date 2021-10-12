package client

import (
	"fmt"
	"github.com/imroc/req"
)

type AppDClient struct {
	BaseUrl string
	Secret  string
}

func (c *AppDClient) createUrl(applicationId int) string {
	return fmt.Sprintf("%s/controller/alerting/rest/v1/applications/%d", c.BaseUrl, applicationId)
}

func (c *AppDClient) createAuthHeader() req.Header {
	return req.Header{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", c.Secret),
		"Accept":        "application/json",
	}
}
