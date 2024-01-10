resource "awsteam_approvers_account" "this" {
  account_number = 123456789011
  ou_name        = "my-account"
  approvers = [
    "my-group-approvers@contoso.com"
  ]
  group_ids = [
    "d78686b5-bb78-471c-8b2f-817e70e3158b"
  ]
}
