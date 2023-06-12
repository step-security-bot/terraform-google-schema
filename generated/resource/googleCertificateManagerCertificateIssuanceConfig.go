package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerCertificateIssuanceConfig = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The creation timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds with up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "One or more paragraphs of text description of a CertificateIssuanceConfig.",
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
      "key_algorithm": {
        "description": "Key algorithm to use when generating the private key. Possible values: [\"RSA_2048\", \"ECDSA_P256\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "'Set of label tags associated with the CertificateIssuanceConfig resource.\n An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "lifetime": {
        "description": "Lifetime of issued certificates. A duration in seconds with up to nine fractional digits, ending with 's'.\nExample: \"1814400s\". Valid values are from 21 days (1814400s) to 30 days (2592000s)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The Certificate Manager location. If not specified, \"global\" is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the certificate issuance config.\nCertificateIssuanceConfig names must be unique globally.",
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
      "rotation_window_percentage": {
        "description": "It specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate.\nMust be a number between 1-99, inclusive.\nYou must set the rotation window percentage in relation to the certificate lifetime so that certificate renewal occurs at least 7 days after\nthe certificate has been issued and at least 7 days before it expires.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of a CertificateIssuanceConfig. Timestamp is in RFC3339 UTC \"Zulu\" format,\naccurate to nanoseconds with up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "certificate_authority_config": {
        "block": {
          "block_types": {
            "certificate_authority_service_config": {
              "block": {
                "attributes": {
                  "ca_pool": {
                    "description": "A CA pool resource used to issue a certificate.\nThe CA pool string has a relative resource path following the form\n\"projects/{project}/locations/{location}/caPools/{caPool}\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Defines a CertificateAuthorityServiceConfig.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleCertificateManagerCertificateIssuanceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerCertificateIssuanceConfig), &result)
	return &result
}
