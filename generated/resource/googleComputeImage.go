package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeImage = `{
  "block": {
    "attributes": {
      "archive_size_bytes": {
        "computed": true,
        "description": "Size of the image tar.gz archive stored in Google Cloud Storage (in\nbytes).",
        "description_kind": "plain",
        "type": "number"
      },
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
      "disk_size_gb": {
        "computed": true,
        "description": "Size of the image when restored onto a persistent disk (in GB).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "family": {
        "description": "The name of the image family to which this image belongs. You can\ncreate disks by specifying an image family instead of a specific\nimage name. The image family always returns its latest image that is\nnot deprecated. The name of the image family must comply with\nRFC1035.",
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
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource. Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this Image.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means\nthe first character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the\nlast character, which cannot be a dash.",
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
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk": {
        "description": "The source disk to create this image based on.\nYou must provide either this property or the\nrawDisk.source property but not both to create an image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_image": {
        "description": "URL of the source image used to create this image. In order to create an image, you must provide the full or partial\nURL of one of the following:\n\n* The selfLink URL\n* This property\n* The rawDisk.source URL\n* The sourceDisk URL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_snapshot": {
        "description": "URL of the source snapshot used to create this image.\n\nIn order to create an image, you must provide the full or partial URL of one of the following:\n\n* The selfLink URL\n* This property\n* The sourceImage URL\n* The rawDisk.source URL\n* The sourceDisk URL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_locations": {
        "computed": true,
        "description": "Cloud Storage bucket storage location of the image\n(regional or multi-regional).\nReference link: https://cloud.google.com/compute/docs/reference/rest/v1/images",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
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
          "description": "A list of features to enable on the guest operating system.\nApplicable only for bootable images.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "image_encryption_key": {
        "block": {
          "attributes": {
            "kms_key_self_link": {
              "description": "The self link of the encryption key that is stored in Google Cloud\nKMS.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_service_account": {
              "description": "The service account being used for the encryption request for the\ngiven KMS key. If absent, the Compute Engine default service\naccount is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Encrypts the image using a customer-supplied encryption key.\n\nAfter you encrypt an image with a customer-supplied key, you must\nprovide the same key if you use the image later (e.g. to create a\ndisk from the image)",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "raw_disk": {
        "block": {
          "attributes": {
            "container_type": {
              "description": "The format used to encode and transmit the block device, which\nshould be TAR. This is just a container and transmission format\nand not a runtime format. Provided by the client when the disk\nimage is created. Default value: \"TAR\" Possible values: [\"TAR\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sha1": {
              "description": "An optional SHA1 checksum of the disk image before unpackaging.\nThis is provided by the client when the disk image is created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "description": "The full Google Cloud Storage URL where disk storage is stored\nYou must provide either this property or the sourceDisk property\nbut not both.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The parameters of the raw disk image.",
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

func GoogleComputeImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeImage), &result)
	return &result
}
