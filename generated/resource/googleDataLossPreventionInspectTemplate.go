package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionInspectTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of the inspect template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User set display name of the inspect template.",
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
        "computed": true,
        "description": "The resource name of the inspect template. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the inspect template in any of the following formats:\n\n* 'projects/{{project}}'\n* 'projects/{{project}}/locations/{{location}}'\n* 'organizations/{{organization_id}}'\n* 'organizations/{{organization_id}}/locations/{{location}}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description": "The template id can contain uppercase and lowercase letters, numbers, and hyphens;\nthat is, it must match the regular expression: [a-zA-Z\\d-_]+. The maximum length is\n100 characters. Can be empty to allow the system to generate one.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "inspect_config": {
        "block": {
          "attributes": {
            "content_options": {
              "description": "List of options defining data content to scan. If empty, text, images, and other content will be included. Possible values: [\"CONTENT_TEXT\", \"CONTENT_IMAGE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
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
                          "description": "Version name for this InfoType.",
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
                    "required": true,
                    "type": "number"
                  },
                  "max_findings_per_request": {
                    "description": "Max number of findings that will be returned per request/job. The maximum returned is 2000.",
                    "description_kind": "plain",
                    "required": true,
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
                          "required": true,
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
                                "description": "Version name for this InfoType.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does\nnot have an infoType, the DLP API applies the limit against all infoTypes that are found but not\nspecified in another InfoTypeLimit.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
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
                          "description": "Version name for this InfoType.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "List of infoTypes this rule set is applied to.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
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
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Regular expression pattern defining what qualifies as a hotword.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "proximity": {
                                      "block": {
                                        "attributes": {
                                          "window_after": {
                                            "description": "Number of characters after the finding to consider.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "window_before": {
                                            "description": "Number of characters before the finding to consider.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          }
                                        },
                                        "description": "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot\nexceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be\nused to match substrings of the finding itself. For example, the certainty of a phone number regex\n'(\\d{3}) \\d{3}-\\d{4}' could be adjusted upwards if the area code is known to be the local area code of a company\noffice using the hotword regex '(xxx)', where 'xxx' is the area code in question.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Drop if the hotword rule is contained in the proximate context.\nFor tabular data, the context includes the column name.",
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
                                            "description": "Version name for this InfoType.",
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
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "Regular expression pattern defining what qualifies as a hotword.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
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
                                "min_items": 1,
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
                                "min_items": 1,
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

func GoogleDataLossPreventionInspectTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionInspectTemplate), &result)
	return &result
}
