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
		Summary: "Compute the Cartesian product of two string maps.",
		Description: "Accepts two `map(string)` arguments and a separator string. " +
			"Returns a `map(list(string))` whose keys are every combination of input keys joined by the separator, " +
			"and whose values are two-element lists containing the corresponding values from each input map. " +
			"For example, `map_product({a=\"v1\"}, {b=\"v2\"}, \"|\")` returns `{\"a|b\": [\"v1\", \"v2\"]}`.",
		Parameters: []function.Parameter{
			function.MapParameter{
				Name:        "first",
				Description: "The first map(string). Its keys form the left-hand side of each product key.",
				ElementType: types.StringType,
			},
			function.MapParameter{
				Name:        "second",
				Description: "The second map(string). Its keys form the right-hand side of each product key.",
				ElementType: types.StringType,
			},
			function.StringParameter{
				Name:        "separator",
				Description: "The string used to join a key from `first` and a key from `second`, e.g. `\"|\"` produces `\"a|b\"`.",
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
