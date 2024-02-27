resource "awsteam_eligibility_user" "example" {
  user_name         = "my-user@contoso.com"
  approval_required = true
  duration          = 5
  accounts = [
    {
      account_id   = "123456789012"
      account_name = "My-aws-account"
    }
  ]
  ous = [
    {
      ou_id   = "ou-cxt3-2782ty5g"
      ou_name = "my-ou"
    }
  ]
  permissions = [
    {
      permission_arn  = "arn:aws:sso:::permissionSet/ssoins-4334d1f197f50907/ps-f5ge203d3d2428d3"
      permission_name = "elevated-permission"
    }
  ]
}
