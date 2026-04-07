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
var _ function.Function = &ValidateHostCidrFunction{}

type ValidateHostCidrFunction struct{}

func NewValidateHostCidrFunction() function.Function {
	return &ValidateHostCidrFunction{}
}

func (f *ValidateHostCidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "validate_hostcidr"
}

func (f *ValidateHostCidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Validate that an IPv4 CIDR string is a valid host address.",
		Description: "Accepts a string and returns `true` if it is a valid IPv4 CIDR notation AND the address is a usable host address: " +
			"it must not be the network address (all host bits zero, e.g. `\"10.10.10.0/24\"`) " +
			"nor the broadcast address (all host bits one, e.g. `\"10.10.10.255/24\"`).",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "The string to validate as an IPv4 host CIDR notation.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *ValidateHostCidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var cidr string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))
	if resp.Error != nil {
		return
	}

	ip, network, err := net.ParseCIDR(cidr)

	var result bool
	if err == nil && ip.To4() != nil {
		// Compute the broadcast address: network address with all host bits set to 1.
		broadcast := make(net.IP, len(network.IP))
		for i := range network.IP {
			broadcast[i] = network.IP[i] | ^network.Mask[i]
		}
		result = !ip.Equal(network.IP) && !ip.Equal(broadcast)
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
