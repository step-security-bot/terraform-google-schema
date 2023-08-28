package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCertificateManagerCertificate = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
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
        "description": "Set of label tags associated with the Certificate resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The Certificate Manager location. If not specified, \"global\" is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "A user-defined name of the certificate. Certificate names must be unique\nThe name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,\nand all following characters must be a dash, underscore, letter or digit.",
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
      "scope": {
        "description": "The scope of the certificate.\n\nDEFAULT: Certificates with default scope are served from core Google data centers.\nIf unsure, choose this option.\n\nEDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,\nserved from non-core Google data centers.\n\nALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).\nsee https://cloud.google.com/compute/docs/regions-zones",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "managed": {
        "block": {
          "attributes": {
            "authorization_attempt_info": {
              "computed": true,
              "description": "Detailed state of the latest authorization attempt for each domain\nspecified for this Managed Certificate.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "details": "string",
                    "domain": "string",
                    "failure_reason": "string",
                    "state": "string"
                  }
                ]
              ]
            },
            "dns_authorizations": {
              "description": "Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "domains": {
              "description": "The domains for which a managed SSL certificate will be generated.\nWildcard domains are only supported with DNS challenge resolution",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "issuance_config": {
              "description": "The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects/*/locations/*/certificateIssuanceConfigs/*.\nIf this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.\nEither issuanceConfig or dnsAuthorizations should be specificed, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioning_issue": {
              "computed": true,
              "description": "Information about issues with provisioning this Managed Certificate.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "details": "string",
                    "reason": "string"
                  }
                ]
              ]
            },
            "state": {
              "computed": true,
              "description": "A state of this Managed Certificate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Configuration and state of a Managed Certificate.\nCertificate Manager provisions and renews Managed Certificates\nautomatically, for as long as it's authorized to do so.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "self_managed": {
        "block": {
          "attributes": {
            "certificate_pem": {
              "deprecated": true,
              "description": "The certificate chain in PEM-encoded form.\n\nLeaf certificate comes first, followed by intermediate ones if any.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "pem_certificate": {
              "description": "The certificate chain in PEM-encoded form.\n\nLeaf certificate comes first, followed by intermediate ones if any.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pem_private_key": {
              "description": "The private key of the leaf certificate in PEM-encoded form.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "private_key_pem": {
              "deprecated": true,
              "description": "The private key of the leaf certificate in PEM-encoded form.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Certificate data for a SelfManaged Certificate.\nSelfManaged Certificates are uploaded by the user. Updating such\ncertificates before they expire remains the user's responsibility.",
          "description_kind": "plain"
        },
        "max_items": 1,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleCertificateManagerCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCertificateManagerCertificate), &result)
	return &result
}
