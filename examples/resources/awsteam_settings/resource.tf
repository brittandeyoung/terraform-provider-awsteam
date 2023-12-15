resource "awsteam_settings" "this" {
  duration           = 5
  expiry             = 3
  team_admin_group   = "My-Team-Admin-Group"
  team_auditor_group = "My-Team-Auditor-Group"
}
