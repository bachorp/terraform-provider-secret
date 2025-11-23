package secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type SecretProvider struct {
	version string
}

func (p *SecretProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "secret"
	resp.Version = p.version
}

func (p *SecretProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *SecretProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *SecretProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource {
			return &SecretResource{}
		},
	}
}

func (p *SecretProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SecretProvider{
			version,
		}
	}
}
