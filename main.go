package main

import (
	"context"
	"flag"
	"log"

	"github.com/bachorp/terraform-provider-secret/secret"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	err := providerserver.Serve(
		context.Background(),
		secret.New(version),
		providerserver.ServeOpts{Address: "registry.terraform.io/bachorp/secret", Debug: debug, ProtocolVersion: 6},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
