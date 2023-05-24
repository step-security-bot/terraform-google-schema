package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerGcpUserAccessBinding = `{
  "block": {
    "attributes": {
      "access_levels": {
        "description": "Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: \"accessPolicies/9522/accessLevels/device_trusted\"",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "group_key": {
        "description": "Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See \"id\" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: \"01d520gv4vjcrht\"",
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
      "name": {
        "computed": true,
        "description": "Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example: \"organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N\"",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_id": {
        "description": "Required. ID of the parent organization.",
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

func GoogleAccessContextManagerGcpUserAccessBindingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerGcpUserAccessBinding), &result)
	return &result
}
