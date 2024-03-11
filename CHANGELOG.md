# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com), adheres to [Semantic Versioning](https://semver.org), and uses [Conventional Commit](https://www.conventionalcommits.org) syntax.

## Unreleased

[Compare: v1.3.1 → `HEAD`](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.3.1..HEAD)

### :books: Documentation

* [`371ea5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/371ea5db284ad4b814bb3d19aa16527f967f401a): Updated the copyright block. ([#183](https://github.com/northwood-labs/terraform-provider-corefunc/issues/183)) ([@skyzyx](https://github.com/skyzyx))

### :books: Documentation

* [`371ea5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/371ea5db284ad4b814bb3d19aa16527f967f401a): Updated the copyright block. ([#183](https://github.com/northwood-labs/terraform-provider-corefunc/issues/183)) ([@skyzyx](https://github.com/skyzyx))
* [`1911a71`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1911a718ab20dc7d3cb62b7106caeaa8b0ebaaca): Generated an updated CHANGELOG. ([@skyzyx](https://github.com/skyzyx))
* [`e2be61d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e2be61d2837b3915055153566867435a97828ca2): Added Homedir functions. ([@skyzyx](https://github.com/skyzyx))
* [`8110aea`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8110aeadbdc88d6d4e0aa7171c2245b4018b514c): Cleaned up some documentation. ([@skyzyx](https://github.com/skyzyx))
* [`fab52de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fab52deff79538fa4db72df32636a937212195d7): Updated docs around OpenTofu and TF 1.8. ([@skyzyx](https://github.com/skyzyx))
* [`eb651e9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/eb651e9accfca5b83854ff2b42e3294f4fedf47d): Updated the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`04e92d0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/04e92d06f63fee92231c349439e2bd616fbf0ddb): Updated the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`44caacf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/44caacf70040b6707fb072e998996972ef1cc3fa): Small fixes to the docs. ([@skyzyx](https://github.com/skyzyx))
* [`a7f932f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a7f932f06e7b3c589344ae0cf2068763023d8dd4): Uncomment the YAML frontmatter in the Markdown pages. ([@skyzyx](https://github.com/skyzyx))
* [`5be683c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5be683caeed4d696d026c395a9837ee6281fc495): Polishing up the provider function documentation. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`7b33cc0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7b33cc0a804cc1e5866eada269f62559a7e861a9): Updated the Go dependencies and mod file. ([@skyzyx](https://github.com/skyzyx))
* [`7be699c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7be699ca5a3f9168ec17f3b5437aaa8efefd60f4): **deps**: Bump `trufflesecurity/trufflehog` from 3.68.4 to 3.68.5 ([#184](https://github.com/northwood-labs/terraform-provider-corefunc/issues/184)) ([@dependabot](https://github.com/dependabot))
* [`c88ce13`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c88ce13c9b718ccdc24ed3ae37860a6e348b18be): **deps**: Bump `trufflesecurity/trufflehog` from 3.68.5 to 3.69.0 ([#187](https://github.com/northwood-labs/terraform-provider-corefunc/issues/187)) ([@dependabot](https://github.com/dependabot))

### :test_tube: Testing

* [`c37a8d5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c37a8d5169bab412f5ba4d94410332c94377dcb7): Make some adjustments for the test suite. ([@skyzyx](https://github.com/skyzyx))
* [`32ccd04`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/32ccd04409301dfb4d94407a366a3b6fe6cb39db): Add Terraform 1.8-beta1 to the test matrix. ([@skyzyx](https://github.com/skyzyx))

### :tractor: Refactor

* [`ad94772`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ad94772b930b6c9da0f13ecd1130c8b8dbb5b8f9): Moved the provider.Functions function to this branch. ([@skyzyx](https://github.com/skyzyx))
* [`cea02c3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cea02c3b9ee6699d8f7669f35a78c2839ce1f6eb): Add more support for provider functions. ([@skyzyx](https://github.com/skyzyx))
* [`a982bad`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a982bad51250b3b6fc0a8df3ce2d0fbed7127fd0): Switch Description to MarkdownDescription. Fix str_iterative_replace. ([@skyzyx](https://github.com/skyzyx))

### <!-- 0 -->:rocket: Features

* [`e5efd1c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e5efd1c243f81571932e5129619f62694cb3a635): Implemented `str_snake()` as a function. ([@skyzyx](https://github.com/skyzyx))
* [`5599a3c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5599a3cefa5112e6705e7739104540d4f369a9e9): First feature-complete function implementation. ([@skyzyx](https://github.com/skyzyx))
* [`01c202d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/01c202d63d317c5515f11b54c3b11d8e25590faf): Ported runtime functions to functions. ([@skyzyx](https://github.com/skyzyx))
* [`66dade4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/66dade48cb26a9d1d1a4e4c059a8163b27c51abf): Added Homedir functions. ([@skyzyx](https://github.com/skyzyx))
* [`6de9499`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6de949954d0fc1573f5c54521afb7dda8ffe88be): Implementation of string formatting as funcs. ([@skyzyx](https://github.com/skyzyx))
* [`238263e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/238263eaeac7c37157bf3b966d7cdabaf4e18ac4): Implemented leftpads as functions. ([@skyzyx](https://github.com/skyzyx))
* [`098169b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/098169bb1dd9eda481be898ff92d28da3c9e9907): Partial implementation of str_iterative_replace. ([@skyzyx](https://github.com/skyzyx))
* [`b1fa98d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b1fa98dc0290452404cdcd4319e3809fed7e3962): Implemented `str_truncate_label()` as a function. ([@skyzyx](https://github.com/skyzyx))
* [`2a3d46e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2a3d46e6e307ad90f7ba5dd3095675f24b6c3753): Implemented `url_parse()` as a function. ([@skyzyx](https://github.com/skyzyx))

## 1.3.1 — 2024-03-07

[Compare: v1.3.0 → v1.3.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.3.0...v1.3.1)

## 1.3.0 — 2024-02-05

[Compare: v1.2.1 → v1.3.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.2.1...v1.3.0)

### :books: Documentation

* [`00411fe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/00411fee4350e0be6dcf78bbf2ade87ee0c5b03b): Updated docs to remove the 'id' attribute. ([@skyzyx](https://github.com/skyzyx))
* [`e819279`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e8192790d23b54e8120055b610b17ccc1682f1ee): Fixed a few typos. ([@skyzyx](https://github.com/skyzyx))
* [`88db492`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/88db4929008d062f86dc7bd65c36d7b52b997d30): Fixed some links in `corefunc_url_parse`. ([@skyzyx](https://github.com/skyzyx))
* [`ca6d540`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca6d5402046aa39a56f060f8fb4165fc35f688be): Fixed some typos in `corefunc_url_parse`. ([@skyzyx](https://github.com/skyzyx))
* [`d02b3e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d02b3e7a913824b09ebfb9411d810f1473b91073): Fixed some typos in `corefunc_url_parse`. ([@skyzyx](https://github.com/skyzyx))
* [`d6938de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6938de38a824a648ddf9842dec9e0e3cd99cbbf): Added link to Google Safe Browsing rules. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`4cde8d8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4cde8d89e62d2201dcfb05a80b709f7b1258cf3f): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#130](https://github.com/northwood-labs/terraform-provider-corefunc/issues/130)) ([@dependabot](https://github.com/dependabot))
* [`1a298cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1a298cb78cc4d9e6144a9b4282e71311a348b144): **deps**: Bump `github/codeql-action` from 3.23.2 to 3.24.0 ([#131](https://github.com/northwood-labs/terraform-provider-corefunc/issues/131)) ([@dependabot](https://github.com/dependabot))
* [`68a3d4b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/68a3d4b110cd6f4b36d4f12bb0df3160372847a1): **deps**: Bump `trufflesecurity/trufflehog` from 3.66.3 to 3.67.1 ([#132](https://github.com/northwood-labs/terraform-provider-corefunc/issues/132)) ([@dependabot](https://github.com/dependabot))
* [`e548051`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e54805173ee44bd7b13fac10406dd0d0e1968511): Small updates to the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`d4496cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d4496cbae6c11de0c55511cdbf20d5a5b5b2c908): Small updates to the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`87f2772`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/87f277226fc1f3c7593438f78e590692dac45764): Small updates to the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`b4232d2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b4232d2f0fa6f63ea05208730525cd0392adbebb): Small updates to the Makefile. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`5d992f3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d992f31b4d5566e37bb12a0d133400d75d1763f): Regenerate test artifacts. ([@skyzyx](https://github.com/skyzyx))
* [`c4459ad`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c4459ad7664dc6a1c93b3585705dc46d72cb76d2): Updated BATS tests with Shellcheck feedback. ([@skyzyx](https://github.com/skyzyx))
* [`6611f3f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6611f3fbd38af4997fb156c22d4157024bd8d76f): Generated updated test artifacts. ([@skyzyx](https://github.com/skyzyx))
* [`d6f3437`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6f34376297dc934e195c379786053729b8625bb): Added Terratest tests for `corefunc_url_parse`. ([@skyzyx](https://github.com/skyzyx))

### <!-- 0 -->:rocket: Features

* [`1752c6e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1752c6e94daba7aec9a1b6c4740e16b7b3764be6): Implement new runtime data sources. ([@skyzyx](https://github.com/skyzyx))
* [`1da303c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1da303cebdfc414f62a5907bb1f2526dc66b6067): Implemented `corefunc_url_parse`. ([@skyzyx](https://github.com/skyzyx))
* [`e179a2f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e179a2fdd814662fc2dff30ce3dea1db06cbb27a): Added support for Google Safe Browsing canonicalization. ([@skyzyx](https://github.com/skyzyx))

### <!-- 1 -->:bug: Bug Fixes

* [`fc46b7f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fc46b7f8bf4542d86b8bfe7c7c1eb9ac2843dc9d): Updated testing framework to not require the 'id' attribute. ([@skyzyx](https://github.com/skyzyx))
* [`50affa9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/50affa91e6aa6a6c0b59c529070a6d296afd0823): Lower provider logging messages from INFO to DEBUG. ([@skyzyx](https://github.com/skyzyx))
* [`6582fec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6582fecf49b9e400ba7c0d5bbf8b3d483a89aaf4): Fix examples and Shellcheck warnings. ([@skyzyx](https://github.com/skyzyx))

## 1.2.1 — 2024-02-02

[Compare: v1.2.0 → v1.2.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.2.0...v1.2.1)

### :books: Documentation

* [`b35d3f8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b35d3f85fd3cb8ed0a8355b6f631878484f3be98): Adjust the git-cliff format. ([@skyzyx](https://github.com/skyzyx))
* [`6189ded`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6189ded118f391007e0efda00aafa90636caee3a): Resolved an issue with Markdown generation build more consistent docs. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`558c756`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/558c756c6a0b91d40233ee446d3dece24a1c3b53): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#128](https://github.com/northwood-labs/terraform-provider-corefunc/issues/128)) ([@dependabot](https://github.com/dependabot))

### :test_tube: Testing

* [`cc77ed9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cc77ed99085f9c81281dd7a6dbcc74298d4cdbe1): Improve CHANGELOG generation. ([@skyzyx](https://github.com/skyzyx))
* [`0182c79`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0182c79239bb888e61e7affcd3ddfaeced41c0d2): Improve CHANGELOG generation. ([@skyzyx](https://github.com/skyzyx))
* [`acee3ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/acee3ef09ce457f074bb25f8aada3d029e474027): Add auto-cancelation for multiple CI runs of the same test. ([@skyzyx](https://github.com/skyzyx))
* [`85cfb6c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/85cfb6c8dfed8d61bf84d246ba58eacbf91cc650): Setup Terratest testing for OpenTofu from the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`57b1e0e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/57b1e0eb5f91a87e5450dba5d6f9492fb2c0b374): Improve switching between Terraform/OpenTofu in Terratest. ([@skyzyx](https://github.com/skyzyx))
* [`590c3bc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/590c3bce781674b2627e494b203570b142555844): Improve switching between Terraform/OpenTofu in Terratest. ([@skyzyx](https://github.com/skyzyx))

### <!-- 1 -->:bug: Bug Fixes

* [`0533928`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/05339289d72baeb1a72b6983c0f6132e264ba1c8): Remove a cache file from the repo as it changes every time. ([@skyzyx](https://github.com/skyzyx))
* [`1ad7a2b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1ad7a2b160a44b4534883d9ce84969eb7cef519f): Improve the version selection automation. ([@skyzyx](https://github.com/skyzyx))
* [`e952382`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e95238280e794773f9826abb60480382a9855d5c): Improve the version selection automation. ([@skyzyx](https://github.com/skyzyx))
* [`3e1a07a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3e1a07a433aa6e0be88ed139ee45cde343e5b1d6): Improve the version selection automation. ([@skyzyx](https://github.com/skyzyx))

## 1.2.0 — 2024-02-02

[Compare: v1.1.1 → v1.2.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.1...v1.2.0)

### :books: Documentation

* [`c30e191`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c30e1912d5e273e8238703d3f7bc67a492017e66): Added information about compatibility testing. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`92c7065`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92c7065ba587b0851aa9352a1422ac7d2d68e8fa): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#94](https://github.com/northwood-labs/terraform-provider-corefunc/issues/94)) ([@dependabot](https://github.com/dependabot))
* [`96a741f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/96a741f1f36af211f80934694295736df522aed7): **deps**: Bump `crazy-max/ghaction-import-gpg` from 6.0.0 to 6.1.0 ([#97](https://github.com/northwood-labs/terraform-provider-corefunc/issues/97)) ([@dependabot](https://github.com/dependabot))
* [`73ac9c5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73ac9c52d6275e2c7ea25b14a7677a6a91174cc2): **deps**: Bump `github/codeql-action` from 1.1.39 to 3.22.12 ([#95](https://github.com/northwood-labs/terraform-provider-corefunc/issues/95)) ([@dependabot](https://github.com/dependabot))
* [`528dfcd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/528dfcdd689857b97757f887c96a34ca4369d83c): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.5 to 3.63.7 ([#96](https://github.com/northwood-labs/terraform-provider-corefunc/issues/96)) ([@dependabot](https://github.com/dependabot))
* [`523b701`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/523b7012e3cc18de7879f30a3537136cb11a05a9): **deps**: Bump `github.com/chanced/caps` from 1.0.1 to 1.0.2 ([#98](https://github.com/northwood-labs/terraform-provider-corefunc/issues/98)) ([@dependabot](https://github.com/dependabot))
* [`b311501`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b3115010b0e84d6adcc5d3a75610f1f07697b5a7): **deps**: Bump `github.com/chanced/caps` in /generator ([#99](https://github.com/northwood-labs/terraform-provider-corefunc/issues/99)) ([@dependabot](https://github.com/dependabot))
* [`fc859eb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fc859ebd974529d165cd107fd1f7ab33a55826e3): **deps**: Bump `actions/dependency-review-action` from 3.1.4 to 3.1.5 ([#100](https://github.com/northwood-labs/terraform-provider-corefunc/issues/100)) ([@dependabot](https://github.com/dependabot))
* [`1ec1ac2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1ec1ac2488f25d39289247b882044a51335d0125): **deps**: Bump `github.com/cloudflare/circl` from 1.3.6 to 1.3.7 ([#101](https://github.com/northwood-labs/terraform-provider-corefunc/issues/101)) ([@dependabot](https://github.com/dependabot))
* [`7f40e69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7f40e69e02c6314e0728c776e3376492e9e1ccbe): **deps**: Bump `actions/go-dependency-submission` from 1.0.3 to 2.0.0 ([#105](https://github.com/northwood-labs/terraform-provider-corefunc/issues/105)) ([@dependabot](https://github.com/dependabot))
* [`1317bab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1317babb5379cd55528a2b932905acfbea9253de): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#106](https://github.com/northwood-labs/terraform-provider-corefunc/issues/106)) ([@dependabot](https://github.com/dependabot))
* [`8d22b24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d22b24cd8cc52a6ebdabc524a61f2f8ee785627): **deps**: Bump `github.com/gruntwork-io/terratest` in /terratest ([#110](https://github.com/northwood-labs/terraform-provider-corefunc/issues/110)) ([@dependabot](https://github.com/dependabot))
* [`104b541`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/104b541701e30608c39913de362524454f73a925): **deps**: Bump `github/codeql-action` from 3.22.12 to 3.23.1 ([#111](https://github.com/northwood-labs/terraform-provider-corefunc/issues/111)) ([@dependabot](https://github.com/dependabot))
* [`0063c0c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0063c0ce2b277b5201d14ab7cf6862c9963a3443): **deps**: Bump `actions/dependency-review-action` from 3.1.5 to 4.0.0 ([#113](https://github.com/northwood-labs/terraform-provider-corefunc/issues/113)) ([@dependabot](https://github.com/dependabot))
* [`2ca888e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2ca888ed39eae52a622fbe9049d6ecabdf2041f9): **deps**: Bump `actions/upload-artifact` from 4.0.0 to 4.3.0 ([#116](https://github.com/northwood-labs/terraform-provider-corefunc/issues/116)) ([@dependabot](https://github.com/dependabot))
* [`1750875`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1750875ff6005e48b8d8405c2b45ee486b116fab): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.7 to 3.64.0 ([#117](https://github.com/northwood-labs/terraform-provider-corefunc/issues/117)) ([@dependabot](https://github.com/dependabot))
* [`0844a2e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0844a2eb4e5d6ee49b69c5a6c9b5a8d7bc510c4a): **deps**: Bump `github.com/hashicorp/terraform-plugin-docs` ([#118](https://github.com/northwood-labs/terraform-provider-corefunc/issues/118)) ([@dependabot](https://github.com/dependabot))
* [`680a5d4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/680a5d4b3e437d1f21cc636995b924990cb1d6f3): **deps**: Bump `github/codeql-action` from 3.23.1 to 3.23.2 ([#120](https://github.com/northwood-labs/terraform-provider-corefunc/issues/120)) ([@dependabot](https://github.com/dependabot))
* [`771756b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/771756b2138312dadefebb7df8a651b6b9f6f93c): **deps**: Bump `github.com/hashicorp/terraform-plugin-go` ([#121](https://github.com/northwood-labs/terraform-provider-corefunc/issues/121)) ([@dependabot](https://github.com/dependabot))
* [`f35b806`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f35b806374aa31334d04075f0dd92e54e08b59ef): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#123](https://github.com/northwood-labs/terraform-provider-corefunc/issues/123)) ([@dependabot](https://github.com/dependabot))
* [`5396f2d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5396f2d698dfacaa6edc402d2c0e38b90fa34570): **deps**: Bump `step-security/harden-runner` from 2.6.1 to 2.7.0 ([#124](https://github.com/northwood-labs/terraform-provider-corefunc/issues/124)) ([@dependabot](https://github.com/dependabot))
* [`d0df49d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d0df49d95113e5fbe29ba592d3f51717798472e6): **deps**: Bump `trufflesecurity/trufflehog` from 3.64.0 to 3.66.3 ([#126](https://github.com/northwood-labs/terraform-provider-corefunc/issues/126)) ([@dependabot](https://github.com/dependabot))

### :soap: Linting

* [`f2c349f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f2c349f8c33e07c6fd2dcd6ac5e55825f9eaf53b): Ran the linter against the codebase after changes. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`ea785fa`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ea785fa9dc992c8df52756501798240d63dacf91): Disable GoSec Sarif uploading until we know why it's failing. ([@skyzyx](https://github.com/skyzyx))
* [`d2198f3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d2198f3152d878657f9bb77b3ab81affd15dc45e): Add Terraform 1.7 to the test matrix. ([@skyzyx](https://github.com/skyzyx))
* [`73b83b1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73b83b1982933903b8a4e512a2af06f38607c624): Enable workflow_dispatch to trigger tests manually. ([@skyzyx](https://github.com/skyzyx))
* [`4205631`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4205631e3b504b8448c31fde9a1603e1565128b1): Discovered issue in calling GoSec. ([@skyzyx](https://github.com/skyzyx))
* [`9486823`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9486823382c8c36f3895ffaeb58bfb325c1cf344): Discovered issue in calling GoSec. ([@skyzyx](https://github.com/skyzyx))
* [`472cb4c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/472cb4cb3ff48742d01a4d97a656e653d54dd835): Discovered issue in calling GoSec. ([@skyzyx](https://github.com/skyzyx))
* [`8686958`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8686958d2d46f52438279d249c1455e4a038ea04): Disable GoSec workflow for the time being. ([@skyzyx](https://github.com/skyzyx))
* [`bf75f94`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bf75f945ffdbc872d138c9f6effaf29d64639b7a): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`6672a01`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6672a0174d74b913a5462e1ea5856c3e37c900b2): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`834be43`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/834be43f30a1e68d1da701887e11ee3db623d124): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`40b77b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/40b77b88163ee6a711bd5770985f7a6fc5528ee4): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`e1e7b69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e1e7b695b4fc564b4b56a5ece28d674df3bb86a0): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`9749f52`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9749f529e519a8f6978e40f55eeb3b1f4fdbd308): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`6949371`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6949371a0332b730dad43c48c529dfd23714310c): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`157efbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/157efbe38e925b53deb4d5b729b650ecbfa014bc): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`9c09697`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9c096972030ce0d870d2df5100bab7bfb2d37a71): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`7ddd3e8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7ddd3e86d64d49528e310aa42456420b7f41d2a5): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`d7d77f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d7d77f292f2f8c2b56fd35419d9d9e01ad774dec): Auto-update `AUTHORS.md` and CHANGELOG on commit. ([@skyzyx](https://github.com/skyzyx))
* [`2c282c8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2c282c8e863ab3527dcc605ac3a74ce3a28c026e): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`a20e1bf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a20e1bfaa2eea228abca619bc81840b4ee303ef1): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`82583e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/82583e7bb9128268e702e65b5a9a90ed1f69b6f7): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`5d9cd26`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d9cd26136cebc600044b527874a97d6847ff641): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`2758ec4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2758ec4de8a6cfa8c99859786f43e526c83fa843): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`92c5e44`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92c5e440d9425fb83c11bf9ca403f2b85b74c6a5): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`ca1650f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca1650f2f2e9737a6635851bd19795e2d1b22908): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`ae243c3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ae243c36dba8429a3606b223495d2f94ec5b82e9): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`56501ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/56501ef5e2e26ec6ca6c1c06a1aa1eb7987cb280): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`4a1e899`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4a1e899979c9a3593fd741484bdf1e8b3febfe4c): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`a8b9724`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a8b97249bf45c2fc71fc6d32ffc6470a5b3dc5c3): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`d374a8c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d374a8c42fa2e3c20f410762f7369d41088bad6e): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`44c7648`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/44c76484bc48005d0484162a43bb5efa2170f724): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`b5f2b8e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b5f2b8e0382900ef4a85f0e5e0052c066b9da624): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`6be7398`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6be7398843e0a757ab4f0308a175843d8603c829): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))
* [`66abdf2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/66abdf213d2f88f3ee0e3d44b26dd0b557b25005): Work on adding automated OpenTofu testing. ([@skyzyx](https://github.com/skyzyx))

### <!-- 1 -->:bug: Bug Fixes

* [`32dffbb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/32dffbb12f962543d923f2ea39000675d0ff8714): Resolve issues in the `go.sum` file. ([@skyzyx](https://github.com/skyzyx))

## 1.1.1 — 2023-12-20

[Compare: v1.1.0 → v1.1.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.0...v1.1.1)

### :books: Documentation

* [`f3c07b2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f3c07b267eacc0cc4c3fabd81628112f8ea974f2): Updated copyright in all headers to include 2024. ([@skyzyx](https://github.com/skyzyx))
* [`1d00129`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1d0012982e3c1760735a4fa53a6603781bad6930): Updated the `README.md` for the Terratest tests. ([@skyzyx](https://github.com/skyzyx))
* [`1c2bc3c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1c2bc3c6f91b6326cd188320d8978972e2aa78dc): Updated the templates used by tfplugindocs. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`0c431c2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0c431c21046806396bfe392f32fdfa5ba67743f3): Remove the restriction to pre-1.6 versions of Terraform in examples. ([@skyzyx](https://github.com/skyzyx))
* [`27b9e60`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/27b9e60796e0bfdc3b972cae11bca284c4cfc8d2): Updated Go dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`bc39137`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bc391374137cb742fa719bd3790288cfa4a0c7cd): Pin the version of the GoSec action. ([@skyzyx](https://github.com/skyzyx))
* [`aa4e217`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/aa4e2178e06786ed2a0adb21e0095326854904fb): Pin the version of the Trufflehog action. ([@skyzyx](https://github.com/skyzyx))
* [`1d01f26`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1d01f2686604a931602bdd0880a640ac13d312ce): **deps**: Bump `golang.org/x/crypto` in /terratest ([#90](https://github.com/northwood-labs/terraform-provider-corefunc/issues/90)) ([@dependabot](https://github.com/dependabot))
* [`2f77df9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2f77df906e466c58ba3d4eb7d328684da4aab20e): **deps**: Bump `golang.org/x/crypto` from 0.3.0 to 0.17.0 in /generator ([#89](https://github.com/northwood-labs/terraform-provider-corefunc/issues/89)) ([@dependabot](https://github.com/dependabot))
* [`07df7b9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/07df7b911593bf1b65fbe060b005aae5b0e4faa6): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.4 to 3.63.5 ([#91](https://github.com/northwood-labs/terraform-provider-corefunc/issues/91)) ([@dependabot](https://github.com/dependabot))
* [`956339c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/956339c396e8b199a35d2f4eeff058dd11175be3): **deps**: Bump `golang.org/x/crypto` from 0.16.0 to 0.17.0 ([#88](https://github.com/northwood-labs/terraform-provider-corefunc/issues/88)) ([@dependabot](https://github.com/dependabot))
* [`543eb0a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/543eb0ab95e3918f034c76a99099ff0f8fcdd490): Update the `go.sum` file. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`e0915df`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e0915dfb90f9d2824be769dcc629501db99d1b8c): Added tests with Terratest to compare the Terraform provider with the Go library. ([@skyzyx](https://github.com/skyzyx))
* [`e5a810a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e5a810a8f9a5ffe2737928429df0be3a9a0a25c3): Make the GoSec scanning more robust. ([@skyzyx](https://github.com/skyzyx))

## 1.1.0 — 2023-12-18

[Compare: v1.0.3 → v1.1.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.3...v1.1.0)

### :books: Documentation

* [`3b7d8ca`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b7d8cac9ddc38904505c5c9fbbd47fc15f54df5): Updated the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`d41ad2e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d41ad2e1dafbfac54dacf3ac9d49b50032122089): Updated the Terraform documentation to include the lockfile. ([@skyzyx](https://github.com/skyzyx))
* [`00b7206`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/00b7206578b826912a4e7563d7d6caa2588432a3): Updated the CHANGELOG. ([@skyzyx](https://github.com/skyzyx))
* [`92a73de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92a73de97de90a269afd08f4369ee4c5ad4aee59): Ran Markdownlint on the CHANGELOG. ([@skyzyx](https://github.com/skyzyx))
* [`6b7b0a3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6b7b0a375e9b17be41055a4f8a2c2dce6401b674): Added a note about base10 to `corefunc_int_leftpad`. ([@skyzyx](https://github.com/skyzyx))
* [`b83580e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b83580eeac1b986d8b0fd7d667da31356f7d8bd5): Regenerated the provider documentation. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`de4027c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/de4027ce4f48f09525b5f0ce4c4000ebd6c75de1): **deps**: Bump `github/codeql-action` from 2.22.7 to 2.22.8 ([#72](https://github.com/northwood-labs/terraform-provider-corefunc/issues/72)) ([@dependabot](https://github.com/dependabot))
* [`98d8612`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/98d86121b239c9d15a1a3f429e90b07cf750399f): **deps**: Bump `actions/dependency-review-action` from 3.1.3 to 3.1.4 ([#74](https://github.com/northwood-labs/terraform-provider-corefunc/issues/74)) ([@dependabot](https://github.com/dependabot))
* [`874d704`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/874d7046af06c46ed86ea818936e846188d612d6): **deps**: Bump `actions/setup-go` from 4.1.0 to 5.0.0 ([#75](https://github.com/northwood-labs/terraform-provider-corefunc/issues/75)) ([@dependabot](https://github.com/dependabot))
* [`ca8c440`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca8c4409fb32db9b83647e085839b73923ded5ec): **deps**: Bump `github/codeql-action` from 2.22.8 to 2.22.9 ([#76](https://github.com/northwood-labs/terraform-provider-corefunc/issues/76)) ([@dependabot](https://github.com/dependabot))
* [`232c76e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/232c76e6ecaaa95e43e78182e02c259f2c8bb3b0): **deps**: Bump `github/codeql-action` from 2.22.9 to 3.22.11 ([#79](https://github.com/northwood-labs/terraform-provider-corefunc/issues/79)) ([@dependabot](https://github.com/dependabot))
* [`415d3d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/415d3d3bf5a648cb3168ab1015059613d53046f4): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#82](https://github.com/northwood-labs/terraform-provider-corefunc/issues/82)) ([@dependabot](https://github.com/dependabot))
* [`78ec578`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/78ec578eb53836a96aaf93ae8b78cb69217941b8): **deps**: Bump `actions/upload-artifact` from 3.1.3 to 4.0.0 ([#80](https://github.com/northwood-labs/terraform-provider-corefunc/issues/80)) ([@dependabot](https://github.com/dependabot))

### :soap: Linting

* [`89b97d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/89b97d37ea0495e56fefe181ef7f473680e73a0b): Updated the cache of dependency license information. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`57a7b85`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/57a7b85c17040f880dcaf6546c11cb9acd0a7615): Lookup and run fuzz tests programmatically. ([@skyzyx](https://github.com/skyzyx))

### :tractor: Refactor

* [`2204ce8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2204ce86c8d36127cf84163559254549e819590f): Added wrapper corefunc functions for all caps functions. ([@skyzyx](https://github.com/skyzyx))

### <!-- 0 -->:rocket: Features

* [`1135f94`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1135f9439b1966e3da849c2b9272f84eb89cfd36): Begin implementations of StrLeftPad and IntLeftPad. ([@skyzyx](https://github.com/skyzyx))
* [`92fbf6f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92fbf6f5e7f4fd17b1a33135a851143be1fd276e): Implemented the provider side of leftpad functions. ([#73](https://github.com/northwood-labs/terraform-provider-corefunc/issues/73)) ([@skyzyx](https://github.com/skyzyx))
* [`c047ee1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c047ee16b8db3378bcb76dc430d258be9730568c): Added support for `corefunc.Homedir()` and `corefunc_homedir_get`. ([#77](https://github.com/northwood-labs/terraform-provider-corefunc/issues/77)) ([@skyzyx](https://github.com/skyzyx))
* [`e3985ba`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e3985ba1d576e273a14fe930eefb81999a695c53): Added initial support for Homedir lookups. ([@skyzyx](https://github.com/skyzyx))
* [`290e1cc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/290e1cc1a5138e533935d1d8160d86edecc5b065): Added `corefunc_homedir_expand`. ([@skyzyx](https://github.com/skyzyx))

## 1.0.3 — 2023-11-21

[Compare: v1.0.0 → v1.0.3](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.0...v1.0.3)

### :dependabot: Building and Dependencies

* [`14759b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/14759b8792c77530a0538d2fdb3ed2cacddd8d73): Bump `the` version number to 1.0.3. No changes from 1.0.0. ([@skyzyx](https://github.com/skyzyx))

## 1.0.0 — 2023-11-21

<p>Generated on {{ now() | date(format="%Y-%m-%d") }}.</p>
