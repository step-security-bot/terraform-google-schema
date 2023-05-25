package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIapClient = `{
  "block": {
    "attributes": {
      "brand": {
        "description": "Identifier of the brand to which this client\nis attached to. The format is\n'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_id": {
        "description": "Output only. Unique identifier of the OAuth client.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Human-friendly name given to the OAuth client.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secret": {
        "computed": true,
        "description": "Output only. Client secret of the OAuth client.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleIapClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIapClient), &result)
	return &result
}
