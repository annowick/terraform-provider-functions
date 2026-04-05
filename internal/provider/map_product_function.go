// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &MapProductFunction{}

type MapProductFunction struct{}

func NewMapProductFunction() function.Function {
	return &MapProductFunction{}
}

func (f *MapProductFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "map_product"
}

func (f *MapProductFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Create a product of two maps",
		Description: "Given two map(string) arguments, return a map with keys that are a product of input keys joined by `separator`, and value is a two-element list containing original values",
		Parameters: []function.Parameter{
			function.MapParameter{
				Name:        "first",
				Description: "First map",
				ElementType: types.StringType,
			},
			function.MapParameter{
				Name:        "second",
				Description: "Second map",
				ElementType: types.StringType,
			},
			function.StringParameter{
				Name:        "separator",
				Description: "separator to join keys on",
			},
		},
		Return: function.MapReturn{
			ElementType: types.ListType{ElemType: types.StringType},
		},
	}
}

func (f *MapProductFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var first map[string]string
	var second map[string]string
	var separator string

	// Read Terraform argument data into the variables
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &first, &second, &separator))

	result := make(map[string][]attr.Value, len(first)*len(second))
	for firstKey, firstValue := range first {
		for secondKey, secondValue := range second {
			key := strings.Join([]string{firstKey, secondKey}, separator)
			result[key] = []attr.Value{
				types.StringValue(firstValue),
				types.StringValue(secondValue),
			}
		}
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
