package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleWorkflowsWorkflow = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp of when the workflow was created in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key_name": {
        "description": "The KMS key used to encrypt workflow and execution data.\n\nFormat: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the workflow provided by the user. Must be at most 1000 unicode characters long.",
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
        "description": "A set of key/value label pairs to assign to this Workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Name of the Workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region of the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The revision of the workflow. A new one is generated if the service account or source contents is changed.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "Name of the service account associated with the latest workflow version. This service\naccount represents the identity of the workflow and determines what permissions the workflow has.\nFormat: projects/{project}/serviceAccounts/{account} or {account}.\nUsing - as a wildcard for the {project} or not providing one at all will infer the project from the account.\nThe {account} value can be the email address or the unique_id of the service account.\nIf not provided, workflow will use the project's default service account.\nModifying this field for an existing workflow results in a new workflow revision.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_contents": {
        "description": "Workflow code to be executed. The size limit is 32KB.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the workflow deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp of when the workflow was last updated in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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
  "version": 1
}`

func GoogleWorkflowsWorkflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleWorkflowsWorkflow), &result)
	return &result
}
