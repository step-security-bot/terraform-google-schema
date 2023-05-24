package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSnapshot = `{
  "block": {
    "attributes": {
      "chain_name": {
        "computed": true,
        "description": "Creates the new snapshot in the snapshot chain labeled with the\nspecified name. The chain name must be 1-63 characters long and\ncomply with RFC1035. This is an uncommon option only for advanced\nservice owners who needs to create separate snapshot chains, for\nexample, for chargeback tracking.  When you describe your snapshot\nresource, this field is visible only if it has a non-empty value.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "disk_size_gb": {
        "computed": true,
        "description": "Size of the snapshot, specified in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "filter": {
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
        "computed": true,
        "description": "Labels to apply to this Snapshot.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "licenses": {
        "computed": true,
        "description": "A list of public visible licenses that apply to this snapshot. This\ncan be because the original image had licenses attached (such as a\nWindows image).  snapshotEncryptionKey nested object Encrypts the\nsnapshot using a customer-supplied encryption key.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "Name of the resource; provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_encryption_key": {
        "computed": true,
        "description": "Encrypts the snapshot using a customer-supplied encryption key.\n\nAfter you encrypt a snapshot using a customer-supplied key, you must\nprovide the same key if you use the snapshot later. For example, you\nmust provide the encryption key when you create a disk from the\nencrypted snapshot in a future request.\n\nCustomer-supplied encryption keys do not protect access to metadata of\nthe snapshot.\n\nIf you do not provide an encryption key when creating the snapshot,\nthen the snapshot will be encrypted using an automatically generated\nkey and you do not need to provide a key to use the snapshot later.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_self_link": "string",
              "kms_key_service_account": "string",
              "raw_key": "string",
              "sha256": "string"
            }
          ]
        ]
      },
      "snapshot_id": {
        "computed": true,
        "description": "The unique identifier for the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_disk": {
        "computed": true,
        "description": "A reference to the disk used to create this snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_disk_encryption_key": {
        "computed": true,
        "description": "The customer-supplied encryption key of the source snapshot. Required\nif the source snapshot is protected by a customer-supplied encryption\nkey.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_service_account": "string",
              "raw_key": "string"
            }
          ]
        ]
      },
      "storage_bytes": {
        "computed": true,
        "description": "A size of the storage used by the snapshot. As snapshots share\nstorage, this number is expected to change with snapshot\ncreation/deletion.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_locations": {
        "computed": true,
        "description": "Cloud Storage bucket storage location of the snapshot (regional or multi-regional).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "zone": {
        "computed": true,
        "description": "A reference to the zone where the disk is hosted.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSnapshot), &result)
	return &result
}
