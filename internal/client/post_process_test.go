package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPostProcess_Get(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/post_processes/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "post_processes",
				"attributes": {
					"name": "test-post-process",
					"group_id": 1,
					"service": "email",
					"shared_by_group": false,
					"parameters": {
						"recipients": [
							"test@example.com"
						]
					},
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}`))
	})

	token := "token"
	c, _ := New(&token, WithAPIEndpoint(server.URL))

	postProcess, _, err := c.GetPostProcess("1000")
	if err != nil {
		t.Fatal(err)
	}

	sharedByGroup := false
	expect := &PostProcess{
		Id:            "1000",
		Name:          "test-post-process",
		Service:       "email",
		SharedByGroup: &sharedByGroup,
		GroupId:       1,
		Parameters: map[string]interface{}{
			"email_recipient": "test@example.com",
		},
	}

	testEqual(t, expect, postProcess)
}

func TestPostProcess_List(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/post_processes", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")

		if len(r.URL.Query().Get("page[number]")) == 1 {
			_, _ = w.Write([]byte(`
			{
				"data": [
					{
						"id": "2000",
						"type": "post_processes",
						"attributes": {
							"name": "test-post-process",
							"group_id": 1,
							"service": "email",
							"shared_by_group": false,
							"parameters": {
								"recipients": [
									"test@example.com"
								]
							},
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
						"type": "post_processes",
						"attributes": {
							"name": "test-post-process",
							"group_id": 1,
							"service": "email",
							"shared_by_group": true,
							"parameters": {
								"recipients": [
									"test@example.com"
								]
							},
							"created_at": "2000-01-01T00:00:00.000+09:00",
							"updated_at": "2000-01-01T00:00:00.000+09:00"
						}
					}
				],
				"links": {
					"self": null,
					"first": null,
					"prev": null,
					"next": "%s/post_processes?page[number]=1&page[size]=5",
					"last": null
				}
			}
			`, server.URL)))
		}

	})

	token := "token"
	c, _ := New(&token, WithAPIEndpoint(server.URL))

	postProcesses, _, err := c.GetPostProcesses()
	if err != nil {
		t.Fatal(err)
	}

	sharedByGroupTrue := true
	sharedByGroupFalse := false
	expect := &[]PostProcess{
		{
			Id:            "1000",
			Name:          "test-post-process",
			Service:       "email",
			SharedByGroup: &sharedByGroupTrue,
			GroupId:       1,
			Parameters: map[string]interface{}{
				"email_recipient": "test@example.com",
			},
		},
		{
			Id:            "2000",
			Name:          "test-post-process",
			Service:       "email",
			SharedByGroup: &sharedByGroupFalse,
			GroupId:       1,
			Parameters: map[string]interface{}{
				"email_recipient": "test@example.com",
			},
		},
	}

	testEqual(t, expect, postProcesses)
}

func TestPostProcess_Create(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/post_processes", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "POST")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "post_processes",
				"attributes": {
					"name": "test-post-process",
					"group_id": 1,
					"service": "email",
					"shared_by_group": false,
					"parameters": {
						"recipients": [
							"test@example.com"
						]
					},
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}
		`))
	})

	token := "token"
	c, _ := New(&token, WithAPIEndpoint(server.URL))

	postProcess := new(PostProcess)
	createdPostProcess, _, err := c.CreatePostProcess(postProcess)

	sharedByGroup := false
	expect := &PostProcess{
		Id:            "1000",
		Name:          "test-post-process",
		Service:       "email",
		SharedByGroup: &sharedByGroup,
		GroupId:       1,
		Parameters: map[string]interface{}{
			"email_recipient": "test@example.com",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, createdPostProcess)
}

func TestPostProcess_Update(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/post_processes/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "PATCH")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "post_processes",
				"attributes": {
					"name": "test-post-process",
					"group_id": 1,
					"service": "email",
					"shared_by_group": false,
					"parameters": {
						"recipients": [
							"test@example.com"
						]
					},
					"created_at": "2000-01-01T00:00:00.000+09:00",
					"updated_at": "2000-01-01T00:00:00.000+09:00"
				}
			}
		}
		`))
	})

	token := "token"
	c, _ := New(&token, WithAPIEndpoint(server.URL))

	postProcess := &PostProcess{
		Id:   "1000",
		Name: "updated-post-process",
	}
	updatedPostProcess, _, err := c.UpdatePostProcess(postProcess)

	sharedByGroupFalse := false
	expect := &PostProcess{
		Id:            "1000",
		Name:          "test-post-process",
		Service:       "email",
		SharedByGroup: &sharedByGroupFalse,
		GroupId:       1,
		Parameters: map[string]interface{}{
			"email_recipient": "test@example.com",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, updatedPostProcess)
}

func TestPostProcess_Delete(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/post_processes/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "DELETE")
	})

	token := "token"
	c, _ := New(&token, WithAPIEndpoint(server.URL))

	_, err := c.DeleteJob("1000")
	if err != nil {
		t.Fatal(err)
	}
}
