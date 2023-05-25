package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleActiveDirectoryDomainTrust = `{
  "block": {
    "attributes": {
      "domain": {
        "description": "The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions, \nhttps://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "selective_authentication": {
        "description": "Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "target_dns_ip_addresses": {
        "description": "The target DNS server IP addresses which can resolve the remote domain involved in the trust.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_domain_name": {
        "description": "The fully qualified target domain name which will be in trust with the current domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trust_direction": {
        "description": "The trust direction, which decides if the current domain is trusted, trusting, or both. Possible values: [\"INBOUND\", \"OUTBOUND\", \"BIDIRECTIONAL\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trust_handshake_secret": {
        "description": "The trust secret used for the handshake with the target domain. This will not be stored.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "trust_type": {
        "description": "The type of trust represented by the trust resource. Possible values: [\"FOREST\", \"EXTERNAL\"]",
        "description_kind": "plain",
        "required": true,
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

func GoogleActiveDirectoryDomainTrustSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleActiveDirectoryDomainTrust), &result)
	return &result
}
