package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVpcAccessConnector = `{
  "block": {
    "attributes": {
      "connected_projects": {
        "computed": true,
        "description": "List of projects using the connector.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_cidr_range": {
        "computed": true,
        "description": "The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.",
        "description_kind": "plain",
        "type": "string"
      },
      "machine_type": {
        "computed": true,
        "description": "Machine type of VM Instance underlying connector. Default is e2-micro",
        "description_kind": "plain",
        "type": "string"
      },
      "max_instances": {
        "computed": true,
        "description": "Maximum value of instances in autoscaling group underlying the connector.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_throughput": {
        "computed": true,
        "description": "Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 1000.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_instances": {
        "computed": true,
        "description": "Minimum value of instances in autoscaling group underlying the connector.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_throughput": {
        "computed": true,
        "description": "Minimum throughput of the connector in Mbps. Default and min is 200.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The name of the resource (Max 25 characters).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "Name or self_link of the VPC network. Required if 'ip_cidr_range' is set.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
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
      },
      "subnet": {
        "computed": true,
        "description": "The subnet in which to house the connector",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "project_id": "string"
            }
          ]
        ]
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
