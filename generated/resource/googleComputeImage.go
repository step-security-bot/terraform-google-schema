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
      }
    },
    "block_types": {
      "guest_os_features": {
        "block": {
          "attributes": {
            "type": {
              "description": "The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "raw_disk": {
        "block": {
          "attributes": {
            "container_type": {
              "description": "The format used to encode and transmit the block device, which\nshould be TAR. This is just a container and transmission format\nand not a runtime format. Provided by the client when the disk\nimage is created.",
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
