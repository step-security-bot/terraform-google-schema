package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeHaVpnGateway = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
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
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035.  Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The network this VPN gateway is accepting traffic for.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region this gateway should sit in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_type": {
        "computed": true,
        "description": "The stack type for this VPN gateway to identify the IP protocols that are enabled.\nIf not specified, IPV4_ONLY will be used. Default value: \"IPV4_ONLY\" Possible values: [\"IPV4_ONLY\", \"IPV4_IPV6\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_interfaces": {
        "computed": true,
        "description": "A list of interfaces on this VPN gateway.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "number",
              "interconnect_attachment": "string",
              "ip_address": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeHaVpnGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeHaVpnGateway), &result)
	return &result
}
