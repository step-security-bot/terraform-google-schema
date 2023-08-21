package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsManagedZone = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time that this resource was created on the server.\nThis is in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
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
      "force_destroy": {
        "description": "Set this true to delete all records in the zone.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "managed_zone_id": {
        "computed": true,
        "description": "Unique identifier for the resource; defined by the server.",
        "description_kind": "plain",
        "type": "number"
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
        "description": "The zone's visibility: public zones are exposed to the Internet,\nwhile private zones are visible only to Virtual Private Cloud resources. Default value: \"public\" Possible values: [\"private\", \"public\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cloud_logging_config": {
        "block": {
          "attributes": {
            "enable_logging": {
              "description": "If set, enable query logging for this ManagedZone. False by default, making logging opt-in.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "Cloud logging configuration",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
              "description": "Specifies the mechanism used to provide authenticated denial-of-existence responses.\nnon_existence can only be updated when the state is 'off'. Possible values: [\"nsec\", \"nsec3\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "state": {
              "description": "Specifies whether DNSSEC is enabled, and what mode it is in Possible values: [\"off\", \"on\", \"transfer\"]",
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
                    "description": "String mnemonic specifying the DNSSEC algorithm of this key Possible values: [\"ecdsap256sha256\", \"ecdsap384sha384\", \"rsasha1\", \"rsasha256\", \"rsasha512\"]",
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
                    "description": "Specifies whether this is a key signing key (KSK) or a zone\nsigning key (ZSK). Key signing keys have the Secure Entry\nPoint flag set and, when active, will only be used to sign\nresource record sets of type DNSKEY. Zone signing keys do\nnot have the Secure Entry Point flag set and will be used\nto sign all other types of resource record sets. Possible values: [\"keySigning\", \"zoneSigning\"]",
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
                "description": "Specifies parameters that will be used for generating initial DnsKeys\nfor this ManagedZone. If you provide a spec for keySigning or zoneSigning,\nyou must also provide one for the other.\ndefault_key_specs can only be updated when the state is 'off'.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "DNSSEC configuration",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "forwarding_config": {
        "block": {
          "block_types": {
            "target_name_servers": {
              "block": {
                "attributes": {
                  "forwarding_path": {
                    "description": "Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding\ndecision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go\nto the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: [\"default\", \"private\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ipv4_address": {
                    "description": "IPv4 address of a target name server.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "List of target name servers to forward to. Cloud DNS will\nselect the best available name server if more than\none target is given.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "The presence for this field indicates that outbound forwarding is enabled\nfor this zone. The value of this field contains the set of destinations\nto forward to.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "peering_config": {
        "block": {
          "block_types": {
            "target_network": {
              "block": {
                "attributes": {
                  "network_url": {
                    "description": "The id or fully qualified URL of the VPC network to forward queries to.\nThis should be formatted like 'projects/{project}/global/networks/{network}' or\n'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The network with which to peer.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The presence of this field indicates that DNS Peering is enabled for this\nzone. The value of this field contains the network to peer with.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_visibility_config": {
        "block": {
          "block_types": {
            "gke_clusters": {
              "block": {
                "attributes": {
                  "gke_cluster_name": {
                    "description": "The resource name of the cluster to bind this ManagedZone to.\nThis should be specified in the format like\n'projects/*/locations/*/clusters/*'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of Google Kubernetes Engine clusters that can see this zone.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "networks": {
              "block": {
                "attributes": {
                  "network_url": {
                    "description": "The id or fully qualified URL of the VPC network to bind to.\nThis should be formatted like 'projects/{project}/global/networks/{network}' or\n'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of VPC networks that can see this zone. Until the provider updates to use the Terraform 0.12 SDK in a future release, you\nmay experience issues with this resource while updating. If you've defined a 'networks' block and\nadd another 'networks' block while keeping the old block, Terraform will see an incorrect diff\nand apply an incorrect update to the resource. If you encounter this issue, remove all 'networks'\nblocks in an update and then apply another update adding all of them back simultaneously.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "For privately visible zones, the set of Virtual Private Cloud\nresources that the zone is visible from. At least one of 'gke_clusters' or 'networks' must be specified.",
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
