data "corefunc_url_decode" "url" {
  encoded_url = "mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject"
}

#=> mailto:email?subject=this+is+my+subject
