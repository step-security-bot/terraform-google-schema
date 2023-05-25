package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the resource, provided by the client when initially creating\nthe resource. The resource name must be 1-63 characters long, and comply\nwith RFC1035. Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the\nfirst character must be a lowercase letter, and all following characters\nmust be a dash, lowercase letter, or digit, except the last character,\nwhich cannot be a dash.",
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
      "region": {
        "computed": true,
        "description": "Region where resource policy resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "snapshot_schedule_policy": {
        "block": {
          "block_types": {
            "retention_policy": {
              "block": {
                "attributes": {
                  "max_retention_days": {
                    "description": "Maximum age of the snapshot that is allowed to be kept.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "on_source_disk_delete": {
                    "description": "Specifies the behavior to apply to scheduled snapshots when\nthe source disk is deleted.\nValid options are KEEP_AUTO_SNAPSHOTS and APPLY_RETENTION_POLICY",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule": {
              "block": {
                "block_types": {
                  "daily_schedule": {
                    "block": {
                      "attributes": {
                        "days_in_cycle": {
                          "description": "The number of days between snapshots.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "start_time": {
                          "description": "This must be in UTC format that resolves to one of\n00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,\nboth 13:00-5 and 08:00 are valid.",
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
                  "hourly_schedule": {
                    "block": {
                      "attributes": {
                        "hours_in_cycle": {
                          "description": "The number of hours between snapshots.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "start_time": {
                          "description": "Time within the window to start the operations.\nIt must be in format \"HH:MM\",\nwhere HH : [00-23] and MM : [00-00] GMT.",
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
                  "weekly_schedule": {
                    "block": {
                      "block_types": {
                        "day_of_weeks": {
                          "block": {
                            "attributes": {
                              "day": {
                                "description": "The day of the week to create the snapshot. e.g. MONDAY",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "start_time": {
                                "description": "Time within the window to start the operations.\nIt must be in format \"HH:MM\", where HH : [00-23] and MM : [00-00] GMT.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 7,
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "snapshot_properties": {
              "block": {
                "attributes": {
                  "guest_flush": {
                    "description": "Whether to perform a 'guest aware' snapshot.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "labels": {
                    "description": "A set of key-value pairs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "storage_locations": {
                    "description": "GCS bucket location in which to store the snapshot (regional or multi-regional).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
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

func GoogleComputeResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeResourcePolicy), &result)
	return &result
}
