package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivatecaCaPool = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels with user-defined metadata.\n\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\":\n\"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the CaPool. A full list of valid locations can be found by\nrunning 'gcloud privateca locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name for this CaPool.",
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
      "tier": {
        "description": "The Tier of this CaPool. Possible values: [\"ENTERPRISE\", \"DEVOPS\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "issuance_policy": {
        "block": {
          "attributes": {
            "maximum_lifetime": {
              "description": "The maximum lifetime allowed for issued Certificates. Note that if the issuing CertificateAuthority\nexpires before a Certificate's requested maximumLifetime, the effective lifetime will be explicitly truncated to match it.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "allowed_issuance_modes": {
              "block": {
                "attributes": {
                  "allow_config_based_issuance": {
                    "description": "When true, allows callers to create Certificates by specifying a CertificateConfig.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "allow_csr_based_issuance": {
                    "description": "When true, allows callers to create Certificates by specifying a CSR.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description": "IssuanceModes specifies the allowed ways in which Certificates may be requested from this CaPool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "allowed_key_types": {
              "block": {
                "block_types": {
                  "elliptic_curve": {
                    "block": {
                      "attributes": {
                        "signature_algorithm": {
                          "description": "The algorithm used. Possible values: [\"ECDSA_P256\", \"ECDSA_P384\", \"EDDSA_25519\"]",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents an allowed Elliptic Curve key type.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "rsa": {
                    "block": {
                      "attributes": {
                        "max_modulus_size": {
                          "description": "The maximum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the\nservice will not enforce an explicit upper bound on RSA modulus sizes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "min_modulus_size": {
                          "description": "The minimum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the\nservice-level min RSA modulus size will continue to apply.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Describes an RSA key that may be used in a Certificate issued from a CaPool.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "If any AllowedKeyType is specified, then the certificate request's public key must match one of the key types listed here.\nOtherwise, any key may be used.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "baseline_values": {
              "block": {
                "attributes": {
                  "aia_ocsp_servers": {
                    "description": "Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the\n\"Authority Information Access\" extension in the certificate.",
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
                          "description": "Indicates whether or not this extension is critical (i.e., if the client does not know how to\nhandle this extension, the client should consider this to be an error).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "value": {
                          "description": "The value of this X.509 extension. A base64-encoded string.",
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
                                "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              }
                            },
                            "description": "Describes values that are relevant in a CA certificate.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "ca_options": {
                    "block": {
                      "attributes": {
                        "is_ca": {
                          "description": "Refers to the \"CA\" X.509 extension, which is a boolean value. When this value is missing,\nthe extension will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "max_issuer_path_length": {
                          "description": "Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of\nsubordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this\nvalue is missing, the max path length will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Describes values that are relevant in a CA certificate.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
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
                          "min_items": 1,
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
                            "description": "Describes high-level ways in which a key may be used.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "unknown_extended_key_usages": {
                          "block": {
                            "attributes": {
                              "object_id_path": {
                                "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              }
                            },
                            "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Indicates the intended use for keys that correspond to a certificate.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "policy_ids": {
                    "block": {
                      "attributes": {
                        "object_id_path": {
                          "description": "An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.",
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "number"
                          ]
                        }
                      },
                      "description": "Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A set of X.509 values that will be applied to all certificates issued through this CaPool. If a certificate request\nincludes conflicting values for the same properties, they will be overwritten by the values defined here. If a certificate\nrequest uses a CertificateTemplate that defines conflicting predefinedValues for the same properties, the certificate\nissuance request will fail.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "identity_constraints": {
              "block": {
                "attributes": {
                  "allow_subject_alt_names_passthrough": {
                    "description": "If this is set, the SubjectAltNames extension may be copied from a certificate request into the signed certificate.\nOtherwise, the requested SubjectAltNames will be discarded.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "allow_subject_passthrough": {
                    "description": "If this is set, the Subject field may be copied from a certificate request into the signed certificate.\nOtherwise, the requested Subject will be discarded.",
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
                          "description": "Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "description": "Textual representation of an expression in Common Expression Language syntax.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "location": {
                          "description": "String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "title": {
                          "description": "Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a\ncertificate is signed. To see the full allowed syntax and some examples,\nsee https://cloud.google.com/certificate-authority-service/docs/cel-guide",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Describes constraints on identities that may appear in Certificates issued through this CaPool.\nIf this is omitted, then this CaPool will not add restrictions on a certificate's identity.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The IssuancePolicy to control how Certificates will be issued from this CaPool.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "publishing_options": {
        "block": {
          "attributes": {
            "publish_ca_cert": {
              "description": "When true, publishes each CertificateAuthority's CA certificate and includes its URL in the \"Authority Information Access\"\nX.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding\nX.509 extension will not be written in issued certificates.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "publish_crl": {
              "description": "When true, publishes each CertificateAuthority's CRL and includes its URL in the \"CRL Distribution Points\" X.509 extension\nin all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not\nbe written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are\nalso rebuilt shortly after a certificate is revoked.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description": "The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.",
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

func GooglePrivatecaCaPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivatecaCaPool), &result)
	return &result
}
