output "url_parse_fn" {
  description = "This returns a parsed URL."
  value       = provider::corefunc::url_parse("HTTP://u:p@example.com:80/foo?q=1#bar")
}

#=> normalized = "http://u:p@example.com/foo?q=1#bar"
#=> normalized_nofrag = "http://u:p@example.com/foo?q=1"
#=> protocol = "http:"
#=> scheme = "http"
#=> username = "u"
#=> password = "p"
#=> host = "example.com"
#=> hostname = "example.com"
#=> port = ""
#=> decoded_port = 80
#=> path = "/foo"
#=> search = "?q=1"
#=> query = "q=1"
#=> hash = "#bar"
#=> fragment = "bar"
