data "corefunc_str_iterative_replace" "replacements" {
  string = "This is a string for testing replacements. New Relic. Set-up."
  replacements = [
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
}
#=> this_is_a_string_for_testing_replacements_datadog_setup
