package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdsEndpoint = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation timestamp in RFC 3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_forwarding_rule": {
        "computed": true,
        "description": "URL of the endpoint's network address to which traffic is to be sent by Packet Mirroring.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_ip": {
        "computed": true,
        "description": "Internal IP address of the endpoint's network entry point.",
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
        "description": "The location for the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like \"src-net\") or the full URL to the network (like \"projects/{project_id}/global/networks/src-net\").",
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
      "severity": {
        "description": "The minimum alert severity level that is reported by the endpoint. Possible values: [\"INFORMATIONAL\", \"LOW\", \"MEDIUM\", \"HIGH\", \"CRITICAL\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "threat_exceptions": {
        "description": "Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Last update timestamp in RFC 3339 text format.",
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

func GoogleCloudIdsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdsEndpoint), &result)
	return &result
}
