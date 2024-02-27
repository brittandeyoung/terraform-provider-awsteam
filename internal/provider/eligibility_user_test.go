package provider

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEligibilityUserResource_basic(t *testing.T) {
	resourceName := "awsteam_eligibility_user.test"
	user1 := gofakeit.Email()
	user2 := gofakeit.Email()
	approval1 := true
	approval2 := false
	duration := fmt.Sprint(gofakeit.Number(1, 10))
	ticketNo := gofakeit.BS()
	accountId := gofakeit.DigitN(12)
	accountName := gofakeit.BS()
	ouId := "ou-cxt3-2782ty5g" // hard coded fake ou id
	ouName := gofakeit.BS()
	permissionArn := "arn:aws:sso:::permissionSet/ssoins-4334d1f197f50907/ps-f5ge203d3d2428d3" // hard coded fake arn
	permissionName := "elevated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEligibilityUserResourceConfig(user1, approval1, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "user_name", user1),
					resource.TestCheckResourceAttr(resourceName, "approval_required", fmt.Sprint(approval1)),
					resource.TestCheckResourceAttr(resourceName, "duration", duration),
					resource.TestCheckResourceAttr(resourceName, "ticket_no", ticketNo),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "accounts.*",
						map[string]string{
							"account_id":   accountId,
							"account_name": accountName,
						}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ous.*",
						map[string]string{
							"ou_id":   ouId,
							"ou_name": ouName,
						}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "permissions.*",
						map[string]string{
							"permission_arn":  permissionArn,
							"permission_name": permissionName,
						}),
					// resource.TestCheckTypeSetElemAttr(resourceName, "group_ids.*", groupId1),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccEligibilityUserResourceConfig(user2, approval2, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "user_name", user2),
					resource.TestCheckResourceAttr(resourceName, "approval_required", fmt.Sprint(approval2)),
				),
			},
		},
	})
}

func testAccEligibilityUserResourceConfig(user string, approvalRequired bool, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName string) string {
	return fmt.Sprintf(`
resource "awsteam_eligibility_user" "test" {
	user_name         = "%s"
	approval_required = %t
	duration          = %s
	ticket_no = "%s"
	accounts = [
		{
		account_id   = "%s"
		account_name = "%s"
		}
	]
	ous = [
		{
		ou_id   = "%s"
		ou_name = "%s"
		}
	]
	permissions = [
		{
		permission_arn   = "%s"
		permission_name = "%s"
		}
	]
}`, user, approvalRequired, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName)
}
