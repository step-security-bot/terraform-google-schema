package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionJobTrigger = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The creation timestamp of an inspectTemplate. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the job trigger.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User set display name of the job trigger.",
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
      "last_run_time": {
        "computed": true,
        "description": "The timestamp of the last time this trigger executed.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the job trigger. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the trigger, either in the format 'projects/{{project}}'\nor 'projects/{{project}}/locations/{{location}}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description": "Whether the trigger is currently active. Default value: \"HEALTHY\" Possible values: [\"PAUSED\", \"HEALTHY\", \"CANCELLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of an inspectTemplate. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "inspect_job": {
        "block": {
          "attributes": {
            "inspect_template_name": {
              "description": "The name of the template to run when this job is triggered.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "deidentify": {
                    "block": {
                      "attributes": {
                        "cloud_storage_output": {
                          "description": "User settable Cloud Storage bucket and folders to store de-identified files.\n\nThis field must be set for cloud storage deidentification.\n\nThe output Cloud Storage bucket must be different from the input bucket.\n\nDe-identified files will overwrite files in the output path.\n\nForm of: gs://bucket/folder/ or gs://bucket",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "file_types_to_transform": {
                          "description": "List of user-specified file type groups to transform. If specified, only the files with these filetypes will be transformed.\n\nIf empty, all supported files will be transformed. Supported types may be automatically added over time.\n\nIf a file type is set in this field that isn't supported by the Deidentify action then the job will fail and will not be successfully created/started. Possible values: [\"IMAGE\", \"TEXT_FILE\", \"CSV\", \"TSV\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "transformation_config": {
                          "block": {
                            "attributes": {
                              "deidentify_template": {
                                "description": "If this template is specified, it will serve as the default de-identify template.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "image_redact_template": {
                                "description": "If this template is specified, it will serve as the de-identify template for images.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "structured_deidentify_template": {
                                "description": "If this template is specified, it will serve as the de-identify template for structured content such as delimited files and tables.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "User specified deidentify templates and configs for structured, unstructured, and image files.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "transformation_details_storage_config": {
                          "block": {
                            "block_types": {
                              "table": {
                                "block": {
                                  "attributes": {
                                    "dataset_id": {
                                      "description": "The ID of the dataset containing this table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "project_id": {
                                      "description": "The ID of the project containing this table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "table_id": {
                                      "description": "The ID of the table. The ID must contain only letters (a-z,\nA-Z), numbers (0-9), or underscores (_). The maximum length\nis 1,024 characters.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The BigQuery table in which to store the output.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Config for storing transformation details.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Create a de-identified copy of the requested table or files.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "job_notification_emails": {
                    "block": {
                      "description": "Sends an email when the job completes. The email goes to IAM project owners and technical Essential Contacts.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "pub_sub": {
                    "block": {
                      "attributes": {
                        "topic": {
                          "description": "Cloud Pub/Sub topic to send notifications to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Publish a message into a given Pub/Sub topic when the job completes.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "publish_findings_to_cloud_data_catalog": {
                    "block": {
                      "description": "Publish findings of a DlpJob to Data Catalog.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "publish_summary_to_cscc": {
                    "block": {
                      "description": "Publish the result summary of a DlpJob to the Cloud Security Command Center.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "publish_to_stackdriver": {
                    "block": {
                      "description": "Enable Stackdriver metric dlp.googleapis.com/findingCount.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "save_findings": {
                    "block": {
                      "block_types": {
                        "output_config": {
                          "block": {
                            "attributes": {
                              "output_schema": {
                                "description": "Schema used for writing the findings for Inspect jobs. This field is only used for\nInspect and must be unspecified for Risk jobs. Columns are derived from the Finding\nobject. If appending to an existing table, any columns from the predefined schema\nthat are missing will be added. No columns in the existing table will be deleted.\n\nIf unspecified, then all available columns will be used for a new table or an (existing)\ntable with no schema, and no changes will be made to an existing table that has a schema.\nOnly for use with external storage. Possible values: [\"BASIC_COLUMNS\", \"GCS_COLUMNS\", \"DATASTORE_COLUMNS\", \"BIG_QUERY_COLUMNS\", \"ALL_COLUMNS\"]",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "table": {
                                "block": {
                                  "attributes": {
                                    "dataset_id": {
                                      "description": "Dataset ID of the table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "project_id": {
                                      "description": "The Google Cloud Platform project ID of the project containing the table.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "table_id": {
                                      "description": "Name of the table. If is not set a new one will be generated for you with the following format:\n'dlp_googleapis_yyyy_mm_dd_[dlp_job_id]'. Pacific timezone will be used for generating the date details.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Information on the location of the target BigQuery Table.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Information on where to store output",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "If set, the detailed findings will be persisted to the specified OutputStorageConfig. Only a single instance of this action can be specified. Compatible with: Inspect, Risk",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "A task to execute on the completion of a job.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "inspect_config": {
              "block": {
                "attributes": {
                  "exclude_info_types": {
                    "description": "When true, excludes type information of the findings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "include_quote": {
                    "description": "When true, a contextual quote from the data that triggered a finding is included in the response.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "min_likelihood": {
                    "description": "Only returns findings equal or above this threshold. See https://cloud.google.com/dlp/docs/likelihood for more info Default value: \"POSSIBLE\" Possible values: [\"VERY_UNLIKELY\", \"UNLIKELY\", \"POSSIBLE\", \"LIKELY\", \"VERY_LIKELY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "custom_info_types": {
                    "block": {
                      "attributes": {
                        "exclusion_type": {
                          "description": "If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching. Possible values: [\"EXCLUSION_TYPE_EXCLUDE\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "likelihood": {
                          "description": "Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria\nspecified by the rule. Default value: \"VERY_LIKELY\" Possible values: [\"VERY_UNLIKELY\", \"UNLIKELY\", \"POSSIBLE\", \"LIKELY\", \"VERY_LIKELY\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "dictionary": {
                          "block": {
                            "block_types": {
                              "cloud_storage_path": {
                                "block": {
                                  "attributes": {
                                    "path": {
                                      "description": "A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "word_list": {
                                "block": {
                                  "attributes": {
                                    "words": {
                                      "description": "Words or phrases defining the dictionary. The dictionary must contain at least one\nphrase and every phrase must contain at least 2 characters that are letters or digits.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "List of words or phrases to search for.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Dictionary which defines the rule.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "info_type": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names\nlisted at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description": "Version of the information type to use. By default, the version is set to stable.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing\ninfoTypes and that infoType is specified in 'info_types' field. Specifying the latter adds findings to the\none detected by the system. If built-in info type is not specified in 'info_types' list then the name is\ntreated as a custom info type.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "regex": {
                          "block": {
                            "attributes": {
                              "group_indexes": {
                                "description": "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              },
                              "pattern": {
                                "description": "Pattern defining the regular expression.\nIts syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Regular expression which defines the rule.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "stored_type": {
                          "block": {
                            "attributes": {
                              "create_time": {
                                "computed": true,
                                "description": "The creation timestamp of an inspectTemplate. Set by the server.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "name": {
                                "description": "Resource name of the requested StoredInfoType, for example 'organizations/433245324/storedInfoTypes/432452342'\nor 'projects/project-id/storedInfoTypes/432452342'.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A reference to a StoredInfoType to use with scanning.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "surrogate_type": {
                          "block": {
                            "description": "Message for detecting output from deidentification transformations that support reversing.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Custom info types to be used. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "info_types": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed\nat https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version": {
                          "description": "Version of the information type to use. By default, the version is set to stable",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Restricts what infoTypes to look for. The values must correspond to InfoType values returned by infoTypes.list\nor listed at https://cloud.google.com/dlp/docs/infotypes-reference.\n\nWhen no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run.\nBy default this may be all types, but may change over time as detectors are updated.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "limits": {
                    "block": {
                      "attributes": {
                        "max_findings_per_item": {
                          "description": "Max number of findings that will be returned for each item scanned. The maximum returned is 2000.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_findings_per_request": {
                          "description": "Max number of findings that will be returned per request/job. The maximum returned is 2000.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "max_findings_per_info_type": {
                          "block": {
                            "attributes": {
                              "max_findings": {
                                "description": "Max findings limit for the given infoType.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "info_type": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed\nat https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "version": {
                                      "description": "Version of the information type to use. By default, the version is set to stable",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does\nnot have an infoType, the DLP API applies the limit against all infoTypes that are found but not\nspecified in another InfoTypeLimit.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Configuration of findings limit given for specified infoTypes.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration to control the number of findings returned.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "rule_set": {
                    "block": {
                      "block_types": {
                        "info_types": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed\nat https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description": "Version of the information type to use. By default, the version is set to stable.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "List of infoTypes this rule set is applied to.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "rules": {
                          "block": {
                            "block_types": {
                              "exclusion_rule": {
                                "block": {
                                  "attributes": {
                                    "matching_type": {
                                      "description": "How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: [\"MATCHING_TYPE_FULL_MATCH\", \"MATCHING_TYPE_PARTIAL_MATCH\", \"MATCHING_TYPE_INVERSE_MATCH\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "dictionary": {
                                      "block": {
                                        "block_types": {
                                          "cloud_storage_path": {
                                            "block": {
                                              "attributes": {
                                                "path": {
                                                  "description": "A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "word_list": {
                                            "block": {
                                              "attributes": {
                                                "words": {
                                                  "description": "Words or phrases defining the dictionary. The dictionary must contain at least one\nphrase and every phrase must contain at least 2 characters that are letters or digits.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description": "List of words or phrases to search for.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Dictionary which defines the rule.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "exclude_by_hotword": {
                                      "block": {
                                        "block_types": {
                                          "hotword_regex": {
                                            "block": {
                                              "attributes": {
                                                "group_indexes": {
                                                  "description": "The index of the submatch to extract as findings. When not specified,\nthe entire match is returned. No more than 3 may be included.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "number"
                                                  ]
                                                },
                                                "pattern": {
                                                  "description": "Pattern defining the regular expression. Its syntax\n(https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Regular expression pattern defining what qualifies as a hotword.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "proximity": {
                                            "block": {
                                              "attributes": {
                                                "window_after": {
                                                  "description": "Number of characters after the finding to consider. Either this or window_before must be specified",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "window_before": {
                                                  "description": "Number of characters before the finding to consider. Either this or window_after must be specified",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot\nexceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be\nused to match substrings of the finding itself. For example, the certainty of a phone number regex\n'(\\d{3}) \\d{3}-\\d{4}' could be adjusted upwards if the area code is known to be the local area code of a company\noffice using the hotword regex '(xxx)', where 'xxx' is the area code in question.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Drop if the hotword rule is contained in the proximate context.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "exclude_info_types": {
                                      "block": {
                                        "block_types": {
                                          "info_types": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed\nat https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "version": {
                                                  "description": "Version of the information type to use. By default, the version is set to stable.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "If a finding is matched by any of the infoType detectors listed here, the finding will be excluded from the scan results.",
                                              "description_kind": "plain"
                                            },
                                            "min_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Set of infoTypes for which findings would affect this rule.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "regex": {
                                      "block": {
                                        "attributes": {
                                          "group_indexes": {
                                            "description": "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "number"
                                            ]
                                          },
                                          "pattern": {
                                            "description": "Pattern defining the regular expression.\nIts syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Regular expression which defines the rule.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "The rule that specifies conditions when findings of infoTypes specified in InspectionRuleSet are removed from results.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "hotword_rule": {
                                "block": {
                                  "block_types": {
                                    "hotword_regex": {
                                      "block": {
                                        "attributes": {
                                          "group_indexes": {
                                            "description": "The index of the submatch to extract as findings. When not specified,\nthe entire match is returned. No more than 3 may be included.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "number"
                                            ]
                                          },
                                          "pattern": {
                                            "description": "Pattern defining the regular expression. Its syntax\n(https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Regular expression pattern defining what qualifies as a hotword.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "likelihood_adjustment": {
                                      "block": {
                                        "attributes": {
                                          "fixed_likelihood": {
                                            "description": "Set the likelihood of a finding to a fixed value. Either this or relative_likelihood can be set. Possible values: [\"VERY_UNLIKELY\", \"UNLIKELY\", \"POSSIBLE\", \"LIKELY\", \"VERY_LIKELY\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "relative_likelihood": {
                                            "description": "Increase or decrease the likelihood by the specified number of levels. For example,\nif a finding would be POSSIBLE without the detection rule and relativeLikelihood is 1,\nthen it is upgraded to LIKELY, while a value of -1 would downgrade it to UNLIKELY.\nLikelihood may never drop below VERY_UNLIKELY or exceed VERY_LIKELY, so applying an\nadjustment of 1 followed by an adjustment of -1 when base likelihood is VERY_LIKELY\nwill result in a final likelihood of LIKELY. Either this or fixed_likelihood can be set.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description": "Likelihood adjustment to apply to all matching findings.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "proximity": {
                                      "block": {
                                        "attributes": {
                                          "window_after": {
                                            "description": "Number of characters after the finding to consider. Either this or window_before must be specified",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "window_before": {
                                            "description": "Number of characters before the finding to consider. Either this or window_after must be specified",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description": "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot\nexceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be\nused to match substrings of the finding itself. For example, the certainty of a phone number regex\n'(\\d{3}) \\d{3}-\\d{4}' could be adjusted upwards if the area code is known to be the local area code of a company\noffice using the hotword regex '(xxx)', where 'xxx' is the area code in question.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Hotword-based detection rule.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Set of rules to be applied to infoTypes. The rules are applied in order.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end,\nother rules are executed in the order they are specified for each info type.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The core content of the template.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "storage_config": {
              "block": {
                "block_types": {
                  "big_query_options": {
                    "block": {
                      "attributes": {
                        "rows_limit": {
                          "description": "Max number of rows to scan. If the table has more rows than this value, the rest of the rows are omitted.\nIf not set, or if set to 0, all rows will be scanned. Only one of rowsLimit and rowsLimitPercent can be\nspecified. Cannot be used in conjunction with TimespanConfig.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "rows_limit_percent": {
                          "description": "Max percentage of rows to scan. The rest are omitted. The number of rows scanned is rounded down.\nMust be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of\nrowsLimit and rowsLimitPercent can be specified. Cannot be used in conjunction with TimespanConfig.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "sample_method": {
                          "description": "How to sample rows if not all rows are scanned. Meaningful only when used in conjunction with either\nrowsLimit or rowsLimitPercent. If not specified, rows are scanned in the order BigQuery reads them. Default value: \"TOP\" Possible values: [\"TOP\", \"RANDOM_START\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "identifying_fields": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of a BigQuery field to be returned with the findings.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Specifies the BigQuery fields that will be returned with findings.\nIf not specified, no identifying fields will be returned for findings.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "table_reference": {
                          "block": {
                            "attributes": {
                              "dataset_id": {
                                "description": "The dataset ID of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "The Google Cloud Platform project ID of the project containing the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "table_id": {
                                "description": "The name of the table.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Set of files to scan.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining BigQuery table and row identifiers.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "cloud_storage_options": {
                    "block": {
                      "attributes": {
                        "bytes_limit_per_file": {
                          "description": "Max number of bytes to scan from a file. If a scanned file's size is bigger than this value\nthen the rest of the bytes are omitted.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "bytes_limit_per_file_percent": {
                          "description": "Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down.\nMust be between 0 and 100, inclusively. Both 0 and 100 means no limit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "file_types": {
                          "description": "List of file type groups to include in the scan. If empty, all files are scanned and available data\nformat processors are applied. In addition, the binary content of the selected files is always scanned as well.\nImages are scanned only as binary if the specified region does not support image inspection and no fileTypes were specified. Possible values: [\"BINARY_FILE\", \"TEXT_FILE\", \"IMAGE\", \"WORD\", \"PDF\", \"AVRO\", \"CSV\", \"TSV\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "files_limit_percent": {
                          "description": "Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down.\nMust be between 0 and 100, inclusively. Both 0 and 100 means no limit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "sample_method": {
                          "description": "How to sample bytes if not all bytes are scanned. Meaningful only when used in conjunction with bytesLimitPerFile.\nIf not specified, scanning would start from the top. Possible values: [\"TOP\", \"RANDOM_START\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "file_set": {
                          "block": {
                            "attributes": {
                              "url": {
                                "description": "The Cloud Storage url of the file(s) to scan, in the format 'gs://\u003cbucket\u003e/\u003cpath\u003e'. Trailing wildcard\nin the path is allowed.\n\nIf the url ends in a trailing slash, the bucket or directory represented by the url will be scanned\nnon-recursively (content in sub-directories will not be scanned). This means that 'gs://mybucket/' is\nequivalent to 'gs://mybucket/*', and 'gs://mybucket/directory/' is equivalent to 'gs://mybucket/directory/*'.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "regex_file_set": {
                                "block": {
                                  "attributes": {
                                    "bucket_name": {
                                      "description": "The name of a Cloud Storage bucket.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "exclude_regex": {
                                      "description": "A list of regular expressions matching file paths to exclude. All files in the bucket that match at\nleast one of these regular expressions will be excluded from the scan.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "include_regex": {
                                      "description": "A list of regular expressions matching file paths to include. All files in the bucket\nthat match at least one of these regular expressions will be included in the set of files,\nexcept for those that also match an item in excludeRegex. Leaving this field empty will\nmatch all files by default (this is equivalent to including .* in the list)",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "description": "The regex-filtered set of files to scan.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Set of files to scan.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining a file or a set of files within a Google Cloud Storage bucket.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "datastore_options": {
                    "block": {
                      "block_types": {
                        "kind": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The name of the Datastore kind.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "A representation of a Datastore kind.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "partition_id": {
                          "block": {
                            "attributes": {
                              "namespace_id": {
                                "description": "If not empty, the ID of the namespace to which the entities belong.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "project_id": {
                                "description": "The ID of the project to which the entities belong.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Datastore partition ID. A partition ID identifies a grouping of entities. The grouping\nis always by project and namespace, however the namespace ID may be empty.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Options defining a data set within Google Cloud Datastore.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "hybrid_options": {
                    "block": {
                      "attributes": {
                        "description": {
                          "description": "A short description of where the data is coming from. Will be stored once in the job. 256 max length.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "labels": {
                          "description": "To organize findings, these labels will be added to each finding.\n\nLabel keys must be between 1 and 63 characters long and must conform to the following regular expression: '[a-z]([-a-z0-9]*[a-z0-9])?'.\n\nLabel values must be between 0 and 63 characters long and must conform to the regular expression '([a-z]([-a-z0-9]*[a-z0-9])?)?'.\n\nNo more than 10 labels can be associated with a given finding.\n\nExamples:\n* '\"environment\" : \"production\"'\n* '\"pipeline\" : \"etl\"'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "required_finding_label_keys": {
                          "description": "These are labels that each inspection request must include within their 'finding_labels' map. Request\nmay contain others, but any missing one of these will be rejected.\n\nLabel keys must be between 1 and 63 characters long and must conform to the following regular expression: '[a-z]([-a-z0-9]*[a-z0-9])?'.\n\nNo more than 10 keys can be required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "table_options": {
                          "block": {
                            "block_types": {
                              "identifying_fields": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name describing the field.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The columns that are the primary keys for table objects included in ContentItem. A copy of this\ncell's value will stored alongside alongside each finding so that the finding can be traced to\nthe specific row it came from. No more than 3 may be provided.",
                                  "description_kind": "plain"
                                },
                                "nesting_mode": "list"
                              }
                            },
                            "description": "If the container is a table, additional information to make findings meaningful such as the columns that are primary keys.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration to control jobs where the content being inspected is outside of Google Cloud Platform.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "timespan_config": {
                    "block": {
                      "attributes": {
                        "enable_auto_population_of_timespan_config": {
                          "description": "When the job is started by a JobTrigger we will automatically figure out a valid startTime to avoid\nscanning files that have not been modified since the last time the JobTrigger executed. This will\nbe based on the time of the execution of the last run of the JobTrigger.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "end_time": {
                          "description": "Exclude files or rows newer than this value. If set to zero, no upper time limit is applied.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start_time": {
                          "description": "Exclude files or rows older than this value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "timestamp_field": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery.\n\nFor BigQuery: Required to filter out rows based on the given start and end times. If not specified and the table was\nmodified between the given start and end times, the entire table will be scanned. The valid data types of the timestamp\nfield are: INTEGER, DATE, TIMESTAMP, or DATETIME BigQuery column.\n\nFor Datastore. Valid data types of the timestamp field are: TIMESTAMP. Datastore entity will be scanned if the\ntimestamp property does not exist or its value is empty or invalid.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "Information on where to inspect",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Information on where to inspect",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Information on where to inspect",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Controls what and how to inspect for findings.",
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
      "triggers": {
        "block": {
          "block_types": {
            "manual": {
              "block": {
                "description": "For use with hybrid jobs. Jobs must be manually created and finished.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule": {
              "block": {
                "attributes": {
                  "recurrence_period_duration": {
                    "description": "With this option a job is started a regular periodic basis. For example: every day (86400 seconds).\n\nA scheduled start time will be skipped if the previous execution has not ended when its scheduled time occurs.\n\nThis value must be set to a time duration greater than or equal to 1 day and can be no longer than 60 days.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Schedule for triggered jobs",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "What event needs to occur for a new job to be started.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDataLossPreventionJobTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionJobTrigger), &result)
	return &result
}
