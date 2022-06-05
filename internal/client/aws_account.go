package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AwsAccount struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type AwsAccountGetResponse struct {
	Data json.RawMessage `json:"data"`
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

	getResponse := new(AwsAccountGetResponse)
	resp, err := c.Do(req, &getResponse)
	if err != nil {
		return nil, resp, err
	}

	awsAccount := new(AwsAccount)
	json.Unmarshal(getResponse.Data, &awsAccount)

	return awsAccount, resp, nil
}

func (a *AwsAccount) UnmarshalJSON(data []byte) error {
	ra := RawAwsAccountData{}
	if err := json.Unmarshal(data, &ra); err != nil {
		return errors.New("unmarshall failed")
	}

	a.Id = ra.Id
	a.Name = ra.Attributes.Name

	return nil
}
