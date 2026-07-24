package cloudautomator

import (
	"testing"

	"terraform-provider-cloudautomator/internal/client"
)

// import 時（state が空で API レスポンスを土台にする経路）に、API 由来の
// additional_tags が破棄されず state に反映されることを確認する。API レスポンス
// の additional_tags は []interface{} 形式のため、*schema.Set への型アサーション
// だけを前提にすると破棄されてしまう欠陥を突く。
func TestBuildActionValue_ImportKeepsAdditionalTags(t *testing.T) {
	// state を空にして API レスポンス経路を通す
	d := resourceJob().TestResourceData()

	job := &client.Job{
		ActionType: "create_image",
		ActionValue: map[string]interface{}{
			"specify_image_instance": "tag",
			"tag_key":                "Name",
			"tag_value":              "example",
			"image_name":             "example-backup",
			"additional_tags": []interface{}{
				map[string]interface{}{"key": "Name", "value": "example-backup"},
			},
		},
	}

	actionValue := buildActionValue(d, job)

	tags, ok := actionValue["additional_tags"]
	if !ok {
		t.Fatal("additional_tags が破棄されている")
	}

	list, ok := tags.([]interface{})
	if !ok {
		t.Fatalf("additional_tags が []interface{} ではない: %T", tags)
	}

	if len(list) != 1 {
		t.Fatalf("additional_tags の要素数が想定と異なる: want 1, got %d", len(list))
	}
}
