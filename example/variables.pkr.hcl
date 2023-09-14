# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

locals {
  foo = data.infisical-my-datasource.mock-data.foo
  bar = data.infisical-my-datasource.mock-data.bar
}