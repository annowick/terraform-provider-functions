// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestMapProductFunction_Basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::map_product({"a" = "v1"}, {"b" = "v2"}, "|")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValueAtPath(
						"test",
						tfjsonpath.New("a|b"),
						knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("v1"),
							knownvalue.StringExact("v2"),
						}),
					),
				},
			},
		},
	})
}

func TestMapProductFunction_EmptyMaps(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::map_product({}, {}, "|")
        }
        `,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.MapExact(map[string]knownvalue.Check{}),
					),
				},
			},
		},
	})
}

func TestMapProductFunction_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
        output "test" {
          value = provider::andrzejtest::map_product(null, {}, "|")
        }
        `,
				// The parameter does not enable AllowNullValue
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}
