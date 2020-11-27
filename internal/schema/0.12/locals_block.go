package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/terraform-schema/internal/schema/refscope"
)

var localsBlockSchema = &schema.BlockSchema{
	Description: lang.Markdown("Local values assigning names to expressions, so you can use these multiple times without repetition\n" +
		"e.g. `service_name = \"forum\"`"),
	Body: &schema.BodySchema{
		AnyAttribute: &schema.AttributeSchema{
			Expr: schema.ExprSchema{
				schema.LiteralValueExpr{},
				schema.ScopeTraversalExpr{ScopeId: refscope.LocalAttr},
			},
			Reference: &schema.AttrReference{
				ScopeId: refscope.LocalAttr,
				Type: schema.ReferenceTypes{
					&schema.RefTypeFromConstraint{AttrName: "type"},
					&schema.InferredRefType{AttrName: "default"},
				},
				Address: schema.Address{
					schema.StaticStep{Value: "local"},
					schema.AttrNameStep{},
				},
			},
		},
	},
}
