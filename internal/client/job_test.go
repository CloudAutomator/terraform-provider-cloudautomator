package client

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestJob_Get(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job",
					"aws_account_id": 1,
					"rule_type": "cron",
					"rule_value": {
						"hour": "9",
						"minutes": "0",
						"time_zone": "Tokyo",
						"schedule_type": "one_time",
						"next_occurrence": "2100-01-01 00:00:00 UTC",
						"one_time_schedule": "2100/01/01"
					},
					"action_type": "delay",
					"action_value": {
						"delay_minutes": 60
					},
					"active": true,
					"group_id": 1,
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00",
					"effective_date": null,
					"expiration_date": null,
					"completed_post_process_id": [
						10
					],
					"failed_post_process_id": [
						20
					]
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job, _, err := c.GetJob("1000")
	if err != nil {
		t.Fatal(err)
	}

	expect := &Job{
		Id:           "1000",
		Name:         "test-job",
		Active:       true,
		GroupId:      1,
		AwsAccountId: 1,
		RuleType:     "cron",
		RuleValue: map[string]interface{}{
			"hour":              "9",
			"minutes":           "0",
			"one_time_schedule": "2100/01/01",
			"schedule_type":     "one_time",
			"time_zone":         "Tokyo",
		},
		ActionType: "delay",
		ActionValue: map[string]interface{}{
			"delay_minutes": float64(60),
		},
		AllowRuntimeActionValues: nil,
		EffectiveDate:            "",
		ExpirationDate:           "",
		CompletedPostProcessId:   []int{10},
		FailedPostProcessId:      []int{20},
	}

	testEqual(t, expect, job)
}

func TestJob_List(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")

		if len(r.URL.Query().Get("page[number]")) == 1 {
			_, _ = w.Write([]byte(`
			{
				"data": [
					{
						"id": "2000",
						"type": "trigger_jobs",
						"attributes": {
							"name": "test-job",
							"aws_account_id": 1,
							"rule_type": "cron",
							"rule_value": {
								"hour": "9",
								"minutes": "0",
								"time_zone": "Tokyo",
								"schedule_type": "one_time",
								"next_occurrence": "2100-01-01 00:00:00 UTC",
								"one_time_schedule": "2100/01/01"
							},
							"action_type": "delay",
							"action_value": {
								"delay_minutes": 60
							},
							"active": true,
							"group_id": 1,
							"created_at": "2000-01-01T00:00:00.000+09:00",
							"updated_at": "2000-01-01T00:00:00.000+09:00",
							"effective_date": null,
							"expiration_date": null,
							"completed_post_process_id": [
								30
							],
							"failed_post_process_id": [
								40
							]
						}
					}
				],
				"links": {
					"self": null,
					"first": null,
					"prev": null,
					"next": null,
					"last": null
				}
			}
			`))
		} else {
			_, _ = w.Write([]byte(fmt.Sprintf(`
			{
				"data": [
					{
						"id": "1000",
						"type": "trigger_jobs",
						"attributes": {
							"name": "test-job",
							"aws_account_id": 1,
							"rule_type": "cron",
							"rule_value": {
								"hour": "9",
								"minutes": "0",
								"time_zone": "Tokyo",
								"schedule_type": "one_time",
								"next_occurrence": "2100-01-01 00:00:00 UTC",
								"one_time_schedule": "2100/01/01"
							},
							"action_type": "delay",
							"action_value": {
								"delay_minutes": 60
							},
							"active": true,
							"group_id": 1,
							"created_at": "2000-01-01T00:00:00.000+09:00",
							"updated_at": "2000-01-01T00:00:00.000+09:00",
							"effective_date": null,
							"expiration_date": null,
							"completed_post_process_id": [
								10
							],
							"failed_post_process_id": [
								20
							]
						}
					}
				],
				"links": {
					"self": null,
					"first": null,
					"prev": null,
					"next": "%s/jobs?page[number]=1&page[size]=5",
					"last": null
				}
			}
			`, server.URL)))
		}

	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	jobs, _, err := c.GetJobs()
	if err != nil {
		t.Fatal(err)
	}

	expect := &[]Job{
		{
			Id:           "1000",
			Name:         "test-job",
			Active:       true,
			GroupId:      1,
			AwsAccountId: 1,
			RuleType:     "cron",
			RuleValue: map[string]interface{}{
				"hour":              "9",
				"minutes":           "0",
				"one_time_schedule": "2100/01/01",
				"schedule_type":     "one_time",
				"time_zone":         "Tokyo",
			},
			ActionType: "delay",
			ActionValue: map[string]interface{}{
				"delay_minutes": float64(60),
			},
			AllowRuntimeActionValues: nil,
			EffectiveDate:            "",
			ExpirationDate:           "",
			CompletedPostProcessId:   []int{10},
			FailedPostProcessId:      []int{20},
		},
		{
			Id:           "2000",
			Name:         "test-job",
			Active:       true,
			GroupId:      1,
			AwsAccountId: 1,
			RuleType:     "cron",
			RuleValue: map[string]interface{}{
				"hour":              "9",
				"minutes":           "0",
				"one_time_schedule": "2100/01/01",
				"schedule_type":     "one_time",
				"time_zone":         "Tokyo",
			},
			ActionType: "delay",
			ActionValue: map[string]interface{}{
				"delay_minutes": float64(60),
			},
			AllowRuntimeActionValues: nil,
			EffectiveDate:            "",
			ExpirationDate:           "",
			CompletedPostProcessId:   []int{30},
			FailedPostProcessId:      []int{40},
		},
	}

	testEqual(t, expect, jobs)
}

