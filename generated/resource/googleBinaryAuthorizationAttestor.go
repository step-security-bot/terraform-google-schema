package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBinaryAuthorizationAttestor = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A descriptive comment. This field may be updated. The field may be\ndisplayed in chooser dialogs.",
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
      "name": {
        "description": "The resource name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "attestation_authority_note": {
        "block": {
          "attributes": {
            "delegation_service_account_email": {
              "computed": true,
              "description": "This field will contain the service account email address that\nthis Attestor will use as the principal when querying Container\nAnalysis. Attestor administrators must grant this service account\nthe IAM role needed to read attestations from the noteReference in\nContainer Analysis (containeranalysis.notes.occurrences.viewer).\nThis email address is fixed for the lifetime of the Attestor, but\ncallers should not make any other assumptions about the service\naccount email; future versions may use an email based on a\ndifferent naming pattern.",
              "description_kind": "plain",
              "type": "string"
            },
            "note_reference": {
              "description": "The resource name of a ATTESTATION_AUTHORITY Note, created by the\nuser. If the Note is in a different project from the Attestor, it\nshould be specified in the format 'projects/*/notes/*' (or the legacy\n'providers/*/notes/*'). This field may not be updated.\nAn attestation by this attestor is stored as a Container Analysis\nATTESTATION_AUTHORITY Occurrence that names a container image\nand that links to this Note.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "public_keys": {
              "block": {
                "attributes": {
                  "ascii_armored_pgp_public_key": {
                    "description": "ASCII-armored representation of a PGP public key, as the\nentire output by the command\n'gpg --export --armor foo@example.com' (either LF or CRLF\nline endings). When using this field, id should be left\nblank. The BinAuthz API handlers will calculate the ID\nand fill it in automatically. BinAuthz computes this ID\nas the OpenPGP RFC4880 V4 fingerprint, represented as\nupper-case hex. If id is provided by the caller, it will\nbe overwritten by the API-calculated ID.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "comment": {
                    "description": "A descriptive comment. This field may be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description": "The ID of this public key. Signatures verified by BinAuthz\nmust include the ID of the public key that can be used to\nverify them, and that ID must match the contents of this\nfield exactly. Additional restrictions on this field can\nbe imposed based on which public key type is encapsulated.\nSee the documentation on publicKey cases below for details.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "pkix_public_key": {
                    "block": {
                      "attributes": {
                        "public_key_pem": {
                          "description": "A PEM-encoded public key, as described in\n'https://tools.ietf.org/html/rfc7468#section-13'",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "signature_algorithm": {
                          "description": "The signature algorithm used to verify a message against\na signature using this key. These signature algorithm must\nmatch the structure and any object identifiers encoded in\npublicKeyPem (i.e. this algorithm must match that of the\npublic key).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A raw PKIX SubjectPublicKeyInfo format public key.\n\nNOTE: id may be explicitly provided by the caller when using this\ntype of public key, but it MUST be a valid RFC3986 URI. If id is left\nblank, a default one will be computed based on the digest of the DER\nencoding of the public key.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Public keys that verify attestations signed by this attestor. This\nfield may be updated.\nIf this field is non-empty, one of the specified public keys must\nverify that an attestation was signed by this attestor for the\nimage specified in the admission request.\nIf this field is empty, this attestor always returns that no valid\nattestations exist.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.",
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

func GoogleBinaryAuthorizationAttestorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBinaryAuthorizationAttestor), &result)
	return &result
}
