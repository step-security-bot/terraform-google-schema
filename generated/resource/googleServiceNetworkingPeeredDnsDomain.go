package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceNetworkingPeeredDnsDomain = `{
  "block": {
    "attributes": {
      "dns_suffix": {
        "description": "The DNS domain name suffix of the peered DNS domain.",
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
      "name": {
        "description": "Name of the peered DNS domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "Network in the consumer project to peer with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "computed": true,
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
      "service": {
        "description": "The name of the service to create a peered DNS domain for, e.g. servicenetworking.googleapis.com",
        "description_kind": "plain",
        "optional": true,
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
            "read": {
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

func GoogleServiceNetworkingPeeredDnsDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceNetworkingPeeredDnsDomain), &result)
	return &result
}
