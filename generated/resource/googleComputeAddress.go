package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeAddress = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description": "The static external IP address represented by this resource. Only\nIPv4 is supported. An address may only be specified for INTERNAL\naddress types. The IP address must be inside the specified subnetwork,\nif any.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_type": {
        "description": "The type of address to reserve, either INTERNAL or EXTERNAL.\nIf unspecified, defaults to EXTERNAL.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
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
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'\nwhich means the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "The networking tier used for configuring this address. This field can\ntake the following values: PREMIUM or STANDARD. If this field is not\nspecified, it is assumed to be PREMIUM.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purpose": {
        "computed": true,
        "description": "The purpose of this resource, which can be one of the following values:\n\n- GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.\n\nThis should only be set when using an Internal address.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created address should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "The URL of the subnetwork in which to reserve the address. If an IP\naddress is specified, it must be within the subnetwork's IP range.\nThis field can only be used with INTERNAL type with\nGCE_ENDPOINT/DNS_RESOLVER purposes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "The URLs of the resources that are using this address.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func GoogleComputeAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeAddress), &result)
	return &result
}
