package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstanceGroupManager = `{
  "block": {
    "attributes": {
      "auto_healing_policies": {
        "computed": true,
        "description": "The autohealing policies for this managed instance group. You can specify only one value.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "health_check": "string",
              "initial_delay_sec": "number"
            }
          ]
        ]
      },
      "base_instance_name": {
        "computed": true,
        "description": "The base instance name to use for instances in this group. The value must be a valid RFC1035 name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional textual description of the instance group manager.",
        "description_kind": "plain",
        "type": "string"
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
        "computed": true,
        "description": "Pagination behavior of the listManagedInstances API method for this managed instance group. Valid values are: \"PAGELESS\", \"PAGINATED\". If PAGELESS (default), Pagination is disabled for the group's listManagedInstances API method. maxResults and pageToken query parameters are ignored and all instances are returned in a single response. If PAGINATED, pagination is enabled, maxResults and pageToken query parameters are respected.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the instance group manager. Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "named_port": {
        "computed": true,
        "description": "The named port configuration.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "port": "number"
            }
          ]
        ]
      },
      "operation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "description": "The URL of the created resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stateful_disk": {
        "computed": true,
        "description": "Disks created on the instances that will be preserved on instance delete, update, etc.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "delete_rule": "string",
              "device_name": "string"
            }
          ]
        ]
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
        "computed": true,
        "description": "The full URL of all target pools to which new instances in the group are added. Updating the target pools attribute does not affect existing instances.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "target_size": {
        "computed": true,
        "description": "The target number of running instances for this managed instance group. This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.",
        "description_kind": "plain",
        "type": "number"
      },
      "update_policy": {
        "computed": true,
        "description": "The update policy for this managed instance group.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_surge_fixed": "number",
              "max_surge_percent": "number",
              "max_unavailable_fixed": "number",
              "max_unavailable_percent": "number",
              "minimal_action": "string",
              "most_disruptive_allowed_action": "string",
              "replacement_method": "string",
              "type": "string"
            }
          ]
        ]
      },
      "version": {
        "computed": true,
        "description": "Application versions managed by this instance group. Each version deals with a specific instance template, allowing canary release scenarios.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instance_template": "string",
              "name": "string",
              "target_size": [
                "list",
                [
                  "object",
                  {
                    "fixed": "number",
                    "percent": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "wait_for_instances": {
        "computed": true,
        "description": "Whether to wait for all instances to be created/updated before returning. Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.",
        "description_kind": "plain",
        "type": "bool"
      },
      "wait_for_instances_status": {
        "computed": true,
        "description": "When used with wait_for_instances specifies the status to wait for. When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective as well as all instances to be stable before returning.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "The zone that instances in this group should be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeInstanceGroupManagerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstanceGroupManager), &result)
	return &result
}
