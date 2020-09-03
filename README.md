# Terraform provider for Rollbar

## License

This is proprietary software.  **NO LICENSE WHATSOEVER is granted to this
software without written contract between author and licensee.**


## Status

![Build](https://github.com/jmcvetta/terraform-provider-rollbar/workflows/Build/badge.svg)


## Debugging

Enable writing debug log to `/tmp/terraform-provider-rollbar.log` by setting an
environment variable:

```
export TERRAFORM_PROVIDER_ROLLBAR_DEBUG=1
terraform apply   # or any command that calls the Rollbar provider
```

### Dev Script

Running `make dev` will build and install the provider, then run `terraform
apply` in the `examples` folder with debug logging enabled, and display the
logs on completion.
