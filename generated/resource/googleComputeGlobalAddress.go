package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeGlobalAddress = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description": "The IP address or beginning of the address range represented by this\nresource. This can be supplied as an input to reserve a specific\naddress or omitted to allow GCP to choose a valid one for you.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_type": {
        "description": "The type of the address to reserve.\n\n* EXTERNAL indicates public/external single IP address.\n* INTERNAL indicates internal IP ranges belonging to some network. Default value: \"EXTERNAL\" Possible values: [\"EXTERNAL\", \"INTERNAL\"]",
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
      "ip_version": {
        "description": "The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: [\"IPV4\", \"IPV6\"]",
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
        "description": "The URL of the network in which to reserve the IP range. The IP range\nmust be in RFC1918 space. The network cannot be deleted if there are\nany reserved IP ranges referring to it.\n\nThis should only be set when using an Internal address.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prefix_length": {
        "description": "The prefix length of the IP range. If not present, it means the\naddress field is a single IP address.\n\nThis field is not applicable to addresses with addressType=EXTERNAL,\nor addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purpose": {
        "description": "The purpose of the resource. Possible values include:\n\n* VPC_PEERING - for peer networks\n\n* PRIVATE_SERVICE_CONNECT - for ([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) Private Service Connect networks",
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

func GoogleComputeGlobalAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeGlobalAddress), &result)
	return &result
}
