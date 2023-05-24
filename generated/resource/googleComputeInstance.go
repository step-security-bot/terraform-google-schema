package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstance = `{
  "block": {
    "attributes": {
      "allow_stopping_for_update": {
        "description": "If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "can_ip_forward": {
        "description": "Whether sending and receiving of packets with non-matching source or destination IPs is allowed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cpu_platform": {
        "computed": true,
        "description": "The CPU platform used by this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_status": {
        "computed": true,
        "description": "\n\t\t\t\t\tCurrent status of the instance.\n\t\t\t\t\tThis could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED.\n\t\t\t\t\tFor more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether deletion protection is enabled on this instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "A brief description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_status": {
        "description": "Desired status of the instance. Either \"RUNNING\" or \"TERMINATED\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_display": {
        "description": "Whether the instance has virtual displays enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "guest_accelerator": {
        "computed": true,
        "description": "List of the type and count of accelerator cards attached to the instance.",
        "description_kind": "plain",
        "optional": true,
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
      "hostname": {
        "description": "A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.",
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
        "computed": true,
        "description": "The server-assigned unique identifier of this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The unique fingerprint of the labels.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs assigned to the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "description": "The machine type to create.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "description": "Metadata key/value pairs made available within the instance.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Metadata startup scripts made available within the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "min_cpu_platform": {
        "computed": true,
        "description": "The minimum CPU platform specified for the VM instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the instance. One of name or self_link must be provided.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_policies": {
        "description": "A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "self_link": {
        "computed": true,
        "description": "The URI of the created resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description": "The list of tags attached to the instance.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "zone": {
        "computed": true,
        "description": "The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "advanced_machine_features": {
        "block": {
          "attributes": {
            "enable_nested_virtualization": {
              "description": "Whether to enable nested virtualization or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "threads_per_core": {
              "description": "The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "visible_core_count": {
              "description": "The number of physical cores to expose to an instance. Multiply by the number of threads per core to compute the total number of virtual CPUs to expose to the instance. If unset, the number of cores is inferred from the instance\\'s nominal CPU count and the underlying platform\\'s SMT width.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Controls for advanced machine-related behavior features.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "attached_disk": {
        "block": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description": "Name with which the attached disk is accessible under /dev/disk/by-id/",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_key_raw": {
              "description": "A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "disk_encryption_key_sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_self_link": {
              "computed": true,
              "description": "The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description": "Read/write mode for the disk. One of \"READ_ONLY\" or \"READ_WRITE\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "description": "The name or self_link of the disk attached to this instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "List of disks attached to the instance",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "boot_disk": {
        "block": {
          "attributes": {
            "auto_delete": {
              "description": "Whether the disk will be auto-deleted when the instance is deleted.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "device_name": {
              "computed": true,
              "description": "Name with which attached disk will be accessible under /dev/disk/by-id/",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disk_encryption_key_raw": {
              "description": "A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "disk_encryption_key_sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_self_link": {
              "computed": true,
              "description": "The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description": "Read/write mode for the disk. One of \"READ_ONLY\" or \"READ_WRITE\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "computed": true,
              "description": "The name or self_link of the disk attached to this instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "initialize_params": {
              "block": {
                "attributes": {
                  "image": {
                    "computed": true,
                    "description": "The image from which this disk was initialised.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "labels": {
                    "computed": true,
                    "description": "A set of key/value label pairs assigned to the disk.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "size": {
                    "computed": true,
                    "description": "The size of the image in gigabytes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Parameters with which a disk was created alongside the instance.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The boot disk for the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "confidential_instance_config": {
        "block": {
          "attributes": {
            "enable_confidential_compute": {
              "description": "Defines whether the instance should have confidential compute enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "ipv6_access_type": {
              "computed": true,
              "description": "One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the interface",
              "description_kind": "plain",
              "type": "string"
            },
            "network": {
              "computed": true,
              "description": "The name or self_link of the network attached to this interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network_ip": {
              "computed": true,
              "description": "The private IP address assigned to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nic_type": {
              "description": "The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "queue_count": {
              "description": "The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stack_type": {
              "computed": true,
              "description": "The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork": {
              "computed": true,
              "description": "The name or self_link of the subnetwork attached to this interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnetwork_project": {
              "computed": true,
              "description": "The project in which the subnetwork belongs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "access_config": {
              "block": {
                "attributes": {
                  "nat_ip": {
                    "computed": true,
                    "description": "The IP address that is be 1:1 mapped to the instance's network ip.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_tier": {
                    "computed": true,
                    "description": "The networking tier used for configuring this instance. One of PREMIUM or STANDARD.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "public_ptr_domain_name": {
                    "description": "The DNS domain name for the public PTR record.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Access configurations, i.e. IPs via which this instance can be accessed via the Internet.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "alias_ip_range": {
              "block": {
                "attributes": {
                  "ip_cidr_range": {
                    "description": "The IP CIDR range represented by this alias IP range.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnetwork_range_name": {
                    "description": "The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An array of alias IP ranges for this network interface.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ipv6_access_config": {
              "block": {
                "attributes": {
                  "external_ipv6": {
                    "computed": true,
                    "description": "The first IPv6 address of the external IPv6 range associated with this instance, prefix length is stored in externalIpv6PrefixLength in ipv6AccessConfig. The field is output only, an IPv6 address from a subnetwork associated with the instance will be allocated dynamically.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "external_ipv6_prefix_length": {
                    "computed": true,
                    "description": "The prefix length of the external IPv6 range.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_tier": {
                    "description": "The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "public_ptr_domain_name": {
                    "description": "The domain name to be used when creating DNSv6 records for the external IPv6 ranges.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The networks attached to the instance.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "reservation_affinity": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of reservation from which this instance can consume resources.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "specific_reservation": {
              "block": {
                "attributes": {
                  "key": {
                    "description": "Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "Corresponds to the label values of a reservation resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Specifies the label selector for the reservation to use.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies the reservations that this instance can consume from.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scheduling": {
        "block": {
          "attributes": {
            "automatic_restart": {
              "description": "Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "instance_termination_action": {
              "description": "Specifies the action GCE should take when SPOT VM is preempted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_node_cpus": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "on_host_maintenance": {
              "computed": true,
              "description": "Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preemptible": {
              "description": "Whether the instance is preemptible.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "provisioning_model": {
              "computed": true,
              "description": "Whether the instance is spot. If this is set as SPOT.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "node_affinities": {
              "block": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "The scheduling strategy being used by the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "scratch_disk": {
        "block": {
          "attributes": {
            "interface": {
              "description": "The disk interface used for attaching this disk. One of SCSI or NVME.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "size": {
              "description": "The size of the disk in gigabytes. One of 375 or 3000.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "The scratch disks attached to the instance.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "service_account": {
        "block": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "The service account e-mail address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scopes": {
              "description": "A list of service scopes.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "The service account to attach to the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "shielded_instance_config": {
        "block": {
          "attributes": {
            "enable_integrity_monitoring": {
              "description": "Whether integrity monitoring is enabled for the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_secure_boot": {
              "description": "Whether secure boot is enabled for the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_vtpm": {
              "description": "Whether the instance uses vTPM.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The shielded vm config being used by the instance.",
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
  "version": 6
}`

func GoogleComputeInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeInstance), &result)
	return &result
}
