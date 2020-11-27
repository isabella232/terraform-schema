package schema

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

func terraformBlockSchema(v *version.Version) *schema.BlockSchema {
	bs := &schema.BlockSchema{
		Description: lang.Markdown("Terraform block used to configure some high-level behaviors of Terraform"),
		Body: &schema.BodySchema{
			Attributes: map[string]*schema.AttributeSchema{
				"required_version": {
					Expr: schema.ExprSchema{
						schema.LiteralValueExpr{Type: cty.String},
					},
					IsOptional: true,
					Description: lang.Markdown("Constraint to specify which versions of Terraform can be used " +
						"with this configuration, e.g. `~> 0.12`"),
				},
			},
			Blocks: map[string]*schema.BlockSchema{
				"backend": {
					Description: lang.Markdown("Backend configuration which defines exactly where and how " +
						"operations are performed, where state snapshots are stored, etc."),
					Labels: []*schema.LabelSchema{
						{
							Name:        "type",
							Description: lang.Markdown("Backend Type"),
							IsDepKey:    true,
						},
					},
					MaxItems: 1,
				},
				"required_providers": {
					Description: lang.Markdown("What provider version to use within this configuration"),
					Body: &schema.BodySchema{
						AnyAttribute: &schema.AttributeSchema{
							Expr: schema.ExprSchema{
								schema.LiteralValueExpr{Type: cty.String},
							},
							Description: lang.Markdown("Version constraint"),
						},
					},
					MaxItems: 1,
				},
			},
		},
	}

	if v.GreaterThanOrEqual(v0_12_18) {
		bs.Body.Attributes["experiments"] = &schema.AttributeSchema{
			Expr: schema.ExprSchema{
				schema.TupleExpr{
					// TODO: StaticTraversalExpr{}?
				},
			},
			IsOptional:  true,
			Description: lang.Markdown("A tuple of experimental language features to enable"),
		}
	}

	if v.GreaterThanOrEqual(v0_12_20) {
		bs.Body.Blocks["required_providers"].Body = &schema.BodySchema{
			AnyAttribute: &schema.AttributeSchema{
				Expr: schema.ExprSchema{
					schema.ObjectExpr{
						Attributes: map[string]schema.ExprSchema{
							"version": {
								schema.LiteralValueExpr{Type: cty.String},
							},
						},
					},
					schema.LiteralValueExpr{Type: cty.String},
				},
				Description: lang.Markdown("Version constraint"),
			},
		}
	}

	return bs
}
