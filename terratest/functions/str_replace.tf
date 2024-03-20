output "str_iterative_replace_fn" {
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
        old = "New_Relic"
        new = "datadog"
      },
      {
        old = "This"
        new = "this"
      },
      {
        old = "Set_up"
        new = "setup"
      },
    ]
  )
}
