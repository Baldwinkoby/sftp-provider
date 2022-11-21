package main

import (
	"context"

	pritunl_wrapper "terraform-provider-pritunl-wrapper/pritunl-wrapper"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), pritunl_wrapper.New, tfsdk.ServeOpts{
		Name: "SFTPGO",
	})
}
