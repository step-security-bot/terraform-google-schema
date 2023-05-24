package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkConnectivityHub = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time the hub was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of the hub.",
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
        "description": "Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Immutable. The name of the hub. Hub names must be unique. They use the following form: ` + "`" + `projects/{project_number}/locations/global/hubs/{hub_id}` + "`" + `",
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
      "routing_vpcs": {
        "computed": true,
        "description": "The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "uri": "string"
            }
          ]
        ]
      },
      "state": {
        "computed": true,
        "description": "Output only. The current lifecycle state of this hub. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING",
        "description_kind": "plain",
        "type": "string"
      },
      "unique_id": {
        "computed": true,
        "description": "Output only. The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time the hub was last updated.",
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
  "version": 0
}`

func GoogleNetworkConnectivityHubSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkConnectivityHub), &result)
	return &result
}
