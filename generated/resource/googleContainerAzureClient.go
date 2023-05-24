package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAzureClient = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The Azure Active Directory Application ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description": "Output only. The PEM encoded x509 certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which this resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of this resource.",
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
      "tenant_id": {
        "description": "The Azure Active Directory Tenant ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A globally unique identifier for the client.",
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

func GoogleContainerAzureClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAzureClient), &result)
	return &result
}
