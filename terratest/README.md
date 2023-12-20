# Terratest

Validates that the Terraform provider produces the same result as the Go library.

## Initialize

Initialize while checking for (and installing) updates.

### Terraform

```bash
terraform init -upgrade && terraform apply -auto-approve
```

### OpenTofu

```bash
tofu init -upgrade && tofu apply -auto-approve
```

## Run tests

```bash
go test -count 1
```
