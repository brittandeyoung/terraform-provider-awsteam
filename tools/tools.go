// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build tools

package tools

import (
	// Changelog Management
	_ "github.com/hashicorp/go-changelog/cmd/changelog-build"
	_ "github.com/hashicorp/go-changelog/cmd/changelog-check"
	_ "github.com/hashicorp/go-changelog/cmd/changelog-entry"

	// Document Generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
