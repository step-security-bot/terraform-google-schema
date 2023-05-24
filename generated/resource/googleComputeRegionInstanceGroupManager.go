package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionInstanceGroupManager = `{
  "block": {
    "attributes": {
      "base_instance_name": {
        "description": "The base instance name to use for instances in this group. The value must be a valid RFC1035 name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "An optional textual description of the instance group manager.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distribution_policy_target_shape": {
        "computed": true,
        "description": "The shape to which the group converges either proactively or on resize events (depending on the value set in updatePolicy.instanceRedistributionType).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distribution_policy_zones": {
        "computed": true,
        "description": "The distribution policy for this managed instance group. You can specify one or more values.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "fingerprint": {
        "computed": true,
        "description": "The fingerprint of the instance group manager.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_group": {
        "computed": true,
        "description": "The full URL of the instance group created by the manager.",
        "description_kind": "plain",
        "type": "string"
      },
      "list_managed_instances_results": {
        "description": "Pagination behavior of the listManagedInstances API method for this managed instance group. Valid values are: \"PAGELESS\", \"PAGINATED\". If PAGELESS (default), Pagination is disabled for the group's listManagedInstances API method. maxResults and pageToken query parameters are ignored and all instances are returned in a single response. If PAGINATED, pagination is enabled, maxResults and pageToken query parameters are respected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the instance group manager. Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region where the managed instance group resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description": "The URL of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of this managed instance group.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "is_stable": "bool",
              "stateful": [
                "list",
                [
                  "object",
                  {
                    "has_stateful_config": "bool",
                    "per_instance_configs": [
                      "list",
                      [
                        "object",
                        {
                          "all_effective": "bool"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "version_target": [
                "list",
                [
                  "object",
                  {
                    "is_reached": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "target_pools": {
        "description": "The full URL of all target pools to which new instances in the group are added. Updating the target pools attribute does not affect existing instances.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "target_size": {
        "computed": true,
        "description": "The target number of running instances for this managed instance group. This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "wait_for_instances": {
        "description": "Whether to wait for all instances to be created/updated before returning. Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "wait_for_instances_status": {
        "description": "When used with wait_for_instances specifies the status to wait for. When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective as well as all instances to be stable before returning.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_healing_policies": {
        "block": {
          "attributes": {
            "health_check": {
              "description": "The health check resource that signals autohealing.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "initial_delay_sec": {
              "description": "The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "The autohealing policies for this managed instance group. You can specify only one value.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "named_port": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name of the port.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "The port number.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "The named port configuration.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "stateful_disk": {
        "block": {
          "attributes": {
            "delete_rule": {
              "description": "A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION. NEVER - detach the disk when the VM is deleted, but do not delete the disk. ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently deleted from the instance group. The default is NEVER.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_name": {
              "description": "The device name of the disk to be attached.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the official documentation. Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the update_policy.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      "update_policy": {
        "block": {
          "attributes": {
            "instance_redistribution_type": {
              "description": "The instance redistribution policy for regional managed instance groups. Valid values are: \"PROACTIVE\", \"NONE\". If PROACTIVE (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If NONE, proactive redistribution is disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_surge_fixed": {
              "computed": true,
              "description": "The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with max_surge_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_surge_percent": {
              "description": "The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with max_surge_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_fixed": {
              "computed": true,
              "description": "The maximum number of instances that can be unavailable during the update process. Conflicts with max_unavailable_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_percent": {
              "description": "The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with max_unavailable_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimal_action": {
              "description": "Minimal action to be taken on an instance. You can specify either REFRESH to update without stopping instances, RESTART to restart existing instances or REPLACE to delete and create new instances from the target template. If you specify a REFRESH, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "most_disruptive_allowed_action": {
              "description": "Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "replacement_method": {
              "description": "The instance replacement method for regional managed instance groups. Valid values are: \"RECREATE\", \"SUBSTITUTE\". If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "The type of update process. You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The update policy for this managed instance group.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "version": {
        "block": {
          "attributes": {
            "instance_template": {
              "description": "The full URL to an instance template from which all new instances of this version will be created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "Version name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "target_size": {
              "block": {
                "attributes": {
                  "fixed": {
                    "description": "The number of instances which are managed for this version. Conflicts with percent.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "percent": {
                    "description": "The number of instances (calculated as percentage) which are managed for this version. Conflicts with fixed. Note that when using percent, rounding will be in favor of explicitly set target_size values; a managed instance group with 2 instances and 2 versions, one of which has a target_size.percent of 60 will create 2 instances of that version.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The number of instances calculated as a fixed number or a percentage depending on the settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Application versions managed by this instance group. Each version deals with a specific instance template, allowing canary release scenarios.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeRegionInstanceGroupManagerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionInstanceGroupManager), &result)
	return &result
}
