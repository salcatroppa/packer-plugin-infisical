# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "infisical-my-datasource" "test" {
  mock = "mock-config"
}

locals {
  foo = data.infisical-my-datasource.test.foo
  bar = data.infisical-my-datasource.test.bar
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo foo: ${local.foo}",
      "echo bar: ${local.bar}",
    ]
  }
}
