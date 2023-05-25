package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetwork = `{
  "block": {
    "attributes": {
      "auto_create_subnetworks": {
        "description": "When set to 'true', the network is created in \"auto subnet mode\" and\nit will create a subnet for each region automatically across the\n'10.128.0.0/9' address range.\n\nWhen set to 'false', the network is created in \"custom subnet mode\" so\nthe user can explicitly connect subnetwork resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delete_default_routes_on_create": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "An optional description of this resource. The resource must be\nrecreated to modify this field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gateway_ipv4": {
        "computed": true,
        "description": "The gateway address for default routing out of the network. This value\nis selected by GCP.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mtu": {
        "computed": true,
        "description": "Maximum Transmission Unit in bytes. The minimum value for this field is 1460\nand the maximum value is 1500 bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
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
      "routing_mode": {
        "computed": true,
        "description": "The network-wide routing mode to use. If set to 'REGIONAL', this\nnetwork's cloud routers will only advertise routes with subnetworks\nof this network in the same region as the router. If set to 'GLOBAL',\nthis network's cloud routers will advertise routes with all\nsubnetworks of this network, across regions. Possible values: [\"REGIONAL\", \"GLOBAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
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

func GoogleComputeNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetwork), &result)
	return &result
}
