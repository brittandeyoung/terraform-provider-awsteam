resource "awsteam_approvers_ou" "this" {
  ou_id   = "ou-cxt3-2782ty5g"
  ou_name = "my-ou"
  approvers = [
    "my-group-approvers@contoso.com"
  ]
  group_ids = [
    "d78686b5-bb78-471c-8b2f-817e70e3158b"
  ]
}
