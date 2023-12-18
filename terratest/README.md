# Terratest

Validates that the Terraform provider produces the same result as the Go library.

## Initialize

Initialize while checking for (and installing) updates.

```bash
terraform init -upgrade
```

## Run tests

```bash
go test
```
