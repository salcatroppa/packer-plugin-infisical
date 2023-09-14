# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "infisical-secrets" "test" {
  folder_path = "/"
  env_slug    = "dev"
}

locals {
  secrets = data.infisical-secrets.test.secrets
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
      "echo secret_key: ${local.secrets["SECRET_KEY"]}",
    ]
  }
}
