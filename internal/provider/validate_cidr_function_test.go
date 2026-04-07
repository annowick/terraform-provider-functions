// Copyright IBM Corp. 2021, 2025
// Copyright github.com/annowick 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestValidateCidrFunction_ValidHostCidr(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("192.168.1.5/24")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestValidateCidrFunction_ValidNetworkCidr(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("10.0.0.0/8")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
				},
			},
		},
	})
}

func TestValidateCidrFunction_InvalidNoCidr(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("192.168.1.5")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestValidateCidrFunction_InvalidPrefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("192.168.1.5/33")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestValidateCidrFunction_InvalidIPv6(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("2001:db8::/32")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
				},
			},
		},
	})
}

func TestValidateCidrFunction_EmptyString(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::validate_cidr("")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
				},
			},
		},
	})
}
