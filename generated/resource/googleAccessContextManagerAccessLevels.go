package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerAccessLevels = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent": {
        "description": "The AccessPolicy this AccessLevel lives in.\nFormat: accessPolicies/{policy_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "access_levels": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description of the AccessLevel and its use. Does not affect behavior.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "Resource name for the Access Level. The short_name component must begin\nwith a letter and only include alphanumeric and '_'.\nFormat: accessPolicies/{policy_id}/accessLevels/{short_name}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "title": {
              "description": "Human readable title. Must be unique within the Policy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "basic": {
              "block": {
                "attributes": {
                  "combining_function": {
                    "description": "How the conditions list should be combined to determine if a request\nis granted this AccessLevel. If AND is used, each Condition in\nconditions must be satisfied for the AccessLevel to be applied. If\nOR is used, at least one Condition in conditions must be satisfied\nfor the AccessLevel to be applied. Default value: \"AND\" Possible values: [\"AND\", \"OR\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "conditions": {
                    "block": {
                      "attributes": {
                        "ip_subnetworks": {
                          "description": "A list of CIDR block IP subnetwork specification. May be IPv4\nor IPv6.\nNote that for a CIDR IP address block, the specified IP address\nportion must be properly truncated (i.e. all the host bits must\nbe zero) or the input is considered malformed. For example,\n\"192.0.2.0/24\" is accepted but \"192.0.2.1/24\" is not. Similarly,\nfor IPv6, \"2001:db8::/32\" is accepted whereas \"2001:db8::1/32\"\nis not. The originating IP of a request must be in one of the\nlisted subnets in order for this Condition to be true.\nIf empty, all IP addresses are allowed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "members": {
                          "description": "An allowed list of members (users, service accounts).\nUsing groups is not supported yet.\n\nThe signed-in user originating the request must be a part of one\nof the provided members. If not specified, a request may come\nfrom any user (logged in/not logged in, not present in any\ngroups, etc.).\nFormats: 'user:{emailid}', 'serviceAccount:{emailid}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "negate": {
                          "description": "Whether to negate the Condition. If true, the Condition becomes\na NAND over its non-empty fields, each field must be false for\nthe Condition overall to be satisfied. Defaults to false.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "regions": {
                          "description": "The request must originate from one of the provided\ncountries/regions.\nFormat: A valid ISO 3166-1 alpha-2 code.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "required_access_levels": {
                          "description": "A list of other access levels defined in the same Policy,\nreferenced by resource name. Referencing an AccessLevel which\ndoes not exist is an error. All access levels listed must be\ngranted for the Condition to be true.\nFormat: accessPolicies/{policy_id}/accessLevels/{short_name}",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "device_policy": {
                          "block": {
                            "attributes": {
                              "allowed_device_management_levels": {
                                "description": "A list of allowed device management levels.\nAn empty list allows all management levels. Possible values: [\"MANAGEMENT_UNSPECIFIED\", \"NONE\", \"BASIC\", \"COMPLETE\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_encryption_statuses": {
                                "description": "A list of allowed encryptions statuses.\nAn empty list allows all statuses. Possible values: [\"ENCRYPTION_UNSPECIFIED\", \"ENCRYPTION_UNSUPPORTED\", \"UNENCRYPTED\", \"ENCRYPTED\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "require_admin_approval": {
                                "description": "Whether the device needs to be approved by the customer admin.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "require_corp_owned": {
                                "description": "Whether the device needs to be corp owned.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "require_screen_lock": {
                                "description": "Whether or not screenlock is required for the DevicePolicy\nto be true. Defaults to false.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "os_constraints": {
                                "block": {
                                  "attributes": {
                                    "minimum_version": {
                                      "description": "The minimum allowed OS version. If not set, any version\nof this OS satisfies the constraint.\nFormat: \"major.minor.patch\" such as \"10.5.301\", \"9.2.1\".",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "os_type": {
                                      "description": "The operating system type of the device. Possible values: [\"OS_UNSPECIFIED\", \"DESKTOP_MAC\", \"DESKTOP_WINDOWS\", \"DESKTOP_LINUX\", \"DESKTOP_CHROME_OS\", \"ANDROID\", \"IOS\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A list of allowed OS versions.\nAn empty list allows all types and all versions.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Device specific restrictions, all restrictions must hold for\nthe Condition to be true. If not specified, all devices are\nallowed.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "A set of requirements for the AccessLevel to be granted.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A set of predefined conditions for the access level and a combining function.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "custom": {
              "block": {
                "block_types": {
                  "expr": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "Description of the expression",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "location": {
                          "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "title": {
                          "description": "Title for the expression, i.e. a short string describing its purpose.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.\nThis page details the objects and attributes that are used to the build the CEL expressions for \ncustom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request. \nSee CEL spec at: https://github.com/google/cel-spec.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The desired Access Levels that should replace all existing Access Levels in the Access Policy.",
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

func GoogleAccessContextManagerAccessLevelsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerAccessLevels), &result)
	return &result
}
