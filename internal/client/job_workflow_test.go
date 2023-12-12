package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestJobWorkflow_Get(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/job_workflows/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "job_workflows",
				"attributes": {
					"name": "test-job-workflow",
					"active": true,
					"group_id": 1,
					"first_job_id": 1,
					"following_job_ids": [
						2,
						3
					],
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job_workflow, _, err := c.GetJobWorkflow("1000")
	if err != nil {
		t.Fatal(err)
	}

	trueVal := true
	expect := &JobWorkflow{
		Id:              "1000",
		Active:          &trueVal,
		Name:            "test-job-workflow",
		GroupId:         1,
		FirstJobId:      1,
		FollowingJobIds: []int{2, 3},
	}

	testEqual(t, expect, job_workflow)
}

func TestJobWorkflow_List(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/job_workflows", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")

		if len(r.URL.Query().Get("page[number]")) == 1 {
			_, _ = w.Write([]byte(`
			{
				"data": [
					{
						"id": "2000",
						"type": "job_workflows",
						"attributes": {
							"name": "test-job-workflow-2",
							"active": true,
							"group_id": 1,
							"first_job_id": 4,
							"following_job_ids": [
								5,
								6
							],
							"created_at": "2000-01-01T00:00:00.000+09:00",
							"updated_at": "2000-01-01T00:00:00.000+09:00"
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
						"type": "job_workflows",
						"attributes": {
							"name": "test-job-workflow-1",
							"active": true,
							"group_id": 1,
							"first_job_id": 1,
							"following_job_ids": [
								2,
								3
							],
							"created_at": "2000-01-01T00:00:00.000+09:00",
							"updated_at": "2000-01-01T00:00:00.000+09:00"
						}
					}
				],
				"links": {
					"self": null,
					"first": null,
					"prev": null,
					"next": "%s/job_workflows?page[number]=1&page[size]=5",
					"last": null
				}
			}
			`, server.URL)))
		}

	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	jobWorkflows, _, err := c.GetJobWorkflows()
	if err != nil {
		t.Fatal(err)
	}

	trueVal := true
	expect := &[]JobWorkflow{
		{
			Id:              "1000",
			Active:          &trueVal,
			Name:            "test-job-workflow-1",
			GroupId:         1,
			FirstJobId:      1,
			FollowingJobIds: []int{2, 3},
		},
		{
			Id:              "2000",
			Active:          &trueVal,
			Name:            "test-job-workflow-2",
			GroupId:         1,
			FirstJobId:      4,
			FollowingJobIds: []int{5, 6},
		},
	}

	testEqual(t, expect, jobWorkflows)
}

func TestJobWorkflow_Create(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/job_workflows", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "POST")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job-workflow-1",
					"active": true,
					"group_id": 1,
					"first_job_id": 1,
					"following_job_ids": [
						2,
						3
					],
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}
		`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job_workflow := new(JobWorkflow)
	createdJobWorkflow, _, err := c.CreateJobWorkflow(job_workflow)

	trueVal := true
	expect := &JobWorkflow{
		Id:              "1000",
		Active:          &trueVal,
		Name:            "test-job-workflow-1",
		GroupId:         1,
		FirstJobId:      1,
		FollowingJobIds: []int{2, 3},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, createdJobWorkflow)
}

func TestJobWorkflow_Update(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/job_workflows/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "PATCH")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"name": "test-job-workflow-1",
					"active": true,
					"group_id": 1,
					"first_job_id": 1,
					"following_job_ids": [
						2,
						3
					],
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}
		`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	job_workflow := &JobWorkflow{
		Id: "1000",
	}
	updatedJobWorkflow, _, err := c.UpdateJobWorkflow(job_workflow)

	trueVal := true
	expect := &JobWorkflow{
		Id:              "1000",
		Active:          &trueVal,
		Name:            "test-job-workflow-1",
		GroupId:         1,
		FirstJobId:      1,
		FollowingJobIds: []int{2, 3},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, updatedJobWorkflow)
}

func TestJobWorkflow_Delete(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/job_workflows/1", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "DELETE")
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	_, err := c.DeleteJobWorkflow("1")
	if err != nil {
		t.Fatal(err)
	}
}
