package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDocumentAiProcessor = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "The display name. Must be unique.",
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
      "kms_key_name": {
        "description": "The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the processor.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleDocumentAiProcessorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDocumentAiProcessor), &result)
	return &result
}
