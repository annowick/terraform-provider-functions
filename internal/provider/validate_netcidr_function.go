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
var _ function.Function = &ValidateNetCidrFunction{}

type ValidateNetCidrFunction struct{}

func NewValidateNetCidrFunction() function.Function {
	return &ValidateNetCidrFunction{}
}

func (f *ValidateNetCidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "validate_netcidr"
}

func (f *ValidateNetCidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Validate that an IPv4 CIDR string is a true network address (no host bits set).",
		Description: "Accepts a string and returns `true` if it is a valid IPv4 CIDR notation AND the address has no host bits set " +
			"(i.e. it is a proper network address). For example, `\"11.12.13.0/24\"` returns `true`, " +
			"but `\"11.12.13.0/16\"` returns `false` because the `.13.0` portion lies outside the /16 network boundary.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "The string to validate as an IPv4 network CIDR notation.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *ValidateNetCidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var cidr string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))
	if resp.Error != nil {
		return
	}

	ip, network, err := net.ParseCIDR(cidr)
	result := err == nil && ip.To4() != nil && ip.Equal(network.IP)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
