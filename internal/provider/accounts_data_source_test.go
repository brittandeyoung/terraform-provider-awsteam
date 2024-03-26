package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccountsDataSource_basic(t *testing.T) {
	dataSourceName := "data.awsteam_accounts.test"

	// This environment variable should be set to the number of accounts that are expected to be returned by the API
	expectedAccountsCountVar := "AWSTEAM_TESTS_EXPECTED_ACCOUNTS_COUNT"
	expectedAccountsCount := os.Getenv(expectedAccountsCountVar)
	if expectedAccountsCount == "" {
		t.Skipf("Skipping Accounts Tests, Environment variable %s is not set.", expectedAccountsCountVar)
	}

	// This environment variable should be set to a name of one of the accounts expected to be returned.
	expectedAccountsNameVar := "AWSTEAM_TESTS_EXPECTED_ACCOUNT_NAME"
	expectedAccountsName := os.Getenv(expectedAccountsNameVar)
	if expectedAccountsName == "" {
		t.Skipf("Skipping Accounts Tests, Environment variable %s is not set.", expectedAccountsNameVar)
	}

	// This environment variable should be set to the account ID of the account name provided expected to be returned.
	expectedAccountsIdVar := "AWSTEAM_TESTS_EXPECTED_ACCOUNT_ID"
	expectedAccountsId := os.Getenv(expectedAccountsIdVar)
	if expectedAccountsId == "" {
		t.Skipf("Skipping Accounts Tests, Environment variable %s is not set.", expectedAccountsIdVar)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountsDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "id", "accounts"),
					resource.TestCheckResourceAttr(dataSourceName, "accounts.#", expectedAccountsCount),
					resource.TestCheckTypeSetElemNestedAttrs(dataSourceName, "accounts.*",
						map[string]string{
							"id":   expectedAccountsId,
							"name": expectedAccountsName,
						}),
				),
			},
		},
	})
}

func testAccAccountsDataSourceConfig() string {
	return `data "awsteam_accounts" "test" {}`
}
