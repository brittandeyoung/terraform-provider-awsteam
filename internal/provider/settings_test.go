package provider

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSettings_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Resource": {
			"basic": testAccSettingsResource_basic,
		},
		"DataSource": {
			"basic": testAccSettingsDataSource_basic,
		},
	}

	RunSerialTests2Levels(t, testCases, 0)
}

func testAccSettingsResource_basic(t *testing.T) {
	resourceName := "awsteam_settings.test"
	teamAdminGroup1 := "Team-Admin-Group"
	teamAuditorGroup1 := "Team-Auditor-Group"
	teamAdminGroup2 := "Team-Admin-Group"
	teamAuditorGroup2 := "Team-Auditor-Group"
	duration := rand.Intn(10)
	expiry := rand.Intn(10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccSettingsResourceConfig(teamAdminGroup1, teamAuditorGroup1, duration, expiry),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "team_admin_group", teamAdminGroup1),
					resource.TestCheckResourceAttr(resourceName, "team_auditor_group", teamAuditorGroup1),
					resource.TestCheckResourceAttr(resourceName, "approval", "false"),
					resource.TestCheckResourceAttr(resourceName, "comments", "false"),
					resource.TestCheckResourceAttr(resourceName, "ses_notifications_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "sns_notifications_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "slack_notifications_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ticket_no", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccSettingsResourceConfig(teamAdminGroup2, teamAuditorGroup2, duration, expiry),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "team_admin_group", teamAdminGroup2),
					resource.TestCheckResourceAttr(resourceName, "team_auditor_group", teamAuditorGroup2)),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccSettingsResourceConfig(teamAdminGroup string, teamAuditorGroup string, duration int, expiry int) string {
	return fmt.Sprintf(`
resource "awsteam_settings" "test" {
  team_admin_group = %[1]q
  team_auditor_group = %[2]q
  duration = %d
  expiry = %d
}
`, teamAdminGroup, teamAuditorGroup, duration, expiry)
}
