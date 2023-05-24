package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetworkFirewallPolicyRule = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The Action to perform when the client connection triggers the rule. Can currently be either \"allow\" or \"deny()\" where valid values for status are 403, 404, and 502.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "An optional description for this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "direction": {
        "description": "The direction in which this rule applies. Possible values: INGRESS, EGRESS",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disabled": {
        "description": "Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_logging": {
        "description": "Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on \"goto_next\" rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "firewall_policy": {
        "description": "The firewall policy of the resource.",
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
      "kind": {
        "computed": true,
        "description": "Type of the resource. Always ` + "`" + `compute#firewallPolicyRule` + "`" + ` for firewall policy rules",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "description": "An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_name": {
        "description": "An optional name for the rule. This field is not a unique identifier and can be updated.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_tuple_count": {
        "computed": true,
        "description": "Calculation of the complexity of a single firewall policy rule.",
        "description_kind": "plain",
        "type": "number"
      },
      "target_service_accounts": {
        "description": "A list of service accounts indicating the sets of instances that are applied with this rule.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "match": {
        "block": {
          "attributes": {
            "dest_address_groups": {
              "description": "Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_fqdns": {
              "description": "Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_ip_ranges": {
              "description": "CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_region_codes": {
              "description": "The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dest_threat_intelligences": {
              "description": "Name of the Google Cloud Threat Intelligence list.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_address_groups": {
              "description": "Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_fqdns": {
              "description": "Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_ip_ranges": {
              "description": "CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_region_codes": {
              "description": "The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "src_threat_intelligences": {
              "description": "Name of the Google Cloud Threat Intelligence list.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "layer4_configs": {
              "block": {
                "attributes": {
                  "ip_protocol": {
                    "description": "The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `esp` + "`" + `, ` + "`" + `ah` + "`" + `, ` + "`" + `ipip` + "`" + `, ` + "`" + `sctp` + "`" + `), or the IP protocol number.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "ports": {
                    "description": "An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Pairs of IP protocols and ports that the rule should match.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "src_secure_tags": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "[Output Only] State of the secure tag, either ` + "`" + `EFFECTIVE` + "`" + ` or ` + "`" + `INEFFECTIVE` + "`" + `. A secure tag is ` + "`" + `INEFFECTIVE` + "`" + ` when it is deleted or its network is deleted.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the \u003ccode\u003esrcSecureTag\u003c/code\u003e are INEFFECTIVE, and there is no \u003ccode\u003esrcIpRange\u003c/code\u003e, this rule will be ignored. Maximum number of source tag values allowed is 256.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "target_secure_tags": {
        "block": {
          "attributes": {
            "name": {
              "description": "Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "[Output Only] State of the secure tag, either ` + "`" + `EFFECTIVE` + "`" + ` or ` + "`" + `INEFFECTIVE` + "`" + `. A secure tag is ` + "`" + `INEFFECTIVE` + "`" + ` when it is deleted or its network is deleted.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "A list of secure tags that controls which instances the firewall rule applies to. If \u003ccode\u003etargetSecureTag\u003c/code\u003e are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE state, then this rule will be ignored. \u003ccode\u003etargetSecureTag\u003c/code\u003e may not be set at the same time as \u003ccode\u003etargetServiceAccounts\u003c/code\u003e. If neither \u003ccode\u003etargetServiceAccounts\u003c/code\u003e nor \u003ccode\u003etargetSecureTag\u003c/code\u003e are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.",
          "description_kind": "plain"
        },
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

func GoogleComputeNetworkFirewallPolicyRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetworkFirewallPolicyRule), &result)
	return &result
}
