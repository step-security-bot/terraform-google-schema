package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLookerInstance = `{
  "block": {
    "attributes": {
      "consumer_network": {
        "description": "Network name in the consumer project in the format of: projects/{project}/global/networks/{network}\nNote that the consumer network may be in a different GCP project than the consumer\nproject that is hosting the Looker Instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the instance was created in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_public_ip": {
        "computed": true,
        "description": "Public Egress IP (IPv4).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ingress_private_ip": {
        "computed": true,
        "description": "Private Ingress IP (IPv4).",
        "description_kind": "plain",
        "type": "string"
      },
      "ingress_public_ip": {
        "computed": true,
        "description": "Public Ingress IP (IPv4).",
        "description_kind": "plain",
        "type": "string"
      },
      "looker_uri": {
        "computed": true,
        "description": "Looker instance URI which can be used to access the Looker Instance UI.",
        "description_kind": "plain",
        "type": "string"
      },
      "looker_version": {
        "computed": true,
        "description": "The Looker version that the instance is using.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The ID of the instance or a fully qualified identifier for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform_edition": {
        "description": "Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:\n- LOOKER_CORE_TRIAL: trial instance\n- LOOKER_CORE_STANDARD: pay as you go standard instance\n- LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance\n- LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance\n- LOOKER_CORE_EMBED_ANNUAL: subscription embed instance\n- LOOKER_MODELER: standalone modeling service Default value: \"LOOKER_CORE_TRIAL\" Possible values: [\"LOOKER_CORE_TRIAL\", \"LOOKER_CORE_STANDARD\", \"LOOKER_CORE_STANDARD_ANNUAL\", \"LOOKER_CORE_ENTERPRISE_ANNUAL\", \"LOOKER_CORE_EMBED_ANNUAL\", \"LOOKER_MODELER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_enabled": {
        "description": "Whether private IP is enabled on the Looker instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_enabled": {
        "description": "Whether public IP is enabled on the Looker instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "region": {
        "computed": true,
        "description": "The name of the Looker region of the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_range": {
        "description": "Name of a reserved IP address range within the consumer network, to be used for\nprivate service access connection. User may or may not specify this in a request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the instance was updated in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "admin_settings": {
        "block": {
          "attributes": {
            "allowed_email_domains": {
              "description": "Email domain allowlist for the instance.\n\nDefine the email domains to which your users can deliver Looker (Google Cloud core) content.\nUpdating this list will restart the instance. Updating the allowed email domains from terraform\nmeans the value provided will be considered as the entire list and not an amendment to the\nexisting list of allowed email domains.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Looker instance Admin settings.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "deny_maintenance_period": {
        "block": {
          "block_types": {
            "end_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0\nto specify a year by itself or a year and month where the day isn't significant.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a\nmonth and day.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without\na year.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Required. Start date of the deny maintenance period",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "start_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0\nto specify a year by itself or a year and month where the day isn't significant.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "month": {
                    "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a\nmonth and day.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "year": {
                    "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without\na year.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Required. Start date of the deny maintenance period",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "time": {
              "block": {
                "attributes": {
                  "hours": {
                    "description": "Hours of day in 24 hour format. Should be from 0 to 23.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "Minutes of hour of day. Must be from 0 to 59.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "nanos": {
                    "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Seconds of minutes of the time. Must normally be from 0 to 59.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Required. Start time of the window in UTC time.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Maintenance denial period for this instance.\n\nYou must allow at least 14 days of maintenance availability\nbetween any two deny maintenance periods.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "Name of the customer managed encryption key (CMEK) in KMS.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_name_version": {
              "computed": true,
              "description": "Full name and version of the CMEK key currently in use to encrypt Looker data.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_state": {
              "computed": true,
              "description": "Status of the customer managed encryption key (CMEK) in KMS.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Looker instance encryption settings.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_window": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description": "Required. Day of the week for this MaintenanceWindow (in UTC).\n\n- MONDAY: Monday\n- TUESDAY: Tuesday\n- WEDNESDAY: Wednesday\n- THURSDAY: Thursday\n- FRIDAY: Friday\n- SATURDAY: Saturday\n- SUNDAY: Sunday Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "start_time": {
              "block": {
                "attributes": {
                  "hours": {
                    "description": "Hours of day in 24 hour format. Should be from 0 to 23.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "Minutes of hour of day. Must be from 0 to 59.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "nanos": {
                    "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description": "Seconds of minutes of the time. Must normally be from 0 to 59.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Required. Start time of the window in UTC time.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Maintenance window for an instance.\n\nMaintenance of your instance takes place once a month, and will require\nyour instance to be restarted during updates, which will temporarily\ndisrupt service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oauth_config": {
        "block": {
          "attributes": {
            "client_id": {
              "description": "The client ID for the Oauth config.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_secret": {
              "description": "The client secret for the Oauth config.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Looker Instance OAuth login settings.",
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
      "user_metadata": {
        "block": {
          "attributes": {
            "additional_developer_user_count": {
              "description": "Number of additional Developer Users to allocate to the Looker Instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "additional_standard_user_count": {
              "description": "Number of additional Standard Users to allocate to the Looker Instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "additional_viewer_user_count": {
              "description": "Number of additional Viewer Users to allocate to the Looker Instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Metadata about users for a Looker instance.\n\nThese settings are only available when platform edition LOOKER_CORE_STANDARD is set.\n\nThere are ten Standard and two Developer users included in the cost of the product.\nYou can allocate additional Standard, Viewer, and Developer users for this instance.\nIt is an optional step and can be modified later.\n\nWith the Standard edition of Looker (Google Cloud core), you can provision up to 50\ntotal users, distributed across Viewer, Standard, and Developer.",
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

func GoogleLookerInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLookerInstance), &result)
	return &result
}
