package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeInstance = `{
  "block": {
    "attributes": {
      "allow_stopping_for_update": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "attached_disk": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "device_name": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "kms_key_self_link": "string",
              "mode": "string",
              "source": "string"
            }
          ]
        ]
      },
      "boot_disk": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_delete": "bool",
              "device_name": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "initialize_params": [
                "list",
                [
                  "object",
                  {
                    "image": "string",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "size": "number",
                    "type": "string"
                  }
                ]
              ],
              "kms_key_self_link": "string",
              "mode": "string",
              "source": "string"
            }
          ]
        ]
      },
      "can_ip_forward": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cpu_platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_delete": "bool",
              "device_name": "string",
              "disk": "string",
              "disk_encryption_key_raw": "string",
              "disk_encryption_key_sha256": "string",
              "image": "string",
              "scratch": "bool",
              "size": "number",
              "type": "string"
            }
          ]
        ]
      },
      "enable_display": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "guest_accelerator": {
        "computed": true,
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
      "hostname": {
        "computed": true,
        "description_kind": "plain",
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
        "description_kind": "plain",
        "type": "string"
      },
      "label_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "metadata_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_startup_script": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "min_cpu_platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface": {
        "computed": true,
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
                    "assigned_nat_ip": "string",
                    "nat_ip": "string",
                    "network_tier": "string",
                    "public_ptr_domain_name": "string"
                  }
                ]
              ],
              "address": "string",
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
              "name": "string",
              "network": "string",
              "network_ip": "string",
              "subnetwork": "string",
              "subnetwork_project": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduling": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_restart": "bool",
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
              "preemptible": "bool"
            }
          ]
        ]
      },
      "scratch_disk": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "interface": "string"
            }
          ]
        ]
      },
      "self_link": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account": {
        "computed": true,
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
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
