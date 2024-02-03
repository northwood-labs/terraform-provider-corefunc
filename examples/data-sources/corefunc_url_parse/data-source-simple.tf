data "corefunc_url_parse" "simple" {
  url = "https://example.com"
}
#=> .normalized        = "https://example.com/"
#=> .normalized_nofrag = "https://example.com/"
#=> .protocol          = "https:"
#=> .scheme            = "https"
#=> .host              = "example.com"
#=> .path              = "/"
#=> .decoded_port      = 443
