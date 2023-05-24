package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googlePrivatecaCertificateAuthority = `{
  "block": {
    "attributes": {
      "access_urls": {
        "computed": true,
        "description": "URLs for accessing content published by this CA, such as the CA certificate and CRLs.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ca_certificate_access_url": "string",
              "crl_access_urls": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "certificate_authority_id": {
        "description": "The user provided Resource ID for this Certificate Authority.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time at which this CertificateAuthority was created.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether or not to allow Terraform to destroy the CertificateAuthority. Unless this field is set to false\nin Terraform state, a 'terraform destroy' or 'terraform apply' that would delete the instance will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "desired_state": {
        "description": "Desired state of the CertificateAuthority. Set this field to 'STAGED' to create a 'STAGED' root CA.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "gcs_bucket": {
        "description": "The name of a Cloud Storage bucket where this CertificateAuthority will publish content,\nsuch as the CA certificate and CRLs. This must be a bucket name, without any prefixes\n(such as 'gs://') or suffixes (such as '.googleapis.com'). For example, to use a bucket named\nmy-bucket, you would simply specify 'my-bucket'. If not specified, a managed bucket will be\ncreated.",
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
      "ignore_active_certificates_on_deletion": {
        "description": "This field allows the CA to be deleted even if the CA has active certs. Active certs include both unrevoked and unexpired certs.\nUse with care. Defaults to 'false'.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "lifetime": {
        "description": "The desired lifetime of the CA certificate. Used to create the \"notBeforeTime\" and\n\"notAfterTime\" fields inside an X.509 certificate. A duration in seconds with up to nine\nfractional digits, terminated by 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "Location of the CertificateAuthority. A full list of valid locations can be found by\nrunning 'gcloud privateca locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for this CertificateAuthority in the format\nprojects/*/locations/*/certificateAuthorities/*.",
        "description_kind": "plain",
        "type": "string"
      },
      "pem_ca_certificate": {
        "description": "The signed CA certificate issued from the subordinated CA's CSR. This is needed when activating the subordiante CA with a third party issuer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pem_ca_certificates": {
        "computed": true,
        "description": "This CertificateAuthority's certificate chain, including the current\nCertificateAuthority's certificate. Ordered such that the root issuer is the final\nelement (consistent with RFC 5246). For a self-signed CA, this will only list the current\nCertificateAuthority's certificate.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "pool": {
        "description": "The name of the CaPool this Certificate Authority belongs to.",
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
      "skip_grace_period": {
        "description": "If this flag is set, the Certificate Authority will be deleted as soon as\npossible without a 30-day grace period where undeletion would have been\nallowed. If you proceed, there will be no way to recover this CA.\nUse with care. Defaults to 'false'.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The State for this CertificateAuthority.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The Type of this CertificateAuthority.\n\n~\u003e **Note:** For 'SUBORDINATE' Certificate Authorities, they need to\nbe activated before they can issue certificates. Default value: \"SELF_SIGNED\" Possible values: [\"SELF_SIGNED\", \"SUBORDINATE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which this CertificateAuthority was updated.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "block_types": {
            "subject_config": {
              "block": {
                "block_types": {
                  "subject": {
                    "block": {
                      "attributes": {
                        "common_name": {
                          "description": "The common name of the distinguished name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "country_code": {
                          "description": "The country code of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "locality": {
                          "description": "The locality or city of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "organization": {
                          "description": "The organization of the subject.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "organizational_unit": {
                          "description": "The organizational unit of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "postal_code": {
                          "description": "The postal code of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "province": {
                          "description": "The province, territory, or regional state of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "street_address": {
                          "description": "The street address of the subject.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Contains distinguished name fields such as the location and organization.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "subject_alt_name": {
                    "block": {
                      "attributes": {
                        "dns_names": {
                          "description": "Contains only valid, fully-qualified host names.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "email_addresses": {
                          "description": "Contains only valid RFC 2822 E-mail addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "ip_addresses": {
                          "description": "Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "uris": {
                          "description": "Contains only valid RFC 3986 URIs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "The subject alternative name fields.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies some of the values in a certificate that are related to the subject.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "x509_config": {
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
                          "description": "When true, the \"CA\" in Basic Constraints extension will be set to true.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "max_issuer_path_length": {
                          "description": "Refers to the \"path length constraint\" in Basic Constraints extension. For a CA certificate, this value describes the depth of\nsubordinate CA certificates that are allowed. If this value is less than 0, the request will fail. Setting the value to 0\nrequires setting 'zero_max_issuer_path_length = true'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "non_ca": {
                          "description": "When true, the \"CA\" in Basic Constraints extension will be set to false.\nIf both 'is_ca' and 'non_ca' are unset, the extension will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "zero_max_issuer_path_length": {
                          "description": "When true, the \"path length constraint\" in Basic Constraints extension will be set to 0.\nIf both 'max_issuer_path_length' and 'zero_max_issuer_path_length' are unset,\nthe max path length will be omitted from the CA certificate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
                  "name_constraints": {
                    "block": {
                      "attributes": {
                        "critical": {
                          "description": "Indicates whether or not the name constraints are marked critical.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "excluded_dns_names": {
                          "description": "Contains excluded DNS names. Any DNS name that can be\nconstructed by simply adding zero or more labels to\nthe left-hand side of the name satisfies the name constraint.\nFor example, 'example.com', 'www.example.com', 'www.sub.example.com'\nwould satisfy 'example.com' while 'example1.com' does not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_email_addresses": {
                          "description": "Contains the excluded email addresses. The value can be a particular\nemail address, a hostname to indicate all email addresses on that host or\na domain with a leading period (e.g. '.example.com') to indicate\nall email addresses in that domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_ip_ranges": {
                          "description": "Contains the excluded IP ranges. For IPv4 addresses, the ranges\nare expressed using CIDR notation as specified in RFC 4632.\nFor IPv6 addresses, the ranges are expressed in similar encoding as IPv4\naddresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "excluded_uris": {
                          "description": "Contains the excluded URIs that apply to the host part of the name.\nThe value can be a hostname or a domain with a\nleading period (like '.example.com')",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_dns_names": {
                          "description": "Contains permitted DNS names. Any DNS name that can be\nconstructed by simply adding zero or more labels to\nthe left-hand side of the name satisfies the name constraint.\nFor example, 'example.com', 'www.example.com', 'www.sub.example.com'\nwould satisfy 'example.com' while 'example1.com' does not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_email_addresses": {
                          "description": "Contains the permitted email addresses. The value can be a particular\nemail address, a hostname to indicate all email addresses on that host or\na domain with a leading period (e.g. '.example.com') to indicate\nall email addresses in that domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_ip_ranges": {
                          "description": "Contains the permitted IP ranges. For IPv4 addresses, the ranges\nare expressed using CIDR notation as specified in RFC 4632.\nFor IPv6 addresses, the ranges are expressed in similar encoding as IPv4\naddresses.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "permitted_uris": {
                          "description": "Contains the permitted URIs that apply to the host part of the name.\nThe value can be a hostname or a domain with a\nleading period (like '.example.com')",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Describes the X.509 name constraints extension.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
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
                "description": "Describes how some of the technical X.509 fields in a certificate should be populated.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The config used to create a self-signed X.509 certificate or CSR.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "key_spec": {
        "block": {
          "attributes": {
            "algorithm": {
              "description": "The algorithm to use for creating a managed Cloud KMS key for a for a simplified\nexperience. All managed keys will be have their ProtectionLevel as HSM. Possible values: [\"SIGN_HASH_ALGORITHM_UNSPECIFIED\", \"RSA_PSS_2048_SHA256\", \"RSA_PSS_3072_SHA256\", \"RSA_PSS_4096_SHA256\", \"RSA_PKCS1_2048_SHA256\", \"RSA_PKCS1_3072_SHA256\", \"RSA_PKCS1_4096_SHA256\", \"EC_P256_SHA256\", \"EC_P384_SHA384\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cloud_kms_key_version": {
              "description": "The resource name for an existing Cloud KMS CryptoKeyVersion in the format\n'projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority\nis a self-signed CertificateAuthority, this key is also used to sign the self-signed CA\ncertificate. Otherwise, it is used to sign a CSR.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "subordinate_config": {
        "block": {
          "attributes": {
            "certificate_authority": {
              "description": "This can refer to a CertificateAuthority that was used to create a\nsubordinate CertificateAuthority. This field is used for information\nand usability purposes only. The resource name is in the format\n'projects/*/locations/*/caPools/*/certificateAuthorities/*'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "pem_issuer_chain": {
              "block": {
                "attributes": {
                  "pem_certificates": {
                    "description": "Expected to be in leaf-to-root order according to RFC 5246.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Contains the PEM certificate chain for the issuers of this CertificateAuthority,\nbut not pem certificate for this CA itself.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "If this is a subordinate CertificateAuthority, this field will be set\nwith the subordinate configuration, which describes its issuers.",
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

func GooglePrivatecaCertificateAuthoritySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googlePrivatecaCertificateAuthority), &result)
	return &result
}
