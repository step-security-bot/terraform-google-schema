package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIapClient = `{
  "block": {
    "attributes": {
      "brand": {
        "description": "Identifier of the brand to which this client\nis attached to. The format is\n'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_id": {
        "computed": true,
        "description": "Output only. Unique identifier of the OAuth client.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Human-friendly name given to the OAuth client.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret": {
        "computed": true,
        "description": "Output only. Client secret of the OAuth client.",
        "description_kind": "plain",
        "sensitive": true,
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

func GoogleIapClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIapClient), &result)
	return &result
}
