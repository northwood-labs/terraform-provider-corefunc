data "corefunc_url_parse" "gsb" {
  url           = "HTTP://u:p@example.com:80/foo?q=1#bar"
  canonicalizer = "google_safe_browsing"
}

#=> .normalized        = "http://u:p@example.com/foo?q=1"
#=> .normalized_nofrag = "http://u:p@example.com/foo?q=1"
#=> .protocol          = "http:"
#=> .scheme            = "http"
#=> .username          = "u"
#=> .password          = "p"
#=> .host              = "example.com"
#=> .port              = ""
#=> .path              = "/foo"
#=> .search            = "?q=1"
#=> .query             = "q=1"
#=> .hash              = ""
#=> .fragment          = ""
#=> .decoded_port      = 80
