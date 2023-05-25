package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataflowJob = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_workers": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_delete": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_email": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "temp_gcs_location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_gcs_path": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataflowJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataflowJob), &result)
	return &result
}
