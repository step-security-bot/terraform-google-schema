package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpAppGateway = `{
  "block": {
    "attributes": {
      "allocated_connections": {
        "computed": true,
        "description": "A list of connections allocated for the Gateway.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ingress_port": "number",
              "psc_uri": "string"
            }
          ]
        ]
      },
      "display_name": {
        "computed": true,
        "description": "An arbitrary user-provided name for the AppGateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_type": {
        "computed": true,
        "description": "The type of hosting used by the AppGateway. Default value: \"HOST_TYPE_UNSPECIFIED\" Possible values: [\"HOST_TYPE_UNSPECIFIED\", \"GCP_REGIONAL_MIG\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "ID of the AppGateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region of the AppGateway.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Represents the different states of a AppGateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of network connectivity used by the AppGateway. Default value: \"TYPE_UNSPECIFIED\" Possible values: [\"TYPE_UNSPECIFIED\", \"TCP_PROXY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "Server-defined URI for this resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBeyondcorpAppGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBeyondcorpAppGateway), &result)
	return &result
}
