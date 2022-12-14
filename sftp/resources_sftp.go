package sftp

import (
	"context"
	"time"

	//"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/models"
)

type resourceAdminType struct{}

// Admin Resource schema
func (r resourceAdminType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"status": {
				Type:     types.Int64Type,
				Required: true,
			},
			"username": {
				Type:     types.StringType,
				Required: true,
				//ForceNew: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"password": {
				Type:      types.StringType,
				Required:  true,
				Sensitive: true,
				//StateFunc: true,
			},
			"email": {
				Type:     types.StringType,
				Computed: true,
			},
			"permissions": {
				Required: true,
				Type:     types.ListType.ElementType(types.ListType{}),
			},
			"filters": {
				Required: true,
				//Type:     types.ListType.ElementType(types.ListType{}),
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"allow_list": {
						Description: "Allowed IP's",
						Type:        types.StringType,
						Required:    true,
					},
				}),
			},
			"additional_info": {
				Required: true,
				Type:     types.StringType},
		},
	}, nil
}

/*
// New resource instance
func (r resourceAdminType) NewResource(_ context.Context, p tfsdk.Schema) (tfsdk.Resource, diag.Diagnostics) {
	return resourceUser{
		p: *(p.(*Sftpgoprovider)),
	}, nil
}
*/

type resourceUser struct {
	p SftpgoProvider
}

// Create a new resource
func (r resourceUser) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var state *models.User

	// Retrieve values from plan
	var plan *models.User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

	// // Create new user
	// var user = models.User{
	// 	Id:       plan.Id,
	// 	Username: plan.Username,
	// 	Password: plan.Password,
	// }

	// createdUser, err := r.p.client.CreateUser(ctx, user)
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Error creating user",
	// 		"Could not create user, unexpected error: "+err.Error(),
	// 	)
	// 	return
	// }

	// Generate resource state struct
	var result = models.User{
		Id:                types.Int64{},
		Status:            types.Int64{},
		Username:          types.String{},
		Description:       types.String{},
		ExpirationDate:    types.Float64{},
		Password:          types.String{},
		PublicKeys:        []types.String{},
		HomeDir:           types.String{},
		Uid:               types.Int64{},
		Gid:               types.Int64{},
		MaxSessions:       types.Int64{},
		QuotaSize:         types.Float64{},
		QuotaFiles:        types.Int64{},
		VirtualFolders:    []models.VirtualFolder{},
		UploadBandwidth:   types.Int64{},
		DownloadBandwidth: types.Int64{},
		Filters:           &models.Filters{},
		Filesystem:        &models.Filesystem{},
		AdditionalInfo:    types.String{},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceUser) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *resourceUser

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Sftpgo, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// UPDATE.GO
func (r resourceUser) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state models.User

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}

// DELETE.GO USERS
func (r resourceUser) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.User

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	//Get User ID from API
	UserID := state.Id.String()
	//Delete Existing resource
	err := r.p.client.DeleteUser(ctx, UserID)

	if err != nil {
		return
	}
	resp.State.RemoveResource(ctx)

}

func (r resourceUser) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
