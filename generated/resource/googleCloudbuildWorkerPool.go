package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudbuildWorkerPool = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Time at which the request to create the ` + "`" + `WorkerPool` + "`" + ` was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "Output only. Time at which the request to delete the ` + "`" + `WorkerPool` + "`" + ` was received.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "A user-specified, human-readable name for the ` + "`" + `WorkerPool` + "`" + `. If provided, this value must be 1-63 characters.",
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
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "User-defined name of the ` + "`" + `WorkerPool` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. ` + "`" + `WorkerPool` + "`" + ` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. A unique identifier for the ` + "`" + `WorkerPool` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Time at which the request to update the ` + "`" + `WorkerPool` + "`" + ` was received.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "network_config": {
        "block": {
          "attributes": {
            "peered_network": {
              "description": "Required. Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to ` + "`" + `WorkerPool.project_id` + "`" + ` on the service producer network. Must be in the format ` + "`" + `projects/{project}/global/networks/{network}` + "`" + `, where ` + "`" + `{project}` + "`" + ` is a project number, such as ` + "`" + `12345` + "`" + `, and ` + "`" + `{network}` + "`" + ` is the name of a VPC network in the project. See [Understanding network configuration options](https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peered_network_ip_range": {
              "description": "Optional. Immutable. Subnet IP range within the peered network. This is specified in CIDR notation with a slash and the subnet prefix size. You can optionally specify an IP address before the subnet prefix value. e.g. ` + "`" + `192.168.0.0/29` + "`" + ` would specify an IP range starting at 192.168.0.0 with a prefix size of 29 bits. ` + "`" + `/16` + "`" + ` would specify a prefix size of 16 bits, with an automatically determined IP within the peered VPC. If unspecified, a value of ` + "`" + `/24` + "`" + ` will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Network configuration for the ` + "`" + `WorkerPool` + "`" + `.",
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
      },
      "worker_config": {
        "block": {
          "attributes": {
            "disk_size_gb": {
              "description": "Size of the disk attached to the worker, in GB. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If ` + "`" + `0` + "`" + ` is specified, Cloud Build will use a standard disk size.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "machine_type": {
              "description": "Machine type of a worker, such as ` + "`" + `n1-standard-1` + "`" + `. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use ` + "`" + `n1-standard-1` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "no_external_ip": {
              "computed": true,
              "description": "If true, workers are created without any public address, which prevents network egress to public IPs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configuration to be used for a creating workers in the ` + "`" + `WorkerPool` + "`" + `.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudbuildWorkerPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudbuildWorkerPool), &result)
	return &result
}
