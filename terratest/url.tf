# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/url_parse
data "corefunc_url_parse" "url" {
  url = "HTTP://u:p@example.com:80/foo?q=1#bar"
}
