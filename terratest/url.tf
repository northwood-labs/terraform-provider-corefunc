# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
data "corefunc_url_parse" "url" {
  url = "HTTP://u:p@example.com:80/foo?q=1#bar"
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
data "corefunc_url_parse" "gsb" {
  url           = "HTTP://u:p@example.com:80/foo?q=1#bar"
  canonicalizer = "google_safe_browsing"
}
