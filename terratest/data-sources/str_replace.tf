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

output "str_iterative_replace_ds" {
  description = "This is a replacement output."
  value       = data.corefunc_str_iterative_replace.replacements.value
}
