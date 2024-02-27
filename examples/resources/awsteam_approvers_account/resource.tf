resource "awsteam_approvers_account" "example" {
  account_id   = "123456789011"
  account_name = "my-account"
  approvers = [
    "my-group-approvers@contoso.com"
  ]
  group_ids = [
    "d78686b5-bb78-471c-8b2f-817e70e3158b"
  ]
}
