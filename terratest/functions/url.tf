output "url_parse_fn" {
  description = "This returns a parsed URL."
  value       = provider::corefunc::url_parse("HTTP://u:p@example.com:80/foo?q=1#bar")
}

output "url_parse_gsb_fn" {
  description = "This returns a parsed URL canonicalized for Google Safe Browsing."
  value       = provider::corefunc::url_parse("HTTP://u:p@example.com:80/foo?q=1#bar", "google_safe_browsing")
}
