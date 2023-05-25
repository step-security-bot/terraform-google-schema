package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNotebooksInstance = `{
  "block": {
    "attributes": {
      "boot_disk_size_gb": {
        "description": "The size of the boot disk in GB attached to this instance,\nup to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.\nIf not specified, this defaults to 100.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "boot_disk_type": {
        "description": "Possible disk types for notebook instances. Possible values: [\"DISK_TYPE_UNSPECIFIED\", \"PD_STANDARD\", \"PD_SSD\", \"PD_BALANCED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Instance creation time",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_gpu_driver_path": {
        "description": "Specify a custom Cloud Storage path where the GPU driver is stored.\nIf not specified, we'll automatically choose from official GPU drivers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_disk_size_gb": {
        "description": "The size of the data disk in GB attached to this instance,\nup to a maximum of 64000 GB (64 TB).\nYou can choose the size of the data disk based on how big your notebooks and data are.\nIf not specified, this defaults to 100.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_disk_type": {
        "description": "Possible disk types for notebook instances. Possible values: [\"DISK_TYPE_UNSPECIFIED\", \"PD_STANDARD\", \"PD_SSD\", \"PD_BALANCED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disk_encryption": {
        "description": "Disk encryption method used on the boot and data disks, defaults to GMEK. Possible values: [\"DISK_ENCRYPTION_UNSPECIFIED\", \"GMEK\", \"CMEK\"]",
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
      "install_gpu_driver": {
        "description": "Whether the end user authorizes Google Cloud to install GPU driver\non this instance. If this field is empty or set to false, the GPU driver\nwon't be installed. Only applicable to instances with GPUs.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance_owners": {
        "description": "The list of owners of this instance after creation.\nFormat: alias@example.com.\nCurrently supports one owner only.\nIf not specified, all of the service account users of\nyour VM instance's service account can use the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "kms_key": {
        "description": "The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.\nFormat: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Labels to apply to this instance. These can be later modified by the setLabels method.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "A reference to the zone where the machine resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "machine_type": {
        "description": "A reference to a machine type which defines VM kind.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "description": "Custom metadata to apply to this instance.\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "The name specified for the Notebook instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The name of the VPC that this instance is in.\nFormat: projects/{project_id}/global/networks/{network_id}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "no_proxy_access": {
        "description": "The notebook instance will not register with the proxy..",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "no_public_ip": {
        "description": "No public IP will be assigned to this instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "no_remove_data_disk": {
        "description": "If true, the data disk will not be auto deleted when deleting the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "post_startup_script": {
        "description": "Path to a Bash script that automatically runs after a\nnotebook instance fully boots up. The path must be a URL\nor Cloud Storage path (gs://path-to-file/file-name).",
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
      "proxy_uri": {
        "computed": true,
        "description": "The proxy endpoint that is used to access the Jupyter notebook.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_account": {
        "computed": true,
        "description": "The service account on this instance, giving access to other\nGoogle Cloud services. You can use any service account within\nthe same project, but you must have the service account user\npermission to use the instance. If not specified,\nthe Compute Engine default service account is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_scopes": {
        "description": "Optional. The URIs of service account scopes to be included in Compute Engine instances.\nIf not specified, the following scopes are defined:\n- https://www.googleapis.com/auth/cloud-platform\n- https://www.googleapis.com/auth/userinfo.email",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "The state of this instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet": {
        "computed": true,
        "description": "The name of the subnet that this instance is in.\nFormat: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description": "The Compute Engine tags to add to runtime.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Instance update time.",
        "description_kind": "plain",
        "optional": true,
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
              "required": true,
              "type": "number"
            },
            "type": {
              "description": "Type of this accelerator. Possible values: [\"ACCELERATOR_TYPE_UNSPECIFIED\", \"NVIDIA_TESLA_K80\", \"NVIDIA_TESLA_P100\", \"NVIDIA_TESLA_V100\", \"NVIDIA_TESLA_P4\", \"NVIDIA_TESLA_T4\", \"NVIDIA_TESLA_T4_VWS\", \"NVIDIA_TESLA_P100_VWS\", \"NVIDIA_TESLA_P4_VWS\", \"NVIDIA_TESLA_A100\", \"TPU_V2\", \"TPU_V3\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The hardware accelerator used on this instance. If you use accelerators,\nmake sure that your configuration has enough vCPUs and memory to support the\nmachineType you have selected.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "container_image": {
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
          "description": "Use a container image to start the notebook instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "shielded_instance_config": {
        "block": {
          "attributes": {
            "enable_integrity_monitoring": {
              "description": "Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the\nboot integrity of the instance. The attestation is performed against the integrity policy baseline.\nThis baseline is initially derived from the implicitly trusted boot image when the instance is created.\nEnabled by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_secure_boot": {
              "description": "Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs\nauthentic software by verifying the digital signature of all boot components, and halting the boot process\nif signature verification fails.\nDisabled by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_vtpm": {
              "description": "Defines whether the instance has the vTPM enabled.\nEnabled by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "A set of Shielded Instance options. Check [Images using supported Shielded VM features]\nNot all combinations are valid",
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
      "vm_image": {
        "block": {
          "attributes": {
            "image_family": {
              "description": "Use this VM image family to find the image; the newest image in this family will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_name": {
              "description": "Use VM image name to find the image.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project": {
              "description": "The name of the Google Cloud project that this VM image belongs to.\nFormat: projects/{project_id}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Use a Compute Engine VM image to start the notebook instance.",
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

func GoogleNotebooksInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNotebooksInstance), &result)
	return &result
}
