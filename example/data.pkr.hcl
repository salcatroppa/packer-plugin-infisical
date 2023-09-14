# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "infisical-secrets" "test" {
  folder_path = "/"
  env_slug    = "dev"
}