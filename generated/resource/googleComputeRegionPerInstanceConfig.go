package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionPerInstanceConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "minimal_action": {
        "description": "The minimal action to perform on the instance during an update.\nDefault is 'NONE'. Possible values are:\n* REPLACE\n* RESTART\n* REFRESH\n* NONE",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "most_disruptive_allowed_action": {
        "description": "The most disruptive action to perform on the instance during an update.\nDefault is 'REPLACE'. Possible values are:\n* REPLACE\n* RESTART\n* REFRESH\n* NONE",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name for this per-instance config and its corresponding instance.",
        "description_kind": "plain",
        "required": true,
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
        "description": "Region where the containing instance group manager is located",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_instance_group_manager": {
        "description": "The region instance group manager this instance config is part of.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remove_instance_state_on_destroy": {
        "description": "When true, deleting this config will immediately remove any specified state from the underlying instance.\nWhen false, deleting this config will *not* immediately remove any state from the underlying instance.\nState will be removed on the next instance recreation or update.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "preserved_state": {
        "block": {
          "attributes": {
            "metadata": {
              "description": "Preserved metadata defined for this instance. This is a list of key-\u003evalue pairs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "disk": {
              "block": {
                "attributes": {
                  "delete_rule": {
                    "description": "A value that prescribes what should happen to the stateful disk when the VM instance is deleted.\nThe available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.\n'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.\n'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently\ndeleted from the instance group. Default value: \"NEVER\" Possible values: [\"NEVER\", \"ON_PERMANENT_INSTANCE_DELETION\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "device_name": {
                    "description": "A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "mode": {
                    "description": "The mode of the disk. Default value: \"READ_WRITE\" Possible values: [\"READ_ONLY\", \"READ_WRITE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source": {
                    "description": "The URI of an existing persistent disk to attach under the specified device-name in the format\n'projects/project-id/zones/zone/disks/disk-name'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Stateful disks for the instance.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The preserved state for this instance.",
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

func GoogleComputeRegionPerInstanceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionPerInstanceConfig), &result)
	return &result
}
