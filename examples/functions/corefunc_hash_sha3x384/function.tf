output "hash_sha3x384" {
  value = provider::corefunc::hash_sha3x384("hello world")
}

#=> 83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b
