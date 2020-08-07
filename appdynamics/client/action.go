package client

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func (c *AppDClient) CreateAction(action *Action, applicationId int) (*Action, error) {

	resp, err := req.Post(c.createActionsUrl(applicationId), c.createAuthHeader(), req.BodyJSON(&action))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Action: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Action{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) UpdateAction(action *Action, applicationId int) (*Action, error) {

	resp, err := req.Put(c.createActionUrl(action.ID, applicationId), c.createAuthHeader(), req.BodyJSON(&action))
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Action: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Action{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) DeleteAction(applicationId int, actionId int) error {

	resp, err := req.Delete(c.createActionUrl(actionId, applicationId), c.createAuthHeader())
	if err != nil {
		return err
	}

	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Action: %d, %s", resp.Response().StatusCode, respString))
	}

	return nil
}

func (c *AppDClient) GetAction(actionId int, applicationId int) (*Action, error) {

	resp, err := req.Get(c.createActionUrl(actionId, applicationId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error deleting Action: %d, %s", resp.Response().StatusCode, respString))
	}

	updated := Action{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (c *AppDClient) createActionsUrl(applicationId int) string {
	return fmt.Sprintf("%s/%s", c.createUrl(applicationId), "actions")
}

func (c *AppDClient) createActionUrl(actionId int, applicationId int) string {
	return fmt.Sprintf("%s/%d", c.createActionsUrl(applicationId), actionId)
}

type ActionVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Action struct {
	ID                      int               `json:"id"`
	ActionType              string            `json:"actionType"`
	Name                    string            `json:"name"`
	Emails                  []interface{}     `json:"emails"`
	PhoneNumber             string            `json:"phoneNumber"`
	HttpRequestTemplateName string            `json:"httpRequestTemplateName"`
	CustomTemplateVariables []*ActionVariable `json:"customTemplateVariables"`
}
