package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccount = `{
  "block": {
    "attributes": {
      "account_id": {
        "description": "The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035. Changing this forces a new service account to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "A text description of the service account. Must be less than or equal to 256 UTF-8 bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disabled": {
        "description": "Whether the service account is disabled. Defaults to false",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "description": "The display name for the service account. Can be updated without creating a new resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "email": {
        "computed": true,
        "description": "The e-mail address of the service account. This value should be referenced from any google_iam_policy data sources that would grant the service account privileges.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The fully-qualified name of the service account.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project that the service account will be created in. Defaults to the provider project configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "unique_id": {
        "computed": true,
        "description": "The unique id of the service account.",
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

func GoogleServiceAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccount), &result)
	return &result
}
