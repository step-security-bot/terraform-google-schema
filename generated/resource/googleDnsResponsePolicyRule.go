package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsResponsePolicyRule = `{
  "block": {
    "attributes": {
      "dns_name": {
        "description": "The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.",
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
      "response_policy": {
        "description": "Identifies the response policy addressed by this request.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
        "description": "An identifier for this rule. Must be unique with the ResponsePolicy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "local_data": {
        "block": {
          "block_types": {
            "local_datas": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "For example, www.example.com.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rrdatas": {
                    "description": "As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ttl": {
                    "description": "Number of seconds that this ResourceRecordSet can be cached by\nresolvers.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "description": "One of valid DNS resource types. Possible values: [\"A\", \"AAAA\", \"CAA\", \"CNAME\", \"DNSKEY\", \"DS\", \"HTTPS\", \"IPSECVPNKEY\", \"MX\", \"NAPTR\", \"NS\", \"PTR\", \"SOA\", \"SPF\", \"SRV\", \"SSHFP\", \"SVCB\", \"TLSA\", \"TXT\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "All resource record sets for this selector, one per resource record type. The name must match the dns_name.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;\nin particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
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

func GoogleDnsResponsePolicyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsResponsePolicyRule), &result)
	return &result
}
