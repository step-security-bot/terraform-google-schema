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
      "machine_type": {
        "description": "Machine type of VM Instance underlying connector. Default is e2-micro",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_instances": {
        "computed": true,
        "description": "Maximum value of instances in autoscaling group underlying the connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_throughput": {
        "description": "Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_instances": {
        "computed": true,
        "description": "Minimum value of instances in autoscaling group underlying the connector.",
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
        "computed": true,
        "description": "Name or self_link of the VPC network. Required if 'ip_cidr_range' is set.",
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
      "subnet": {
        "block": {
          "attributes": {
            "name": {
              "description": "Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is\nhttps://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}\"",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_id": {
              "computed": true,
              "description": "Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The subnet in which to house the connector",
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
