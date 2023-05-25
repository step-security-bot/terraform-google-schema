package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryTable = `{
  "block": {
    "attributes": {
      "clustering": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "friendly_name": {
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
      "labels": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "num_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "num_long_term_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "num_rows": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "external_data_configuration": {
        "block": {
          "attributes": {
            "autodetect": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "compression": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_unknown_values": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_bad_records": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "source_format": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_uris": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "csv_options": {
              "block": {
                "attributes": {
                  "allow_jagged_rows": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "allow_quoted_newlines": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "encoding": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "field_delimiter": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "quote": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_sheets_options": {
              "block": {
                "attributes": {
                  "range": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "skip_leading_rows": {
                    "description_kind": "plain",
                    "optional": true,
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
        "nesting_mode": "list"
      },
      "time_partitioning": {
        "block": {
          "attributes": {
            "expiration_ms": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "field": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_partition_filter": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
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
      "view": {
        "block": {
          "attributes": {
            "query": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_legacy_sql": {
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
  "version": 0
}`

func GoogleBigqueryTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryTable), &result)
	return &result
}
