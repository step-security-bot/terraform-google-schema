package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsManagedZone = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A textual description field.",
        "description_kind": "markdown",
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description": "The fully qualified DNS name of this zone.",
        "description_kind": "markdown",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "DNS managed zone identifier",
        "description_kind": "markdown",
        "type": "string"
      },
      "managed_zone_id": {
        "computed": true,
        "description": "Unique identifier for the resource; defined by the server.",
        "description_kind": "markdown",
        "type": "number"
      },
      "name": {
        "description": "A unique name for the resource.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      },
      "name_servers": {
        "computed": true,
        "description": "The list of nameservers that will be authoritative for this domain. Use NS records to redirect from your DNS provider to these names, thus making Google Cloud DNS authoritative for this zone.",
        "description_kind": "markdown",
        "type": [
          "list",
          "string"
        ]
      },
      "project": {
        "description": "The ID of the project for the Google Cloud.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description": "The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.",
        "description_kind": "markdown",
        "type": "string"
      }
    },
    "description": "Provides access to a zone's attributes within Google Cloud DNS",
    "description_kind": "markdown"
  },
  "version": 0
}`

func GoogleDnsManagedZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsManagedZone), &result)
	return &result
}
