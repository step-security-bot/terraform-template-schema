package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-template-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestTemplateDirSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.TemplateDirSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
