output "replacements" {
  value = provider::corefunc::str_iterative_replace(
    "This is a string for testing replacements. New Relic. Set-up.",
    [
      {
        old = "."
        new = ""
      },
      {
        old = " "
        new = "_"
      },
      {
        old = "-"
        new = "_"
      },
      {
        old = "New Relic"
        new = "datadog"
      },
      {
        old = "This"
        new = "this"
      },
      {
        old = "Set-up"
        new = "setup"
      },
    ],
  )
}

#=> this_is_a_string_for_testing_replacements_datadog_setup
