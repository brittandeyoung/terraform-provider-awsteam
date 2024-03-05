package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEligibilityGroupResource_basic(t *testing.T) {
	ctx := context.Background()
	resourceName := "awsteam_eligibility_group.test"
	group1 := gofakeit.Email()
	group2 := gofakeit.Email()
	groupId1 := gofakeit.UUID()
	groupId2 := gofakeit.UUID()
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
				Config: testAccEligibilityGroupResourceConfig(group1, groupId1, approval1, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccEligibilityResourceExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "id", groupId1),
					resource.TestCheckResourceAttr(resourceName, "group_id", groupId1),
					resource.TestCheckResourceAttr(resourceName, "group_name", group1),
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
				Config: testAccEligibilityGroupResourceConfig(group2, groupId2, approval2, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "group_name", group2),
					resource.TestCheckResourceAttr(resourceName, "approval_required", fmt.Sprint(approval2)),
				),
			},
		},
	})
}

func TestAccEligibilityGroupResource_disappears(t *testing.T) {
	ctx := context.Background()
	resourceName := "awsteam_eligibility_group.test"
	group1 := gofakeit.Email()
	groupId := gofakeit.UUID()
	approval1 := true
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
				Config: testAccEligibilityGroupResourceConfig(group1, groupId, approval1, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName),
				Check: resource.ComposeTestCheckFunc(
					testAccEligibilityResourceExists(ctx, resourceName),
					testAccEligibilityResourceDisappears(ctx, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccEligibilityGroupResourceConfig(group string, groupId string, approvalRequired bool, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName string) string {
	return fmt.Sprintf(`
resource "awsteam_eligibility_group" "test" {
	group_name         = "%s"
	group_id = "%s"
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
}`, group, groupId, approvalRequired, duration, ticketNo, accountId, accountName, ouId, ouName, permissionArn, permissionName)
}
