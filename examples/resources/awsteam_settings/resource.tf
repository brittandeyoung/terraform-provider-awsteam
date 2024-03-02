resource "awsteam_settings" "example" {
  duration           = 5
  expiry             = 3
  team_admin_group   = "My-Team-Admin-Group"
  team_auditor_group = "My-Team-Auditor-Group"
}

// Import the existing settings on a fresh install of AWS TEAM
import {
  to = awsteam_settings.example
  id = "settings"
}
