package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/terraform-schema/internal/schema/refscope"
	"github.com/zclconf/go-cty/cty"
)

var outputBlockSchema = &schema.BlockSchema{
	Labels: []*schema.LabelSchema{
		{
			Name:        "name",
			Description: lang.PlainText("Output Name"),
		},
	},
	Description: lang.PlainText("Output value for consumption by another module or a human interacting via the UI"),
	Body: &schema.BodySchema{
		Attributes: map[string]*schema.AttributeSchema{
			"description": {
				Expr: schema.ExprSchema{
					schema.LiteralValueExpr{Type: cty.String},
				},
				IsOptional:  true,
				Description: lang.PlainText("Human-readable description of the output (for documentation and UI)"),
			},
			"value": {
				Expr: schema.ExprSchema{
					schema.ScopeTraversalExpr{ScopeId: refscope.VariableBlock},
					schema.ScopeTraversalExpr{ScopeId: refscope.LocalAttr},
					schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock},
					schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock},
					schema.ScopeTraversalExpr{ScopeId: refscope.ModuleBlock},
					schema.LiteralValueExpr{},
				},
				IsRequired:  true,
				Description: lang.PlainText("Value, typically a reference to an attribute of a resource or a data source"),
			},
			"sensitive": {
				Expr: schema.ExprSchema{
					schema.LiteralValueExpr{Type: cty.Bool},
				},
				IsOptional:  true,
				Description: lang.PlainText("Whether the output contains sensitive material and should be hidden in the UI"),
			},
			"depends_on": {
				Expr: schema.ExprSchema{
					schema.TupleExpr{
						Exprs: []schema.Expr{
							schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock},
							schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock},
							schema.ScopeTraversalExpr{ScopeId: refscope.ModuleBlock},
						},
					},
				},
				IsOptional:  true,
				Description: lang.PlainText("Set of references to hidden dependencies (e.g. resources or data sources)"),
			},
		},
	},
	Reference: &schema.BlockReference{
		ScopeId: refscope.OutputBlock,
		Type: schema.ReferenceTypes{
			&schema.InferredRefType{AttrName: "value"},
		},
		Address: schema.Address{
			schema.StaticStep{Value: "output"},
			schema.LabelValueStep{Index: 0},
		},
	},
}
