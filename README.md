![tests](https://github.com/MyPureCloud/terraform-provider-genesyscloud/workflows/Tests/badge.svg?branch=main)
# Genesys Cloud Terraform Provider
<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider by running `make build`

## Using the provider

When using the Terraform CLI, you can run [`terraform init`](https://www.terraform.io/docs/commands/init.html) in the directory containing your provider configuration and Terraform will automatically install the required provider. The Genesys Cloud provider must be configured with an authorized OAuth client ID and secret to call the SDK.

```hcl
terraform {
  required_version = "~> 0.14.0"
  required_providers {
    genesyscloud = {
      source  = "mypurecloud/genesyscloud"
      version = "~> 1.0.0"
    }
  }
}

provider "genesyscloud" {
  oauthclient_id = "<client-id>"
  oauthclient_secret = "<client-secret>"
  aws_region = "<aws-region>"
}

```

The following environment variables may be set to avoid hardcoding OAuth Client information into your Terraform files:

```
GENESYSCLOUD_OAUTHCLIENT_ID
GENESYSCLOUD_OAUTHCLIENT_SECRET
GENESYSCLOUD_REGION
```

*Note:* The provider makes Public API calls to perform all of the CRUD operations necessary to manage Genesys Cloud resources. All of these API calls require specific permissions and OAuth scopes. Therefore it is important that you verify your OAuth Client is authorized for all necessary scopes and is assigned an admin role capable of creating, reading, updating, and deleting all resources that your Terraform configuration will manage.

For any issues, questions, or suggestions for the provider, visit the [Genesys Cloud Developer Forum](https://developer.mypurecloud.com/forum/)

### Data Sources

There may be cases where you want to reference existing resources in a Terraform configuration file but do not want those resources to be managed by Terraform. This provider supports several data source types that can act as a read-only resource for existing objects in your org. To include one in your configuration, add a `data` block to your configuration file with one of the supported data source types:
```hcl
data "genesyscloud_auth_role" "employee" {
  name = "employee"
}
```
The example above will attempt to find a role named "employee" which can be referenced elsewhere in the config. By default, all data sources will allow you to access the `id` attribute which is useful for setting reference attributes that require IDs. Additional attributes may be added to data sources as needs arise.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

In order to run the full suite of Acceptance tests, run `make testacc`. You can also specify individual tests using the TESTARGS variable:

```sh
$ make testacc TESTARGS="-run TestAccResourceUserBasic"
```

All new resources must have passing acceptance tests and docs in order to be merged. Most of the docs are generated automatically from the schema. An example `resource.tf` file for the resource should be added to the examples folder along with an `apis.md` file listing all of the APIs the resource uses. To generate or update documentation, run `go generate`.

*Note:* Acceptance tests create real resources and require an OAuth Client authorized to create, update, and delete all resources in your org. The OAuth Client information must be set in the `GENESYSCLOUD_*` environment variables prior to running the tests.

### Using the Provider locally

In order to use a locally compiled version of the provider, the correct binary for your system must be copied to the local `~/.terraform.d/plugins` folder. Run `make sideload` to build the provider and copy it to the correct folder. In your Terraform config file, specify version `0.1.0` and set the provider source to `genesys.com/mypurecloud/genesyscloud`. Run `terraform init` and verify that it finds the local version.

### Branches

Branch names should begin with `feat/` for new features or `bug/` for bug fixes. This ensures that the PR for this branch is correctly labeled and added to the changelog in the next release.

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to the provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

### Releases

A GitHub release action will be triggered when pushing version tags starting with 'v'. The release number must follow the [Semantic Versioning](https://semver.org/spec/v2.0.0.html) spec.

```
$ git tag -a v1.1.1 -m "Release v1.1.1"
$ git push origin v1.1.1
```

This action will build binaries, generate a changelog from labeled PRs, and create a draft GitHub release. The GitHub release draft will be reviewed and released manually.