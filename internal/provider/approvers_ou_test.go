package provider

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApproversOUResource_basic(t *testing.T) {
	resourceName := "awsteam_approvers_ou.test"
	ouId := "ou-cxt3-2782ty5g"
	ouName := gofakeit.BS()
	approver1 := gofakeit.Email()
	groupId1 := gofakeit.UUID()
	approver2 := gofakeit.Email()
	groupId2 := gofakeit.UUID()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApproversOUResourceConfig(ouId, ouName, approver1, groupId1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ou_id", ouId),
					resource.TestCheckResourceAttr(resourceName, "id", ouId),
					resource.TestCheckResourceAttr(resourceName, "ou_name", ouName),
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
				Config: testAccApproversOUResourceConfig(ouId, ouName, approver2, groupId2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckTypeSetElemAttr(resourceName, "approvers.*", approver2),
					resource.TestCheckTypeSetElemAttr(resourceName, "group_ids.*", groupId2),
				),
			},
		},
	})
}

func testAccApproversOUResourceConfig(ouId, ouName, approver, groupId string) string {
	return fmt.Sprintf(`
resource "awsteam_approvers_ou" "test" {
	ou_id   = %[1]q
	ou_name = %[2]q
	approvers = [
		%[3]q
	]
	group_ids = [
		%[4]q
	]
}`, ouId, ouName, approver, groupId)
}
