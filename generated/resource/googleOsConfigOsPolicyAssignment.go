package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOsConfigOsPolicyAssignment = `{
  "block": {
    "attributes": {
      "baseline": {
        "computed": true,
        "description": "Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision.\nFor a given OS policy assignment, there is only one revision with a value of 'true' for this field.",
        "description_kind": "plain",
        "type": "bool"
      },
      "deleted": {
        "computed": true,
        "description": "Output only. Indicates that this revision deletes the OS policy assignment.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "description": "OS policy assignment description. Length of the description is limited to 1024 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. Indicates that reconciliation is in progress for the revision. This value is 'true' when the 'rollout_state' is one of:\n* IN_PROGRESS\n* CANCELLING",
        "description_kind": "plain",
        "type": "bool"
      },
      "revision_create_time": {
        "computed": true,
        "description": "Output only. The timestamp that the revision was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment",
        "description_kind": "plain",
        "type": "string"
      },
      "rollout_state": {
        "computed": true,
        "description": "Output only. OS policy assignment rollout state",
        "description_kind": "plain",
        "type": "string"
      },
      "skip_await_rollout": {
        "description": "Set to true to skip awaiting rollout during resource creation and update.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "uid": {
        "computed": true,
        "description": "Output only. Server generated unique id for the OS policy assignment resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "instance_filter": {
        "block": {
          "attributes": {
            "all": {
              "description": "Target all VMs in the project. If true, no other criteria is permitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "exclusion_labels": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Labels are identified by key/value pairs in this map. A VM should contain all the key/value pairs specified in this map to be selected.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "List of label sets used for VM exclusion.\nIf the list has more than one label set, the VM is excluded if any of the label sets are applicable for the VM.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "inclusion_labels": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Labels are identified by key/value pairs in this map. A VM should contain all the key/value pairs specified in this map to be selected.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description": "List of label sets used for VM inclusion.\nIf the list has more than one 'LabelSet', the VM is included if any of the label sets are applicable for the VM.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "inventories": {
              "block": {
                "attributes": {
                  "os_short_name": {
                    "description": "The OS short name",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "os_version": {
                    "description": "The OS version Prefix matches are supported if asterisk(*) is provided as the last character. For example, to match all versions with a major version of '7', specify the following value for this field '7.*' An empty string matches all OS versions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of inventories to select VMs.\nA VM is selected if its inventory data matches at least one of the following inventories.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Filter to select VMs.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "os_policies": {
        "block": {
          "attributes": {
            "allow_no_resource_group_match": {
              "description": "This flag determines the OS policy compliance status when none of the resource groups within the policy are applicable for a VM. Set this value to 'true' if the policy needs to be reported as compliant even if the policy has nothing to validate or enforce.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "description": "Policy description. Length of the description is limited to 1024 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description": "The id of the OS policy with the following restrictions:\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the assignment.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mode": {
              "description": "Policy mode Possible values: [\"MODE_UNSPECIFIED\", \"VALIDATION\", \"ENFORCEMENT\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "resource_groups": {
              "block": {
                "block_types": {
                  "inventory_filters": {
                    "block": {
                      "attributes": {
                        "os_short_name": {
                          "description": "The OS short name",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "os_version": {
                          "description": "The OS version\nPrefix matches are supported if asterisk(*) is provided as the last character. For example, to match all versions with a major version of '7', specify the following value for this field '7.*'\nAn empty string matches all OS versions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of inventory filters for the resource group.\nThe resources in this resource group are applied to the target VM if it satisfies at least one of the following inventory filters.\nFor example, to apply this resource group to VMs running either 'RHEL' or 'CentOS' operating systems, specify 2 items for the list with following values: inventory_filters[0].os_short_name='rhel' and inventory_filters[1].os_short_name='centos'\nIf the list is empty, this resource group will be applied to the target VM unconditionally.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "resources": {
                    "block": {
                      "attributes": {
                        "id": {
                          "description": "The id of the resource with the following restrictions:\n* Must contain only lowercase letters, numbers, and hyphens.\n* Must start with a letter.\n* Must be between 1-63 characters.\n* Must end with a number or a letter.\n* Must be unique within the OS policy.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "exec": {
                          "block": {
                            "block_types": {
                              "enforce": {
                                "block": {
                                  "attributes": {
                                    "args": {
                                      "description": "Optional arguments to pass to the source during execution.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "interpreter": {
                                      "description": "The script interpreter to use. Possible values: [\"INTERPRETER_UNSPECIFIED\", \"NONE\", \"SHELL\", \"POWERSHELL\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "output_file_path": {
                                      "description": "Only recorded for enforce Exec. Path to an output file (that is created by this Exec) whose content will be recorded in OSPolicyResourceCompliance after a successful run. Absence or failure to read this file will result in this ExecResource being non-compliant. Output file size is limited to 100K bytes.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "script": {
                                      "description": "An inline script. The size of the script is limited to 1024 characters.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "file": {
                                      "block": {
                                        "attributes": {
                                          "allow_insecure": {
                                            "description": "Defaults to false. When false, files are subject to validations based on the file type: Remote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "local_path": {
                                            "description": "A local path within the VM to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "gcs": {
                                            "block": {
                                              "attributes": {
                                                "bucket": {
                                                  "description": "Bucket of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "generation": {
                                                  "description": "Generation number of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "object": {
                                                  "description": "Name of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A Cloud Storage object.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "remote": {
                                            "block": {
                                              "attributes": {
                                                "sha256_checksum": {
                                                  "description": "SHA256 checksum of the remote file.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A generic remote file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A remote or local file.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "What to run to bring this resource into the desired state. An exit code of 100 indicates \"success\", any other exit code indicates a failure running enforce.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "validate": {
                                "block": {
                                  "attributes": {
                                    "args": {
                                      "description": "Optional arguments to pass to the source during execution.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "interpreter": {
                                      "description": "The script interpreter to use. Possible values: [\"INTERPRETER_UNSPECIFIED\", \"NONE\", \"SHELL\", \"POWERSHELL\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "output_file_path": {
                                      "description": "Only recorded for enforce Exec. Path to an output file (that is created by this Exec) whose content will be recorded in OSPolicyResourceCompliance after a successful run. Absence or failure to read this file will result in this ExecResource being non-compliant. Output file size is limited to 100K bytes.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "script": {
                                      "description": "An inline script. The size of the script is limited to 1024 characters.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "file": {
                                      "block": {
                                        "attributes": {
                                          "allow_insecure": {
                                            "description": "Defaults to false. When false, files are subject to validations based on the file type:\nRemote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "local_path": {
                                            "description": "A local path within the VM to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "gcs": {
                                            "block": {
                                              "attributes": {
                                                "bucket": {
                                                  "description": "Bucket of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "generation": {
                                                  "description": "Generation number of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "object": {
                                                  "description": "Name of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A Cloud Storage object.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "remote": {
                                            "block": {
                                              "attributes": {
                                                "sha256_checksum": {
                                                  "description": "SHA256 checksum of the remote file.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A generic remote file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A remote or local file.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "What to run to validate this resource is in the desired state. An exit code of 100 indicates \"in desired state\", and exit code of 101 indicates \"not in desired state\". Any other exit code indicates a failure running validate.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Exec resource",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "file": {
                          "block": {
                            "attributes": {
                              "content": {
                                "description": "A a file with this content. The size of the content is limited to 1024 characters.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "description": "The absolute path of the file within the VM.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "permissions": {
                                "computed": true,
                                "description": "Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one bit corresponds to the execute permission. Default behavior is 755.\nBelow are some examples of permissions and their associated values: read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "state": {
                                "description": "Desired state of the file. Possible values: [\"DESIRED_STATE_UNSPECIFIED\", \"PRESENT\", \"ABSENT\", \"CONTENTS_MATCH\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "file": {
                                "block": {
                                  "attributes": {
                                    "allow_insecure": {
                                      "description": "Defaults to false. When false, files are subject to validations based on the file type: Remote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "local_path": {
                                      "description": "A local path within the VM to use.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "gcs": {
                                      "block": {
                                        "attributes": {
                                          "bucket": {
                                            "description": "Bucket of the Cloud Storage object.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "generation": {
                                            "description": "Generation number of the Cloud Storage object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "object": {
                                            "description": "Name of the Cloud Storage object.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A Cloud Storage object.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "remote": {
                                      "block": {
                                        "attributes": {
                                          "sha256_checksum": {
                                            "description": "SHA256 checksum of the remote file.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "uri": {
                                            "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A generic remote file.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A remote or local source.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "File resource",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "pkg": {
                          "block": {
                            "attributes": {
                              "desired_state": {
                                "description": "The desired state the agent should maintain for this package. Possible values: [\"DESIRED_STATE_UNSPECIFIED\", \"INSTALLED\", \"REMOVED\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "apt": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Package name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A package managed by Apt.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "deb": {
                                "block": {
                                  "attributes": {
                                    "pull_deps": {
                                      "description": "Whether dependencies should also be installed. - install when false: 'dpkg -i package' - install when true: 'apt-get update \u0026\u0026 apt-get -y install package.deb'",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "source": {
                                      "block": {
                                        "attributes": {
                                          "allow_insecure": {
                                            "description": "Defaults to false. When false, files are subject to validations based on the file type:\nRemote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "local_path": {
                                            "description": "A local path within the VM to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "gcs": {
                                            "block": {
                                              "attributes": {
                                                "bucket": {
                                                  "description": "Bucket of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "generation": {
                                                  "description": "Generation number of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "object": {
                                                  "description": "Name of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A Cloud Storage object.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "remote": {
                                            "block": {
                                              "attributes": {
                                                "sha256_checksum": {
                                                  "description": "SHA256 checksum of the remote file.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A generic remote file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "A deb package.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "A deb package file.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "googet": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Package name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A package managed by GooGet.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "msi": {
                                "block": {
                                  "attributes": {
                                    "properties": {
                                      "description": "Additional properties to use during installation. This should be in the format of Property=Setting. Appended to the defaults of 'ACTION=INSTALL REBOOT=ReallySuppress'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "source": {
                                      "block": {
                                        "attributes": {
                                          "allow_insecure": {
                                            "description": "Defaults to false. When false, files are subject to validations based on the file type:\nRemote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "local_path": {
                                            "description": "A local path within the VM to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "gcs": {
                                            "block": {
                                              "attributes": {
                                                "bucket": {
                                                  "description": "Bucket of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "generation": {
                                                  "description": "Generation number of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "object": {
                                                  "description": "Name of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A Cloud Storage object.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "remote": {
                                            "block": {
                                              "attributes": {
                                                "sha256_checksum": {
                                                  "description": "SHA256 checksum of the remote file.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A generic remote file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The MSI package.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "An MSI package.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "rpm": {
                                "block": {
                                  "attributes": {
                                    "pull_deps": {
                                      "description": "Whether dependencies should also be installed. - install when false: 'rpm --upgrade --replacepkgs package.rpm' - install when true: 'yum -y install package.rpm' or 'zypper -y install package.rpm'",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "source": {
                                      "block": {
                                        "attributes": {
                                          "allow_insecure": {
                                            "description": "Defaults to false. When false, files are subject to validations based on the file type:\nRemote: A checksum must be specified. Cloud Storage: An object generation number must be specified.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "local_path": {
                                            "description": "A local path within the VM to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "gcs": {
                                            "block": {
                                              "attributes": {
                                                "bucket": {
                                                  "description": "Bucket of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "generation": {
                                                  "description": "Generation number of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "object": {
                                                  "description": "Name of the Cloud Storage object.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A Cloud Storage object.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "remote": {
                                            "block": {
                                              "attributes": {
                                                "sha256_checksum": {
                                                  "description": "SHA256 checksum of the remote file.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "uri": {
                                                  "description": "URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "A generic remote file.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "An rpm package.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "An rpm package file.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "yum": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Package name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A package managed by YUM.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "zypper": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Package name.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A package managed by Zypper.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Package resource",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "repository": {
                          "block": {
                            "block_types": {
                              "apt": {
                                "block": {
                                  "attributes": {
                                    "archive_type": {
                                      "description": "Type of archive files in this repository. Possible values: [\"ARCHIVE_TYPE_UNSPECIFIED\", \"DEB\", \"DEB_SRC\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "components": {
                                      "description": "List of components for this repository. Must contain at least one item.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "distribution": {
                                      "description": "Distribution of this repository.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "gpg_key": {
                                      "description": "URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "uri": {
                                      "description": "URI for this repository.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "An Apt Repository.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "goo": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "The name of the repository.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "url": {
                                      "description": "The url of the repository.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A Goo Repository.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "yum": {
                                "block": {
                                  "attributes": {
                                    "base_url": {
                                      "description": "The location of the repository directory.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "display_name": {
                                      "description": "The display name of the repository.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "gpg_keys": {
                                      "description": "URIs of GPG keys.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "id": {
                                      "description": "A one word, unique name for this repository. This is the 'repo id' in the yum config file and also the 'display_name' if 'display_name' is omitted. This id is also used as the unique identifier when checking for resource conflicts.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A Yum Repository.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "zypper": {
                                "block": {
                                  "attributes": {
                                    "base_url": {
                                      "description": "The location of the repository directory.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "display_name": {
                                      "description": "The display name of the repository.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "gpg_keys": {
                                      "description": "URIs of GPG keys.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "id": {
                                      "description": "A one word, unique name for this repository. This is the 'repo id' in the zypper config file and also the 'display_name' if 'display_name' is omitted. This id is also used as the unique identifier when checking for GuestPolicy conflicts.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "A Zypper Repository.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Package repository resource",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "List of resources configured for this resource group. The resources are executed in the exact order specified here.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "List of resource groups for the policy. For a particular VM, resource groups are evaluated in the order specified and the first resource group that is applicable is selected and the rest are ignored.\nIf none of the resource groups are applicable for a VM, the VM is considered to be non-compliant w.r.t this policy. This behavior can be toggled by the flag 'allow_no_resource_group_match'",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of OS policies to be applied to the VMs.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "rollout": {
        "block": {
          "attributes": {
            "min_wait_duration": {
              "description": "This determines the minimum duration of time to wait after the configuration changes are applied through the current rollout. A VM continues to count towards the 'disruption_budget' at least until this duration of time has passed after configuration changes are applied.",
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
                  "percent": {
                    "description": "Specifies the relative value defined as a percentage, which will be multiplied by a reference value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The maximum number (or percentage) of VMs per zone to disrupt at any given moment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleOsConfigOsPolicyAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOsConfigOsPolicyAssignment), &result)
	return &result
}
