package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"terraform-provider-cloudautomator/internal/utils"
)

type Job struct {
	Id                       string                 `json:"id,omitempty"`
	Name                     string                 `json:"name"`
	GroupId                  int                    `json:"group_id"`
	AwsAccountId             int                    `json:"aws_account_id"`
	RuleType                 string                 `json:"rule_type"`
	RuleValue                map[string]interface{} `json:"rule_value"`
	ActionType               string                 `json:"action_type"`
	ActionValue              map[string]interface{} `json:"action_value"`
	AllowRuntimeActionValues *bool                  `json:"allow_runtime_action_values,omitempty"`
	EffectiveDate            string                 `json:"effective_date,omitempty"`
	ExpirationDate           string                 `json:"expiration_date,omitempty"`
	CompletedPostProcessId   []int                  `json:"completed_post_process_id,omitempty"`
	FailedPostProcessId      []int                  `json:"failed_post_process_id,omitempty"`
}

type GetJobResponse struct {
	Data json.RawMessage `json:"data"`
}

type PostJobResponse struct {
	Data json.RawMessage `json:"data"`
}

type PatchJobResponse struct {
	Data json.RawMessage `json:"data"`
}

type ListJobResponse struct {
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

type RawJobData struct {
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Attributes JobAttributes `json:"attributes"`
}

type JobAttributes struct {
	Name                     string                 `json:"name"`
	Active                   bool                   `json:"active"`
	GroupID                  int                    `json:"group_id"`
	AwsAccountId             int                    `json:"aws_account_id"`
	RuleType                 string                 `json:"rule_type"`
	RuleValue                map[string]interface{} `json:"rule_value"`
	ActionType               string                 `json:"action_type"`
	ActionValue              map[string]interface{} `json:"action_value"`
	AllowRuntimeActionValues *bool                  `json:"allow_runtime_action_values"`
	EffectiveDate            string                 `json:"effective_date"`
	ExpirationDate           string                 `json:"expiration_date"`
	CompletedPostProcessId   []int                  `json:"completed_post_process_id"`
	FailedPostProcessId      []int                  `json:"failed_post_process_id"`
	CreatedAt                time.Time              `json:"created_at"`
	UpdatedAt                time.Time              `json:"updated_at"`
}

var TRACE_STATUS_NOT_SUPPORTED_ACTION_TYPES = []string{
	"authorize_security_group_ingress",
	"change_instance_type",
	"deregister_instances",
	"deregister_target_instances",
	"reboot_rds_instances",
	"register_instances",
	"register_target_instances",
	"restore_from_cluster_snapshot",
	"restore_rds_cluster",
	"revoke_security_group_ingress",
	"start_workspaces",
	"update_record_set",
	"windows_update",
}

func (c *Client) GetJob(jobId string) (*Job, *http.Response, error) {
	requestUrl := fmt.Sprintf("jobs/%s", jobId)
	req, err := c.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	getResponse := new(GetJobResponse)
	resp, err := c.Do(req, &getResponse)
	if err != nil {
		return nil, resp, err
	}

	job := new(Job)
	json.Unmarshal(getResponse.Data, &job)

	return job, resp, nil
}

func (c *Client) GetJobs() (*[]Job, *http.Response, error) {
	jobs := []Job{}
	requestUrl := "jobs"

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

		listResponse := new(ListJobResponse)
		resp, err := c.Do(req, listResponse)
		if err != nil {
			return nil, resp, err
		}

		for _, r := range listResponse.Data {
			job := new(Job)
			if err := json.Unmarshal(r, &job); err != nil {
				return nil, nil, errors.New("unmarshal failed")
			}
			jobs = append(jobs, *job)
		}

		requestUrl = listResponse.Links.Next
	}

	return &jobs, nil, nil
}

func (c *Client) CreateJob(job *Job) (*Job, *http.Response, error) {
	requestUrl := "jobs"
	req, err := c.NewRequest("POST", requestUrl, job)
	if err != nil {
		return nil, nil, err
	}

	postResponse := new(PostJobResponse)
	resp, err := c.Do(req, &postResponse)
	if err != nil {
		return nil, resp, err
	}

	j := new(Job)
	if err := json.Unmarshal(postResponse.Data, &j); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

	return j, resp, nil
}

