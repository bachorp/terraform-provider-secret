package secret

import (
	"errors"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSecretResource(t *testing.T) {
	var must_be_imported, err = regexp.Compile("must be imported")
	if err != nil {
		t.Error(err)
		return
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"secret": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config:      `resource "secret_resource" "test" {}`,
				ExpectError: must_be_imported,
			},
			{
				ResourceName:  "secret_resource.test",
				ImportState:   true,
				ImportStateId: "xxx",
				Config:        `resource "secret_resource" "test" {}`,
				ImportStateCheck: func(state []*terraform.InstanceState) error {
					if value := state[0].Attributes["value"]; value != "xxx" {
						return errors.New("Attribute `value` invalid")
					}
					return nil
				},
			},
		},
	})
}
