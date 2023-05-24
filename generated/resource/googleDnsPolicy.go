package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A textual description field. Defaults to 'Managed by Terraform'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_inbound_forwarding": {
        "description": "Allows networks bound to this policy to receive DNS queries sent\nby VMs or applications over VPN connections. When enabled, a\nvirtual IP address will be allocated from each of the sub-networks\nthat are bound to this policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_logging": {
        "description": "Controls whether logging is enabled for the networks bound to this policy.\nDefaults to no logging if not set.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "User assigned name for this policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "alternative_name_server_config": {
        "block": {
          "block_types": {
            "target_name_servers": {
              "block": {
                "attributes": {
                  "forwarding_path": {
                    "description": "Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding\ndecision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go\nto the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: [\"default\", \"private\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ipv4_address": {
                    "description": "IPv4 address to forward to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Sets an alternative name server for the associated networks. When specified,\nall DNS queries are forwarded to a name server that you choose. Names such as .internal\nare not available when an alternative name server is specified.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "Sets an alternative name server for the associated networks.\nWhen specified, all DNS queries are forwarded to a name server that you choose.\nNames such as .internal are not available when an alternative name server is specified.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "networks": {
        "block": {
          "attributes": {
            "network_url": {
              "description": "The id or fully qualified URL of the VPC network to forward queries to.\nThis should be formatted like 'projects/{project}/global/networks/{network}' or\n'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of network names specifying networks to which this policy is applied.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
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

func GoogleDnsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsPolicy), &result)
	return &result
}
