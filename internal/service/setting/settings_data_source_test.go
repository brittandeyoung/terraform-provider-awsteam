package setting_test

import (
	"testing"

	"github.com/brittandeyoung/terraform-provider-awsteam/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSettingSettingsDataSource(t *testing.T) {
	dataSource := "data.awsteam_setting_settings.test"
	resource.Test(t, resource.TestCase{
		// PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSource, "id", "settings"),
					// resource.TestCheckResourceAttrSet(dataSource, "approval"),
					// resource.TestCheckResourceAttrSet(dataSource, "comments"),
					// resource.TestCheckResourceAttrSet(dataSource, "max_request_duration"),
					// resource.TestCheckResourceAttrSet(dataSource, "request_expiry_timeout"),
					// resource.TestCheckResourceAttrSet(dataSource, "ses_notifications_enabled"),
					// resource.TestCheckResourceAttrSet(dataSource, "sns_notifications_enabled"),
					// resource.TestCheckResourceAttrSet(dataSource, "slack_notifications_enabled"),
					// resource.TestCheckResourceAttrSet(dataSource, "ses_source_email"),
					// resource.TestCheckResourceAttrSet(dataSource, "ses_source_arn"),
					// resource.TestCheckResourceAttrSet(dataSource, "slack_token"),
					// resource.TestCheckResourceAttrSet(dataSource, "team_admin_group"),
					// resource.TestCheckResourceAttrSet(dataSource, "team_auditor_group"),
					// resource.TestCheckResourceAttrSet(dataSource, "ticket_no"),
					// resource.TestCheckResourceAttrSet(dataSource, "modified_by"),
					// resource.TestCheckResourceAttrSet(dataSource, "created_at"),
					// resource.TestCheckResourceAttrSet(dataSource, "updated_at"),
				),
			},
		},
	})
}

const testAccExampleDataSourceConfig = `
data "awsteam_setting_settings" "test" {}
`
