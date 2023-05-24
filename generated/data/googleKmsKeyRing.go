package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsKeyRing = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the KeyRing.\nA full list of valid locations can be found by running 'gcloud kms locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name for the KeyRing.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsKeyRingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsKeyRing), &result)
	return &result
}
