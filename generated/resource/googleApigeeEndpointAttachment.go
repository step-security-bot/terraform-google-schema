package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEndpointAttachment = `{
  "block": {
    "attributes": {
      "connection_state": {
        "computed": true,
        "description": "State of the endpoint attachment connection to the service attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_attachment_id": {
        "description": "ID of the endpoint attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host": {
        "computed": true,
        "description": "Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location of the endpoint attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the Endpoint Attachment in the following format:\norganizations/{organization}/endpointAttachments/{endpointAttachment}.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_attachment": {
        "description": "Format: projects/*/regions/*/serviceAttachments/*",
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

func GoogleApigeeEndpointAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEndpointAttachment), &result)
	return &result
}
