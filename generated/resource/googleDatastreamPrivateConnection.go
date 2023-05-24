package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatastreamPrivateConnection = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Display name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "error": {
        "computed": true,
        "description": "The PrivateConnection error in case of failure.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "details": [
                "map",
                "string"
              ],
              "message": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location this private connection is located in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource's name.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_connection_id": {
        "description": "The private connectivity identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the PrivateConnection.",
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
      },
      "vpc_peering_config": {
        "block": {
          "attributes": {
            "subnet": {
              "description": "A free subnet for peering. (CIDR of /29)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc": {
              "description": "Fully qualified name of the VPC that Datastream will peer to.\nFormat: projects/{project}/global/{networks}/{name}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The VPC Peering configuration is used to create VPC peering\nbetween Datastream and the consumer's VPC.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDatastreamPrivateConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatastreamPrivateConnection), &result)
	return &result
}
