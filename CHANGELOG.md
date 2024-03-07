# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com), adheres to [Semantic Versioning](https://semver.org), and uses [Conventional Commit](https://www.conventionalcommits.org) syntax.

## Unreleased

[Compare: v1.3.1 → `HEAD`](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.3.1..HEAD)

### :dependabot: Building and Dependencies

* [`789a93c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/789a93c2557434b571361bb9601f7f7195e9105b): **deps**: Bump `trufflesecurity/trufflehog` from 3.68.3 to 3.68.4 ([#178](https://github.com/northwood-labs/terraform-provider-corefunc/issues/178)) ([@dependabot](https://github.com/dependabot))
* [`9416af7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9416af70c4c4d29924ac356cf170a8556082401c): **deps**: Bump `github.com/stretchr/testify` in /terratest ([#179](https://github.com/northwood-labs/terraform-provider-corefunc/issues/179)) ([@dependabot](https://github.com/dependabot))
* [`0f3695f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0f3695f668737a3a81deb338e39d21c7680786f5): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#180](https://github.com/northwood-labs/terraform-provider-corefunc/issues/180)) ([@dependabot](https://github.com/dependabot))
* [`f679757`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f679757feabde6f9cf5dd243a0140dcc8324325b): **deps**: Bump `github.com/hashicorp/terraform-plugin-testing` ([#181](https://github.com/northwood-labs/terraform-provider-corefunc/issues/181)) ([@dependabot](https://github.com/dependabot))
* [`1947c93`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1947c93b28fea3a6e322723dfa644128fd7ded85): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#182](https://github.com/northwood-labs/terraform-provider-corefunc/issues/182)) ([@dependabot](https://github.com/dependabot))
* [`80c5b50`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/80c5b50cefc2e864eebfd346c284613cc3f5eaca): Updated the Go dependencies. ([@skyzyx](https://github.com/skyzyx))

### :soap: Linting

* [`9a28c7f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9a28c7f5b6ccdb418249b2a4c251d2a60192daf4): Exclude cliff.toml from EditorConfig rules. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`1359e49`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1359e4958872a963f8572d1772516672671a0533): Updated minimum Go version to v1.22.1, which addresses vulns. ([@skyzyx](https://github.com/skyzyx))
* [`c075bdc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c075bdc220405e560b60727305fa397a72bbcdde): Allow connections to api.github.com. ([@skyzyx](https://github.com/skyzyx))
* [`7d27121`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7d2712186b50efb6ccfe8fe8b806833dbfa86e46): Allow tests to connect to objects.githubusercontent.com:443. ([@skyzyx](https://github.com/skyzyx))
* [`1bb42ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1bb42ef146ffc645ba3c5f09dc9630f7822c6d66): Allow tests to connect to golang.org:443. ([@skyzyx](https://github.com/skyzyx))

## 1.3.1 — 2024-03-01

[Compare: v1.3.0 → v1.3.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.3.0...v1.3.1)

### :books: Documentation

* [`0f5913a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0f5913a747a9de59024fc0a1822b955227a283f2): Clarified documentation for `corefunc_url_parse`.canonicalizer. ([@skyzyx](https://github.com/skyzyx))
* [`e9e3b7a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e9e3b7a7df92d425b8af1165856b6137b14073b8): Clarified documentation for `corefunc_url_parse`.canonicalizer. ([@skyzyx](https://github.com/skyzyx))
* [`6dfac39`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6dfac39a87673043a8a0f78362c1b2c8b5170ebb): Added note about different kinds of providers. ([@skyzyx](https://github.com/skyzyx))
* [`761d230`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/761d230f025d6aa43d69d21a9a0f818453a3544d): Renamed `AUTHORS.md` to `CONTRIBUTORS.md` to better align with established patterns on the web. ([@skyzyx](https://github.com/skyzyx))
* [`2758a7e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2758a7e9a60c6f9909c326a71550ad5e5cefe755): Create a new `ACKNOWLEDGMENTS.md` file. ([@skyzyx](https://github.com/skyzyx))
* [`877ed3f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/877ed3f137eba9f4a1706ab1e33c1037edf6b7ed): Updated the Contribution Guide. ([@skyzyx](https://github.com/skyzyx))
* [`c1cfa03`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c1cfa0370008cb03771449d060c98107b401bbed): Updated the charts for test coverage and binsize. ([@skyzyx](https://github.com/skyzyx))
* [`9dffe13`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9dffe130440ae850818a32bcb2776184c518d44e): Updated the Contribution Guide with more clarity. ([@skyzyx](https://github.com/skyzyx))
* [`95ae2e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/95ae2e791526e0a640cc24bffa08860ee79b7581): Lots of tweaking to the CHANGELOG format to produce more consistent results. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`08799f4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/08799f414450923a2596b2914d748618ab641f36): Added a specific 'replace' clause to `go.mod`. ([@skyzyx](https://github.com/skyzyx))
* [`e72bd58`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e72bd58d488f38ae0abffcc3353fca47bbfb66fd): **deps**: Bump `actions/upload-artifact` from 4.3.0 to 4.3.1 ([#145](https://github.com/northwood-labs/terraform-provider-corefunc/issues/145)) ([@dependabot](https://github.com/dependabot))
* [`4812eb9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4812eb9a1cb294502840db8ee757ad8752ab4461): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#147](https://github.com/northwood-labs/terraform-provider-corefunc/issues/147)) ([@dependabot](https://github.com/dependabot))
* [`d5368da`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d5368da178ed5596d12b6ff0f0e4431499926daf): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.1 to 3.67.4 ([#148](https://github.com/northwood-labs/terraform-provider-corefunc/issues/148)) ([@dependabot](https://github.com/dependabot))
* [`c49b3f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c49b3f2356e2c49b923ab191862e094c4275d0d0): **deps**: Bump `actions/setup-node` from 4.0.1 to 4.0.2 ([#149](https://github.com/northwood-labs/terraform-provider-corefunc/issues/149)) ([@dependabot](https://github.com/dependabot))
* [`d8e97c0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d8e97c01e838855d2c1c19476ce0483959e295d5): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.4 to 3.67.5 ([#150](https://github.com/northwood-labs/terraform-provider-corefunc/issues/150)) ([@dependabot](https://github.com/dependabot))
* [`1912070`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/19120703de4526419d7509c7d4051757e14e2804): **deps**: Bump `golangci/golangci-lint-action` from 3.7.0 to 3.7.1 ([#151](https://github.com/northwood-labs/terraform-provider-corefunc/issues/151)) ([@dependabot](https://github.com/dependabot))
* [`2dc57ba`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2dc57ba40f5b2196f3e4ed1ee057d70e2e1166a7): Prepend sudo to chag installer. ([@skyzyx](https://github.com/skyzyx))
* [`6b11f0f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6b11f0f14aa4c37f5f275eb969a9d04b2af03038): Remove former dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`091b22e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/091b22e9d80efb729bdf22c7e0b338d7bcdb1c98): Keep default.pgo in the repo. ([@skyzyx](https://github.com/skyzyx))
* [`db40302`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/db40302a0763c2f3c052b294295ac3c39ef86531): Added the GPG Public Key used for signing. ([@skyzyx](https://github.com/skyzyx))
* [`77f1ea9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/77f1ea9f970165e552e0183b2cf760a8345c3cc7): **deps**: Bump `github/codeql-action` from 3.24.0 to 3.24.1 ([#152](https://github.com/northwood-labs/terraform-provider-corefunc/issues/152)) ([@dependabot](https://github.com/dependabot))
* [`71866d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/71866d3c9570ef4249e69139182dc6cce7a06f9f): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.5 to 3.67.6 ([#153](https://github.com/northwood-labs/terraform-provider-corefunc/issues/153)) ([@dependabot](https://github.com/dependabot))
* [`3b76f93`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b76f93ade1290f2895c2635aad30bddd2c8084e): Add support for running single benchmarks. ([@skyzyx](https://github.com/skyzyx))
* [`0533fbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0533fbe8131c03d1e322727cfde0129733ed236e): Updated the dependencies to their latest minor/patch releases. ([@skyzyx](https://github.com/skyzyx))
* [`23003f6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/23003f6f9f9de0a8ecb1caf913e788986524fb65): **deps**: Bump `actions/dependency-review-action` from 4.0.0 to 4.1.0 ([#162](https://github.com/northwood-labs/terraform-provider-corefunc/issues/162)) ([@dependabot](https://github.com/dependabot))
* [`211ef6f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/211ef6f57831f2fcdb1f4a86b3ae079b0cf10319): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.6 to 3.67.7 ([#168](https://github.com/northwood-labs/terraform-provider-corefunc/issues/168)) ([@dependabot](https://github.com/dependabot))
* [`fa4eef0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fa4eef06823503760efb0057c407faf3c5fd9480): **deps**: Bump `github/codeql-action` from 3.24.1 to 3.24.3 ([#164](https://github.com/northwood-labs/terraform-provider-corefunc/issues/164)) ([@dependabot](https://github.com/dependabot))
* [`339904e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/339904e7ac5af4a460525b1419874a773c105923): **deps**: Bump `actions/go-dependency-submission` from 2.0.0 to 2.0.1 ([#165](https://github.com/northwood-labs/terraform-provider-corefunc/issues/165)) ([@dependabot](https://github.com/dependabot))
* [`6c924d8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6c924d85af3efbd7bdc2e2bcd911b15f52e81c34): **deps**: Bump `actions/dependency-review-action` from 4.1.0 to 4.1.3 ([#167](https://github.com/northwood-labs/terraform-provider-corefunc/issues/167)) ([@dependabot](https://github.com/dependabot))
* [`b19428b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b19428b0fa575ef09806d2406489194ed62f9e22): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.7 to 3.68.0 ([#169](https://github.com/northwood-labs/terraform-provider-corefunc/issues/169)) ([@dependabot](https://github.com/dependabot))
* [`5cef9c0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5cef9c0aeb100119208dd771c9d84687484d6168): **deps**: Bump `golang/govulncheck-action` from 1.0.1 to 1.0.2 ([#170](https://github.com/northwood-labs/terraform-provider-corefunc/issues/170)) ([@dependabot](https://github.com/dependabot))
* [`355da49`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/355da49dd1be6f41dbdb4f5bbf1d5130e2e7d367): **deps**: Bump `github/codeql-action` from 3.24.3 to 3.24.4 ([#171](https://github.com/northwood-labs/terraform-provider-corefunc/issues/171)) ([@dependabot](https://github.com/dependabot))
* [`0346ed9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0346ed99f2dc4ef5c6d68bcd7522044eefe0c9f1): **deps**: Bump `Go` to 1.22. ([@skyzyx](https://github.com/skyzyx))
* [`88ad2e8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/88ad2e8a9359147972ef935397a90f067ea26576): **deps**: Bump `github/codeql-action` from 3.24.4 to 3.24.6 ([#177](https://github.com/northwood-labs/terraform-provider-corefunc/issues/177)) ([@dependabot](https://github.com/dependabot))
* [`deb3174`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/deb31740b532d4311051c74a39abbcc286dd43d4): **deps**: Bump `trufflesecurity/trufflehog` from 3.68.0 to 3.68.3 ([#176](https://github.com/northwood-labs/terraform-provider-corefunc/issues/176)) ([@dependabot](https://github.com/dependabot))
* [`c4624c0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c4624c0586b3781ec97895a99a1553e5540ff5dd): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#175](https://github.com/northwood-labs/terraform-provider-corefunc/issues/175)) ([@dependabot](https://github.com/dependabot))
* [`617991e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/617991e78fbf88dc0b109ff2bfd4d7106ce9e19f): **deps**: Bump `github.com/hashicorp/terraform-plugin-go` ([#172](https://github.com/northwood-labs/terraform-provider-corefunc/issues/172)) ([@dependabot](https://github.com/dependabot))
* [`088f2f7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/088f2f7e3e6881eae6eed63eae2e6da279a3d35f): Cleanup the version when tagging a release in the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`388b043`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/388b0430c5816de40358407a8e138363f8bd4372): **deps**: Updated archstring dependency. ([@skyzyx](https://github.com/skyzyx))
* [`876b473`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/876b47312d9dbd244b6e7f64262e6d11dec6bfa9): Update PGO data. ([@skyzyx](https://github.com/skyzyx))

### :soap: Linting

* [`113cb5f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/113cb5f21e9cbf19d8e7c8cd9a48ab30491cce4c): Resolved issues caught by the Go linter. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`68b0221`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/68b02214644de4a22ceb5338287c9d4dc2f5969b): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`e8adfe2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e8adfe28f690e43d221808bfbdf749acb538a710): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`6755d2a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6755d2aef20e433d56225333fa94754ab713c5fc): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`8d2248e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d2248ea3fa3b32cb3749a71afe0517b0bbb7ef4): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`3b1289a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b1289a48abaf53e84dc5e9c0673770123de44ab): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`1174f10`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1174f10e33f178acc449896aa8651ec00c5bb3cc): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`25e7a69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/25e7a69eb1e3167a7b47bb3baffcd3ecdf9cdb9e): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`0190c0b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0190c0bdf9b6b52e92d31d319d53445a67e7532b): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`e9ed6e1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e9ed6e13f0e0345036a324eeaf7d763b7803ee66): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`28306c5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/28306c57cf0d812ff6ecf38bdb823a5c180fb29f): Allow more network connections for test suite. ([@skyzyx](https://github.com/skyzyx))
* [`237a6ee`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/237a6eea432d109504b98d10d851eb0d7c898935): Allow more outgoing connections from the runners. ([@skyzyx](https://github.com/skyzyx))
* [`0f7246f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0f7246fec05d3cf18ff796f605b910eb64cf8d2a): Allow more outgoing connections from the runners. ([@skyzyx](https://github.com/skyzyx))
* [`e223a24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e223a247d6b90fdc6d8ceb4dd1b871be0390df83): Allow outbound connections to goreleaser.com. ([@skyzyx](https://github.com/skyzyx))
* [`d6e3832`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6e3832643759fb7e868265d63b72d8203d8855a): Allow outbound connections to api.gumroad.com. ([@skyzyx](https://github.com/skyzyx))
* [`e6f4441`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e6f4441137647e01947e1316932f16b39a167344): Allow outbound connections to *.golang.org. ([@skyzyx](https://github.com/skyzyx))
* [`1c007e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1c007e7ba57acc7f70149c5e20c266fc20f46260): Allow outbound connections to *.golang.org. ([@skyzyx](https://github.com/skyzyx))
* [`08f3b6f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/08f3b6f3aa27a575392b5f1675b71bafecd0d01a): Allow outbound connections to *.golang.org. ([@skyzyx](https://github.com/skyzyx))

## 1.3.0 — 2024-02-05

[Compare: v1.2.1 → v1.3.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.2.1...v1.3.0)

## 1.2.1 — 2024-02-02

[Compare: v1.2.0 → v1.2.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.2.0...v1.2.1)

## 1.2.0 — 2024-02-02

[Compare: v1.1.1 → v1.2.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.1...v1.2.0)

## 1.1.1 — 2023-12-20

[Compare: v1.1.0 → v1.1.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.0...v1.1.1)

## 1.1.0 — 2023-12-18

[Compare: v1.0.3 → v1.1.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.3...v1.1.0)

## 1.0.3 — 2023-11-21

[Compare: v1.0.0 → v1.0.3](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.0...v1.0.3)

## 1.0.0 — 2023-11-21

### :art: Styling

* [`d451600`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d451600b1b601c242979cae99f7c915f86742ec4): Minor adjustments to the Makefile and author script. ([@skyzyx](https://github.com/skyzyx))

### :books: Documentation

* [`e9e2ad6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e9e2ad680320dedd9449bdce46343910399a7830): **terraform**: Updated the documentation. ([@skyzyx](https://github.com/skyzyx))
* [`cf79e84`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cf79e84019cbcb7c854a1dbef4dfd751d85fb61f): **terraform**: More documentation edits. ([@skyzyx](https://github.com/skyzyx))
* [`f4341db`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4341dba1db6bdbd3bdef234fdc81f08a41ddcb7): Updated the Contribution Guide. ([@skyzyx](https://github.com/skyzyx))
* [`540584d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/540584ddf3f7fc6f0c420c237e79fba4347b1350): Updated the Go package documentation. ([@skyzyx](https://github.com/skyzyx))
* [`4e9c027`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4e9c027a434c98501d9b7527336008da0eb9947f): Preparing support for the Terraform Registry. ([@skyzyx](https://github.com/skyzyx))
* [`f4278de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4278de00eac06cffb5d4b6112784312b3897f83): Added an example to `TruncateLabel()`. ([@skyzyx](https://github.com/skyzyx))
* [`822bdf9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/822bdf9d6b966034a00a76218690c3c569a1fcee): Updated the output of a make task. ([@skyzyx](https://github.com/skyzyx))
* [`c12ccfc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c12ccfcf4f0209997a9e4f750ae66db3d655775d): Update the formatting of the provider examples. ([@skyzyx](https://github.com/skyzyx))
* [`56b682a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/56b682aa6fc8143e874b734fabc07d82aa6cf9ec): Mark a paragraph as a note in the Terraform Registry docs. ([@skyzyx](https://github.com/skyzyx))
* [`b2a5e4e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b2a5e4e82fcd75c094e6b9650672fd37e7089a4b): Update the list of planned features. ([@skyzyx](https://github.com/skyzyx))
* [`16f8552`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/16f8552bff32fd211c3f90076c7e897e92050889): **terraform**: Updated examples. ([@skyzyx](https://github.com/skyzyx))
* [`52efecc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/52efecc177335987ff16113c7cfe1b88f8ffa848): **readme**: Update `README.md` ([@skyzyx](https://github.com/skyzyx))
* [`b163ca6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b163ca6f8086afc3582a443b7252c7d377af533d): Removed the list of to-dos from the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`a5b1aab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a5b1aab169ce6338525810178bc4670bd14a06ff): **provider**: Split the examples for env_ensure into separate code blocks. ([@skyzyx](https://github.com/skyzyx))
* [`8670d50`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8670d50ba1961e6a95e0d415e7ad4b964e1637b7): **provider**: Split the examples for env_ensure into separate code blocks. ([@skyzyx](https://github.com/skyzyx))
* [`e204b4a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e204b4ac6f53579ae5eec37a624308d2f6a5c540): Updated `corefunc_str_snake` docs. ([@skyzyx](https://github.com/skyzyx))
* [`918eb60`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/918eb60b43194bc93dbc72987e4c8879fc770eb5): Updated the provider example. ([@skyzyx](https://github.com/skyzyx))
* [`6664868`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/666486822d8d220654c028ae7c739ba7ceb757b7): Clarified the `CONTRIBUTING.md` guide. ([@skyzyx](https://github.com/skyzyx))
* [`ff5fa12`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ff5fa1291e92e13ed621db7d93877b57a2202683): Added Git LFS to the documentation. ([@skyzyx](https://github.com/skyzyx))
* [`7325224`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73252242e32ec5113248a2063584c8b6fda80fe9): Updated the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`c1e0640`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c1e064071c814286baa79dd9fe9be24a3e5921ec): Added documentation about BATS tests. ([@skyzyx](https://github.com/skyzyx))
* [`0e1f3ec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0e1f3ecdcd4293ea263cafc20666071f432284af): Define a  security policy. ([@skyzyx](https://github.com/skyzyx))
* [`ba01f24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ba01f2403d3892fc6d0e59391b8cad364c49061a): Renamed OpenTF → OpenTofu in the `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`81dfde8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/81dfde8d7b22940fe4543d83b8de668d1a05ae4d): Add badges to `README.md`. ([@skyzyx](https://github.com/skyzyx))
* [`5b173d1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5b173d15cea08cbc091738dca21d7a2201886171): Updated the generator docs. ([@skyzyx](https://github.com/skyzyx))
* [`5a014fb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5a014fb0ec891a3ef782aa36c2ee38d28d6c48c4): Updated the `README.md` and `AUTHORS.md`. ([@skyzyx](https://github.com/skyzyx))
* [`42e1e7d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/42e1e7d7f9274250178e71434e791c659a166802): Fixed issue link to `pkg.go.dev`. ([@skyzyx](https://github.com/skyzyx))
* [`be02409`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/be02409405351f0619bba8f165510c2bb1acd356): Added link to chanced/caps documentation. ([@skyzyx](https://github.com/skyzyx))
* [`1e7c19b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1e7c19bbca8de371380a52d4ffcdcfa65786d7dd): Removing the Markdown version of the license. ([@skyzyx](https://github.com/skyzyx))
* [`0be7a78`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0be7a78e18bd718a1eee181d8050de988f4a65e4): Trying to get both `pkg.go.dev` and GitHub to understand the license. ([@skyzyx](https://github.com/skyzyx))
* [`21deae6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/21deae652b7d6839c9c1069f019d3873308cf1da): Messing with the license again. ([@skyzyx](https://github.com/skyzyx))

### :dependabot: Building and Dependencies

* [`d6f79f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6f79f2273c1935f2f2bc9a9b80123ccf87589c1): **deps**: Updated the Go dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`44c1b50`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/44c1b50f7b0811c2a77a731ea447556479b3eb83): **deps**: Added testify as a test dependency. ([@skyzyx](https://github.com/skyzyx))
* [`0a66652`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0a66652c694115790b9af5228c81f0cae29e5377): **deps**: Updated GitHub project templates. ([@skyzyx](https://github.com/skyzyx))
* [`6f1057e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6f1057ee91498774d943b15cec01509edf8b745e): Get licensing in order. ([@skyzyx](https://github.com/skyzyx))
* [`5cefbd5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5cefbd569b23d4e95b1d6c371c8d861825a0fe14): **git**: Added a git hook to check commit message. ([@skyzyx](https://github.com/skyzyx))
* [`0a3d40c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0a3d40c4f9296dd0a0c15400e1543b176450b53a): **deps**: Setup some basic project governance. ([@skyzyx](https://github.com/skyzyx))
* [`ea42c5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ea42c5df77ba2c2a50f5c5964fe39b0c051cd4e9): **deps**: Added Trufflehog. ([@skyzyx](https://github.com/skyzyx))
* [`d47d060`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d47d060ff8ed07db91dcc15f6fe38f307b8b5bcb): Introduce PGO. ([@skyzyx](https://github.com/skyzyx))
* [`8d42454`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d4245463e3c787b17aef21061d638a7bc457143): **deps**: Updated Go dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`a0c8c7f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a0c8c7f1895165775c92f34ffb955a7957141a5d): **deps**: Integrate Dependabot. ([@skyzyx](https://github.com/skyzyx))
* [`394d0b0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/394d0b03061787989880f14879954356ef2e8ef6): **deps**: Integrate OSSF Scorecard. ([@skyzyx](https://github.com/skyzyx))
* [`6e12657`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6e12657fa729c3f80e7fd3ed4a703df39a17a66e): **deps**: Integrate CodeQL. ([@skyzyx](https://github.com/skyzyx))
* [`f03f919`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f03f9198d695ee7503c26b8dc3ab54d121d4c1e8): **deps**: Bump `golang.org/x/net` from 0.16.0 to 0.17.0 ([#28](https://github.com/northwood-labs/terraform-provider-corefunc/issues/28)) ([@dependabot](https://github.com/dependabot))
* [`97822cf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/97822cf3a60e3b9901fb96ee036bde540c6bbb16): **deps**: Bump `github.com/google/go-cmp` from 0.5.9 to 0.6.0 ([#30](https://github.com/northwood-labs/terraform-provider-corefunc/issues/30)) ([@dependabot](https://github.com/dependabot))
* [`35b7b63`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/35b7b63c5679b752093e18ce2aa4492b1cffe5a8): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#29](https://github.com/northwood-labs/terraform-provider-corefunc/issues/29)) ([@dependabot](https://github.com/dependabot))
* [`ac6dce2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ac6dce25c3a615e2d292df0e516eb7fdc44a75eb): `Pkg.go.dev` having trouble detecting a license. ([@skyzyx](https://github.com/skyzyx))
* [`41827d2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/41827d2b0c83c0c4ed5a27c215ef2bf0ca477b1d): Add more security scanning. ([@skyzyx](https://github.com/skyzyx))
* [`a93efad`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a93efadc92f392fb9f0e3214d3ad34eac0d61abe): **deps**: Bump `actions/upload-artifact` from 3.1.0 to 3.1.3 ([#33](https://github.com/northwood-labs/terraform-provider-corefunc/issues/33)) ([@dependabot](https://github.com/dependabot))
* [`f79b5da`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f79b5da624f6351b604bf38dfcf136fcba858042): **deps**: Bump `ossf/scorecard-action` from 2.1.2 to 2.3.0 ([#34](https://github.com/northwood-labs/terraform-provider-corefunc/issues/34)) ([@dependabot](https://github.com/dependabot))
* [`be9c9c3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/be9c9c333ef0cba5de5a0491ffd4f2ab8ca98833): **deps**: Bump `github/codeql-action` from 2.2.4 to 2.22.4 ([#35](https://github.com/northwood-labs/terraform-provider-corefunc/issues/35)) ([@dependabot](https://github.com/dependabot))
* [`62c5b97`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/62c5b97cee358f502ff95586ffe076360001ab0b): **deps**: Bump `actions/dependency-review-action` from 2.5.1 to 3.1.0 ([#36](https://github.com/northwood-labs/terraform-provider-corefunc/issues/36)) ([@dependabot](https://github.com/dependabot))
* [`a50fe00`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a50fe00c4d6cc183ed837a6b02ff3719eeb6590f): Update `AUTHORS.md`. ([@skyzyx](https://github.com/skyzyx))
* [`3aad86b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3aad86b042e0145fce576487fd9ce250513499d7): Remove devcontainer for the time being. ([@skyzyx](https://github.com/skyzyx))
* [`5307c33`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5307c33ef37bc8373fd87247332f5a1a655cb30b): Add golangci-lint to CI. ([@skyzyx](https://github.com/skyzyx))
* [`6c752ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6c752aeb05727c68fe3be30878b2afca16d9439d): **deps**: Bump `google.golang.org/grpc` from 1.58.2 to 1.58.3 ([#40](https://github.com/northwood-labs/terraform-provider-corefunc/issues/40)) ([@dependabot](https://github.com/dependabot))
* [`ef529fb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ef529fb5b1da2e7f5223f1103296163c52a743eb): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#39](https://github.com/northwood-labs/terraform-provider-corefunc/issues/39)) ([@dependabot](https://github.com/dependabot))
* [`0981387`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/09813878b0270e56aebb03f7aeb742931871d2f7): **deps**: Bump `ossf/scorecard-action` from 2.3.0 to 2.3.1 ([#38](https://github.com/northwood-labs/terraform-provider-corefunc/issues/38)) ([@dependabot](https://github.com/dependabot))
* [`3792e1c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3792e1c465f6c8ab36a6f806f256356ebe4c0752): **deps**: Bump `github/codeql-action` from 2.22.4 to 2.22.5 ([#53](https://github.com/northwood-labs/terraform-provider-corefunc/issues/53)) ([@dependabot](https://github.com/dependabot))
* [`7db0e5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7db0e5ddc66ff91525ce08185046ee05aec303d7): Testing Sigstore Gitsign commit signing. ([@skyzyx](https://github.com/skyzyx))
* [`b9846fd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b9846fd5dab730e513d106c48321eb66cf590721): Testing Sigstore Gitsign commit signing. ([@skyzyx](https://github.com/skyzyx))
* [`bd005ed`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bd005ed4bd4c1200000e6b6f456166e4362b3899): **deps**: Bump `hashicorp/setup-terraform` from 2.0.3 to 3.0.0 ([#56](https://github.com/northwood-labs/terraform-provider-corefunc/issues/56)) ([@dependabot](https://github.com/dependabot))
* [`a0cabe1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a0cabe1d4b2686ab6057c05c3733558528b5a0c2): **deps**: Bump `actions/dependency-review-action` from 3.1.0 to 3.1.3 ([#61](https://github.com/northwood-labs/terraform-provider-corefunc/issues/61)) ([@dependabot](https://github.com/dependabot))
* [`ee012ce`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ee012ce9473571640b7a1b76fe80a551db077bbe): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#60](https://github.com/northwood-labs/terraform-provider-corefunc/issues/60)) ([@dependabot](https://github.com/dependabot))
* [`8d9004c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d9004cc706a78935808621c70934ef0861febf4): **deps**: Bump `github.com/spf13/cobra` from 1.7.0 to 1.8.0 ([#57](https://github.com/northwood-labs/terraform-provider-corefunc/issues/57)) ([@dependabot](https://github.com/dependabot))
* [`b58ec24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b58ec2443a157beec3e68771ab91d1ca7108c2ae): Update the `go.sum` file. ([@skyzyx](https://github.com/skyzyx))
* [`247c257`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/247c25731a9df60daa9dc86f850506090c36404e): Update Go dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`6ee61b7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6ee61b75a6224179ca4e71942f9626d5bd95e6ea): **deps**: Bump `github/codeql-action` from 2.22.5 to 2.22.6 ([#62](https://github.com/northwood-labs/terraform-provider-corefunc/issues/62)) ([@dependabot](https://github.com/dependabot))
* [`85864cd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/85864cd8b4ae635c548f02c9e7372e245e3457f1): Added editorconfig-checker as a Homebrew dependency. ([@skyzyx](https://github.com/skyzyx))
* [`668be53`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/668be53fb2b9a2565a3dfcf2d5939bf9518031e4): Update the Go dependencies. ([@skyzyx](https://github.com/skyzyx))
* [`f68f8f0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f68f8f00fef79acf3c7db82d81a0f7582f01616f): **deps**: Bump `github/codeql-action` from 2.22.6 to 2.22.7 ([#65](https://github.com/northwood-labs/terraform-provider-corefunc/issues/65)) ([@dependabot](https://github.com/dependabot))
* [`47ff449`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/47ff449b6e55e7406e1fde01f51b8170caebbe51): **deps**: Bump `step-security/harden-runner` from 2.6.0 to 2.6.1 ([#64](https://github.com/northwood-labs/terraform-provider-corefunc/issues/64)) ([@dependabot](https://github.com/dependabot))
* [`63775bf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/63775bf0368919d4f0ef403abd3293650f46fb47): Replaced much of Licensei with Trivy for licenses. ([#66](https://github.com/northwood-labs/terraform-provider-corefunc/issues/66)) ([@skyzyx](https://github.com/skyzyx))
* [`e0b5924`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e0b59240e4a8dce4f86499b668b038e8f478d55f): Updated license cache data. ([@skyzyx](https://github.com/skyzyx))
* [`33557b7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/33557b7ca19fce19b28ded1b97c4cb0d62a3734b): Added git-cliff as a dependency for changelogs. ([@skyzyx](https://github.com/skyzyx))
* [`d3d344e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d3d344e370a173613d7cab7bfa80a832c2215707): Remove default.pgo. ([@skyzyx](https://github.com/skyzyx))
* [`b267f1e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b267f1ee74b109510ceb38ef0f0e9e9a49f88d59): Prepping a new test tag to trigger the build. ([@skyzyx](https://github.com/skyzyx))
* [`5d52625`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d52625a44197115de7e90b6a2bf3b9c14f901d8): Prepping a new test tag to trigger the build. ([@skyzyx](https://github.com/skyzyx))
* [`dee279e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/dee279e1b75b230716ccc17517cd41c3b5d6e4e4): Prepping a new test tag to trigger the build. ([@skyzyx](https://github.com/skyzyx))

### :soap: Linting

* [`6770816`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6770816190f2262d71c4f0f482c43ed59b8e09b2): Updated the linter definitions. ([@skyzyx](https://github.com/skyzyx))
* [`83a1a67`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/83a1a674f6930ede51c3703735a4ea707904b4c6): Cleaned up the YAML files. ([@skyzyx](https://github.com/skyzyx))
* [`a4085ee`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a4085ee2139ba9e737ec16791be7d07b634c79f9): Cleaned up the JSON files. ([@skyzyx](https://github.com/skyzyx))
* [`cc22b71`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cc22b718bf9f8e52cfd3039463735e19ebc8dab9): Cleaned up the Markdown files. ([@skyzyx](https://github.com/skyzyx))
* [`e2bc42a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e2bc42a0328d5300fd6537c42668591277ded90d): Updated the Markdownlint rules. ([@skyzyx](https://github.com/skyzyx))
* [`f33108f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f33108fdfeffe97d4efead1d87331fc2d3bbc65c): Removed unnecessary pre-commit rules. ([@skyzyx](https://github.com/skyzyx))
* [`14356ab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/14356abb4e1cfa1a5ed0934857c2e172ff8f439d): Ran the linter and performed a little cleanup. ([@skyzyx](https://github.com/skyzyx))
* [`f4c2504`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4c2504d3e31b70c2e800b3e8fd7f672ac35dc8a): Run yamlfmt. ([@skyzyx](https://github.com/skyzyx))
* [`778d4cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/778d4cb0e637a6b898e5e232348a2f77e993ad41): Run yamlfmt. ([@skyzyx](https://github.com/skyzyx))
* [`de6a00b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/de6a00bbe308feab45c5e2630dfe105a87de3462): Added and applied more linting. ([@skyzyx](https://github.com/skyzyx))
* [`58555a5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/58555a5f351e6b65c773205a7f7eac574df9fa00): Resolved conflicts in linting rules. ([@skyzyx](https://github.com/skyzyx))
* [`2be9512`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2be9512991c7ee4638a5c8f0a48bcee815e89a92): Add CodeQL analysis. ([@skyzyx](https://github.com/skyzyx))
* [`15321a7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/15321a74b62d07a2aa72d5198bb88c02deca9759): Improve CodeQL analysis. ([@skyzyx](https://github.com/skyzyx))
* [`f941cf1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f941cf1a6ca38bcc987298e1713b444944037b86): Remove CodeQL analysis YAML. Rely on built-in feature. ([@skyzyx](https://github.com/skyzyx))
* [`31b5e9a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/31b5e9a3497668bc7405af51223dc6cf3eb83c48): Additional linting. ([@skyzyx](https://github.com/skyzyx))
* [`f0fe79b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f0fe79b289b4d613bf2d9b4b3c5a6ac56029f143): Additional linting. ([@skyzyx](https://github.com/skyzyx))
* [`58c6d30`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/58c6d308a1d334d02223451707fc7521cb6121cc): Fix Go version. ([@skyzyx](https://github.com/skyzyx))
* [`480862d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/480862d2f01247acb8512d39f5b36f62315614be): Update PR review config. ([@skyzyx](https://github.com/skyzyx))
* [`e5e2328`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e5e2328833d15644869aa78be79d5223b9c00608): Move test running into CI. ([@skyzyx](https://github.com/skyzyx))
* [`f334dc4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f334dc40d54c7056ca371fb69b38a911d90b6300): Verified that all files adhere to .editorconfig. ([@skyzyx](https://github.com/skyzyx))

### :test_tube: Testing

* [`c80d484`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c80d484232cdd54545f7aa926c0bb6fd55cf0905): Add a debug mode for acc tests. ([@skyzyx](https://github.com/skyzyx))
* [`15662fc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/15662fc5ecb07db57eec6d7fbbe35866b1eb005a): Added mutation tests. ([@skyzyx](https://github.com/skyzyx))
* [`4546390`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4546390ca53c108cb424614c10ef1d493cb601db): Fix the tests running in CI. ([@skyzyx](https://github.com/skyzyx))
* [`2100e0c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2100e0cf246193cdd808f1c7c32fa584894c8e21): Fix the tests running in CI. ([@skyzyx](https://github.com/skyzyx))
* [`650e1c5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/650e1c5a89f284124d4306bb01dfd36210f8a094): Fix the tests running in CI. ([@skyzyx](https://github.com/skyzyx))
* [`d7df64f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d7df64f8533221b01b98feafe48cc6d1f190e71e): Fix the tests running in CI. ([@skyzyx](https://github.com/skyzyx))
* [`d54db2d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d54db2d14880d822348a9f098828755f38bf5adb): Manually install latest Go for govulncheck. ([@skyzyx](https://github.com/skyzyx))
* [`f965fbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f965fbe6a0246138cf8d80e50ea705a54ddde548): Manually install latest Go for govulncheck. ([@skyzyx](https://github.com/skyzyx))
* [`78966d5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/78966d5edc0e3c57f7ab86827f8d3595b2442072): Allow `storage.googleapis.com:443` from GitHub Actions. ([@skyzyx](https://github.com/skyzyx))

### :tractor: Refactor

* [`d24133f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d24133f28f580002231227dc247daa9a09a189f6): Fix naming of functions and data sources. ([@skyzyx](https://github.com/skyzyx))
* [`b946b69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b946b693e5728b1cbd0d272485f2db880b2db6e5): Split the library code from the provider code. ([@skyzyx](https://github.com/skyzyx))
* [`df2690f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/df2690f01db54e17ae8bbcb547168e133d336303): **make**: Split pre-commit out from lint. ([@skyzyx](https://github.com/skyzyx))
* [`3c7a078`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3c7a078b6376d8b9c76796e52949e3d7c445eb24): Remove line from Makefile; add it to GoReleaser later. ([@skyzyx](https://github.com/skyzyx))

### <!-- 0 -->:rocket: Features

* [`d4a065d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d4a065d012cb4ad546054c774d614c6a3d5582e8): Initial commit of code. We have one function working. ([@skyzyx](https://github.com/skyzyx))
* [`2f0a4ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2f0a4ae04d1b32b78c980dc4df03976ddb9abcbb): Added the Go code for the `EnvEnsure()` function. ([@skyzyx](https://github.com/skyzyx))
* [`b69b67b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b69b67bff73b68fc4b389896ae3d48a712ab5612): Added the Terraform wrapper and docs for `EnvEnsure()`. ([@skyzyx](https://github.com/skyzyx))
* [`5c75b92`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5c75b922b13185dfee646b7bea6a7f839ed96703): Significant improvements to the Makefile. ([@skyzyx](https://github.com/skyzyx))
* [`6f96ddf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6f96ddf400090026fc38561323e5564d6776ab8e): Add support for regex patterns to `EnvEnsure()`. ([@skyzyx](https://github.com/skyzyx))
* [`7f2e3b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7f2e3b8de1e9875b2b32eea2a3f838d5b8ef6d50): Simplify Makefile by reading build data. ([@skyzyx](https://github.com/skyzyx))
* [`27060e4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/27060e41a312bb2b1cc9cfd0850cfc443beca198): Adding some initial checking for conventional commits. ([@skyzyx](https://github.com/skyzyx))
* [`3228245`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/32282453ad3e2d3f58d85d3623c9ed1d728f9f19): Initial attempt at a devcontainer. ([@skyzyx](https://github.com/skyzyx))
* [`b763957`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b763957d7f4412de9f9258783fef50a5afc14fe1): Added `corefunc_str_snake`. ([@skyzyx](https://github.com/skyzyx))
* [`3065ca0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3065ca02b0331733eb417cf3a42a55229fa528a9): Added `corefunc_str_camel`. ([@skyzyx](https://github.com/skyzyx))
* [`690c7f8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/690c7f84d6f0d67f7cad4123fc7f90ac4669e887): Added `corefunc_str_pascal`. ([@skyzyx](https://github.com/skyzyx))
* [`67c33fd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/67c33fd2058f46f16568db3f15b5ca7222460857): Added `corefunc_str_kebab`. ([@skyzyx](https://github.com/skyzyx))
* [`c6f0a0d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c6f0a0d243f450292bf82e4e611d9154e2906b71): Added `corefunc_str_constant`. ([@skyzyx](https://github.com/skyzyx))
* [`02db190`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/02db1904110bf9b0950e97d63cc7034a73620056): Added code generation for new provider functions. ([@skyzyx](https://github.com/skyzyx))
* [`194dde1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/194dde1db6bec67f8c4c82273071a4c73729c41d): Implemented support for `str_iterative_replace`. ([#31](https://github.com/northwood-labs/terraform-provider-corefunc/issues/31)) ([@skyzyx](https://github.com/skyzyx))
* [`280dc21`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/280dc217f860709fbeb82e20e7657cd63e5c91c3): Modernize issue creation. ([@skyzyx](https://github.com/skyzyx))
* [`edd5210`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/edd521023634afd33a9efd38a137bea368aec392): Create issue template for bug reports. ([@skyzyx](https://github.com/skyzyx))
* [`9080b41`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9080b417f2cb4713da83f5e83578a078a5c5e49d): Added a template for feature requests. ([@skyzyx](https://github.com/skyzyx))
* [`261aa0d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/261aa0dcab848277960a1a02e1231096d788d569): Update `feature.yml` ([@skyzyx](https://github.com/skyzyx))
* [`65fa14c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/65fa14cd32b098464d28704b8a3ba88b0b359b67): Added generation for binsize. ([@skyzyx](https://github.com/skyzyx))
* [`0e21531`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0e21531d2ac18bfb0adfe36cb13e95f602dfca5d): Improve code coverage reporting. ([@skyzyx](https://github.com/skyzyx))

### <!-- 1 -->:bug: Bug Fixes

* [`425bcd0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/425bcd0c60b57c99e8811bb91e29fa194904ca96): Fixed the last failing test for `EnvEnsure()`. ([@skyzyx](https://github.com/skyzyx))
* [`52540b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/52540b8caf8794046afdf8f9bb135797e8764138): Removed appName from the build. ([@skyzyx](https://github.com/skyzyx))
* [`edf0671`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/edf06712deb67a69d08764d92e87dcd480acb37d): Some cleanup in consistency. ([@skyzyx](https://github.com/skyzyx))
* [`0ee237a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0ee237a94781673e3c34372048f55f5b8209c565): Allow Terraform 1.6+ after upgrading to terraform-plugin-framework v1.4.1. ([@skyzyx](https://github.com/skyzyx))
* [`3b0c194`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b0c19493005d66c3baad7bdf5c28b3109d23b3c): Update `SECURITY.md` to point to GitHub's tool. ([@skyzyx](https://github.com/skyzyx))
* [`c7cba09`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c7cba0942031eab84b60e8fe3b52dfc94be908ea): Apply security best practices from StepSecurity. ([#32](https://github.com/northwood-labs/terraform-provider-corefunc/issues/32)) ([@step-security-bot](https://github.com/step-security-bot))

### <!-- ZZZ -->:gear: Miscellaneous Tasks

* [`7e8e142`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7e8e14253539a76669a1c251ec6f0d7918808db4): Cleaned up some mis-named jobs. ([@skyzyx](https://github.com/skyzyx))
* [`993304e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/993304ed23f234cf5959c1ff236f9287180b07ea): Cleaned up some mis-named jobs. ([@skyzyx](https://github.com/skyzyx))
* [`d1276b3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d1276b3ffac3ab3d93d031d63f88c42de9393983): Cleaned up some mis-named jobs. ([@skyzyx](https://github.com/skyzyx))
* [`a1dd8ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a1dd8efc0aef98d1269e4e3beccd709cfcec9003): Update the GoReleaser configuration. ([@skyzyx](https://github.com/skyzyx))
* [`07dc031`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/07dc0315da588063861b54862b72e78d8f10fa74): Testing GoReleaser config changes. ([@skyzyx](https://github.com/skyzyx))
* [`e61c254`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e61c254b8008b5b9a8c2e18f7c317024a92be0e7): Update CHANGELOG configuration for GoReleaser. ([@skyzyx](https://github.com/skyzyx))
* [`6901616`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6901616b298e5be1bf159d4d6f7c54e6ec339894): Run GoReleaser on Git tag. ([@skyzyx](https://github.com/skyzyx))
* [`b6447ec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b6447ec9b6848f35c1d0449fe305bc2bc3d6afb9): Improving the security of the workflow. ([@skyzyx](https://github.com/skyzyx))
* [`eb9f7ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/eb9f7aecfdaa7bda9087b9b58229d3943571a863): Update the GoReleaser configuration. ([@skyzyx](https://github.com/skyzyx))
* [`a01b12e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a01b12e6870b821f561624bca79407e4c87c73cf): Trying to get the build to work. ([@skyzyx](https://github.com/skyzyx))

<p>Generated on 2024-03-07.</p>
