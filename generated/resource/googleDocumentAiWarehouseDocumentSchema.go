package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDocumentAiWarehouseDocumentSchema = `{
  "block": {
    "attributes": {
      "display_name": {
        "description": "Name of the schema given by the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "document_is_folder": {
        "description": "Tells whether the document is a folder or a typical document.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the document schema.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_number": {
        "description": "The unique identifier of the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "property_definitions": {
        "block": {
          "attributes": {
            "display_name": {
              "description": "The display-name for the property, used for front-end.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_filterable": {
              "description": "Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_metadata": {
              "description": "Whether the property is user supplied metadata.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_repeatable": {
              "description": "Whether the property can have multiple values.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_required": {
              "description": "Whether the property is mandatory.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_searchable": {
              "description": "Indicates that the property should be included in a global search.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "The name of the metadata property.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retrieval_importance": {
              "description": "Stores the retrieval importance. Possible values: [\"HIGHEST\", \"HIGHER\", \"HIGH\", \"MEDIUM\", \"LOW\", \"LOWEST\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "date_time_type_options": {
              "block": {
                "description": "Date time property. Not supported by CMEK compliant deployment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "enum_type_options": {
              "block": {
                "attributes": {
                  "possible_values": {
                    "description": "List of possible enum values.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "validation_check_disabled": {
                    "description": "Make sure the enum property value provided in the document is in the possile value list during document creation. The validation check runs by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Enum/categorical property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "float_type_options": {
              "block": {
                "description": "Float property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "integer_type_options": {
              "block": {
                "description": "Integer property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "map_type_options": {
              "block": {
                "description": "Map property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "property_type_options": {
              "block": {
                "block_types": {
                  "property_definitions": {
                    "block": {
                      "attributes": {
                        "display_name": {
                          "description": "The display-name for the property, used for front-end.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_filterable": {
                          "description": "Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_metadata": {
                          "description": "Whether the property is user supplied metadata.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_repeatable": {
                          "description": "Whether the property can have multiple values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_required": {
                          "description": "Whether the property is mandatory.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_searchable": {
                          "description": "Indicates that the property should be included in a global search.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "name": {
                          "description": "The name of the metadata property.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "retrieval_importance": {
                          "description": "Stores the retrieval importance. Possible values: [\"HIGHEST\", \"HIGHER\", \"HIGH\", \"MEDIUM\", \"LOW\", \"LOWEST\"]",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "date_time_type_options": {
                          "block": {
                            "description": "Date time property. Not supported by CMEK compliant deployment.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "enum_type_options": {
                          "block": {
                            "attributes": {
                              "possible_values": {
                                "description": "List of possible enum values.",
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "validation_check_disabled": {
                                "description": "Make sure the enum property value provided in the document is in the possile value list during document creation. The validation check runs by default.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description": "Enum/categorical property.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "float_type_options": {
                          "block": {
                            "description": "Float property.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "integer_type_options": {
                          "block": {
                            "description": "Integer property.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "map_type_options": {
                          "block": {
                            "description": "Map property.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "schema_sources": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The schema name in the source.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "processor_type": {
                                "description": "The Doc AI processor type name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The schema source information.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        },
                        "text_type_options": {
                          "block": {
                            "description": "Text property.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        },
                        "timestamp_type_options": {
                          "block": {
                            "description": "Timestamp property. Not supported by CMEK compliant deployment.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Defines the metadata for a schema property.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Nested structured data property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schema_sources": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The schema name in the source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "processor_type": {
                    "description": "The Doc AI processor type name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The schema source information.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "text_type_options": {
              "block": {
                "description": "Text/string property.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "timestamp_type_options": {
              "block": {
                "description": "Timestamp property. Not supported by CMEK compliant deployment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Defines the metadata for a schema property.",
          "description_kind": "plain"
        },
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

func GoogleDocumentAiWarehouseDocumentSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDocumentAiWarehouseDocumentSchema), &result)
	return &result
}
