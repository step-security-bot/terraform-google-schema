package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsManagedZone = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A textual description field. Defaults to 'Managed by Terraform'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name": {
        "description": "The DNS name of this managed zone, for instance \"example.com.\".",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "A set of key/value label pairs to assign to this ManagedZone.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "User assigned name for this resource.\nMust be unique within the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name_servers": {
        "computed": true,
        "description": "Delegate your managed_zone to these virtual name servers;\ndefined by the server",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "visibility": {
        "description": "The zone's visibility: public zones are exposed to the Internet,\nwhile private zones are visible only to Virtual Private Cloud resources.\nMust be one of: 'public', 'private'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "dnssec_config": {
        "block": {
          "attributes": {
            "kind": {
              "description": "Identifies what kind of resource this is",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "non_existence": {
              "computed": true,
              "description": "Specifies the mechanism used to provide authenticated denial-of-existence responses.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description": "Specifies whether DNSSEC is enabled, and what mode it is in",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_key_specs": {
              "block": {
                "attributes": {
                  "algorithm": {
                    "description": "String mnemonic specifying the DNSSEC algorithm of this key",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_length": {
                    "description": "Length of the keys in bits",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "key_type": {
                    "description": "Specifies whether this is a key signing key (KSK) or a zone\nsigning key (ZSK). Key signing keys have the Secure Entry\nPoint flag set and, when active, will only be used to sign\nresource record sets of type DNSKEY. Zone signing keys do\nnot have the Secure Entry Point flag set and will be used\nto sign all other types of resource record sets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kind": {
                    "description": "Identifies what kind of resource this is",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_visibility_config": {
        "block": {
          "block_types": {
            "networks": {
              "block": {
                "attributes": {
                  "network_url": {
                    "description": "The fully qualified URL of the VPC network to bind to.\nThis should be formatted like\n'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
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

func GoogleDnsManagedZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsManagedZone), &result)
	return &result
}
