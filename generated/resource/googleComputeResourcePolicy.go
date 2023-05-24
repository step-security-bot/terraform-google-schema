package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeResourcePolicy = `{
  "block": {
    "attributes": {
      "description": {
        "description": "An optional description of this resource. Provide this property when you create the resource.",
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
      "group_placement_policy": {
        "block": {
          "attributes": {
            "availability_domain_count": {
              "description": "The number of availability domains instances will be spread across. If two instances are in different\navailability domain, they will not be put in the same low latency network",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "collocation": {
              "description": "Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.\nSpecify 'COLLOCATED' to enable collocation. Can only be specified with 'vm_count'. If compute instances are created\nwith a COLLOCATED policy, then exactly 'vm_count' instances must be created at the same time with the resource policy\nattached. Possible values: [\"COLLOCATED\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vm_count": {
              "description": "Number of VMs in this placement group. Google does not recommend that you use this field\nunless you use a compact policy and you want your policy to work only if it contains this\nexact number of VMs.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description": "Resource policy for instances used for placement configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "instance_schedule_policy": {
        "block": {
          "attributes": {
            "expiration_time": {
              "description": "The expiration time of the schedule. The timestamp is an RFC3339 string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "description": "The start time of the schedule. The timestamp is an RFC3339 string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_zone": {
              "description": "Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name\nfrom the tz database: http://en.wikipedia.org/wiki/Tz_database.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "vm_start_schedule": {
              "block": {
                "attributes": {
                  "schedule": {
                    "description": "Specifies the frequency for the operation, using the unix-cron format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the schedule for starting instances.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "vm_stop_schedule": {
              "block": {
                "attributes": {
                  "schedule": {
                    "description": "Specifies the frequency for the operation, using the unix-cron format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Specifies the schedule for stopping instances.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Resource policy for scheduling instance operations.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
                    "description": "Specifies the behavior to apply to scheduled snapshots when\nthe source disk is deleted. Default value: \"KEEP_AUTO_SNAPSHOTS\" Possible values: [\"KEEP_AUTO_SNAPSHOTS\", \"APPLY_RETENTION_POLICY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Retention policy applied to snapshots created by this resource policy.",
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
                      "description": "The policy will execute every nth day at the specified time.",
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
                          "description": "Time within the window to start the operations.\nIt must be in an hourly format \"HH:MM\",\nwhere HH : [00-23] and MM : [00] GMT.\neg: 21:00",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "The policy will execute every nth hour starting at the specified time.",
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
                                "description": "The day of the week to create the snapshot. e.g. MONDAY Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
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
                            "description": "May contain up to seven (one for each day of the week) snapshot times.",
                            "description_kind": "plain"
                          },
                          "max_items": 7,
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description": "Allows specifying a snapshot time for each day of the week.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "snapshot_properties": {
              "block": {
                "attributes": {
                  "chain_name": {
                    "description": "Creates the new snapshot in the snapshot chain labeled with the\nspecified name. The chain name must be 1-63 characters long and comply\nwith RFC1035.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
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
                    "description": "Cloud Storage bucket location to store the auto snapshot\n(regional or multi-regional)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "Properties with which the snapshots are created, such as labels.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Policy for creating snapshots of persistent disks.",
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
