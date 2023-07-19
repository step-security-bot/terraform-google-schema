package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkSecurityGatewaySecurityPolicyRule = `{
  "block": {
    "attributes": {
      "application_matcher": {
        "description": "CEL expression for matching on L7/application level criteria.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "basic_profile": {
        "description": "Profile which tells what the primitive action should be. Possible values are: * ALLOW * DENY. Possible values: [\"BASIC_PROFILE_UNSPECIFIED\", \"ALLOW\", \"DENY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp when the resource was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Free-text description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "description": "Whether the rule is enforced.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "gateway_security_policy": {
        "description": "The name of the gatewat security policy this rule belongs to.",
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
      "location": {
        "description": "The location of the gateway security policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}\nrule should match the pattern: (^a-z?$).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description": "Priority of the rule. Lower number corresponds to higher precedence.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_matcher": {
        "description": "CEL expression for matching on session criteria.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tls_inspection_enabled": {
        "description": "Flag to enable TLS inspection of traffic matching on. Can only be true if the\nparent GatewaySecurityPolicy references a TLSInspectionConfig.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp when the resource was updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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

func GoogleNetworkSecurityGatewaySecurityPolicyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkSecurityGatewaySecurityPolicyRule), &result)
	return &result
}
