package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeInstance = `{
  "block": {
    "attributes": {
      "description": {
        "description": "Description of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_encryption_key_name": {
        "description": "Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.\nUse the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host": {
        "computed": true,
        "description": "Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.",
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
        "description": "Compute Engine location where the instance resides. For trial organization\nsubscriptions, the location must be a Compute Engine zone. For paid organization\nsubscriptions, it should correspond to a Compute Engine region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Resource ID of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering_cidr_range": {
        "description": "The size of the CIDR block range that will be reserved by the instance. Possible values: [\"SLASH_16\", \"SLASH_20\", \"SLASH_22\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "Output only. Port number of the exposed Apigee endpoint.",
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

func GoogleApigeeInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeInstance), &result)
	return &result
}
