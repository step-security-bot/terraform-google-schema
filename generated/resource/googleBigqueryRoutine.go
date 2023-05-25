package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryRoutine = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time when this routine was created, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataset_id": {
        "description": "The ID of the dataset containing this routine",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "definition_body": {
        "description": "The body of the routine. For functions, this is the expression in the AS clause.\nIf language=SQL, it is the substring inside (but excluding) the parentheses.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the routine if defined.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "determinism_level": {
        "description": "The determinism level of the JavaScript UDF if defined. Possible values: [\"DETERMINISM_LEVEL_UNSPECIFIED\", \"DETERMINISTIC\", \"NOT_DETERMINISTIC\"]",
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
      "imported_libraries": {
        "description": "Optional. If language = \"JAVASCRIPT\", this field stores the path of the\nimported JAVASCRIPT libraries.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "language": {
        "description": "The language of the routine. Possible values: [\"SQL\", \"JAVASCRIPT\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time when this routine was modified, in milliseconds since the\nepoch.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "return_type": {
        "description": "A JSON schema for the return type. Optional if language = \"SQL\"; required otherwise.\nIf absent, the return type is inferred from definitionBody at query time in each query\nthat references this routine. If present, then the evaluated result will be cast to\nthe specified returned type at query time. ~\u003e**NOTE**: Because this field expects a JSON\nstring, any changes to the string will create a diff, even if the JSON itself hasn't\nchanged. If the API returns a different value for the same schema, e.g. it switche\nd the order of values or replaced STRUCT field type with RECORD field type, we currently\ncannot suppress the recurring diff this causes. As a workaround, we recommend using\nthe schema as returned by the API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routine_id": {
        "description": "The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routine_type": {
        "description": "The type of routine. Possible values: [\"SCALAR_FUNCTION\", \"PROCEDURE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "arguments": {
        "block": {
          "attributes": {
            "argument_kind": {
              "description": "Defaults to FIXED_TYPE. Default value: \"FIXED_TYPE\" Possible values: [\"FIXED_TYPE\", \"ANY_TYPE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_type": {
              "description": "A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.\n~\u003e**NOTE**: Because this field expects a JSON string, any changes to the string\nwill create a diff, even if the JSON itself hasn't changed. If the API returns\na different value for the same schema, e.g. it switched the order of values\nor replaced STRUCT field type with RECORD field type, we currently cannot\nsuppress the recurring diff this causes. As a workaround, we recommend using\nthe schema as returned by the API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description": "Specifies whether the argument is input or output. Can be set for procedures only. Possible values: [\"IN\", \"OUT\", \"INOUT\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The name of this argument. Can be absent for function return argument.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Input/output argument of a function or a stored procedure.",
          "description_kind": "plain"
        },
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

func GoogleBigqueryRoutineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryRoutine), &result)
	return &result
}
