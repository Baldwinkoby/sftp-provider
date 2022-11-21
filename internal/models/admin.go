package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Admin struct {
	ID             types.Int64    `tfsdk:"id"`
	Status         types.Int64    `tfsdk:"status"`
	Username       types.String   `tfsdk:"username"`
	Description    types.String   `tfsdk:"description"`
	Password       types.String   `tfsdk:"password"`
	Email          types.String   `tfsdk:"email"`
	Permission     []types.String `tfsdk:"permissions"`
	Filters        AdminFilters   `tfsdk:"filters"`
	AdditionalInfo types.String   `tfsdk:"additional_info"`
}

type AdminFilters struct {
	AllowList []string `json:"allow_list"`
}
