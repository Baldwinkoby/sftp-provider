package admin

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	//"github.com/hashicorp/terraform-plugin-go/tftypes"
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

// New resource instance
func (r resourceAdminType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceUser{
		p: *(p.(*provider)),
	}, nil
}

type resourceUser struct {
	p provider
}

// Create a new resource
func (r resourceUser) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from plan
	var plan User
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new user
	var user = pritunl_wrapper.User{
		Name:             plan.Name.Value,
		Email:            plan.Email.Value,
		OrganizationName: plan.OrganizationName.Value,
	}

	createdUser, err := r.p.client.CreateUser(user)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating user",
			"Could not create user, unexpected error: "+err.Error(),
		)
		return
	}

	// Generate resource state struct
	var result = Admin{
		ID:               types.String{Value: createdUser.ID},
		Name:             types.String{Value: createdUser.Name},
		Email:            types.String{Value: createdUser.Email},
		OrganizationName: types.String{Value: createdUser.OrganizationName},
		OrganizationID:   types.String{Value: createdUser.OrganizationID},
		OtpSecret:        types.String{Value: createdUser.OtpSecret},
		LastUpdated:      types.String{Value: string(time.Now().Format(time.RFC850))},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
