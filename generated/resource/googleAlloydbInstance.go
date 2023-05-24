package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbInstance = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "availability_type": {
        "computed": true,
        "description": "Availability type of an Instance. Defaults to REGIONAL for both primary and read instances. Note that primary and read instances can have different availability types. Possible values: [\"AVAILABILITY_TYPE_UNSPECIFIED\", \"ZONAL\", \"REGIONAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster": {
        "description": "Identifies the alloydb cluster. Must be in the format\n'projects/{project}/locations/{location}/clusters/{cluster_id}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the Instance was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_flags": {
        "description": "Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "display_name": {
        "description": "User-settable and human-readable display name for the Instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gce_zone": {
        "description": "The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity.",
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
      "instance_id": {
        "description": "The ID of the alloydb instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "description": "The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the 'depends_on' meta-data attribute. Possible values: [\"PRIMARY\", \"READ_POOL\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The IP address for the Instance. This is the connection endpoint for an end-user application.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for the alloydb instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the instance resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of the alloydb instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "The system-generated UID of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the Instance was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "machine_config": {
        "block": {
          "attributes": {
            "cpu_count": {
              "computed": true,
              "description": "The number of CPU's in the VM instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Configurations for the machines that host the underlying database engine.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "read_pool_config": {
        "block": {
          "attributes": {
            "node_count": {
              "description": "Read capacity, i.e. number of nodes in a read pool instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Read pool specific config.",
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

func GoogleAlloydbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbInstance), &result)
	return &result
}
