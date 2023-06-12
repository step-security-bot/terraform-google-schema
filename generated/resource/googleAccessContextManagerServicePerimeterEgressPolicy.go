package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerServicePerimeterEgressPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "perimeter": {
        "description": "The name of the Service Perimeter to add this resource to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "egress_from": {
        "block": {
          "attributes": {
            "identities": {
              "description": "A list of identities that are allowed access through this 'EgressPolicy'.\nShould be in the format of email address. The email address should\nrepresent individual user or service account only.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "identity_type": {
              "description": "Specifies the type of identities that are allowed access to outside the\nperimeter. If left unspecified, then members of 'identities' field will\nbe allowed access. Possible values: [\"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Defines conditions on the source of a request causing this 'EgressPolicy' to apply.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "egress_to": {
        "block": {
          "attributes": {
            "external_resources": {
              "description": "A list of external resources that are allowed to be accessed. A request\nmatches if it contains an external resource in this list (Example:\ns3://bucket/path). Currently '*' is not allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "description": "A list of resources, currently only projects in the form\n'projects/\u003cprojectnumber\u003e', that match this to stanza. A request matches\nif it contains a resource in this list. If * is specified for resources,\nthen this 'EgressTo' rule will authorize access to all resources outside\nthe perimeter.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "operations": {
              "block": {
                "attributes": {
                  "service_name": {
                    "description": "The name of the API whose methods or permissions the 'IngressPolicy' or\n'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName\nfield set to '*' will allow all methods AND permissions for all services.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "method_selectors": {
                    "block": {
                      "attributes": {
                        "method": {
                          "description": "Value for 'method' should be a valid method name for the corresponding\n'serviceName' in 'ApiOperation'. If '*' used as value for method,\nthen ALL methods and permissions are allowed.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "permission": {
                          "description": "Value for permission should be a valid Cloud IAM permission for the\ncorresponding 'serviceName' in 'ApiOperation'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "API methods or permissions to allow. Method or permission must belong\nto the service specified by 'serviceName' field. A single MethodSelector\nentry with '*' specified for the 'method' field will allow all methods\nAND permissions for the service specified in 'serviceName'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of 'ApiOperations' that this egress rule applies to. A request matches\nif it contains an operation/service in this list.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines the conditions on the 'ApiOperation' and destination resources that\ncause this 'EgressPolicy' to apply.",
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

func GoogleAccessContextManagerServicePerimeterEgressPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerServicePerimeterEgressPolicy), &result)
	return &result
}
