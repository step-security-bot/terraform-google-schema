package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeDisk = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource. Provide this property when\nyou create the resource.",
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
      "image": {
        "description": "The image from which to initialize this disk. This can be\none of: the image's 'self_link', 'projects/{project}/global/images/{image}',\n'projects/{project}/global/images/family/{family}', 'global/images/{image}',\n'global/images/family/{family}', 'family/{family}', '{project}/{family}',\n'{project}/{image}', '{family}', or '{image}'. If referred by family, the\nimages names must include the family name. If they don't, use the\n[google_compute_image data source](/docs/providers/google/d/compute_image.html).\nFor instance, the image 'centos-6-v20180104' includes its family name 'centos-6'.\nThese images can be referred by family name here.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this disk.  A list of key-\u003evalue pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_attach_timestamp": {
        "computed": true,
        "description": "Last attach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_detach_timestamp": {
        "computed": true,
        "description": "Last detach timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "licenses": {
        "computed": true,
        "description": "Any applicable license URI.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "physical_block_size_bytes": {
        "computed": true,
        "description": "Physical block size of the persistent disk, in bytes. If not present\nin a request, a default value is used. Currently supported sizes\nare 4096 and 16384, other sizes may be added in the future.\nIf an unsupported value is requested, the error message will list\nthe supported values for the caller's project.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_iops": {
        "computed": true,
        "description": "Indicates how many IOPS must be provisioned for the disk.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "Size of the persistent disk, specified in GB. You can specify this\nfield when creating a persistent disk using the 'image' or\n'snapshot' parameter, or specify it alone to create an empty\npersistent disk.\n\nIf you specify this field along with 'image' or 'snapshot',\nthe value must not be less than the size of the image\nor the size of the snapshot.\n\n~\u003e**NOTE** If you change the size, Terraform updates the disk size\nif upsizing is detected but recreates the disk if downsizing is requested.\nYou can add 'lifecycle.prevent_destroy' in the config to prevent destroying\nand recreating.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot": {
        "description": "The source snapshot used to create this disk. You can provide this as\na partial or full URL to the resource. If the snapshot is in another\nproject than this disk, you must supply a full URL. For example, the\nfollowing are valid values:\n\n* 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot'\n* 'projects/project/global/snapshots/snapshot'\n* 'global/snapshots/snapshot'\n* 'snapshot'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_disk": {
        "description": "The source disk used to create this disk. You can provide this as a partial or full URL to the resource.\nFor example, the following are valid values:\n\n* https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/disks/{disk}\n* https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/disks/{disk}\n* projects/{project}/zones/{zone}/disks/{disk}\n* projects/{project}/regions/{region}/disks/{disk}\n* zones/{zone}/disks/{disk}\n* regions/{region}/disks/{disk}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_disk_id": {
        "computed": true,
        "description": "The ID value of the disk used to create this image. This value may\nbe used to determine whether the image was taken from the current\nor a previous instance of a given disk name.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_image_id": {
        "computed": true,
        "description": "The ID value of the image used to create this disk. This value\nidentifies the exact image that was used to create this persistent\ndisk. For example, if you created the persistent disk from an image\nthat was later deleted and recreated under the same name, the source\nimage ID would identify the exact version of the image that was used.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_snapshot_id": {
        "computed": true,
        "description": "The unique ID of the snapshot used to create this disk. This value\nidentifies the exact snapshot that was used to create this persistent\ndisk. For example, if you created the persistent disk from a snapshot\nthat was later deleted and recreated under the same name, the source\nsnapshot ID would identify the exact version of the snapshot that was\nused.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "URL of the disk type resource describing which disk type to use to\ncreate the disk. Provide this when creating the disk.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "Links to the users of the disk (attached instances) in form:\nproject/zones/zone/instances/instance",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "zone": {
        "computed": true,
        "description": "A reference to the zone where the disk resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "disk_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to encrypt the disk. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account used for the encryption request for the given KMS key.\nIf absent, the Compute Engine Service Agent service account is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "rsa_encrypted_key": {
              "description": "Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit\ncustomer-supplied encryption key to either encrypt or decrypt\nthis resource. You can provide either the rawKey or the rsaEncryptedKey.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied\nencryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Encrypts the disk using a customer-supplied encryption key.\n\nAfter you encrypt a disk with a customer-supplied key, you must\nprovide the same key if you use the disk later (e.g. to create a disk\nsnapshot or an image, or to attach the disk to a virtual machine).\n\nCustomer-supplied encryption keys do not protect access to metadata of\nthe disk.\n\nIf you do not provide an encryption key when creating the disk, then\nthe disk will be encrypted using an automatically generated key and\nyou do not need to provide a key to use the disk later.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "guest_os_features": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: [\"MULTI_IP_SUBNET\", \"SECURE_BOOT\", \"SEV_CAPABLE\", \"UEFI_COMPATIBLE\", \"VIRTIO_SCSI_MULTIQUEUE\", \"WINDOWS\", \"GVNIC\", \"SEV_LIVE_MIGRATABLE\", \"SEV_SNP_CAPABLE\", \"SUSPEND_RESUME_COMPATIBLE\", \"TDX_CAPABLE\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A list of features to enable on the guest operating system.\nApplicable only for bootable disks.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "source_image_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to encrypt the disk. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account used for the encryption request for the given KMS key.\nIf absent, the Compute Engine Service Agent service account is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied\nencryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The customer-supplied encryption key of the source image. Required if\nthe source image is protected by a customer-supplied encryption key.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_snapshot_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key used to encrypt the disk. Also called KmsKeyName\nin the cloud console. Your project's Compute Engine System service account\n('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com') must have\n'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.\nSee https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account used for the encryption request for the given KMS key.\nIf absent, the Compute Engine Service Agent service account is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "raw_key": {
              "description": "Specifies a 256-bit customer-supplied encryption key, encoded in\nRFC 4648 base64 to either encrypt or decrypt this resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sha256": {
              "computed": true,
              "description": "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied\nencryption key that protects this resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The customer-supplied encryption key of the source snapshot. Required\nif the source snapshot is protected by a customer-supplied encryption\nkey.",
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

func GoogleComputeDiskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeDisk), &result)
	return &result
}
