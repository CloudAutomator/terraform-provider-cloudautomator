package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAwsAccount_Get(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/groups/1000/aws_accounts/1000", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "aws_accounts",
				"attributes": {
					"name": "example-aws-account"
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	postProcess, _, err := c.GetAwsAccount("1000", "1000")
	if err != nil {
		t.Fatal(err)
	}

	expect := &AwsAccount{
		Id:   "1000",
		Name: "example-aws-account",
	}

	testEqual(t, expect, postProcess)
}

func TestAwsAccount_List(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/groups/1000/aws_accounts", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "GET")

		if len(r.URL.Query().Get("page[number]")) == 1 {
			_, _ = w.Write([]byte(`
			{
				"data": [
					{
						"id": "2000",
						"type": "aws_accounts",
						"attributes": {
							"name": "example-aws-account-2"
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
						"type": "aws_accounts",
						"attributes": {
							"name": "example-aws-account-1"
						}
					}
				],
				"links": {
					"self": null,
					"first": null,
					"prev": null,
					"next": "%s/groups/1000/aws_accounts?page[number]=1&page[size]=5",
					"last": null
				}
			}
			`, server.URL)))
		}

	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	postProcesses, _, err := c.GetAwsAccounts("1000")
	if err != nil {
		t.Fatal(err)
	}

	expect := &[]AwsAccount{
		{
			Id:   "1000",
			Name: "example-aws-account-1",
		},
		{
			Id:   "2000",
			Name: "example-aws-account-2",
		},
	}

	testEqual(t, expect, postProcesses)
}
