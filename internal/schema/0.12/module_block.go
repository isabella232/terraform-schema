package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/terraform-schema/internal/schema/refscope"
	"github.com/zclconf/go-cty/cty"
)

var moduleBlockSchema = &schema.BlockSchema{
	Labels: []*schema.LabelSchema{
		{
			Name:        "name",
			Description: lang.PlainText("Reference Name"),
		},
	},
	Description: lang.PlainText("Module block to call a locally or remotely stored module"),
	Body: &schema.BodySchema{
		Attributes: map[string]*schema.AttributeSchema{
			"source": {
				Expr: schema.ExprSchema{
					schema.LiteralValueExpr{Type: cty.String},
				},
				Description: lang.Markdown("Source where to load the module from, " +
					"a local directory (e.g. `./module`) or a remote address - e.g. " +
					"`hashicorp/consul/aws` (Terraform Registry address) or " +
					"`github.com/hashicorp/example` (GitHub)"),
				IsRequired: true,
				IsDepKey:   true,
			},
			"version": {
				Expr: schema.ExprSchema{
					schema.LiteralValueExpr{Type: cty.String},
				},
				IsOptional: true,
				Description: lang.Markdown("Constraint to set the version of the module, e.g. `~> 1.0`." +
					" Only applicable to modules in a module registry."),
			},
			"providers": {
				Expr: schema.ExprSchema{
					schema.MapExpr{
						Key:  schema.LiteralValueExpr{Type: cty.String},
						Expr: schema.ScopeTraversalExpr{ScopeId: refscope.ProviderBlock},
					},
				},
				IsOptional:  true,
				Description: lang.Markdown("Explicit mapping of providers which the module uses"),
			},
		},
	},
	// TODO: DependentBody (based on variables)
	// TODO: DependentReference (based on outputs)
	Reference: &schema.BlockReference{
		ScopeId: refscope.ModuleBlock,
		Address: schema.Address{
			schema.StaticStep{Value: "module"},
			schema.LabelValueStep{Index: 0},
		},
	},
}