func (c *Client) UpdateJob(job *Job) (*Job, *http.Response, error) {
	requestUrl := fmt.Sprintf("jobs/%s", job.Id)
	req, err := c.NewRequest("PATCH", requestUrl, job)
	if err != nil {
		return nil, nil, err
	}

	patchResponse := new(PatchJobResponse)
	resp, err := c.Do(req, &patchResponse)
	if err != nil {
		return nil, resp, err
	}

	j := new(Job)
	if err := json.Unmarshal(patchResponse.Data, &j); err != nil {
		return nil, resp, errors.New("unmarshal failed")
	}

	return j, resp, nil
}

func (c *Client) DeleteJob(jobId string) (*http.Response, error) {
	requestUrl := fmt.Sprintf("jobs/%s", jobId)
	req, err := c.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func readRuleValues(rawJob *JobAttributes) map[string]interface{} {
	ruleValue := rawJob.RuleValue

	switch rawJob.RuleType {
	case "cron":
		switch ruleValue["schedule_type"] {
		case "weekly":
			ruleValue["weekly_schedule"] = utils.StringToSlice(ruleValue["weekly_schedule"].(string))
			ruleValue["dates_to_skip"] = utils.StringToSlice(ruleValue["dates_to_skip"].(string))
		case "monthly":
			ruleValue["dates_to_skip"] = utils.StringToSlice(ruleValue["dates_to_skip"].(string))
		case "monthly_day_of_week":
			ruleValue["dates_to_skip"] = utils.StringToSlice(ruleValue["dates_to_skip"].(string))
			ruleValue["monthly_day_of_week_schedule"] = []interface{}{ruleValue["monthly_schedule"]}

			delete(ruleValue, "monthly_schedule")
		}

		delete(ruleValue, "next_occurrence")
	case "schedule":
		delete(ruleValue, "next_schedule")
	case "sqs_v2":
		awsAccountId, _ := strconv.Atoi(ruleValue["aws_account_id"].(string))
		ruleValue["sqs_aws_account_id"] = awsAccountId
		ruleValue["sqs_region"] = ruleValue["region"]

		delete(ruleValue, "region")
		delete(ruleValue, "aws_account_id")
		delete(ruleValue, "sqs_aws_account_number")
		delete(ruleValue, "time_zone")
	}

	return ruleValue
}

func readActionValues(rawJob *JobAttributes) map[string]interface{} {
	switch rawJob.ActionType {
	case "authorize_security_group_ingress", "revoke_security_group_ingress":
		toPort := rawJob.ActionValue["to_port"].(float64)
		rawJob.ActionValue["to_port"] = strconv.Itoa(int(toPort))
	}

	deleteTraceStatus(rawJob)

	return rawJob.ActionValue
}

// リソースの終了ステータスチェックに対応していないアクションでも、
// レスポンスのJSONに trace_status が含まれる場合があるため
// Unmarshal のタイミングで削除します。
func deleteTraceStatus(rawJob *JobAttributes) {
	if utils.Contains(TRACE_STATUS_NOT_SUPPORTED_ACTION_TYPES, rawJob.ActionType) {
		delete(rawJob.ActionValue, "trace_status")
	}
}

func (j *Job) UnmarshalJSON(data []byte) error {
	rj := RawJobData{}
	if err := json.Unmarshal(data, &rj); err != nil {
		return errors.New("unmarshal failed")
	}

	j.Id = rj.Id
	j.Name = rj.Attributes.Name
	j.GroupId = rj.Attributes.GroupID
	j.AwsAccountId = rj.Attributes.AwsAccountId
	j.RuleType = rj.Attributes.RuleType
	j.RuleValue = readRuleValues(&rj.Attributes)
	j.ActionType = rj.Attributes.ActionType
	j.ActionValue = readActionValues(&rj.Attributes)
	j.AllowRuntimeActionValues = rj.Attributes.AllowRuntimeActionValues
	j.EffectiveDate = rj.Attributes.EffectiveDate
	j.ExpirationDate = rj.Attributes.ExpirationDate
	j.CompletedPostProcessId = rj.Attributes.CompletedPostProcessId
	j.FailedPostProcessId = rj.Attributes.FailedPostProcessId

	return nil
}
