package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const templateCloudinitConfig = `{
  "block": {
    "attributes": {
      "base64_encode": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gzip": {
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
      "rendered": {
        "computed": true,
        "description": "rendered cloudinit configuration",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "part": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "content_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filename": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "merge_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func TemplateCloudinitConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(templateCloudinitConfig), &result)
	return &result
}
