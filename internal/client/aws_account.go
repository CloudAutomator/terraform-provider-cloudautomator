package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type AwsAccount struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type GetAwsAccountResponse struct {
	Data json.RawMessage `json:"data"`
}

type ListAwsAccountResponse struct {
	Data  []json.RawMessage `json:"data"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

type RawAwsAccountData struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes AwsAccountAttributes `json:"attributes"`
}

type AwsAccountAttributes struct {
	Name string `json:"name"`
}

func (c *Client) GetAwsAccount(group_id, awsAccountId string) (*AwsAccount, *http.Response, error) {
	requestUrl := fmt.Sprintf("/groups/%s/aws_accounts/%s", group_id, awsAccountId)
	req, err := c.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	getResponse := new(GetAwsAccountResponse)
	resp, err := c.Do(req, &getResponse)
	if err != nil {
		return nil, resp, err
	}

	awsAccount := new(AwsAccount)
	json.Unmarshal(getResponse.Data, &awsAccount)

	return awsAccount, resp, nil
}

func (c *Client) GetAwsAccounts(group_id string) (*[]AwsAccount, *http.Response, error) {
	awsAccounts := []AwsAccount{}
	requestUrl := fmt.Sprintf("/groups/%s/aws_accounts", group_id)

	for len(requestUrl) > 0 {
		rel, err := url.Parse(requestUrl)
		if err != nil {
			return nil, nil, err
		}

		q := rel.Query()
		q.Set("page[size]", "100")
		rel.RawQuery = q.Encode()

		req, err := c.NewRequest("GET", rel.String(), nil)
		if err != nil {
			return nil, nil, err
		}

		listResponse := new(ListAwsAccountResponse)
		resp, err := c.Do(req, listResponse)
		if err != nil {
			return nil, resp, err
		}

		for _, r := range listResponse.Data {
			awsAccount := new(AwsAccount)
			if err := json.Unmarshal(r, &awsAccount); err != nil {
				return nil, nil, errors.New("unmarshal failed")
			}
			awsAccounts = append(awsAccounts, *awsAccount)
		}

		requestUrl = listResponse.Links.Next
	}

	return &awsAccounts, nil, nil
}

func (a *AwsAccount) UnmarshalJSON(data []byte) error {
	ra := RawAwsAccountData{}
	if err := json.Unmarshal(data, &ra); err != nil {
		return errors.New("unmarshal failed")
	}

	a.Id = ra.Id
	a.Name = ra.Attributes.Name

	return nil
}
