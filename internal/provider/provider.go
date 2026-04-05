// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure AndrzejTestProvider satisfies various provider interfaces.
var _ provider.Provider = &AndrzejTestProvider{}
var _ provider.ProviderWithFunctions = &AndrzejTestProvider{}

// AndrzejTestProvider defines the provider implementation.
type AndrzejTestProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AndrzejTestProviderModel describes the provider data model.
type AndrzejTestProviderModel struct {
	// Endpoint types.String `tfsdk:"description"`
}

func (p *AndrzejTestProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "andrzejtest"
	resp.Version = p.version
}

func (p *AndrzejTestProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Not needed for functions-only provider.
func (p *AndrzejTestProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *AndrzejTestProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// NewExampleResource,
	}
}

func (p *AndrzejTestProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// NewExampleDataSource,
	}
}

func (p *AndrzejTestProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewMapProductFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AndrzejTestProvider{
			version: version,
		}
	}
}
