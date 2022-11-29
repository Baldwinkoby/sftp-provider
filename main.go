package main

import (
	"context"
	"log"

	provider "github.com/hashicorp/terraform-provider-scaffolding-framework/sftp"
	//"google.golang.org/protobuf/internal/version"

	//"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{
		Address:         "registry.terraform.io/hashicorp/scaffolding",
		Debug:           false,
		ProtocolVersion: 0,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
