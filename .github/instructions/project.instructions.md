---
applyTo: "**"
---

# Project Overview

This project is a Terraform provider which provides built-in functions and data sources that are self-contained and do not require network access. It is built in Go, using the Terraform Plugin Framework.

## Folder Structure

* `/cmd`: Contains the source code for the command-line interface enabled by the provider.
* `/corefunc`: Contains the source code for underying Go functions which are later wrapped and called by the provider interfaces.
* `/corefuncprovider`: Contains the source code for implementing a Terraform provider. It uses the Terraform Plugin Framework, and never uses source code from <https://github.com/hashicorp/terraform>.
* `/docs`: Contains the Markdown documentation generated for the Terraform provider.
* `/examples`: Contains the Terraform files which demonstrate how a piece of functionality works.
* `/terratest`: Contains the Terraform files and the Go code which implements the "Terratest" testing framework for Terraform code.
* `/testfixtures`: Contains the Go code which leverages table-driven tests to provide data for unit tests written for the Go code in `/corefunc` and the Terraform provider Go code in `/corefuncprovider`.

## Libraries and Frameworks

* Terraform Plugin Framework for implementing the interfaces for the Terraform provider.
* Source code is written in Go.
* We test against all versions of Terraform starting with 1.0. We also test against all versions of OpenTofu starting with 1.6.
