package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkEndpointGroup = `{
  "block": {
    "attributes": {
      "default_port": {
        "computed": true,
        "description": "The default port used if the port number is not specified in the\nnetwork endpoint.",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The network to which all network endpoints in the NEG belong.\nUses \"default\" project network if unspecified.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_endpoint_type": {
        "computed": true,
        "description": "Type of network endpoints in this network endpoint group. Currently\nthe only supported value is GCE_VM_IP_PORT.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_link": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "Number of network endpoints in the network endpoint group.",
        "description_kind": "plain",
        "type": "number"
      },
      "subnetwork": {
        "computed": true,
        "description": "Optional subnetwork to which all network endpoints in the NEG belong.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "Zone where the network endpoint group is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeNetworkEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkEndpointGroup), &result)
	return &result
}
