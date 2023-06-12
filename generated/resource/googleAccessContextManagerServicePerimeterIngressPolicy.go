package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerServicePerimeterIngressPolicy = `{
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
      "ingress_from": {
        "block": {
          "attributes": {
            "identities": {
              "description": "A list of identities that are allowed access through this ingress policy.\nShould be in the format of email address. The email address should represent\nindividual user or service account only.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "identity_type": {
              "description": "Specifies the type of identities that are allowed access from outside the\nperimeter. If left unspecified, then members of 'identities' field will be\nallowed access. Possible values: [\"ANY_IDENTITY\", \"ANY_USER_ACCOUNT\", \"ANY_SERVICE_ACCOUNT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "sources": {
              "block": {
                "attributes": {
                  "access_level": {
                    "description": "An 'AccessLevel' resource name that allow resources within the\n'ServicePerimeters' to be accessed from the internet. 'AccessLevels' listed\nmust be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent\n'AccessLevel' will cause an error. If no 'AccessLevel' names are listed,\nresources within the perimeter can only be accessed via Google Cloud calls\nwith request origins within the perimeter.\nExample 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.'\nIf * is specified, then all IngressSources will be allowed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource": {
                    "description": "A Google Cloud resource that is allowed to ingress the perimeter.\nRequests from these resources will be allowed to access perimeter data.\nCurrently only projects are allowed. Format 'projects/{project_number}'\nThe project may be in any Google Cloud organization, not just the\norganization that the perimeter is defined in. '*' is not allowed, the case\nof allowing all Google Cloud resources only is not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Sources that this 'IngressPolicy' authorizes access from.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines the conditions on the source of a request causing this 'IngressPolicy'\nto apply.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ingress_to": {
        "block": {
          "attributes": {
            "resources": {
              "description": "A list of resources, currently only projects in the form\n'projects/\u003cprojectnumber\u003e', protected by this 'ServicePerimeter'\nthat are allowed to be accessed by sources defined in the\ncorresponding 'IngressFrom'. A request matches if it contains\na resource in this list. If '*' is specified for resources,\nthen this 'IngressTo' rule will authorize access to all\nresources inside the perimeter, provided that the request\nalso matches the 'operations' field.",
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
                    "description": "The name of the API whose methods or permissions the 'IngressPolicy' or\n'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName'\nfield set to '*' will allow all methods AND permissions for all services.",
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
                          "description": "Value for method should be a valid method name for the corresponding\nserviceName in 'ApiOperation'. If '*' used as value for 'method', then\nALL methods and permissions are allowed.",
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
                      "description": "API methods or permissions to allow. Method or permission must belong to\nthe service specified by serviceName field. A single 'MethodSelector' entry\nwith '*' specified for the method field will allow all methods AND\npermissions for the service specified in 'serviceName'.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'\nare allowed to perform in this 'ServicePerimeter'.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Defines the conditions on the 'ApiOperation' and request destination that cause\nthis 'IngressPolicy' to apply.",
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

func GoogleAccessContextManagerServicePerimeterIngressPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerServicePerimeterIngressPolicy), &result)
	return &result
}
