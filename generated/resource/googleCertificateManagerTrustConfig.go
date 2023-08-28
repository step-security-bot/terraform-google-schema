package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerTrustConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The creation timestamp of a TrustConfig.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "One or more paragraphs of text description of a trust config.",
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
        "description": "Set of label tags associated with the trust config.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The trust config location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the trust config. Trust config names must be unique globally.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a TrustConfig.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
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
      },
      "trust_stores": {
        "block": {
          "block_types": {
            "intermediate_cas": {
              "block": {
                "attributes": {
                  "pem_certificate": {
                    "description": "PEM intermediate certificate used for building up paths for validation.\nEach certificate provided in PEM format may occupy up to 5kB.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "Set of intermediate CA certificates used for the path building phase of chain validation.\nThe field is currently not supported if trust config is used for the workload certificate feature.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "trust_anchors": {
              "block": {
                "attributes": {
                  "pem_certificate": {
                    "description": "PEM root certificate of the PKI used for validation.\nEach certificate provided in PEM format may occupy up to 5kB.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description": "List of Trust Anchors to be used while performing validation against a given TrustStore.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Set of trust stores to perform validation against.\nThis field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCertificateManagerTrustConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerTrustConfig), &result)
	return &result
}
