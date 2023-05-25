package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSecurityPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this security policy. Max size is 2048.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description": "Fingerprint of this resource.",
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
        "description": "The name of the security policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "action": {
              "description": "Action to take when match matches the request. Valid values:   \"allow\" : allow access to target, \"deny(status)\" : deny access to target, returns the HTTP response code specified (valid values are 403, 404 and 502)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "description": "An optional description of this rule. Max size is 64.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preview": {
              "computed": true,
              "description": "When set to true, the action specified above is not enforced. Stackdriver logs for requests that trigger a preview action are annotated as such.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "priority": {
              "description": "An unique positive integer indicating the priority of evaluation for a rule. Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "match": {
              "block": {
                "attributes": {
                  "versioned_expr": {
                    "description": "Predefined rule expression. If this field is specified, config must also be specified. Available options:   SRC_IPS_V1: Must specify the corresponding src_ip_ranges field in config.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "config": {
                    "block": {
                      "attributes": {
                        "src_ip_ranges": {
                          "description": "Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation to match against inbound traffic. There is a limit of 10 IP ranges per rule. A value of '*' matches all IPs (can be used to override the default behavior).",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description": "The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "expr": {
                    "block": {
                      "attributes": {
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax. The application context of the containing message determines which well-known feature set of CEL is supported.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding action is enforced.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match \"*\"). If no rules are provided when creating a security policy, a default rule with action \"allow\" will be added.",
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

func GoogleComputeSecurityPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSecurityPolicy), &result)
	return &result
}
