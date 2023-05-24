package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeAddresses = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "address_type": "string",
              "description": "string",
              "name": "string",
              "region": "string",
              "self_link": "string",
              "status": "string"
            }
          ]
        ]
      },
      "filter": {
        "description": "Filter sets the optional parameter \"filter\": A filter expression that\nfilters resources listed in the response. The expression must specify\nthe field name, an operator, and the value that you want to use for\nfiltering. The value must be a string, a number, or a boolean. The\noperator must be either \"=\", \"!=\", \"\u003e\", \"\u003c\", \"\u003c=\", \"\u003e=\" or \":\". For\nexample, if you are filtering Compute Engine instances, you can\nexclude instances named \"example-instance\" by specifying \"name !=\nexample-instance\". The \":\" operator can be used with string fields to\nmatch substrings. For non-string fields it is equivalent to the \"=\"\noperator. The \":*\" comparison can be used to test whether a key has\nbeen defined. For example, to find all objects with \"owner\" label\nuse: \"\"\" labels.owner:* \"\"\" You can also filter nested fields. For\nexample, you could specify \"scheduling.automaticRestart = false\" to\ninclude instances only if they are not scheduled for automatic\nrestarts. You can use filtering on nested fields to filter based on\nresource labels. To filter on multiple expressions, provide each\nseparate expression within parentheses. For example: \"\"\"\n(scheduling.automaticRestart = true) (cpuPlatform = \"Intel Skylake\")\n\"\"\" By default, each expression is an \"AND\" expression. However, you\ncan include \"AND\" and \"OR\" expressions explicitly. For example: \"\"\"\n(cpuPlatform = \"Intel Skylake\") OR (cpuPlatform = \"Intel Broadwell\")\nAND (scheduling.automaticRestart = true) \"\"\"",
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
      "project": {
        "computed": true,
        "description": "The google project in which addresses are listed. Defaults to provider's configuration if missing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "Region that should be considered to search addresses. All regions are considered if missing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeAddressesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeAddresses), &result)
	return &result
}
