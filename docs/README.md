# Infisical Plugin
The `Infisical` multi-component plugin can be used with HashiCorp [Packer](https://www.packer.io) to read secrets from [Infisical](https://infisical.com/).

For the full list of available features for this plugin see [docs](docs).

The Infisical client code was took 1:1 from the [Terraform provider](https://github.com/Infisical/terraform-provider-infisical).

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    name = {
      version = ">= 1.0.0"
      source  = "github.com/salcatroppa/infisical"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/salcatroppa/packer-plugin-infisical/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `make build` from the root
directory. Upon successful compilation, a `packer-plugin-name` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


## Plugin Contents
### Data Sources

- [secrets](/docs/datasources/secrets.mdx) - The `secrets` data source is used to
  read secrets from [Infisical](https://infisical.com/).