func TestJob_Create(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "POST")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job",
					"aws_account_id": 1,
					"rule_type": "cron",
					"rule_value": {
						"hour": "9",
						"minutes": "0",
						"time_zone": "Tokyo",
						"schedule_type": "one_time",
						"next_occurrence": "2100-01-01 00:00:00 UTC",
						"one_time_schedule": "2100/01/01"
					},
					"action_type": "delay",
					"action_value": {
						"delay_minutes": 60
					},
					"active": true,
					"group_id": 1,
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00",
					"effective_date": null,
					"expiration_date": null,
					"completed_post_process_id": [
						37
					],
					"failed_post_process_id": []
				}
			}
		}
		`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job := new(Job)
	createdJob, _, err := c.CreateJob(job)

	expect := &Job{
		Id:           "1000",
		Name:         "test-job",
		Active:       true,
		GroupId:      1,
		AwsAccountId: 1,
		RuleType:     "cron",
		RuleValue: map[string]interface{}{
			"hour":              "9",
			"minutes":           "0",
			"one_time_schedule": "2100/01/01",
			"schedule_type":     "one_time",
			"time_zone":         "Tokyo",
		},
		ActionType: "delay",
		ActionValue: map[string]interface{}{
			"delay_minutes": float64(60),
		},
		AllowRuntimeActionValues: nil,
		EffectiveDate:            "",
		ExpirationDate:           "",
		CompletedPostProcessId:   []int{37},
		FailedPostProcessId:      []int{},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, createdJob)
}

func TestJob_Update(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "PATCH")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job",
					"aws_account_id": 1,
					"rule_type": "cron",
					"rule_value": {
						"hour": "9",
						"minutes": "0",
						"time_zone": "Tokyo",
						"schedule_type": "one_time",
						"next_occurrence": "2100-01-01 00:00:00 UTC",
						"one_time_schedule": "2100/01/01"
					},
					"action_type": "delay",
					"action_value": {
						"delay_minutes": 60
					},
					"active": true,
					"group_id": 1,
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00",
					"effective_date": null,
					"expiration_date": null,
					"completed_post_process_id": [
						37
					],
					"failed_post_process_id": []
				}
			}
		}
		`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job := &Job{
		Id: "1000",
	}
	updatedJob, _, err := c.UpdateJob(job)

	expect := &Job{
		Id:           "1000",
		Name:         "test-job",
		Active:       true,
		GroupId:      1,
		AwsAccountId: 1,
		RuleType:     "cron",
		RuleValue: map[string]interface{}{
			"hour":              "9",
			"minutes":           "0",
			"one_time_schedule": "2100/01/01",
			"schedule_type":     "one_time",
			"time_zone":         "Tokyo",
		},
		ActionType: "delay",
		ActionValue: map[string]interface{}{
			"delay_minutes": float64(60),
		},
		AllowRuntimeActionValues: nil,
		EffectiveDate:            "",
		ExpirationDate:           "",
		CompletedPostProcessId:   []int{37},
		FailedPostProcessId:      []int{},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, updatedJob)
}

func TestReadRuleValues_Cron_Weekly(t *testing.T) {
	raw := &JobAttributes{
		RuleType: "cron",
		RuleValue: map[string]interface{}{
			"schedule_type":   "weekly",
			"weekly_schedule": `["", "monday"]`,
			"dates_to_skip":   `[]`,
			"next_occurrence": "2100-01-01 00:00:00 UTC",
		},
	}
	rv := readRuleValues(raw)
	got := rv["weekly_schedule"].([]string)
	want := []string{"monday"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("weekly_schedule: want %v, got %v", want, got)
	}
}

func TestReadRuleValues_Cron_Monthly(t *testing.T) {
	raw := &JobAttributes{
		RuleType: "cron",
		RuleValue: map[string]interface{}{
			"schedule_type":   "monthly",
			"dates_to_skip":   `["2024-01-01"]`,
			"next_occurrence": "2100-01-01 00:00:00 UTC",
		},
	}
	rv := readRuleValues(raw)
	got := rv["dates_to_skip"].([]string)
	want := []string{"2024-01-01"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dates_to_skip: want %v, got %v", want, got)
	}
}

func TestJob_Delete(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "DELETE")
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	_, err := c.DeleteJob("1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestJob_Activate(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000/activate", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "POST")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job",
					"active": true,
					"group_id": 1,
					"rule_type": "cron",
					"rule_value": {
						"hour": "9",
						"minutes": "0",
						"time_zone": "Tokyo",
						"schedule_type": "one_time",
						"one_time_schedule": "2100/01/01"
					},
					"action_type": "delay",
					"action_value": {
						"delay_minutes": 60
					},
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job, _, err := c.ActivateJob("1000")
	if err != nil {
		t.Fatal(err)
	}

	if !job.Active {
		t.Errorf("expected job.Active to be true, got false")
	}
}

func TestJob_Deactivate(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000/deactivate", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	_, err := c.DeactivateJob("1000")
	if err != nil {
		t.Fatal(err)
	}
}
