package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type JobWorkflow struct {
	Id              string `json:"id,omitempty"`
	Active          *bool  `json:"active,omitempty"`
	Name            string `json:"name"`
	GroupId         int    `json:"group_id"`
	FirstJobId      int    `json:"first_job_id"`
	FollowingJobIds []int  `json:"following_job_ids"`
}

type JobWorkflowGetResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobWorkflowPostResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobWorkflowPatchResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobWorkflowListResponse struct {
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

type RawJobWorkflowData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes JobWorkflowAttributes `json:"attributes"`
}

type JobWorkflowAttributes struct {
	Name            string    `json:"name"`
	Active          *bool     `json:"active,omitempty"`
	GroupId         int       `json:"group_id"`
	FirstJobId      int       `json:"first_job_id"`
	FollowingJobIds []int     `json:"following_job_ids"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (c *Client) GetJobWorkflow(jobWorkflowId string) (*JobWorkflow, *http.Response, error) {
	requestUrl := fmt.Sprintf("job_workflows/%s", jobWorkflowId)

	getResponse := new(JobWorkflowGetResponse)
	resp, err := c.requestWithRetry("GET", requestUrl, nil, getResponse, 5)
	if err != nil {
		return nil, resp, err
	}

	jobWorkflow := new(JobWorkflow)
	if err := json.Unmarshal(getResponse.Data, &jobWorkflow); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	return jobWorkflow, resp, nil
}

func (c *Client) GetJobWorkflows() (*[]JobWorkflow, *http.Response, error) {
	jobWorkflows := []JobWorkflow{}
	requestUrl := "job_workflows"

	for len(requestUrl) > 0 {
		rel, err := url.Parse(requestUrl)
		if err != nil {
			return nil, nil, err
		}

		q := rel.Query()
		q.Set("page[size]", "100")
		rel.RawQuery = q.Encode()

		listResponse := new(JobWorkflowListResponse)
		resp, err := c.requestWithRetry("GET", rel.String(), nil, listResponse, defaultRetryCount)

		if err != nil {
			return nil, resp, err
		}

		for _, r := range listResponse.Data {
			jobWorkflow := new(JobWorkflow)
			if err := json.Unmarshal(r, &jobWorkflow); err != nil {
				return nil, nil, errors.New("unmarshal failed")
			}
			jobWorkflows = append(jobWorkflows, *jobWorkflow)
		}

		requestUrl = listResponse.Links.Next
	}

	return &jobWorkflows, nil, nil
}

func (c *Client) CreateJobWorkflow(jobWorkflow *JobWorkflow) (*JobWorkflow, *http.Response, error) {
	requestUrl := "job_workflows"
	postResponse := new(JobWorkflowPostResponse)
	resp, err := c.requestWithRetry("POST", requestUrl, jobWorkflow, postResponse, defaultRetryCount)
	if err != nil {
		return nil, resp, err
	}

	jw := new(JobWorkflow)
	if err := json.Unmarshal(postResponse.Data, &jw); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	if jobWorkflow.Active != nil && jw.Active != jobWorkflow.Active {
		jw.Active = jobWorkflow.Active
		if _, _, err := c.UpdateJobWorkflow(jw); err != nil {
			return jw, resp, fmt.Errorf("failed to update active status: %v", err)
		}
	}

	return jw, resp, nil
}

func (c *Client) UpdateJobWorkflow(jobWorkflow *JobWorkflow) (*JobWorkflow, *http.Response, error) {
	requestUrl := fmt.Sprintf("job_workflows/%s", jobWorkflow.Id)
	patchResponse := new(JobWorkflowPatchResponse)
	resp, err := c.requestWithRetry("PATCH", requestUrl, jobWorkflow, patchResponse, defaultRetryCount)
	if err != nil {
		return nil, resp, err
	}

	jw := new(JobWorkflow)
	if err := json.Unmarshal(patchResponse.Data, &jw); err != nil {
		return nil, resp, errors.New("unmarshal failed")
	}

	return jw, resp, nil
}

func (c *Client) DeleteJobWorkflow(jobWorkflowId string) (*http.Response, error) {
	requestUrl := fmt.Sprintf("job_workflows/%s", jobWorkflowId)
	resp, err := c.requestWithRetry("DELETE", requestUrl, nil, nil, defaultRetryCount)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (jw *JobWorkflow) UnmarshalJSON(data []byte) error {
	rjw := RawJobWorkflowData{}
	if err := json.Unmarshal(data, &rjw); err != nil {
		return errors.New("unmarshal failed")
	}

	jw.Id = rjw.Id
	jw.Active = rjw.Attributes.Active
	jw.Name = rjw.Attributes.Name
	jw.GroupId = rjw.Attributes.GroupId
	jw.FirstJobId = rjw.Attributes.FirstJobId
	jw.FollowingJobIds = rjw.Attributes.FollowingJobIds

	return nil
}
