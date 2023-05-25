package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionStoredInfoType = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of the info type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User set display name of the info type.",
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
        "description": "The resource name of the info type. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the info type in any of the following formats:\n\n* 'projects/{{project}}'\n* 'projects/{{project}}/locations/{{location}}'\n* 'organizations/{{organization_id}}'\n* 'organizations/{{organization_id}}/locations/{{location}}'",
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
      "large_custom_dictionary": {
        "block": {
          "block_types": {
            "big_query_field": {
              "block": {
                "block_types": {
                  "field": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "Name describing the field.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Designated field in the BigQuery table.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "table": {
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
                      "description": "Field in a BigQuery table where each cell represents a dictionary phrase.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Field in a BigQuery table where each cell represents a dictionary phrase.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "cloud_storage_file_set": {
              "block": {
                "attributes": {
                  "url": {
                    "description": "The url, in the format 'gs://\u003cbucket\u003e/\u003cpath\u003e'. Trailing wildcard in the path is allowed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Set of files containing newline-delimited lists of dictionary phrases.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "output_path": {
              "block": {
                "attributes": {
                  "path": {
                    "description": "A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.\nIf any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Dictionary which defines the rule.",
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

func GoogleDataLossPreventionStoredInfoTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionStoredInfoType), &result)
	return &result
}
