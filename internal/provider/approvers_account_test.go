package provider

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApproversAccountResource_basic(t *testing.T) {
	resourceName := "awsteam_approvers_account.test"
	accountId := gofakeit.DigitN(12)
	accountName := gofakeit.BS()
	approver1 := gofakeit.Email()
	groupId1 := gofakeit.UUID()
	approver2 := gofakeit.Email()
	groupId2 := gofakeit.UUID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApproversAccountResourceConfig(accountId, accountName, approver1, groupId1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "id", accountId),
					resource.TestCheckResourceAttr(resourceName, "account_name", accountName),
					resource.TestCheckTypeSetElemAttr(resourceName, "approvers.*", approver1),
					resource.TestCheckTypeSetElemAttr(resourceName, "group_ids.*", groupId1),
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
				Config: testAccApproversAccountResourceConfig(accountId, accountName, approver2, groupId2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckTypeSetElemAttr(resourceName, "approvers.*", approver2),
					resource.TestCheckTypeSetElemAttr(resourceName, "group_ids.*", groupId2),
				),
			},
		},
	})
}

func TestAccApproversAccountResource_accountId(t *testing.T) {
	resourceName := "awsteam_approvers_account.test"
	accountIdLeadingZeros := "000000123456"
	accountName := gofakeit.BS()
	approver1 := gofakeit.Email()
	groupId1 := gofakeit.UUID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApproversAccountResourceConfig(accountIdLeadingZeros, accountName, approver1, groupId1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", accountIdLeadingZeros),
					resource.TestCheckResourceAttr(resourceName, "id", accountIdLeadingZeros),
					resource.TestCheckResourceAttr(resourceName, "account_name", accountName),
					resource.TestCheckTypeSetElemAttr(resourceName, "approvers.*", approver1),
					resource.TestCheckTypeSetElemAttr(resourceName, "group_ids.*", groupId1),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccApproversAccountResourceConfig(accountId, accountName, approver, groupId string) string {
	return fmt.Sprintf(`
resource "awsteam_approvers_account" "test" {
	account_id   = %[1]q
	account_name = %[2]q
	approvers = [
		%[3]q
	]
	group_ids = [
		%[4]q
	]
}`, accountId, accountName, approver, groupId)
}
