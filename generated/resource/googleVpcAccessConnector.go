package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVpcAccessConnector = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_cidr_range": {
        "description": "The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_throughput": {
        "description": "Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_throughput": {
        "description": "Minimum throughput of the connector in Mbps. Default and min is 200.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "The name of the resource (Max 25 characters).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "Name of the VPC network. Required if 'ip_cidr_range' is set.",
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
      "region": {
        "computed": true,
        "description": "Region where the VPC Access connector resides. If it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The fully qualified name of this VPC connector",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the VPC access connector.",
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

func GoogleVpcAccessConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVpcAccessConnector), &result)
	return &result
}
