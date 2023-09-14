// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package infisical

import (
	"os"

	infisical "packer-plugin-infisical/client"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Host         string `mapstructure:"host"`
	ServiceToken string `mapstructure:"service_token"`
	FolderPath   string `mapstructure:"folder_path"`
	EnvSlug      string `mapstructure:"env_slug"`
}

type Client struct {
	Client *infisical.Client
}

type Datasource struct {
	config Config
	client Client
}

type InfisicalSecretDetails struct {
	Value      string `mapstructure:"value"`
	Comment    string `mapstructure:"comment"`
	SecretType string `mapstructure:"secret_type"`
}

type DatasourceOutput struct {
	Secrets map[string]string `mapstructure:"secrets"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	host := os.Getenv("INFISICAL_HOST")
	// Set default to cloud infisical if host is empty
	if d.config.Host == "" && host == "" {
		d.config.Host = "https://app.infisical.com"
	}

	serviceToken := os.Getenv("INFISICAL_SERVICE_TOKEN")
	if d.config.ServiceToken == "" {
		d.config.ServiceToken = serviceToken
	}

	client, client_err := infisical.NewClient(infisical.Config{HostURL: d.config.Host, ServiceToken: d.config.ServiceToken})
	d.client.Client = client
	if client_err != nil {
		return client_err
	}

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	plainTextSecrets, _, err := d.client.Client.GetPlainTextSecretsViaServiceToken(d.config.FolderPath, d.config.EnvSlug)
	if err != nil {
		return cty.Value{}, err
	}

	secrets := make(map[string]string)
	for _, secret := range plainTextSecrets {
		secrets[secret.Key] = secret.Value
	}

	output := DatasourceOutput{
		Secrets: secrets,
	}

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
