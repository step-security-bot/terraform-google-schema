package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsRecordSet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "DNS record set identifier",
        "description_kind": "markdown",
        "type": "string"
      },
      "managed_zone": {
        "description": "The Name of the zone.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The DNS name for the resource.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project for the Google Cloud.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "rrdatas": {
        "computed": true,
        "description": "The string data for the records in this record set.",
        "description_kind": "markdown",
        "type": [
          "list",
          "string"
        ]
      },
      "ttl": {
        "computed": true,
        "description": "The time-to-live of this record set (seconds).",
        "description_kind": "markdown",
        "type": "number"
      },
      "type": {
        "description": "The identifier of a supported record type. See the list of Supported DNS record types.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      }
    },
    "description": "A DNS record set within Google Cloud DNS",
    "description_kind": "markdown"
  },
  "version": 0
}`

func GoogleDnsRecordSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsRecordSet), &result)
	return &result
}
