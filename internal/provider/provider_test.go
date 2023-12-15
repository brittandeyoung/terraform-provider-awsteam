package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"awsteam": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	// 	testAccProviderConfigure.Do(func() {
	// 		envvar.FailIfAllEmpty(t, []string{envvar.Profile, envvar.AccessKeyId, envvar.ContainerCredentialsFullURI}, "credentials for running acceptance testing")

	// 		if os.Getenv(envvar.AccessKeyId) != "" {
	// 			envvar.FailIfEmpty(t, envvar.SecretAccessKey, "static credentials value when using "+envvar.AccessKeyId)
	// 		}

	// 		// Setting the AWS_DEFAULT_REGION environment variable here allows all tests to omit
	// 		// a provider configuration with a region. This defaults to us-west-2 for provider
	// 		// developer simplicity and has been in the codebase for a very long time.
	// 		//
	// 		// This handling must be preserved until either:
	// 		//   * AWS_DEFAULT_REGION is required and checked above (should mention us-west-2 default)
	// 		//   * Region is automatically handled via shared AWS configuration file and still verified
	// 		region := Region()
	// 		os.Setenv(envvar.DefaultRegion, region)

	// diags := provider.Configure(ctx, terraformsdk.NewResourceConfigRaw(nil))
	// if err := sdkdiag.DiagnosticsError(diags); err != nil {
	// 	t.Fatalf("configuring provider: %s", err)
	// }
	//		})
	//	}
}
