package secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecretResource struct{}

func (r *SecretResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "secret_resource"
}

func (r *SecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:             0,
		Description:         "Wrapper for a sensitive string `value` to be supplied using `import`",
		MarkdownDescription: "Wrapper for a sensitive string `value` to be supplied using `import`",
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description:         "The contained secret",
				MarkdownDescription: "The contained secret",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

func (r *SecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError("This resource must be imported", "Resources of this type are *computed* and can only be created using `import` (received `Create`).")
}

func (r *SecretResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
}

func (r *SecretResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
}

func (r *SecretResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
}

func (r *SecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("value"), req, resp)
}

func (r *SecretResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var value types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("value"), &value)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if value.IsNull() {
		resp.Diagnostics.AddError("This resource must be imported", "Resources of this type are *computed* and can only be created using `import` (received `ModifyPlan`).")
	}
}
