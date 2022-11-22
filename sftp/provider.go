package sftp

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// New is a helper function to simplify provider server and testing implementation.
func New() tfsdk.Provider {
	return &SftpgoProvider{}
}

// SftpgoProvider defines the provider implementation.
type SftpgoProvider struct {
	configured bool
	//client     *api_pritunl_wrapper.Client
}

// Metadata returns the provider type name.
func (p *SftpgoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "Sftpgo"
}

// GetSchema defines the provider-level schema for configuration data.
func (p *SftpgoProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"host": {
				Type:     types.StringType,
				Optional: true,
			},
			"username": {
				Type:     types.StringType,
				Optional: true,
			},
			"password": {
				Type:      types.StringType,
				Optional:  true,
				Sensitive: true,
			},
		},
	}, nil
}

// SftpgoProviderModel describes the provider data model.
type SftpgoProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *SftpgoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring SFTP Provider")

	// Retrieve provider data from configuration
	var config SftpgoProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("SFTPGO_HOST")
	if host == "" {
		if config.Host.ValueString() == "" {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create client: The host is missing.",
			)
			return
		}
		host = config.Host.ValueString()
	}

	username := os.Getenv("SFTPGO_USERNAME")
	if username == "" {
		if config.Username.ValueString() == "" {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create client: The username is missing.",
			)
			return
		}
		host = config.Username.ValueString()
	}

	password := os.Getenv("SFTPGO_PASSWORD")
	if password == "" {
		if config.Password.ValueString() == "" {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create client: The password is missing.",
			)
			return
		}
		host = config.Password.ValueString()
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing SFTPGO API Host",
			"The provider cannot create the SFTPGO API client as there is a missing or empty value for the SFTPGO API host. "+
				"Set the host value in the configuration or use the SFTPGO_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing SFTPGO API Username",
			"The provider cannot create the SFTPGO API client as there is a missing or empty value for the SFTPGO API username. "+
				"Set the username value in the configuration or use the SFTPGO_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing SFTPGO API Password",
			"The provider cannot create the SFTPGO API client as there is a missing or empty value for the SFTPGO API password. "+
				"Set the password value in the configuration or use the SFTPGO_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "SFTPGO_host", host)
	ctx = tflog.SetField(ctx, "SFTPGO_username", username)
	ctx = tflog.SetField(ctx, "SFTPGO_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "SFTPGO_password")

	tflog.Debug(ctx, "Creating SFTPGO client")

	// Create a new SFTPGO client using the configuration values
	client, err := api.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create SFTPGO API Client",
			"An unexpected error occurred when creating the SFTPGO API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"SFTPGO Client Error: "+err.Error(),
		)
		return
	}

	// Make the SFTPGO client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured SFTPGO client", map[string]any{"success": true})

}

/*
// GetResources - Defines provider resources
func (p *SftpgoProvider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"stfpgo_user": resourceUserType{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *SftpgoProvider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
*/
