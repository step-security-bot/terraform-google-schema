package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSqlSslCert = `{
  "block": {
    "attributes": {
      "cert": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cert_serial_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "common_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_ca_cert": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sha1_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleSqlSslCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlSslCert), &result)
	return &result
}
