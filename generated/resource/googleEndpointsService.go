package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEndpointsService = `{
  "block": {
    "attributes": {
      "apis": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "methods": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "request_type": "string",
                    "response_type": "string",
                    "syntax": "string"
                  }
                ]
              ],
              "name": "string",
              "syntax": "string",
              "version": "string"
            }
          ]
        ]
      },
      "config_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "name": "string"
            }
          ]
        ]
      },
      "grpc_config": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "openapi_config": {
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
      "protoc_output": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protoc_output_base64": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleEndpointsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEndpointsService), &result)
	return &result
}
