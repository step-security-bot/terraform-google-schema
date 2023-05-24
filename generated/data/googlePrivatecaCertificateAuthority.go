package data

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
        "optional": true,
        "type": "string"
      },
      "config": {
        "computed": true,
        "description": "The config used to create a self-signed X.509 certificate or CSR.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "subject_config": [
                "list",
                [
                  "object",
                  {
                    "subject": [
                      "list",
                      [
                        "object",
                        {
                          "common_name": "string",
                          "country_code": "string",
                          "locality": "string",
                          "organization": "string",
                          "organizational_unit": "string",
                          "postal_code": "string",
                          "province": "string",
                          "street_address": "string"
                        }
                      ]
                    ],
                    "subject_alt_name": [
                      "list",
                      [
                        "object",
                        {
                          "dns_names": [
                            "list",
                            "string"
                          ],
                          "email_addresses": [
                            "list",
                            "string"
                          ],
                          "ip_addresses": [
                            "list",
                            "string"
                          ],
                          "uris": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "x509_config": [
                "list",
                [
                  "object",
                  {
                    "additional_extensions": [
                      "list",
                      [
                        "object",
                        {
                          "critical": "bool",
                          "object_id": [
                            "list",
                            [
                              "object",
                              {
                                "object_id_path": [
                                  "list",
                                  "number"
                                ]
                              }
                            ]
                          ],
                          "value": "string"
                        }
                      ]
                    ],
                    "aia_ocsp_servers": [
                      "list",
                      "string"
                    ],
                    "ca_options": [
                      "list",
                      [
                        "object",
                        {
                          "is_ca": "bool",
                          "max_issuer_path_length": "number",
                          "non_ca": "bool",
                          "zero_max_issuer_path_length": "bool"
                        }
                      ]
                    ],
                    "key_usage": [
                      "list",
                      [
                        "object",
                        {
                          "base_key_usage": [
                            "list",
                            [
                              "object",
                              {
                                "cert_sign": "bool",
                                "content_commitment": "bool",
                                "crl_sign": "bool",
                                "data_encipherment": "bool",
                                "decipher_only": "bool",
                                "digital_signature": "bool",
                                "encipher_only": "bool",
                                "key_agreement": "bool",
                                "key_encipherment": "bool"
                              }
                            ]
                          ],
                          "extended_key_usage": [
                            "list",
                            [
                              "object",
                              {
                                "client_auth": "bool",
                                "code_signing": "bool",
                                "email_protection": "bool",
                                "ocsp_signing": "bool",
                                "server_auth": "bool",
                                "time_stamping": "bool"
                              }
                            ]
                          ],
                          "unknown_extended_key_usages": [
                            "list",
                            [
                              "object",
                              {
                                "object_id_path": [
                                  "list",
                                  "number"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "name_constraints": [
                      "list",
                      [
                        "object",
                        {
                          "critical": "bool",
                          "excluded_dns_names": [
                            "list",
                            "string"
                          ],
                          "excluded_email_addresses": [
                            "list",
                            "string"
                          ],
                          "excluded_ip_ranges": [
                            "list",
                            "string"
                          ],
                          "excluded_uris": [
                            "list",
                            "string"
                          ],
                          "permitted_dns_names": [
                            "list",
                            "string"
                          ],
                          "permitted_email_addresses": [
                            "list",
                            "string"
                          ],
                          "permitted_ip_ranges": [
                            "list",
                            "string"
                          ],
                          "permitted_uris": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "policy_ids": [
                      "list",
                      [
                        "object",
                        {
                          "object_id_path": [
                            "list",
                            "number"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The time at which this CertificateAuthority was created.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether or not to allow Terraform to destroy the CertificateAuthority. Unless this field is set to false\nin Terraform state, a 'terraform destroy' or 'terraform apply' that would delete the instance will fail.",
        "description_kind": "plain",
        "type": "bool"
      },
      "desired_state": {
        "computed": true,
        "description": "Desired state of the CertificateAuthority. Set this field to 'STAGED' to create a 'STAGED' root CA.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcs_bucket": {
        "computed": true,
        "description": "The name of a Cloud Storage bucket where this CertificateAuthority will publish content,\nsuch as the CA certificate and CRLs. This must be a bucket name, without any prefixes\n(such as 'gs://') or suffixes (such as '.googleapis.com'). For example, to use a bucket named\nmy-bucket, you would simply specify 'my-bucket'. If not specified, a managed bucket will be\ncreated.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_active_certificates_on_deletion": {
        "computed": true,
        "description": "This field allows the CA to be deleted even if the CA has active certs. Active certs include both unrevoked and unexpired certs.\nUse with care. Defaults to 'false'.",
        "description_kind": "plain",
        "type": "bool"
      },
      "key_spec": {
        "computed": true,
        "description": "Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority\nis a self-signed CertificateAuthority, this key is also used to sign the self-signed CA\ncertificate. Otherwise, it is used to sign a CSR.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "cloud_kms_key_version": "string"
            }
          ]
        ]
      },
      "labels": {
        "computed": true,
        "description": "Labels with user-defined metadata.\n\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\":\n\"1.3kg\", \"count\": \"3\" }.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "lifetime": {
        "computed": true,
        "description": "The desired lifetime of the CA certificate. Used to create the \"notBeforeTime\" and\n\"notAfterTime\" fields inside an X.509 certificate. A duration in seconds with up to nine\nfractional digits, terminated by 's'. Example: \"3.5s\".",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "Location of the CertificateAuthority. A full list of valid locations can be found by\nrunning 'gcloud privateca locations list'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for this CertificateAuthority in the format\nprojects/*/locations/*/certificateAuthorities/*.",
        "description_kind": "plain",
        "type": "string"
      },
      "pem_ca_certificate": {
        "computed": true,
        "description": "The signed CA certificate issued from the subordinated CA's CSR. This is needed when activating the subordiante CA with a third party issuer.",
        "description_kind": "plain",
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
      "pem_csr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool": {
        "description": "The name of the CaPool this Certificate Authority belongs to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "skip_grace_period": {
        "computed": true,
        "description": "If this flag is set, the Certificate Authority will be deleted as soon as\npossible without a 30-day grace period where undeletion would have been\nallowed. If you proceed, there will be no way to recover this CA.\nUse with care. Defaults to 'false'.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The State for this CertificateAuthority.",
        "description_kind": "plain",
        "type": "string"
      },
      "subordinate_config": {
        "computed": true,
        "description": "If this is a subordinate CertificateAuthority, this field will be set\nwith the subordinate configuration, which describes its issuers.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_authority": "string",
              "pem_issuer_chain": [
                "list",
                [
                  "object",
                  {
                    "pem_certificates": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "type": {
        "computed": true,
        "description": "The Type of this CertificateAuthority.\n\n~\u003e **Note:** For 'SUBORDINATE' Certificate Authorities, they need to\nbe activated before they can issue certificates. Default value: \"SELF_SIGNED\" Possible values: [\"SELF_SIGNED\", \"SUBORDINATE\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which this CertificateAuthority was updated.\n\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
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
