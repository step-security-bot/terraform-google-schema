package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleContainerAnalysisNote = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time this note was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
        "description": "Time of expiration for this note. Leave empty if note does not expire.",
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
      "kind": {
        "computed": true,
        "description": "The type of analysis this note describes",
        "description_kind": "plain",
        "type": "string"
      },
      "long_description": {
        "description": "A detailed description of the note",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the note.",
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
      "related_note_names": {
        "description": "Names of other notes related to this note.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "short_description": {
        "description": "A one sentence description of the note.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time this note was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "attestation_authority": {
        "block": {
          "block_types": {
            "hint": {
              "block": {
                "attributes": {
                  "human_readable_name": {
                    "description": "The human readable name of this Attestation Authority, for\nexample \"qa\".",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "This submessage provides human-readable hints about the purpose of\nthe AttestationAuthority. Because the name of a Note acts as its\nresource reference, it is important to disambiguate the canonical\nname of the Note (which might be a UUID for security purposes)\nfrom \"readable\" names more suitable for debug output. Note that\nthese hints should NOT be used to look up AttestationAuthorities\nin security sensitive contexts, such as when looking up\nAttestations to verify.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Note kind that represents a logical attestation \"role\" or \"authority\".\nFor example, an organization might have one AttestationAuthority for\n\"QA\" and one for \"build\". This Note is intended to act strictly as a\ngrouping mechanism for the attached Occurrences (Attestations). This\ngrouping mechanism also provides a security boundary, since IAM ACLs\ngate the ability for a principle to attach an Occurrence to a given\nNote. It also provides a single point of lookup to find all attached\nAttestation Occurrences, even if they don't all live in the same\nproject.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "related_url": {
        "block": {
          "attributes": {
            "label": {
              "description": "Label to describe usage of the URL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description": "Specific URL associated with the resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "URLs associated with this note and related metadata.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleContainerAnalysisNoteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleContainerAnalysisNote), &result)
	return &result
}
