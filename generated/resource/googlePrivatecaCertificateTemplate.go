package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivatecaCertificateTemplate = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The time at which this CertificateTemplate was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Optional. A human-readable description of scenarios this template is intended for.",
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
        "description": "Optional. Labels with user-defined metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location for the resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name for this CertificateTemplate in the format ` + "`" + `projects/*/locations/*/certificateTemplates/*` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project for the resource",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The time at which this CertificateTemplate was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "identity_constraints": {
        "block": {
          "attributes": {
            "allow_subject_alt_names_passthrough": {
              "description": "Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "allow_subject_passthrough": {
              "description": "Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "cel_expression": {
              "block": {
                "attributes": {
                  "description": {
                    "description": "Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression": {
                    "description": "Textual representation of an expression in Common Expression Language syntax.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "location": {
                    "description": "Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "title": {
                    "description": "Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "passthrough_extensions": {
        "block": {
          "attributes": {
            "known_extensions": {
              "description": "Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "additional_extensions": {
              "block": {
                "attributes": {
                  "object_id_path": {
                    "description": "Required. The parts of an OID path. The most significant parts of the path come first.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description": "Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "predefined_values": {
        "block": {
          "attributes": {
            "aia_ocsp_servers": {
              "description": "Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the \"Authority Information Access\" extension in the certificate.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "additional_extensions": {
              "block": {
                "attributes": {
                  "critical": {
                    "description": "Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "value": {
                    "description": "Required. The value of this X.509 extension.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "object_id": {
                    "block": {
                      "attributes": {
                        "object_id_path": {
                          "description": "Required. The parts of an OID path. The most significant parts of the path come first.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        }
                      },
                      "description": "Required. The OID for this X.509 extension.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Describes custom X.509 extensions.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ca_options": {
              "block": {
                "attributes": {
                  "is_ca": {
                    "description": "Optional. Refers to the \"CA\" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "max_issuer_path_length": {
                    "description": "Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Optional. Describes options in this X509Parameters that are relevant in a CA certificate.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "key_usage": {
              "block": {
                "block_types": {
                  "base_key_usage": {
                    "block": {
                      "attributes": {
                        "cert_sign": {
                          "description": "The key may be used to sign certificates.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "content_commitment": {
                          "description": "The key may be used for cryptographic commitments. Note that this may also be referred to as \"non-repudiation\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "crl_sign": {
                          "description": "The key may be used sign certificate revocation lists.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "data_encipherment": {
                          "description": "The key may be used to encipher data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "decipher_only": {
                          "description": "The key may be used to decipher only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "digital_signature": {
                          "description": "The key may be used for digital signatures.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encipher_only": {
                          "description": "The key may be used to encipher only.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_agreement": {
                          "description": "The key may be used in a key agreement protocol.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_encipherment": {
                          "description": "The key may be used to encipher other keys.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Describes high-level ways in which a key may be used.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "extended_key_usage": {
                    "block": {
                      "attributes": {
                        "client_auth": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as \"TLS WWW client authentication\", though regularly used for non-WWW TLS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "code_signing": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as \"Signing of downloadable executable code client authentication\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "email_protection": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as \"Email protection\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "ocsp_signing": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as \"Signing OCSP responses\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "server_auth": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as \"TLS WWW server authentication\", though regularly used for non-WWW TLS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "time_stamping": {
                          "description": "Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as \"Binding the hash of an object to a time\".",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description": "Detailed scenarios in which a key may be used.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "unknown_extended_key_usages": {
                    "block": {
                      "attributes": {
                        "object_id_path": {
                          "description": "Required. The parts of an OID path. The most significant parts of the path come first.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        }
                      },
                      "description": "Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Indicates the intended use for keys that correspond to a certificate.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "policy_ids": {
              "block": {
                "attributes": {
                  "object_id_path": {
                    "description": "Required. The parts of an OID path. The most significant parts of the path come first.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description": "Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.",
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
  "version": 0
}`

func GooglePrivatecaCertificateTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivatecaCertificateTemplate), &result)
	return &result
}
