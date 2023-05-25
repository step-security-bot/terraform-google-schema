package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageTransferJob = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "last_modification_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "schedule": {
        "block": {
          "block_types": {
            "schedule_end_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule_start_date": {
              "block": {
                "attributes": {
                  "day": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "month": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "year": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "start_time_of_day": {
              "block": {
                "attributes": {
                  "hours": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "nanos": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "seconds": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
      "transfer_spec": {
        "block": {
          "block_types": {
            "aws_s3_data_source": {
              "block": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "aws_access_key": {
                    "block": {
                      "attributes": {
                        "access_key_id": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "secret_access_key": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gcs_data_sink": {
              "block": {
                "attributes": {
                  "bucket_name": {
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
            "gcs_data_source": {
              "block": {
                "attributes": {
                  "bucket_name": {
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
            "http_data_source": {
              "block": {
                "attributes": {
                  "list_url": {
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
            "object_conditions": {
              "block": {
                "attributes": {
                  "exclude_prefixes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_prefixes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "max_time_elapsed_since_last_modification": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "min_time_elapsed_since_last_modification": {
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
            "transfer_options": {
              "block": {
                "attributes": {
                  "delete_objects_from_source_after_transfer": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "delete_objects_unique_in_sink": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "overwrite_objects_already_existing_in_sink": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageTransferJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageTransferJob), &result)
	return &result
}
