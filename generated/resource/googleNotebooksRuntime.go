package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNotebooksRuntime = `{
  "block": {
    "attributes": {
      "health_state": {
        "computed": true,
        "description": "The health state of this runtime. For a list of possible output\nvalues, see 'https://cloud.google.com/vertex-ai/docs/workbench/\nreference/rest/v1/projects.locations.runtimes#healthstate'.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "A reference to the zone where the machine resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metrics": {
        "computed": true,
        "description": "Contains Runtime daemon metrics such as Service status and JupyterLab\nstatus",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "system_metrics": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "The name specified for the Notebook runtime.",
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
      "state": {
        "computed": true,
        "description": "The state of this runtime.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "access_config": {
        "block": {
          "attributes": {
            "access_type": {
              "description": "The type of access mode this instance. For valid values, see\n'https://cloud.google.com/vertex-ai/docs/workbench/reference/\nrest/v1/projects.locations.runtimes#RuntimeAccessType'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "proxy_uri": {
              "computed": true,
              "description": "The proxy endpoint that is used to access the runtime.",
              "description_kind": "plain",
              "type": "string"
            },
            "runtime_owner": {
              "description": "The owner of this runtime after creation. Format: 'alias@example.com'.\nCurrently supports one owner only.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The config settings for accessing runtime.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "software_config": {
        "block": {
          "attributes": {
            "custom_gpu_driver_path": {
              "description": "Specify a custom Cloud Storage path where the GPU driver is stored.\nIf not specified, we'll automatically choose from official GPU drivers.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_health_monitoring": {
              "description": "Verifies core internal services are running. Default: True.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "idle_shutdown": {
              "description": "Runtime will automatically shutdown after idle_shutdown_time.\nDefault: True",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "idle_shutdown_timeout": {
              "description": "Time in minutes to wait before shuting down runtime.\nDefault: 180 minutes",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "install_gpu_driver": {
              "description": "Install Nvidia Driver automatically.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "notebook_upgrade_schedule": {
              "description": "Cron expression in UTC timezone for schedule instance auto upgrade.\nPlease follow the [cron format](https://en.wikipedia.org/wiki/Cron).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "post_startup_script": {
              "description": "Path to a Bash script that automatically runs after a notebook instance\nfully boots up. The path must be a URL or\nCloud Storage path (gs://path-to-file/file-name).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "post_startup_script_behavior": {
              "description": "Behavior for the post startup script. Possible values: [\"POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED\", \"RUN_EVERY_START\", \"DOWNLOAD_AND_RUN_EVERY_START\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "upgradeable": {
              "computed": true,
              "description": "Bool indicating whether an newer image is available in an image family.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "block_types": {
            "kernels": {
              "block": {
                "attributes": {
                  "repository": {
                    "description": "The path to the container image repository.\nFor example: gcr.io/{project_id}/{imageName}",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag": {
                    "description": "The tag of the container image. If not specified, this defaults to the latest tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Use a list of container images to use as Kernels in the notebook instance.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The config settings for software inside the runtime.",
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
      "virtual_machine": {
        "block": {
          "attributes": {
            "instance_id": {
              "computed": true,
              "description": "The unique identifier of the Managed Compute Engine instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_name": {
              "computed": true,
              "description": "The user-friendly name of the Managed Compute Engine instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "virtual_machine_config": {
              "block": {
                "attributes": {
                  "guest_attributes": {
                    "computed": true,
                    "description": "The Compute Engine guest attributes. (see [Project and instance\nguest attributes](https://cloud.google.com/compute/docs/\nstoring-retrieving-metadata#guest_attributes)).",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "internal_ip_only": {
                    "description": "If true, runtime will only have internal IP addresses. By default,\nruntimes are not restricted to internal IP addresses, and will\nhave ephemeral external IP addresses assigned to each vm. This\n'internal_ip_only' restriction can only be enabled for subnetwork\nenabled networks, and all dependencies must be configured to be\naccessible without external IP addresses.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "labels": {
                    "computed": true,
                    "description": "The labels to associate with this runtime. Label **keys** must\ncontain 1 to 63 characters, and must conform to [RFC 1035]\n(https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be\nempty, but, if present, must contain 1 to 63 characters, and must\nconform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No\nmore than 32 labels can be associated with a cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "machine_type": {
                    "description": "The Compute Engine machine type used for runtimes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metadata": {
                    "computed": true,
                    "description": "The Compute Engine metadata entries to add to virtual machine.\n(see [Project and instance metadata](https://cloud.google.com\n/compute/docs/storing-retrieving-metadata#project_and_instance\n_metadata)).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "network": {
                    "description": "The Compute Engine network to be used for machine communications.\nCannot be specified with subnetwork. If neither 'network' nor\n'subnet' is specified, the \"default\" network of the project is\nused, if it exists. A full URL or partial URI. Examples:\n  * 'https://www.googleapis.com/compute/v1/projects/[project_id]/\n  regions/global/default'\n  * 'projects/[project_id]/regions/global/default'\nRuntimes are managed resources inside Google Infrastructure.\nRuntimes support the following network configurations:\n  * Google Managed Network (Network \u0026 subnet are empty)\n  * Consumer Project VPC (network \u0026 subnet are required). Requires\n  configuring Private Service Access.\n  * Shared VPC (network \u0026 subnet are required). Requires\n  configuring Private Service Access.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "nic_type": {
                    "description": "The type of vNIC to be used on this interface. This may be gVNIC\nor VirtioNet. Possible values: [\"UNSPECIFIED_NIC_TYPE\", \"VIRTIO_NET\", \"GVNIC\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reserved_ip_range": {
                    "description": "Reserved IP Range name is used for VPC Peering. The\nsubnetwork allocation will use the range *name* if it's assigned.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet": {
                    "description": "The Compute Engine subnetwork to be used for machine\ncommunications. Cannot be specified with network. A full URL or\npartial URI are valid. Examples:\n  * 'https://www.googleapis.com/compute/v1/projects/[project_id]/\n  regions/us-east1/subnetworks/sub0'\n  * 'projects/[project_id]/regions/us-east1/subnetworks/sub0'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "The Compute Engine tags to add to runtime (see [Tagging instances]\n(https://cloud.google.com/compute/docs/\nlabel-or-tag-resources#tags)).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "The zone where the virtual machine is located.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "accelerator_config": {
                    "block": {
                      "attributes": {
                        "core_count": {
                          "description": "Count of cores of this accelerator.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "type": {
                          "description": "Accelerator model. For valid values, see\n'https://cloud.google.com/vertex-ai/docs/workbench/reference/\nrest/v1/projects.locations.runtimes#AcceleratorType'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The Compute Engine accelerator configuration for this runtime.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "container_images": {
                    "block": {
                      "attributes": {
                        "repository": {
                          "description": "The path to the container image repository.\nFor example: gcr.io/{project_id}/{imageName}",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "tag": {
                          "description": "The tag of the container image. If not specified, this defaults to the latest tag.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Use a list of container images to start the notebook instance.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "data_disk": {
                    "block": {
                      "attributes": {
                        "auto_delete": {
                          "computed": true,
                          "description": "Optional. Specifies whether the disk will be auto-deleted\nwhen the instance is deleted (but not when the disk is\ndetached from the instance).",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "boot": {
                          "computed": true,
                          "description": "Optional. Indicates that this is a boot disk. The virtual\nmachine will use the first partition of the disk for its\nroot filesystem.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "device_name": {
                          "computed": true,
                          "description": "Optional. Specifies a unique device name of your choice\nthat is reflected into the /dev/disk/by-id/google-* tree\nof a Linux operating system running within the instance.\nThis name can be used to reference the device for mounting,\nresizing, and so on, from within the instance.\nIf not specified, the server chooses a default device name\nto apply to this disk, in the form persistent-disk-x, where\nx is a number assigned by Google Compute Engine. This field\nis only applicable for persistent disks.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "guest_os_features": {
                          "computed": true,
                          "description": "Indicates a list of features to enable on the guest operating\nsystem. Applicable only for bootable images. To see a list of\navailable features, read 'https://cloud.google.com/compute/docs/\nimages/create-delete-deprecate-private-images#guest-os-features'\noptions. ''",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "index": {
                          "computed": true,
                          "description": "Output only. A zero-based index to this disk, where 0 is\nreserved for the boot disk. If you have many disks attached\nto an instance, each disk would have a unique index number.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "interface": {
                          "description": "\"Specifies the disk interface to use for attaching this disk,\nwhich is either SCSI or NVME. The default is SCSI. Persistent\ndisks must always use SCSI and the request will fail if you attempt\nto attach a persistent disk in any other format than SCSI. Local SSDs\ncan use either NVME or SCSI. For performance characteristics of SCSI\nover NVMe, see Local SSD performance. Valid values: * NVME * SCSI\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kind": {
                          "computed": true,
                          "description": "Type of the resource. Always compute#attachedDisk for attached\ndisks.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "licenses": {
                          "computed": true,
                          "description": "Output only. Any valid publicly visible licenses.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "mode": {
                          "description": "The mode in which to attach this disk, either READ_WRITE\nor READ_ONLY. If not specified, the default is to attach\nthe disk in READ_WRITE mode.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "source": {
                          "description": "Specifies a valid partial or full URL to an existing\nPersistent Disk resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "type": {
                          "description": "Specifies the type of the disk, either SCRATCH or PERSISTENT.\nIf not specified, the default is PERSISTENT.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "initialize_params": {
                          "block": {
                            "attributes": {
                              "description": {
                                "description": "Provide this property when creating the disk.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "disk_name": {
                                "description": "Specifies the disk name. If not specified, the default is\nto use the name of the instance. If the disk with the\ninstance name exists already in the given zone/region, a\nnew name will be automatically generated.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "disk_size_gb": {
                                "description": "Specifies the size of the disk in base-2 GB. If not\nspecified, the disk will be the same size as the image\n(usually 10GB). If specified, the size must be equal to\nor larger than 10GB. Default 100 GB.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "disk_type": {
                                "description": "The type of the boot disk attached to this runtime,\ndefaults to standard persistent disk. For valid values,\nsee 'https://cloud.google.com/vertex-ai/docs/workbench/\nreference/rest/v1/projects.locations.runtimes#disktype'",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "labels": {
                                "computed": true,
                                "description": "Labels to apply to this disk. These can be later modified\nby the disks.setLabels method. This field is only\napplicable for persistent disks.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "description": "Input only. Specifies the parameters for a new disk that will\nbe created alongside the new instance. Use initialization\nparameters to create boot disks or local SSDs attached to the\nnew instance. This property is mutually exclusive with the\nsource property; you can only define one or the other, but not\nboth.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Data disk option configuration settings.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "encryption_config": {
                    "block": {
                      "attributes": {
                        "kms_key": {
                          "description": "The Cloud KMS resource identifier of the customer-managed\nencryption key used to protect a resource, such as a disks.\nIt has the following format:\n'projects/{PROJECT_ID}/locations/{REGION}/keyRings/\n{KEY_RING_NAME}/cryptoKeys/{KEY_NAME}'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Encryption settings for virtual machine data disk.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "shielded_instance_config": {
                    "block": {
                      "attributes": {
                        "enable_integrity_monitoring": {
                          "description": "Defines whether the instance has integrity monitoring enabled.\nEnables monitoring and attestation of the boot integrity of\nthe instance. The attestation is performed against the\nintegrity policy baseline. This baseline is initially derived\nfrom the implicitly trusted boot image when the instance is\ncreated. Enabled by default.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_secure_boot": {
                          "description": "Defines whether the instance has Secure Boot enabled.Secure\nBoot helps ensure that the system only runs authentic software\nby verifying the digital signature of all boot components, and\nhalting the boot process if signature verification fails.\nDisabled by default.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_vtpm": {
                          "description": "Defines whether the instance has the vTPM enabled. Enabled by\ndefault.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Shielded VM Instance configuration settings.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Virtual Machine configuration settings.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Use a Compute Engine VM image to start the managed notebook instance.",
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

func GoogleNotebooksRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNotebooksRuntime), &result)
	return &result
}
