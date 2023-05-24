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
        "description": "The actual certificate data for this client certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "cert_serial_number": {
        "computed": true,
        "description": "The serial number extracted from the certificate data.",
        "description_kind": "plain",
        "type": "string"
      },
      "common_name": {
        "description": "The common name to be used in the certificate to identify the client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the certificate was created in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "computed": true,
        "description": "The time when the certificate expires in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.",
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
        "description": "The name of the Cloud SQL instance. Changing this forces a new resource to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_key": {
        "computed": true,
        "description": "The private key associated with the client certificate.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_ca_cert": {
        "computed": true,
        "description": "The CA cert of the server this client cert was generated from.",
        "description_kind": "plain",
        "type": "string"
      },
      "sha1_fingerprint": {
        "computed": true,
        "description": "The SHA1 Fingerprint of the certificate.",
        "description_kind": "plain",
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
  "version": 1
}`

func GoogleSqlSslCertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSqlSslCert), &result)
	return &result
}
