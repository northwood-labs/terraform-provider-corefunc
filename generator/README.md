# Stub Generator

## Generator

This will accept a name, then generate the correct files needed to implement the provider functionality.

The name should match what the Terraform data source will be called (e.g., `corefunc_example`).

From this `generator/` directory, run the command.

```bash
go run main.go --name corefunc_example
```

After that, search the project for `@TODO` and make the appropriate changes. Exclude `./generator,*.md,.*.yml`.

## Checklist

* [ ] Generate the stub files.
* [ ] In the new `examples/data-sources/*/data-source.tf` file, ensure the example is implemented correctly.
* [ ] In the new `examples/data-sources/*/terraform-plan.example` file, copy content or delete.
* [ ] In the new `templates/data-sources/*.tmpl` file, update the template that will be used by the `tfplugindocs` package.
* [ ] Generate the Terraform provider documentation.
* [ ] In the new `*_data_source.go` file, update the `*SourceModel` struct.
* [ ] In the new `*_data_source.go` file, update the Schema Attributes to match the `*SourceModel` struct.
* [ ] In the new `*_data_source.go` file, implement the remainder of the `Read` function.
* [ ] In the new test fixture file, implement the test cases.
* [ ] In the new `*_data_source_fixture.tftpl` file, ensure the fixture is implemented correctly.
* [ ] In the new `*_data_source_test.go` file, ensure the test references the fixture correctly.
* [ ] Run the integration tests.
* [ ] Run `tfschema data show -format=json DATASOURCE | jq -Mrc '.attributes[]'`.
* [ ] In the new `bats/*.bats.sh` file, update its contents to match the `tfschema` output.
* [ ] Run the bats tests.
