module github.com/northwood-labs/terraform-provider-corefunc/v2

go 1.25.8

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.9.4-0.20241118143825-d1e633264448
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.16
)

require (
	github.com/apparentlymart/go-cidr v1.1.1
	github.com/chanced/caps v1.0.2
	github.com/charmbracelet/fang v1.0.0
	github.com/gtramontina/ooze v0.2.0
	github.com/hashicorp/go-version v1.9.0
	github.com/hashicorp/terraform-plugin-framework v1.19.0
	github.com/hashicorp/terraform-plugin-go v0.31.0
	github.com/hashicorp/terraform-plugin-log v0.10.0
	github.com/hashicorp/terraform-plugin-testing v1.16.0
	github.com/lithammer/dedent v1.1.0
	github.com/mattn/go-runewidth v0.0.23
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nlnwa/whatwg-url v0.6.2
	github.com/northwood-labs/cli-helpers v0.0.0-20250902191730-736c08104235
	github.com/pelletier/go-toml/v2 v2.2.4
	github.com/spf13/cobra v1.10.1
	github.com/pelletier/go-toml/v2 v2.3.1
	github.com/spf13/cobra v1.10.2
	github.com/stretchr/testify v1.11.1
	golang.org/x/crypto v0.50.0
)
