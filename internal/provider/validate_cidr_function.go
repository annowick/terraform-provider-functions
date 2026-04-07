// Copyright IBM Corp. 2021, 2025
// Copyright github.com/annowick 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &ValidateCidrFunction{}

type ValidateCidrFunction struct{}

func NewValidateCidrFunction() function.Function {
	return &ValidateCidrFunction{}
}

func (f *ValidateCidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "validate_cidr"
}

func (f *ValidateCidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Validate an IPv4 CIDR notation string.",
		Description: "Accepts a string and returns `true` if it is a valid IPv4 CIDR notation (e.g. `\"192.168.0.0/24\"`), `false` otherwise.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "The string to validate as an IPv4 CIDR notation.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *ValidateCidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var cidr string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))
	if resp.Error != nil {
		return
	}

	ip, _, err := net.ParseCIDR(cidr)
	result := err == nil && ip.To4() != nil

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
