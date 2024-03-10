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
	ForWorkflow              *bool                  `json:"for_workflow,omitempty"`
	AwsAccountId             int                    `json:"aws_account_id,omitempty"`
	AwsAccountIds            []int                  `json:"aws_account_ids,omitempty"`
	GoogleCloudAccountId     int                    `json:"google_cloud_account_id,omitempty"`
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

type JobGetResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobPostResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobPatchResponse struct {
	Data json.RawMessage `json:"data"`
}

type JobListResponse struct {
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
	AwsAccountId             int                    `json:"aws_account_id,omitempty"`
	AwsAccountIds            []int                  `json:"aws_account_ids,omitempty"`
	GoogleCloudAccountId     int                    `json:"google_cloud_account_id,omitempty"`
	ForWorkflow              *bool                  `json:"for_workflow,omitempty"`
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
	"attach_user_policy",
	"detach_user_policy",
	"authorize_security_group_ingress",
	"copy_rds_cluster_snapshot",
	"create_fsx_backup",
	"stop_ecs_tasks",
	"change_instance_type",
	"deregister_instances",
	"deregister_target_instances",
	"describe_metadata",
	"dynamodb_start_backup_job",
	"google_compute_insert_machine_image",
	"google_compute_stop_vm_instances",
	"google_compute_start_vm_instances",
	"invoke_lambda_function",
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

	getResponse := new(JobGetResponse)
	resp, err := c.requestWithRetry("GET", requestUrl, nil, getResponse, 5)
	if err != nil {
		return nil, resp, err
	}

	job := new(Job)
	if err := json.Unmarshal(getResponse.Data, &job); err != nil {
		return nil, nil, errors.New("unmarshal failed")
	}

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

		listResponse := new(JobListResponse)
		resp, err := c.requestWithRetry("GET", rel.String(), nil, listResponse, defaultRetryCount)

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
	postResponse := new(JobPostResponse)
	resp, err := c.requestWithRetry("POST", requestUrl, job, postResponse, defaultRetryCount)
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
	patchResponse := new(JobPatchResponse)
	resp, err := c.requestWithRetry("PATCH", requestUrl, job, patchResponse, defaultRetryCount)
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
	resp, err := c.requestWithRetry("DELETE", requestUrl, nil, nil, defaultRetryCount)
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
		delete(ruleValue, "schedule_trigger_setting_id")
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
	case "create_ebs_snapshot", "create_image":
		switch rawJob.ActionValue["generation"].(type) {
		case string:
			generation, _ := strconv.Atoi(rawJob.ActionValue["generation"].(string))
			rawJob.ActionValue["generation"] = generation
		}
	case "delay":
		switch rawJob.ActionValue["delay_minutes"].(type) {
		case string:
			delay_minutes, _ := strconv.Atoi(rawJob.ActionValue["delay_minutes"].(string))
			rawJob.ActionValue["delay_minutes"] = delay_minutes
		}
	case "authorize_security_group_ingress", "revoke_security_group_ingress":
		toPort := rawJob.ActionValue["to_port"].(float64)
		rawJob.ActionValue["to_port"] = strconv.Itoa(int(toPort))
	case "rebuild_workspaces":
		delete(rawJob.ActionValue, "specify_workspace")
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
	j.ForWorkflow = rj.Attributes.ForWorkflow
	j.AwsAccountId = rj.Attributes.AwsAccountId
	j.AwsAccountIds = rj.Attributes.AwsAccountIds
	j.GoogleCloudAccountId = rj.Attributes.GoogleCloudAccountId
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
