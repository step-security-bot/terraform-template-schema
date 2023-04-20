package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-template-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestTemplateCloudinitConfigSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.TemplateCloudinitConfigSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
