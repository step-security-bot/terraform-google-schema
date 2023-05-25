package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataLossPreventionDeidentifyTemplate = `{
  "block": {
    "attributes": {
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
      }
    },
    "block_types": {
      "deidentify_config": {
        "block": {
          "block_types": {
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
                                          "character_to_skip": {
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
                                              "description": "Kms wrapped key",
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
                                              "description": "Transient crypto key",
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
                                              "description": "Unwrapped crypto key",
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
                                              "description": "Kms wrapped key",
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
                                              "description": "Transient crypto key",
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
                                              "description": "Unwrapped crypto key",
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
                                        "description": "Replace each input value with a given value.",
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
                              }
                            },
                            "description": "Primitive transformation to apply to the infoType.",
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
                "description": "Specifies free-text based transformations to be applied to the dataset.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
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
