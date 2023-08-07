package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstanceTemplate = `{
  "block": {
    "attributes": {
      "advanced_machine_features": {
        "computed": true,
        "description": "Controls for advanced machine-related behavior features.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_nested_virtualization": "bool",
              "threads_per_core": "number",
              "visible_core_count": "number"
            }
          ]
        ]
      },
      "can_ip_forward": {
        "computed": true,
        "description": "Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "confidential_instance_config": {
        "computed": true,
        "description": "The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail to create.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_confidential_compute": "bool"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "A brief description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "disk": {
        "computed": true,
        "description": "Disks to attach to instances created from this template. This can be specified multiple times for multiple disks.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_delete": "bool",
              "boot": "bool",
              "device_name": "string",
              "disk_encryption_key": [
                "list",
                [
                  "object",
                  {
                    "kms_key_self_link": "string"
                  }
                ]
              ],
              "disk_name": "string",
              "disk_size_gb": "number",
              "disk_type": "string",
              "interface": "string",
              "labels": [
                "map",
                "string"
              ],
              "mode": "string",
              "resource_policies": [
                "list",
                "string"
              ],
              "source": "string",
              "source_image": "string",
              "source_image_encryption_key": [
                "list",
                [
                  "object",
                  {
                    "kms_key_self_link": "string",
                    "kms_key_service_account": "string"
                  }
                ]
              ],
              "source_snapshot": "string",
              "source_snapshot_encryption_key": [
                "list",
                [
                  "object",
                  {
                    "kms_key_self_link": "string",
                    "kms_key_service_account": "string"
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "filter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "guest_accelerator": {
        "computed": true,
        "description": "List of the type and count of accelerator cards attached to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "count": "number",
              "type": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_description": {
        "computed": true,
        "description": "A description of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "A set of key/value label pairs to assign to instances created from this template,",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "computed": true,
        "description": "The machine type to create. To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Metadata key/value pairs to make available from within instances created from this template.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "metadata_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the metadata.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_startup_script": {
        "computed": true,
        "description": "An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.",
        "description_kind": "plain",
        "type": "string"
      },
      "min_cpu_platform": {
        "computed": true,
        "description": "Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.",
        "description_kind": "plain",
        "type": "string"
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description": "Creates a unique name beginning with the specified prefix. Conflicts with name.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface": {
        "computed": true,
        "description": "Networks to attach to instances created from this template. This can be specified multiple times for multiple networks.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_config": [
                "list",
                [
                  "object",
                  {
                    "nat_ip": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ],
              "alias_ip_range": [
                "list",
                [
                  "object",
                  {
                    "ip_cidr_range": "string",
                    "subnetwork_range_name": "string"
                  }
                ]
              ],
              "ipv6_access_config": [
                "list",
                [
                  "object",
                  {
                    "external_ipv6": "string",
                    "external_ipv6_prefix_length": "string",
                    "name": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ],
              "ipv6_access_type": "string",
              "name": "string",
              "network": "string",
              "network_ip": "string",
              "nic_type": "string",
              "queue_count": "number",
              "stack_type": "string",
              "subnetwork": "string",
              "subnetwork_project": "string"
            }
          ]
        ]
      },
      "network_performance_config": {
        "computed": true,
        "description": "Configures network performance settings for the instance. If not specified, the instance will be created with its default network performance configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "total_egress_bandwidth_tier": "string"
            }
          ]
        ]
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom subnetwork resource is tied to a specific region. Defaults to the region of the Provider if no value is given.",
        "description_kind": "plain",
        "type": "string"
      },
      "reservation_affinity": {
        "computed": true,
        "description": "Specifies the reservations that this instance can consume from.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "specific_reservation": [
                "list",
                [
                  "object",
                  {
                    "key": "string",
                    "values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "resource_policies": {
        "computed": true,
        "description": "A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scheduling": {
        "computed": true,
        "description": "The scheduling strategy to use.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_restart": "bool",
              "instance_termination_action": "string",
              "local_ssd_recovery_timeout": [
                "list",
                [
                  "object",
                  {
                    "nanos": "number",
                    "seconds": "number"
                  }
                ]
              ],
              "min_node_cpus": "number",
              "node_affinities": [
                "set",
                [
                  "object",
                  {
                    "key": "string",
                    "operator": "string",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "on_host_maintenance": "string",
              "preemptible": "bool",
              "provisioning_model": "string"
            }
          ]
        ]
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "self_link_unique": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "Service account to attach to the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "email": "string",
              "scopes": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "shielded_instance_config": {
        "computed": true,
        "description": "Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Note: shielded_instance_config can only be used with boot images with shielded vm support.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_integrity_monitoring": "bool",
              "enable_secure_boot": "bool",
              "enable_vtpm": "bool"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags to attach to the instance.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the tags.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleComputeInstanceTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstanceTemplate), &result)
	return &result
}
