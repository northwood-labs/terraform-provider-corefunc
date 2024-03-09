data "corefunc_url_parse" "url" {
  url = "HTTP://u:p@example.com:80/foo?q=1#bar"
}

data "corefunc_url_parse" "gsb" {
  url           = "HTTP://u:p@example.com:80/foo?q=1#bar"
  canonicalizer = "google_safe_browsing"
}

output "url_parse_ds" {
  description = "This returns a parsed URL."
  value       = data.corefunc_url_parse.url
}

output "url_parse_gsb_ds" {
  description = "This returns a parsed URL canonicalized for Google Safe Browsing."
  value       = data.corefunc_url_parse.gsb
}
