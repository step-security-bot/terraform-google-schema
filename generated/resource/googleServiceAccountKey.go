package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceAccountKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keepers": {
        "description": "Arbitrary map of values that, when changed, will trigger recreation of resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "key_algorithm": {
        "description": "The algorithm used to generate the key, used only on create. KEY_ALG_RSA_2048 is the default algorithm. Valid values are: \"KEY_ALG_RSA_1024\", \"KEY_ALG_RSA_2048\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name used for this key pair",
        "description_kind": "plain",
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description": "The private key in JSON format, base64 encoded. This is what you normally get as a file when creating service account keys through the CLI or web console. This is only populated when creating a new key.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "private_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description": "The public key, base64 encoded",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_data": {
        "description": "A field that allows clients to upload their own public key. If set, use this public key data to create a service account key for given service account. Please note, the expected format for this field is a base64 encoded X509_PEM.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_id": {
        "description": "The ID of the parent service account of the key. This can be a string in the format {ACCOUNT} or projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}, where {ACCOUNT} is the email address or unique id of the service account. If the {ACCOUNT} syntax is used, the project will be inferred from the provider's configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "valid_after": {
        "computed": true,
        "description": "The key can be used after this timestamp. A timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "valid_before": {
        "computed": true,
        "description": "The key can be used before this timestamp. A timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds. Example: \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleServiceAccountKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceAccountKey), &result)
	return &result
}
