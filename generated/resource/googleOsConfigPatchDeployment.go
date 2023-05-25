package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOsConfigPatchDeployment = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the patch deployment was created. Timestamp is in RFC3339 text format.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the patch deployment. Length of the description is limited to 1024 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
        "description": "Duration of the patch. After the duration ends, the patch times out.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\"",
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
      "last_execute_time": {
        "computed": true,
        "description": "The last time a patch job was started by this deployment. Timestamp is in RFC3339 text format.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Unique name for the patch deployment resource in a project.\nThe patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.",
        "description_kind": "plain",
        "type": "string"
      },
      "patch_deployment_id": {
        "description": "A name for the patch deployment in the project. When creating a name the following rules apply:\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the project.",
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
      "update_time": {
        "computed": true,
        "description": "Time the patch deployment was last updated. Timestamp is in RFC3339 text format.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "instance_filter": {
        "block": {
          "attributes": {
            "all": {
              "description": "Target all VM instances in the project. If true, no other criteria is permitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "instance_name_prefixes": {
              "description": "Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group\nVMs when targeting configs, for example prefix=\"prod-\".",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "instances": {
              "description": "Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}',\n'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or\n'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "zones": {
              "description": "Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "group_labels": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Compute Engine instance labels that must be present for a VM instance to be targeted by this filter",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "VM instances to patch.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "one_time_schedule": {
        "block": {
          "attributes": {
            "execute_time": {
              "description": "The desired patch job execution time. A timestamp in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Schedule a one-time execution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "patch_config": {
        "block": {
          "attributes": {
            "reboot_config": {
              "description": "Post-patch reboot settings. Possible values: [\"DEFAULT\", \"ALWAYS\", \"NEVER\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "apt": {
              "block": {
                "attributes": {
                  "excludes": {
                    "description": "List of packages to exclude from update. These packages will be excluded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusive_packages": {
                    "description": "An exclusive list of packages to be updated. These are the only packages that will be updated.\nIf these packages are not installed, they will be ignored. This field cannot be specified with\nany other patch configuration fields.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "type": {
                    "description": "By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: [\"DIST\", \"UPGRADE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Apt update settings. Use this setting to override the default apt patch rules.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "goo": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "goo update settings. Use this setting to override the default goo patch rules.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "goo update settings. Use this setting to override the default goo patch rules.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "post_step": {
              "block": {
                "block_types": {
                  "linux_exec_step_config": {
                    "block": {
                      "attributes": {
                        "allowed_success_codes": {
                          "description": "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "interpreter": {
                          "description": "The script interpreter to use to run the script. If no interpreter is specified the script will\nbe executed directly, which will likely only succeed for scripts with shebang lines. Possible values: [\"SHELL\", \"POWERSHELL\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_path": {
                          "description": "An absolute path to the executable on the VM.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "gcs_object": {
                          "block": {
                            "attributes": {
                              "bucket": {
                                "description": "Bucket of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "generation_number": {
                                "description": "Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "object": {
                                "description": "Name of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A Cloud Storage object containing the executable.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The ExecStepConfig for all Linux VMs targeted by the PatchJob.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "windows_exec_step_config": {
                    "block": {
                      "attributes": {
                        "allowed_success_codes": {
                          "description": "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "interpreter": {
                          "description": "The script interpreter to use to run the script. If no interpreter is specified the script will\nbe executed directly, which will likely only succeed for scripts with shebang lines. Possible values: [\"SHELL\", \"POWERSHELL\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_path": {
                          "description": "An absolute path to the executable on the VM.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "gcs_object": {
                          "block": {
                            "attributes": {
                              "bucket": {
                                "description": "Bucket of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "generation_number": {
                                "description": "Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "object": {
                                "description": "Name of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A Cloud Storage object containing the executable.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The ExecStepConfig for all Windows VMs targeted by the PatchJob.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The ExecStep to run after the patch update.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "pre_step": {
              "block": {
                "block_types": {
                  "linux_exec_step_config": {
                    "block": {
                      "attributes": {
                        "allowed_success_codes": {
                          "description": "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "interpreter": {
                          "description": "The script interpreter to use to run the script. If no interpreter is specified the script will\nbe executed directly, which will likely only succeed for scripts with shebang lines. Possible values: [\"SHELL\", \"POWERSHELL\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_path": {
                          "description": "An absolute path to the executable on the VM.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "gcs_object": {
                          "block": {
                            "attributes": {
                              "bucket": {
                                "description": "Bucket of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "generation_number": {
                                "description": "Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "object": {
                                "description": "Name of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A Cloud Storage object containing the executable.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The ExecStepConfig for all Linux VMs targeted by the PatchJob.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "windows_exec_step_config": {
                    "block": {
                      "attributes": {
                        "allowed_success_codes": {
                          "description": "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        },
                        "interpreter": {
                          "description": "The script interpreter to use to run the script. If no interpreter is specified the script will\nbe executed directly, which will likely only succeed for scripts with shebang lines. Possible values: [\"SHELL\", \"POWERSHELL\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "local_path": {
                          "description": "An absolute path to the executable on the VM.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "gcs_object": {
                          "block": {
                            "attributes": {
                              "bucket": {
                                "description": "Bucket of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "generation_number": {
                                "description": "Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "object": {
                                "description": "Name of the Cloud Storage object.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A Cloud Storage object containing the executable.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The ExecStepConfig for all Windows VMs targeted by the PatchJob.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The ExecStep to run before the patch update.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "windows_update": {
              "block": {
                "attributes": {
                  "classifications": {
                    "description": "Only apply updates of these windows update classifications. If empty, all updates are applied. Possible values: [\"CRITICAL\", \"SECURITY\", \"DEFINITION\", \"DRIVER\", \"FEATURE_PACK\", \"SERVICE_PACK\", \"TOOL\", \"UPDATE_ROLLUP\", \"UPDATE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "excludes": {
                    "description": "List of KBs to exclude from update.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusive_patches": {
                    "description": "An exclusive list of kbs to be updated. These are the only patches that will be updated.\nThis field must not be used with other patch configurations.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Windows update settings. Use this setting to override the default Windows patch rules.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "yum": {
              "block": {
                "attributes": {
                  "excludes": {
                    "description": "List of packages to exclude from update. These packages will be excluded.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusive_packages": {
                    "description": "An exclusive list of packages to be updated. These are the only packages that will be updated.\nIf these packages are not installed, they will be ignored. This field cannot be specified with\nany other patch configuration fields.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "minimal": {
                    "description": "Will cause patch to run yum update-minimal instead.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "security": {
                    "description": "Adds the --security flag to yum update. Not supported on all platforms.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Yum update settings. Use this setting to override the default yum patch rules.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "zypper": {
              "block": {
                "attributes": {
                  "categories": {
                    "description": "Install only patches with these categories. Common categories include security, recommended, and feature.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "excludes": {
                    "description": "List of packages to exclude from update.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusive_patches": {
                    "description": "An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.\nThis field must not be used with any other patch configuration fields.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "severities": {
                    "description": "Install only patches with these severities. Common severities include critical, important, moderate, and low.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "with_optional": {
                    "description": "Adds the --with-optional flag to zypper patch.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "with_update": {
                    "description": "Adds the --with-update flag, to zypper patch.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "zypper update settings. Use this setting to override the default zypper patch rules.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Patch configuration that is applied.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "recurring_schedule": {
        "block": {
          "attributes": {
            "end_time": {
              "description": "The end time at which a recurring patch deployment schedule is no longer active.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_execute_time": {
              "computed": true,
              "description": "The time the last patch job ran successfully.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "type": "string"
            },
            "next_execute_time": {
              "computed": true,
              "description": "The time the next patch job is scheduled to run.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "type": "string"
            },
            "start_time": {
              "description": "The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.\nA timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "monthly": {
              "block": {
                "attributes": {
                  "month_day": {
                    "description": "One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.\nMonths without the target day will be skipped. For example, a schedule to run \"every month on the 31st\"\nwill not run in February, April, June, etc.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "week_day_of_month": {
                    "block": {
                      "attributes": {
                        "day_of_week": {
                          "description": "A day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "week_ordinal": {
                          "description": "Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Week day in a month.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Schedule with monthly executions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "time_of_day": {
              "block": {
                "attributes": {
                  "hours": {
                    "description": "Hours of day in 24 hour format. Should be from 0 to 23.\nAn API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
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
                    "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Time of the day to run a recurring deployment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "time_zone": {
              "block": {
                "attributes": {
                  "id": {
                    "description": "IANA Time Zone Database time zone, e.g. \"America/New_York\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description": "IANA Time Zone Database version number, e.g. \"2019a\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are\ndetermined by the chosen time zone.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "weekly": {
              "block": {
                "attributes": {
                  "day_of_week": {
                    "description": "IANA Time Zone Database time zone, e.g. \"America/New_York\". Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Schedule with weekly executions.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Schedule recurring executions.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "rollout": {
        "block": {
          "attributes": {
            "mode": {
              "description": "Mode of the patch rollout. Possible values: [\"ZONE_BY_ZONE\", \"CONCURRENT_ZONES\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "disruption_budget": {
              "block": {
                "attributes": {
                  "fixed": {
                    "description": "Specifies a fixed value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "percentage": {
                    "description": "Specifies the relative value defined as a percentage, which will be multiplied by a reference value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up.\nDuring patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps.\nA VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget.\nFor zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone.\nFor example, if the disruption budget has a fixed value of 10, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Rollout strategy of the patch job.",
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

func GoogleOsConfigPatchDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOsConfigPatchDeployment), &result)
	return &result
}
