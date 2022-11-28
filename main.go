package main

import (
	"context"

	provider "github.com/hashicorp/terraform-provider-scaffolding-framework/sftp"
	//"google.golang.org/protobuf/internal/version"

	//"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
)

func main() {
	providerserver.Serve(context.Background(), provider.SftpgoProvider(), providerserver.ServeOpts{
		//Name: "pritunlwrapper",
	})
}
