package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleProject = `{
  "block": {
    "attributes": {
      "app_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auth_domain": "string",
              "code_bucket": "string",
              "default_bucket": "string",
              "default_hostname": "string",
              "feature_settings": [
                "list",
                [
                  "object",
                  {
                    "split_health_checks": "bool"
                  }
                ]
              ],
              "gcr_domain": "string",
              "location_id": "string",
              "name": "string",
              "serving_status": "string",
              "url_dispatch_rule": [
                "list",
                [
                  "object",
                  {
                    "domain": "string",
                    "path": "string",
                    "service": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "auto_create_network": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "billing_account": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_delete": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleProject), &result)
	return &result
}
