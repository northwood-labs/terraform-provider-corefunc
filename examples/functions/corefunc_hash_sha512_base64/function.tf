output "sha512_base64" {
  value = provider::corefunc::hash_sha512_base64("hello world")
}

#=> MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==
