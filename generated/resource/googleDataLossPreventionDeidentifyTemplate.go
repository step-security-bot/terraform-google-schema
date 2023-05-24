package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionDeidentifyTemplate = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The creation timestamp of an deidentifyTemplate. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User set display name of the template.",
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
        "computed": true,
        "description": "The resource name of the template. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent": {
        "description": "The parent of the template in any of the following formats:\n\n* 'projects/{{project}}'\n* 'projects/{{project}}/locations/{{location}}'\n* 'organizations/{{organization_id}}'\n* 'organizations/{{organization_id}}/locations/{{location}}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last update timestamp of an deidentifyTemplate. Set by the server.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "deidentify_config": {
        "block": {
          "block_types": {
            "image_transformations": {
              "block": {
                "block_types": {
                  "transforms": {
                    "block": {
                      "block_types": {
                        "all_info_types": {
                          "block": {
                            "description": "Apply transformation to all findings not specified in other ImageTransformation's selectedInfoTypes.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "all_text": {
                          "block": {
                            "description": "Apply transformation to all text that doesn't match an infoType.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "redaction_color": {
                          "block": {
                            "attributes": {
                              "blue": {
                                "description": "The amount of blue in the color as a value in the interval [0, 1].",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "green": {
                                "description": "The amount of green in the color as a value in the interval [0, 1].",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "red": {
                                "description": "The amount of red in the color as a value in the interval [0, 1].",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "description": "The color to use when redacting content from an image. If not specified, the default is black.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "selected_info_types": {
                          "block": {
                            "block_types": {
                              "info_types": {
                                "block": {
                                  "attributes": {
                                    "name": {
                                      "description": "Name of the information type.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "version": {
                                      "description": "Version name for this InfoType.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to\nall findings that correspond to infoTypes that were requested in InspectConfig.",
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Apply transformation to the selected infoTypes.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "For determination of how redaction of images should occur.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Treat the dataset as an image and redact.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "info_type_transformations": {
              "block": {
                "block_types": {
                  "transformations": {
                    "block": {
                      "block_types": {
                        "info_types": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name of the information type.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description": "Version name for this InfoType.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to\nall findings that correspond to infoTypes that were requested in InspectConfig.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "primitive_transformation": {
                          "block": {
                            "attributes": {
                              "replace_with_info_type_config": {
                                "description": "Replace each matching finding with the name of the info type.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "block_types": {
                              "character_mask_config": {
                                "block": {
                                  "attributes": {
                                    "masking_character": {
                                      "description": "Character to use to mask the sensitive values—for example, * for an alphabetic string such as a name, or 0 for a numeric string\nsuch as ZIP code or credit card number. This string must have a length of 1. If not supplied, this value defaults to * for\nstrings, and 0 for digits.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "number_to_mask": {
                                      "description": "Number of characters to mask. If not set, all matching chars will be masked. Skipped characters do not count towards this tally.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "reverse_order": {
                                      "description": "Mask characters in reverse order. For example, if masking_character is 0, number_to_mask is 14, and reverse_order is 'false', then the\ninput string '1234-5678-9012-3456' is masked as '00000000000000-3456'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "characters_to_ignore": {
                                      "block": {
                                        "attributes": {
                                          "characters_to_skip": {
                                            "description": "Characters to not transform when masking.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "common_characters_to_ignore": {
                                            "description": "Common characters to not transform when masking. Useful to avoid removing punctuation. Possible values: [\"NUMERIC\", \"ALPHA_UPPER_CASE\", \"ALPHA_LOWER_CASE\", \"PUNCTUATION\", \"WHITESPACE\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Characters to skip when doing de-identification of a value. These will be left alone and skipped.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Partially mask a string by replacing a given number of characters with a fixed character.\nMasking can start from the beginning or end of the string.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "crypto_deterministic_config": {
                                "block": {
                                  "block_types": {
                                    "context": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name describing the field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A context may be used for higher security and maintaining referential integrity such that the same identifier in two different contexts will be given a distinct surrogate. The context is appended to plaintext value being encrypted. On decryption the provided context is validated against the value used during encryption. If a context was provided during encryption, same context must be provided during decryption as well.\n\nIf the context is not set, plaintext would be used as is for encryption. If the context is set but:\n\n1.  there is no record present when transforming a given value or\n2.  the field is not present when transforming a given value,\n\nplaintext would be used as is for encryption.\n\nNote that case (1) is expected when an 'InfoTypeTransformation' is applied to both structured and non-structured 'ContentItem's.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The key used by the encryption function.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "surrogate_info_type": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern '[A-Za-z0-9$-_]{1,64}'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "version": {
                                            "description": "Optional version name for this InfoType.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The custom info type to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom info type followed by the number of characters comprising the surrogate. The following scheme defines the format: {info type name}({surrogate character count}):{surrogate}\n\nFor example, if the name of custom info type is 'MY\\_TOKEN\\_INFO\\_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY\\_TOKEN\\_INFO\\_TYPE(3):abc'\n\nThis annotation identifies the surrogate when inspecting content using the custom info type 'Surrogate'. This facilitates reversal of the surrogate when it occurs in free text.\n\nNote: For record transformations where the entire cell in a table is being transformed, surrogates are not mandatory. Surrogates are used to denote the location of the token and are necessary for re-identification in free form text.\n\nIn order for inspection to work properly, the name of this info type must not occur naturally anywhere in your data; otherwise, inspection may either\n\n*   reverse a surrogate that does not correspond to an actual identifier\n*   be unable to parse the surrogate and result in an error\n\nTherefore, choose your custom info type name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY\\_TOKEN\\_TYPE.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Pseudonymization method that generates deterministic encryption for the given input. Outputs a base64 encoded representation of the encrypted output. Uses AES-SIV based on the RFC [https://tools.ietf.org/html/rfc5297](https://tools.ietf.org/html/rfc5297).",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "crypto_replace_ffx_fpe_config": {
                                "block": {
                                  "attributes": {
                                    "common_alphabet": {
                                      "description": "Common alphabets. Possible values: [\"FFX_COMMON_NATIVE_ALPHABET_UNSPECIFIED\", \"NUMERIC\", \"HEXADECIMAL\", \"UPPER_CASE_ALPHA_NUMERIC\", \"ALPHA_NUMERIC\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "custom_alphabet": {
                                      "description": "This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports. This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range \\[2, 95\\]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is:\n\n''0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~'!@#$%^\u0026*()_-+={[}]|:;\"'\u003c,\u003e.?/''",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "radix": {
                                      "description": "The native way to select the alphabet. Must be in the range \\[2, 95\\].",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "context": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name describing the field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The 'tweak', a context may be used for higher security since the same identifier in two different contexts won't be given the same surrogate. If the context is not set, a default tweak will be used.\n\nIf the context is set but:\n\n1.  there is no record present when transforming a given value or\n2.  the field is not present when transforming a given value,\n\na default tweak will be used.\n\nNote that case (1) is expected when an 'InfoTypeTransformation' is applied to both structured and non-structured 'ContentItem's. Currently, the referenced field may be of value type integer or string.\n\nThe tweak is constructed as a sequence of bytes in big endian byte order such that:\n\n*   a 64 bit integer is encoded followed by a single byte of value 1\n*   a string is encoded in UTF-8 format followed by a single byte of value 2",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The key used by the encryption algorithm.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "surrogate_info_type": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern '[A-Za-z0-9$-_]{1,64}'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "version": {
                                            "description": "Optional version name for this InfoType.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The custom infoType to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom infoType followed by the number of characters comprising the surrogate. The following scheme defines the format: info\\_type\\_name(surrogate\\_character\\_count):surrogate\n\nFor example, if the name of custom infoType is 'MY\\_TOKEN\\_INFO\\_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY\\_TOKEN\\_INFO\\_TYPE(3):abc'\n\nThis annotation identifies the surrogate when inspecting content using the custom infoType ['SurrogateType'](https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#surrogatetype). This facilitates reversal of the surrogate when it occurs in free text.\n\nIn order for inspection to work properly, the name of this infoType must not occur naturally anywhere in your data; otherwise, inspection may find a surrogate that does not correspond to an actual identifier. Therefore, choose your custom infoType name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY\\_TOKEN\\_TYPE",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replaces an identifier with a surrogate using Format Preserving Encryption (FPE) with the FFX mode of operation; however when used in the 'content.reidentify' API method, it serves the opposite function by reversing the surrogate back into the original identifier. The identifier must be encoded as ASCII. For a given crypto key and context, the same identifier will be replaced with the same surrogate. Identifiers must be at least two characters long. In the case that the identifier is the empty string, it will be skipped. See [https://cloud.google.com/dlp/docs/pseudonymization](https://cloud.google.com/dlp/docs/pseudonymization) to learn more.\n\nNote: We recommend using CryptoDeterministicConfig for all use cases which do not require preserving the input alphabet space and size, plus warrant referential integrity.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "replace_config": {
                                "block": {
                                  "block_types": {
                                    "new_value": {
                                      "block": {
                                        "attributes": {
                                          "boolean_value": {
                                            "description": "A boolean value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "day_of_week_value": {
                                            "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "float_value": {
                                            "description": "A float value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "integer_value": {
                                            "description": "An integer value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "string_value": {
                                            "description": "A string value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "timestamp_value": {
                                            "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits.\nExamples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "date_value": {
                                            "block": {
                                              "attributes": {
                                                "day": {
                                                  "description": "Day of month. Must be from 1 to 31 and valid for the year and month, or 0 if specifying a\nyear by itself or a year and month where the day is not significant.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "month": {
                                                  "description": "Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "year": {
                                                  "description": "Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a whole or partial calendar date.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "time_value": {
                                            "block": {
                                              "attributes": {
                                                "hours": {
                                                  "description": "Hours of day in 24 hour format. Should be from 0 to 23.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minutes": {
                                                  "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "nanos": {
                                                  "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "seconds": {
                                                  "description": "Seconds of minutes of the time. Must normally be from 0 to 59.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a time of day.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Replace each input value with a given value.\nThe 'new_value' block must only contain one argument. For example when replacing the contents of a string-type field, only 'string_value' should be set.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replace each input value with a given value.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "replace_dictionary_config": {
                                "block": {
                                  "block_types": {
                                    "word_list": {
                                      "block": {
                                        "attributes": {
                                          "words": {
                                            "description": "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of words to select from for random replacement. The [limits](https://cloud.google.com/dlp/limits) page contains details about the size limits of dictionaries.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replace with a value randomly drawn (with replacement) from a dictionary.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Primitive transformation to apply to the infoType.\nThe 'primitive_transformation' block must only contain one argument, corresponding to the type of transformation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Transformation for each infoType. Cannot specify more than one for a given infoType.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Treat the dataset as free-form text and apply the same free text transformation everywhere",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "record_transformations": {
              "block": {
                "block_types": {
                  "field_transformations": {
                    "block": {
                      "block_types": {
                        "condition": {
                          "block": {
                            "block_types": {
                              "expressions": {
                                "block": {
                                  "attributes": {
                                    "logical_operator": {
                                      "description": "The operator to apply to the result of conditions. Default and currently only supported value is AND Default value: \"AND\" Possible values: [\"AND\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "conditions": {
                                      "block": {
                                        "block_types": {
                                          "conditions": {
                                            "block": {
                                              "attributes": {
                                                "operator": {
                                                  "description": "Operator used to compare the field or infoType to the value. Possible values: [\"EQUAL_TO\", \"NOT_EQUAL_TO\", \"GREATER_THAN\", \"LESS_THAN\", \"GREATER_THAN_OR_EQUALS\", \"LESS_THAN_OR_EQUALS\", \"EXISTS\"]",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field": {
                                                  "block": {
                                                    "attributes": {
                                                      "name": {
                                                        "description": "Name describing the field.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description": "Field within the record this condition is evaluated against.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "min_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "description": "A boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "day_of_week_value": {
                                                        "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "float_value": {
                                                        "description": "A float value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "integer_value": {
                                                        "description": "An integer value (int64 format)",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "description": "A string value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "timestamp_value": {
                                                        "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "date_value": {
                                                        "block": {
                                                          "attributes": {
                                                            "day": {
                                                              "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "month": {
                                                              "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "year": {
                                                              "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            }
                                                          },
                                                          "description": "Represents a whole or partial calendar date.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "time_value": {
                                                        "block": {
                                                          "attributes": {
                                                            "hours": {
                                                              "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "minutes": {
                                                              "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "nanos": {
                                                              "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "seconds": {
                                                              "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            }
                                                          },
                                                          "description": "Represents a time of day.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "Value to compare against.\nThe 'value' block must only contain one argument. For example when a condition is evaluated against a string-type field, only 'string_value' should be set.\nThis argument is mandatory, except for conditions using the 'EXISTS' operator.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A collection of conditions.",
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Conditions to apply to the expression.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "An expression.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Only apply the transformation if the condition evaluates to true for the given RecordCondition. The conditions are allowed to reference fields that are not used in the actual transformation.\nExample Use Cases:\n- Apply a different bucket transformation to an age column if the zip code column for the same record is within a specific range.\n- Redact a field if the date of birth field is greater than 85.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "fields": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "Name describing the field.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "Input field(s) to apply the transformation to. When you have columns that reference their position within a list, omit the index from the FieldId.\nFieldId name matching ignores the index. For example, instead of \"contact.nums[0].type\", use \"contact.nums.type\".",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "primitive_transformation": {
                          "block": {
                            "block_types": {
                              "bucketing_config": {
                                "block": {
                                  "block_types": {
                                    "buckets": {
                                      "block": {
                                        "block_types": {
                                          "max": {
                                            "block": {
                                              "attributes": {
                                                "boolean_value": {
                                                  "description": "A boolean value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "day_of_week_value": {
                                                  "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "float_value": {
                                                  "description": "A float value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "integer_value": {
                                                  "description": "An integer value (int64 format)",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "string_value": {
                                                  "description": "A string value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp_value": {
                                                  "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "date_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "day": {
                                                        "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "month": {
                                                        "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "year": {
                                                        "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a whole or partial calendar date.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "time_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "hours": {
                                                        "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "minutes": {
                                                        "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "nanos": {
                                                        "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "seconds": {
                                                        "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a time of day.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "Upper bound of the range, exclusive; type must match min.\nThe 'max' block must only contain one argument. See the 'bucketing_config' block description for more information about choosing a data type.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "min": {
                                            "block": {
                                              "attributes": {
                                                "boolean_value": {
                                                  "description": "A boolean value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "day_of_week_value": {
                                                  "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "float_value": {
                                                  "description": "A float value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "integer_value": {
                                                  "description": "An integer value (int64 format)",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "string_value": {
                                                  "description": "A string value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp_value": {
                                                  "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "date_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "day": {
                                                        "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "month": {
                                                        "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "year": {
                                                        "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a whole or partial calendar date.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "time_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "hours": {
                                                        "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "minutes": {
                                                        "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "nanos": {
                                                        "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "seconds": {
                                                        "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a time of day.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "Lower bound of the range, inclusive. Type should be the same as max if used.\nThe 'min' block must only contain one argument. See the 'bucketing_config' block description for more information about choosing a data type.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "replacement_value": {
                                            "block": {
                                              "attributes": {
                                                "boolean_value": {
                                                  "description": "A boolean value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "bool"
                                                },
                                                "day_of_week_value": {
                                                  "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "float_value": {
                                                  "description": "A float value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "integer_value": {
                                                  "description": "An integer value (int64 format)",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "string_value": {
                                                  "description": "A string value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp_value": {
                                                  "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "date_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "day": {
                                                        "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "month": {
                                                        "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "year": {
                                                        "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a whole or partial calendar date.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "time_value": {
                                                  "block": {
                                                    "attributes": {
                                                      "hours": {
                                                        "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "minutes": {
                                                        "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "nanos": {
                                                        "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "seconds": {
                                                        "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      }
                                                    },
                                                    "description": "Represents a time of day.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "Replacement value for this bucket.\nThe 'replacement_value' block must only contain one argument.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "min_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Set of buckets. Ranges must be non-overlapping.\nBucket is represented as a range, along with replacement values.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Generalization function that buckets values based on ranges. The ranges and replacement values are dynamically provided by the user for custom behavior, such as 1-30 -\u003e LOW 31-65 -\u003e MEDIUM 66-100 -\u003e HIGH\nThis can be used on data of type: number, long, string, timestamp.\nIf the provided value type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing.\nSee https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "character_mask_config": {
                                "block": {
                                  "attributes": {
                                    "masking_character": {
                                      "description": "Character to use to mask the sensitive values—for example, * for an alphabetic string such as a name, or 0 for a numeric string\nsuch as ZIP code or credit card number. This string must have a length of 1. If not supplied, this value defaults to * for\nstrings, and 0 for digits.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "number_to_mask": {
                                      "description": "Number of characters to mask. If not set, all matching chars will be masked. Skipped characters do not count towards this tally.\nIf number_to_mask is negative, this denotes inverse masking. Cloud DLP masks all but a number of characters. For example, suppose you have the following values:\n- 'masking_character' is *\n- 'number_to_mask' is -4\n- 'reverse_order' is false\n- 'characters_to_ignore' includes -\n- Input string is 1234-5678-9012-3456\n\nThe resulting de-identified string is ****-****-****-3456. Cloud DLP masks all but the last four characters. If reverseOrder is true, all but the first four characters are masked as 1234-****-****-****.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "reverse_order": {
                                      "description": "Mask characters in reverse order. For example, if masking_character is 0, number_to_mask is 14, and reverse_order is 'false', then the\ninput string '1234-5678-9012-3456' is masked as '00000000000000-3456'.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "block_types": {
                                    "characters_to_ignore": {
                                      "block": {
                                        "attributes": {
                                          "characters_to_skip": {
                                            "description": "Characters to not transform when masking.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "common_characters_to_ignore": {
                                            "description": "Common characters to not transform when masking. Useful to avoid removing punctuation. Possible values: [\"NUMERIC\", \"ALPHA_UPPER_CASE\", \"ALPHA_LOWER_CASE\", \"PUNCTUATION\", \"WHITESPACE\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Characters to skip when doing de-identification of a value. These will be left alone and skipped.",
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Partially mask a string by replacing a given number of characters with a fixed character. Masking can start from the beginning or end of the string. This can be used on data of any type (numbers, longs, and so on) and when de-identifying structured data we'll attempt to preserve the original data's type. (This allows you to take a long like 123 and modify it to a string like **3).",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "crypto_deterministic_config": {
                                "block": {
                                  "block_types": {
                                    "context": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name describing the field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "A context may be used for higher security and maintaining referential integrity such that the same identifier in two different contexts will be given a distinct surrogate. The context is appended to plaintext value being encrypted. On decryption the provided context is validated against the value used during encryption. If a context was provided during encryption, same context must be provided during decryption as well.\n\nIf the context is not set, plaintext would be used as is for encryption. If the context is set but:\n\n1. there is no record present when transforming a given value or\n2. the field is not present when transforming a given value,\n\nplaintext would be used as is for encryption.\n\nNote that case (1) is expected when an InfoTypeTransformation is applied to both structured and unstructured ContentItems.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The key used by the encryption function. For deterministic encryption using AES-SIV, the provided key is internally expanded to 64 bytes prior to use.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "surrogate_info_type": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern '[A-Za-z0-9$-_]{1,64}'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "version": {
                                            "description": "Optional version name for this InfoType.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The custom info type to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom info type followed by the number of characters comprising the surrogate. The following scheme defines the format: {info type name}({surrogate character count}):{surrogate}\n\nFor example, if the name of custom info type is 'MY\\_TOKEN\\_INFO\\_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY\\_TOKEN\\_INFO\\_TYPE(3):abc'\n\nThis annotation identifies the surrogate when inspecting content using the custom info type 'Surrogate'. This facilitates reversal of the surrogate when it occurs in free text.\n\nNote: For record transformations where the entire cell in a table is being transformed, surrogates are not mandatory. Surrogates are used to denote the location of the token and are necessary for re-identification in free form text.\n\nIn order for inspection to work properly, the name of this info type must not occur naturally anywhere in your data; otherwise, inspection may either\n\n*   reverse a surrogate that does not correspond to an actual identifier\n*   be unable to parse the surrogate and result in an error\n\nTherefore, choose your custom info type name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY\\_TOKEN\\_TYPE.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Pseudonymization method that generates deterministic encryption for the given input. Outputs a base64 encoded representation of the encrypted output. Uses AES-SIV based on the RFC [https://tools.ietf.org/html/rfc5297](https://tools.ietf.org/html/rfc5297).",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "crypto_hash_config": {
                                "block": {
                                  "block_types": {
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The key used by the encryption function.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Pseudonymization method that generates surrogates via cryptographic hashing. Uses SHA-256. The key size must be either 32 or 64 bytes.\nOutputs a base64 encoded representation of the hashed output (for example, L7k0BHmF1ha5U3NfGykjro4xWi1MPVQPjhMAZbSV9mM=).\nCurrently, only string and integer values can be hashed.\nSee https://cloud.google.com/dlp/docs/pseudonymization to learn more.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "crypto_replace_ffx_fpe_config": {
                                "block": {
                                  "attributes": {
                                    "common_alphabet": {
                                      "description": "Common alphabets. Possible values: [\"FFX_COMMON_NATIVE_ALPHABET_UNSPECIFIED\", \"NUMERIC\", \"HEXADECIMAL\", \"UPPER_CASE_ALPHA_NUMERIC\", \"ALPHA_NUMERIC\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "custom_alphabet": {
                                      "description": "This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports. This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range \\[2, 95\\]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is:\n\n''0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~'!@#$%^\u0026*()_-+={[}]|:;\"'\u003c,\u003e.?/''",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "radix": {
                                      "description": "The native way to select the alphabet. Must be in the range \\[2, 95\\].",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "context": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name describing the field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The 'tweak', a context may be used for higher security since the same identifier in two different contexts won't be given the same surrogate. If the context is not set, a default tweak will be used.\n\nIf the context is set but:\n\n1.  there is no record present when transforming a given value or\n2.  the field is not present when transforming a given value,\n\na default tweak will be used.\n\nNote that case (1) is expected when an 'InfoTypeTransformation' is applied to both structured and non-structured 'ContentItem's. Currently, the referenced field may be of value type integer or string.\n\nThe tweak is constructed as a sequence of bytes in big endian byte order such that:\n\n*   a 64 bit integer is encoded followed by a single byte of value 1\n*   a string is encoded in UTF-8 format followed by a single byte of value 2",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "The key used by the encryption algorithm.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "surrogate_info_type": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern '[A-Za-z0-9$-_]{1,64}'.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "version": {
                                            "description": "Optional version name for this InfoType.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "The custom infoType to annotate the surrogate with. This annotation will be applied to the surrogate by prefixing it with the name of the custom infoType followed by the number of characters comprising the surrogate. The following scheme defines the format: info\\_type\\_name(surrogate\\_character\\_count):surrogate\n\nFor example, if the name of custom infoType is 'MY\\_TOKEN\\_INFO\\_TYPE' and the surrogate is 'abc', the full replacement value will be: 'MY\\_TOKEN\\_INFO\\_TYPE(3):abc'\n\nThis annotation identifies the surrogate when inspecting content using the custom infoType ['SurrogateType'](https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#surrogatetype). This facilitates reversal of the surrogate when it occurs in free text.\n\nIn order for inspection to work properly, the name of this infoType must not occur naturally anywhere in your data; otherwise, inspection may find a surrogate that does not correspond to an actual identifier. Therefore, choose your custom infoType name carefully after considering what your data looks like. One way to select a name that has a high chance of yielding reliable detection is to include one or more unicode characters that are highly improbable to exist in your data. For example, assuming your data is entered from a regular ASCII keyboard, the symbol with the hex code point 29DD might be used like so: ⧝MY\\_TOKEN\\_TYPE",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replaces an identifier with a surrogate using Format Preserving Encryption (FPE) with the FFX mode of operation; however when used in the 'content.reidentify' API method, it serves the opposite function by reversing the surrogate back into the original identifier. The identifier must be encoded as ASCII. For a given crypto key and context, the same identifier will be replaced with the same surrogate. Identifiers must be at least two characters long. In the case that the identifier is the empty string, it will be skipped. See [https://cloud.google.com/dlp/docs/pseudonymization](https://cloud.google.com/dlp/docs/pseudonymization) to learn more.\n\nNote: We recommend using CryptoDeterministicConfig for all use cases which do not require preserving the input alphabet space and size, plus warrant referential integrity.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "date_shift_config": {
                                "block": {
                                  "attributes": {
                                    "lower_bound_days": {
                                      "description": "For example, -5 means shift date to at most 5 days back in the past.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "upper_bound_days": {
                                      "description": "Range of shift in days. Actual shift will be selected at random within this range (inclusive ends). Negative means shift to earlier in time. Must not be more than 365250 days (1000 years) each direction.\n\nFor example, 3 means shift date to at most 3 days into the future.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "context": {
                                      "block": {
                                        "attributes": {
                                          "name": {
                                            "description": "Name describing the field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "description": "Points to the field that contains the context, for example, an entity id.\nIf set, must also set cryptoKey. If set, shift will be consistent for the given context.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "crypto_key": {
                                      "block": {
                                        "block_types": {
                                          "kms_wrapped": {
                                            "block": {
                                              "attributes": {
                                                "crypto_key_name": {
                                                  "description": "The resource name of the KMS CryptoKey to use for unwrapping.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                },
                                                "wrapped_key": {
                                                  "description": "The wrapped data crypto key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "KMS wrapped key.\nInclude to use an existing data crypto key wrapped by KMS. The wrapped key must be a 128-, 192-, or 256-bit key. Authorization requires the following IAM permissions when sending a request to perform a crypto transformation using a KMS-wrapped crypto key: dlp.kms.encrypt\nFor more information, see [Creating a wrapped key](https://cloud.google.com/dlp/docs/create-wrapped-key).\nNote: When you use Cloud KMS for cryptographic operations, [charges apply](https://cloud.google.com/kms/pricing).",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "transient": {
                                            "block": {
                                              "attributes": {
                                                "name": {
                                                  "description": "Name of the key. This is an arbitrary string used to differentiate different keys. A unique key is generated per name: two separate 'TransientCryptoKey' protos share the same generated key if their names are the same. When the data crypto key is generated, this name is not used in any way (repeating the api call will result in a different key being generated).",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Transient crypto key. Use this to have a random data crypto key generated. It will be discarded after the request finishes.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "unwrapped": {
                                            "block": {
                                              "attributes": {
                                                "key": {
                                                  "description": "A 128/192/256 bit key.\n\nA base64-encoded string.",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description": "Unwrapped crypto key. Using raw keys is prone to security risks due to accidentally leaking the key. Choose another type of key if possible.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Causes the shift to be computed based on this key and the context. This results in the same shift for the same context and cryptoKey. If set, must also set context. Can only be applied to table items.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Shifts dates by random number of days, with option to be consistent for the same context. See https://cloud.google.com/dlp/docs/concepts-date-shifting to learn more.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "fixed_size_bucketing_config": {
                                "block": {
                                  "attributes": {
                                    "bucket_size": {
                                      "description": "Size of each bucket (except for minimum and maximum buckets).\nSo if lower_bound = 10, upper_bound = 89, and bucketSize = 10, then the following buckets would be used: -10, 10-20, 20-30, 30-40, 40-50, 50-60, 60-70, 70-80, 80-89, 89+.\nPrecision up to 2 decimals works.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "block_types": {
                                    "lower_bound": {
                                      "block": {
                                        "attributes": {
                                          "boolean_value": {
                                            "description": "A boolean value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "day_of_week_value": {
                                            "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "float_value": {
                                            "description": "A float value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "integer_value": {
                                            "description": "An integer value (int64 format)",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "description": "A string value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "timestamp_value": {
                                            "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "date_value": {
                                            "block": {
                                              "attributes": {
                                                "day": {
                                                  "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "month": {
                                                  "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "year": {
                                                  "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a whole or partial calendar date.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "time_value": {
                                            "block": {
                                              "attributes": {
                                                "hours": {
                                                  "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minutes": {
                                                  "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "nanos": {
                                                  "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "seconds": {
                                                  "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a time of day.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Lower bound value of buckets.\nAll values less than lower_bound are grouped together into a single bucket; for example if lower_bound = 10, then all values less than 10 are replaced with the value \"-10\".\nThe 'lower_bound' block must only contain one argument. See the 'fixed_size_bucketing_config' block description for more information about choosing a data type.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "upper_bound": {
                                      "block": {
                                        "attributes": {
                                          "boolean_value": {
                                            "description": "A boolean value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "day_of_week_value": {
                                            "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "float_value": {
                                            "description": "A float value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "integer_value": {
                                            "description": "An integer value (int64 format)",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "description": "A string value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "timestamp_value": {
                                            "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "date_value": {
                                            "block": {
                                              "attributes": {
                                                "day": {
                                                  "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "month": {
                                                  "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "year": {
                                                  "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a whole or partial calendar date.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "time_value": {
                                            "block": {
                                              "attributes": {
                                                "hours": {
                                                  "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minutes": {
                                                  "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "nanos": {
                                                  "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "seconds": {
                                                  "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a time of day.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Upper bound value of buckets.\nAll values greater than upper_bound are grouped together into a single bucket; for example if upper_bound = 89, then all values greater than 89 are replaced with the value \"89+\".\nThe 'upper_bound' block must only contain one argument. See the 'fixed_size_bucketing_config' block description for more information about choosing a data type.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Buckets values based on fixed size ranges. The Bucketing transformation can provide all of this functionality, but requires more configuration. This message is provided as a convenience to the user for simple bucketing strategies.\n\nThe transformed value will be a hyphenated string of {lower_bound}-{upper_bound}. For example, if lower_bound = 10 and upper_bound = 20, all values that are within this bucket will be replaced with \"10-20\".\n\nThis can be used on data of type: double, long.\n\nIf the bound Value type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing.\n\nSee https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "redact_config": {
                                "block": {
                                  "description": "Redact a given value. For example, if used with an InfoTypeTransformation transforming PHONE_NUMBER, and input 'My phone number is 206-555-0123', the output would be 'My phone number is '.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "replace_config": {
                                "block": {
                                  "block_types": {
                                    "new_value": {
                                      "block": {
                                        "attributes": {
                                          "boolean_value": {
                                            "description": "A boolean value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          },
                                          "day_of_week_value": {
                                            "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "float_value": {
                                            "description": "A float value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "integer_value": {
                                            "description": "An integer value (int64 format)",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "string_value": {
                                            "description": "A string value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "timestamp_value": {
                                            "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "block_types": {
                                          "date_value": {
                                            "block": {
                                              "attributes": {
                                                "day": {
                                                  "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "month": {
                                                  "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "year": {
                                                  "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a whole or partial calendar date.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "time_value": {
                                            "block": {
                                              "attributes": {
                                                "hours": {
                                                  "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minutes": {
                                                  "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "nanos": {
                                                  "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "seconds": {
                                                  "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description": "Represents a time of day.",
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Replace each input value with a given value.\nThe 'new_value' block must only contain one argument. For example when replacing the contents of a string-type field, only 'string_value' should be set.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replace with a specified value.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "replace_dictionary_config": {
                                "block": {
                                  "block_types": {
                                    "word_list": {
                                      "block": {
                                        "attributes": {
                                          "words": {
                                            "description": "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description": "A list of words to select from for random replacement. The [limits](https://cloud.google.com/dlp/limits) page contains details about the size limits of dictionaries.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "Replace with a value randomly drawn (with replacement) from a dictionary.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "time_part_config": {
                                "block": {
                                  "attributes": {
                                    "part_to_extract": {
                                      "description": "The part of the time to keep. Possible values: [\"YEAR\", \"MONTH\", \"DAY_OF_MONTH\", \"DAY_OF_WEEK\", \"WEEK_OF_YEAR\", \"HOUR_OF_DAY\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "For use with Date, Timestamp, and TimeOfDay, extract or preserve a portion of the value.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "Apply the transformation to the entire field.\nThe 'primitive_transformation' block must only contain one argument, corresponding to the type of transformation.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Transform the record by applying various field transformations.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "record_suppressions": {
                    "block": {
                      "block_types": {
                        "condition": {
                          "block": {
                            "block_types": {
                              "expressions": {
                                "block": {
                                  "attributes": {
                                    "logical_operator": {
                                      "description": "The operator to apply to the result of conditions. Default and currently only supported value is AND. Default value: \"AND\" Possible values: [\"AND\"]",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "conditions": {
                                      "block": {
                                        "block_types": {
                                          "conditions": {
                                            "block": {
                                              "attributes": {
                                                "operator": {
                                                  "description": "Operator used to compare the field or infoType to the value. Possible values: [\"EQUAL_TO\", \"NOT_EQUAL_TO\", \"GREATER_THAN\", \"LESS_THAN\", \"GREATER_THAN_OR_EQUALS\", \"LESS_THAN_OR_EQUALS\", \"EXISTS\"]",
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "block_types": {
                                                "field": {
                                                  "block": {
                                                    "attributes": {
                                                      "name": {
                                                        "description": "Name describing the field.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "description": "Field within the record this condition is evaluated against.",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "min_items": 1,
                                                  "nesting_mode": "list"
                                                },
                                                "value": {
                                                  "block": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "description": "A boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "bool"
                                                      },
                                                      "day_of_week_value": {
                                                        "description": "Represents a day of the week. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "float_value": {
                                                        "description": "A float value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "number"
                                                      },
                                                      "integer_value": {
                                                        "description": "An integer value (int64 format)",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "description": "A string value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "timestamp_value": {
                                                        "description": "A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "block_types": {
                                                      "date_value": {
                                                        "block": {
                                                          "attributes": {
                                                            "day": {
                                                              "description": "Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "month": {
                                                              "description": "Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "year": {
                                                              "description": "Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            }
                                                          },
                                                          "description": "Represents a whole or partial calendar date.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      },
                                                      "time_value": {
                                                        "block": {
                                                          "attributes": {
                                                            "hours": {
                                                              "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "minutes": {
                                                              "description": "Minutes of hour of day. Must be from 0 to 59.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "nanos": {
                                                              "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            },
                                                            "seconds": {
                                                              "description": "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "number"
                                                            }
                                                          },
                                                          "description": "Represents a time of day.",
                                                          "description_kind": "plain"
                                                        },
                                                        "max_items": 1,
                                                        "nesting_mode": "list"
                                                      }
                                                    },
                                                    "description": "Value to compare against. [Mandatory, except for EXISTS tests.]",
                                                    "description_kind": "plain"
                                                  },
                                                  "max_items": 1,
                                                  "nesting_mode": "list"
                                                }
                                              },
                                              "description": "A collection of conditions.",
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description": "Conditions to apply to the expression.",
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description": "An expression, consisting of an operator and conditions.",
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "A condition that when it evaluates to true will result in the record being evaluated to be suppressed from the transformed content.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Configuration defining which records get suppressed entirely. Records that match any suppression rule are omitted from the output.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Treat the dataset as structured. Transformations can be applied to specific locations within structured datasets, such as transforming a column within a table.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration of the deidentify template",
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

func GoogleDataLossPreventionDeidentifyTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataLossPreventionDeidentifyTemplate), &result)
	return &result
}
