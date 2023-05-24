package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataplexLake = `{
  "block": {
    "attributes": {
      "asset_status": {
        "computed": true,
        "description": "Output only. Aggregated status of the underlying assets of the lake.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_assets": "number",
              "security_policy_applying_assets": "number",
              "update_time": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time when the lake was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. Description of the lake.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Optional. User friendly display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional. User-defined labels for the lake.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metastore_status": {
        "computed": true,
        "description": "Output only. Metastore status of the lake.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "endpoint": "string",
              "message": "string",
              "state": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The name of the lake.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time when the lake was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "metastore": {
        "block": {
          "attributes": {
            "service": {
              "description": "Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: ` + "`" + `projects/{project_id}/locations/{location_id}/services/{service_id}` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Optional. Settings to manage lake and Dataproc Metastore service instance association.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataplexLakeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataplexLake), &result)
	return &result
}
