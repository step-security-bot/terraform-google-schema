package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAnalysisOccurrence = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time when the repository was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description": "The note kind which explicitly denotes which of the occurrence\ndetails are specified. This field can be used as a filter in list\nrequests.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the occurrence.",
        "description_kind": "plain",
        "type": "string"
      },
      "note_name": {
        "description": "The analysis note associated with this occurrence, in the form of\nprojects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a\nfilter in list requests.",
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
      "remediation": {
        "description": "A description of actions that can be taken to remedy the note.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_uri": {
        "description": "Required. Immutable. A URI that represents the resource for which\nthe occurrence applies. For example,\nhttps://gcr.io/project/image@sha256:123abc for a Docker image.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time when the repository was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "attestation": {
        "block": {
          "attributes": {
            "serialized_payload": {
              "description": "The serialized payload that is verified by one or\nmore signatures. A base64-encoded string.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "signatures": {
              "block": {
                "attributes": {
                  "public_key_id": {
                    "description": "The identifier for the public key that verifies this\nsignature. MUST be an RFC3986 conformant\nURI. * When possible, the key id should be an\nimmutable reference, such as a cryptographic digest.\nExamples of valid values:\n\n* OpenPGP V4 public key fingerprint. See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr\n  for more details on this scheme.\n    * 'openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA'\n* RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER serialization):\n    * \"ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU\"",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "signature": {
                    "description": "The content of the signature, an opaque bytestring.\nThe payload that this signature verifies MUST be\nunambiguously provided with the Signature during\nverification. A wrapper message might provide the\npayload explicitly. Alternatively, a message might\nhave a canonical serialization that can always be\nunambiguously computed to derive the payload.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "One or more signatures over serializedPayload.\nVerifier implementations should consider this attestation\nmessage verified if at least one signature verifies\nserializedPayload. See Signature in common.proto for more\ndetails on signature structure and verification.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "Occurrence that represents a single \"attestation\". The authenticity\nof an attestation can be verified using the attached signature.\nIf the verifier trusts the public key of the signer, then verifying\nthe signature is sufficient to establish trust. In this circumstance,\nthe authority to which this attestation is attached is primarily\nuseful for lookup (how to find this attestation if you already\nknow the authority and artifact to be verified) and intent (for\nwhich authority this attestation was intended to sign.",
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
  "version": 0
}`

func GoogleContainerAnalysisOccurrenceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAnalysisOccurrence), &result)
	return &result
}
