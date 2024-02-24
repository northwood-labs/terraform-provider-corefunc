# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com), adheres to
[Semantic Versioning](https://semver.org), and uses
[Conventional Commit](https://www.conventionalcommits.org) syntax.

## Unreleased

[Compare: v1.3.0 → `HEAD`](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.3.0..HEAD)

### Building and Dependencies

* [`08799f4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/08799f414450923a2596b2914d748618ab641f36): Added a specific 'replace' clause to `go.mod`.
* [`e72bd58`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e72bd58d488f38ae0abffcc3353fca47bbfb66fd): **deps**: Bump `actions/upload-artifact` from 4.3.0 to 4.3.1 ([#145](https://github.com/northwood-labs/terraform-provider-corefunc/issues/145)) (dependabot[bot])
* [`4812eb9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4812eb9a1cb294502840db8ee757ad8752ab4461): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#147](https://github.com/northwood-labs/terraform-provider-corefunc/issues/147)) (dependabot[bot])
* [`d5368da`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d5368da178ed5596d12b6ff0f0e4431499926daf): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.1 to 3.67.4 ([#148](https://github.com/northwood-labs/terraform-provider-corefunc/issues/148)) (dependabot[bot])
* [`c49b3f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c49b3f2356e2c49b923ab191862e094c4275d0d0): **deps**: Bump `actions/setup-node` from 4.0.1 to 4.0.2 ([#149](https://github.com/northwood-labs/terraform-provider-corefunc/issues/149)) (dependabot[bot])
* [`d8e97c0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d8e97c01e838855d2c1c19476ce0483959e295d5): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.4 to 3.67.5 ([#150](https://github.com/northwood-labs/terraform-provider-corefunc/issues/150)) (dependabot[bot])
* [`1912070`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/19120703de4526419d7509c7d4051757e14e2804): **deps**: Bump `golangci/golangci-lint-action` from 3.7.0 to 3.7.1 ([#151](https://github.com/northwood-labs/terraform-provider-corefunc/issues/151)) (dependabot[bot])
* [`2dc57ba`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2dc57ba40f5b2196f3e4ed1ee057d70e2e1166a7): Prepend sudo to chag installer.
* [`6b11f0f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6b11f0f14aa4c37f5f275eb969a9d04b2af03038): Remove former dependencies.
* [`091b22e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/091b22e9d80efb729bdf22c7e0b338d7bcdb1c98): Keep default.pgo in the repo.
* [`db40302`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/db40302a0763c2f3c052b294295ac3c39ef86531): Added the GPG Public Key used for signing.
* [`77f1ea9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/77f1ea9f970165e552e0183b2cf760a8345c3cc7): **deps**: Bump `github/codeql-action` from 3.24.0 to 3.24.1 ([#152](https://github.com/northwood-labs/terraform-provider-corefunc/issues/152)) (dependabot[bot])
* [`71866d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/71866d3c9570ef4249e69139182dc6cce7a06f9f): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.5 to 3.67.6 ([#153](https://github.com/northwood-labs/terraform-provider-corefunc/issues/153)) (dependabot[bot])
* [`3b76f93`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b76f93ade1290f2895c2635aad30bddd2c8084e): Add support for running single benchmarks.
* [`0533fbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0533fbe8131c03d1e322727cfde0129733ed236e): Updated the dependencies to their latest minor/patch releases.
* [`23003f6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/23003f6f9f9de0a8ecb1caf913e788986524fb65): **deps**: Bump `actions/dependency-review-action` from 4.0.0 to 4.1.0 ([#162](https://github.com/northwood-labs/terraform-provider-corefunc/issues/162)) (dependabot[bot])
* [`211ef6f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/211ef6f57831f2fcdb1f4a86b3ae079b0cf10319): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.6 to 3.67.7 ([#168](https://github.com/northwood-labs/terraform-provider-corefunc/issues/168)) (dependabot[bot])
* [`fa4eef0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fa4eef06823503760efb0057c407faf3c5fd9480): **deps**: Bump `github/codeql-action` from 3.24.1 to 3.24.3 ([#164](https://github.com/northwood-labs/terraform-provider-corefunc/issues/164)) (dependabot[bot])
* [`339904e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/339904e7ac5af4a460525b1419874a773c105923): **deps**: Bump `actions/go-dependency-submission` from 2.0.0 to 2.0.1 ([#165](https://github.com/northwood-labs/terraform-provider-corefunc/issues/165)) (dependabot[bot])
* [`6c924d8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6c924d85af3efbd7bdc2e2bcd911b15f52e81c34): **deps**: Bump `actions/dependency-review-action` from 4.1.0 to 4.1.3 ([#167](https://github.com/northwood-labs/terraform-provider-corefunc/issues/167)) (dependabot[bot])
* [`b19428b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b19428b0fa575ef09806d2406489194ed62f9e22): **deps**: Bump `trufflesecurity/trufflehog` from 3.67.7 to 3.68.0 ([#169](https://github.com/northwood-labs/terraform-provider-corefunc/issues/169)) (dependabot[bot])
* [`5cef9c0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5cef9c0aeb100119208dd771c9d84687484d6168): **deps**: Bump `golang/govulncheck-action` from 1.0.1 to 1.0.2 ([#170](https://github.com/northwood-labs/terraform-provider-corefunc/issues/170)) (dependabot[bot])
* [`355da49`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/355da49dd1be6f41dbdb4f5bbf1d5130e2e7d367): **deps**: Bump `github/codeql-action` from 3.24.3 to 3.24.4 ([#171](https://github.com/northwood-labs/terraform-provider-corefunc/issues/171)) (dependabot[bot])
* [`0346ed9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0346ed99f2dc4ef5c6d68bcd7522044eefe0c9f1): **deps**: Bump `Go` to 1.22.

### Documentation

* [`0f5913a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0f5913a747a9de59024fc0a1822b955227a283f2): Clarified documentation for `corefunc_url_parse`.canonicalizer.
* [`e9e3b7a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e9e3b7a7df92d425b8af1165856b6137b14073b8): Clarified documentation for `corefunc_url_parse`.canonicalizer.
* [`6dfac39`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6dfac39a87673043a8a0f78362c1b2c8b5170ebb): Added note about different kinds of providers.
* [`761d230`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/761d230f025d6aa43d69d21a9a0f818453a3544d): Renamed `AUTHORS.md` to CONTRIBUTORS to better align with established patterns on the web.
* [`2758a7e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2758a7e9a60c6f9909c326a71550ad5e5cefe755): Create a new `ACKNOWLEDGMENTS.md` file.
* [`877ed3f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/877ed3f137eba9f4a1706ab1e33c1037edf6b7ed): Updated the Contribution Guide.
* [`c1cfa03`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c1cfa0370008cb03771449d060c98107b401bbed): Updated the charts for test coverage and binsize.
* [`9dffe13`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9dffe130440ae850818a32bcb2776184c518d44e): Updated the Contribution Guide with more clarity.

### Linting

* [`113cb5f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/113cb5f21e9cbf19d8e7c8cd9a48ab30491cce4c): Resolved issues caught by the Go linter.

### Testing

* [`68b0221`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/68b02214644de4a22ceb5338287c9d4dc2f5969b): Allow more network connections for test suite.
* [`e8adfe2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e8adfe28f690e43d221808bfbdf749acb538a710): Allow more network connections for test suite.
* [`6755d2a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6755d2aef20e433d56225333fa94754ab713c5fc): Allow more network connections for test suite.
* [`8d2248e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d2248ea3fa3b32cb3749a71afe0517b0bbb7ef4): Allow more network connections for test suite.
* [`3b1289a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b1289a48abaf53e84dc5e9c0673770123de44ab): Allow more network connections for test suite.
* [`1174f10`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1174f10e33f178acc449896aa8651ec00c5bb3cc): Allow more network connections for test suite.
* [`25e7a69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/25e7a69eb1e3167a7b47bb3baffcd3ecdf9cdb9e): Allow more network connections for test suite.
* [`0190c0b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0190c0bdf9b6b52e92d31d319d53445a67e7532b): Allow more network connections for test suite.
## 1.3.0 — 2024-02-05

[Compare: v1.2.1 → v1.3.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.2.1..v1.3.0)

### Bug Fixes

* [`fc46b7f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fc46b7f8bf4542d86b8bfe7c7c1eb9ac2843dc9d): Updated testing framework to not require the 'id' attribute.
* [`50affa9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/50affa91e6aa6a6c0b59c529070a6d296afd0823): Lower provider logging messages from INFO to DEBUG.
* [`6582fec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6582fecf49b9e400ba7c0d5bbf8b3d483a89aaf4): Fix examples and Shellcheck warnings.

### Building and Dependencies

* [`4cde8d8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4cde8d89e62d2201dcfb05a80b709f7b1258cf3f): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#130](https://github.com/northwood-labs/terraform-provider-corefunc/issues/130)) (dependabot[bot])
* [`1a298cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1a298cb78cc4d9e6144a9b4282e71311a348b144): **deps**: Bump `github/codeql-action` from 3.23.2 to 3.24.0 ([#131](https://github.com/northwood-labs/terraform-provider-corefunc/issues/131)) (dependabot[bot])
* [`68a3d4b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/68a3d4b110cd6f4b36d4f12bb0df3160372847a1): **deps**: Bump `trufflesecurity/trufflehog` from 3.66.3 to 3.67.1 ([#132](https://github.com/northwood-labs/terraform-provider-corefunc/issues/132)) (dependabot[bot])
* [`e548051`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e54805173ee44bd7b13fac10406dd0d0e1968511): Small updates to the Makefile.
* [`d4496cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d4496cbae6c11de0c55511cdbf20d5a5b5b2c908): Small updates to the Makefile.
* [`87f2772`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/87f277226fc1f3c7593438f78e590692dac45764): Small updates to the Makefile.
* [`b4232d2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b4232d2f0fa6f63ea05208730525cd0392adbebb): Small updates to the Makefile.

### Documentation

* [`00411fe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/00411fee4350e0be6dcf78bbf2ade87ee0c5b03b): Updated docs to remove the 'id' attribute.
* [`e819279`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e8192790d23b54e8120055b610b17ccc1682f1ee): Fixed a few typos.
* [`88db492`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/88db4929008d062f86dc7bd65c36d7b52b997d30): Fixed some links in `corefunc_url_parse`.
* [`ca6d540`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca6d5402046aa39a56f060f8fb4165fc35f688be): Fixed some typos in `corefunc_url_parse`.
* [`d02b3e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d02b3e7a913824b09ebfb9411d810f1473b91073): Fixed some typos in `corefunc_url_parse`.
* [`d6938de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6938de38a824a648ddf9842dec9e0e3cd99cbbf): Added link to Google Safe Browsing rules.

### Features

* [`1752c6e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1752c6e94daba7aec9a1b6c4740e16b7b3764be6): Implement new runtime data sources.
* [`1da303c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1da303cebdfc414f62a5907bb1f2526dc66b6067): Implemented `corefunc_url_parse`.
* [`e179a2f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e179a2fdd814662fc2dff30ce3dea1db06cbb27a): Added support for Google Safe Browsing canonicalization.

### Testing

* [`5d992f3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d992f31b4d5566e37bb12a0d133400d75d1763f): Regenerate test artifacts.
* [`c4459ad`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c4459ad7664dc6a1c93b3585705dc46d72cb76d2): Updated BATS tests with Shellcheck feedback.
* [`6611f3f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6611f3fbd38af4997fb156c22d4157024bd8d76f): Generated updated test artifacts.
* [`d6f3437`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6f34376297dc934e195c379786053729b8625bb): Added Terratest tests for `corefunc_url_parse`.
## 1.2.1 — 2024-02-02

[Compare: v1.1.1 → v1.2.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.1..v1.2.1)

### Bug Fixes

* [`32dffbb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/32dffbb12f962543d923f2ea39000675d0ff8714): Resolve issues in the `go.sum` file.
* [`0533928`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/05339289d72baeb1a72b6983c0f6132e264ba1c8): Remove a cache file from the repo as it changes every time.
* [`1ad7a2b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1ad7a2b160a44b4534883d9ce84969eb7cef519f): Improve the version selection automation.
* [`e952382`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e95238280e794773f9826abb60480382a9855d5c): Improve the version selection automation.
* [`3e1a07a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3e1a07a433aa6e0be88ed139ee45cde343e5b1d6): Improve the version selection automation.

### Building and Dependencies

* [`92c7065`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92c7065ba587b0851aa9352a1422ac7d2d68e8fa): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#94](https://github.com/northwood-labs/terraform-provider-corefunc/issues/94)) (dependabot[bot])
* [`96a741f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/96a741f1f36af211f80934694295736df522aed7): **deps**: Bump `crazy-max/ghaction-import-gpg` from 6.0.0 to 6.1.0 ([#97](https://github.com/northwood-labs/terraform-provider-corefunc/issues/97)) (dependabot[bot])
* [`73ac9c5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73ac9c52d6275e2c7ea25b14a7677a6a91174cc2): **deps**: Bump `github/codeql-action` from 1.1.39 to 3.22.12 ([#95](https://github.com/northwood-labs/terraform-provider-corefunc/issues/95)) (dependabot[bot])
* [`528dfcd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/528dfcdd689857b97757f887c96a34ca4369d83c): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.5 to 3.63.7 ([#96](https://github.com/northwood-labs/terraform-provider-corefunc/issues/96)) (dependabot[bot])
* [`523b701`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/523b7012e3cc18de7879f30a3537136cb11a05a9): **deps**: Bump `github.com/chanced/caps` from 1.0.1 to 1.0.2 ([#98](https://github.com/northwood-labs/terraform-provider-corefunc/issues/98)) (dependabot[bot])
* [`b311501`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b3115010b0e84d6adcc5d3a75610f1f07697b5a7): **deps**: Bump `github.com/chanced/caps` in /generator ([#99](https://github.com/northwood-labs/terraform-provider-corefunc/issues/99)) (dependabot[bot])
* [`fc859eb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/fc859ebd974529d165cd107fd1f7ab33a55826e3): **deps**: Bump `actions/dependency-review-action` from 3.1.4 to 3.1.5 ([#100](https://github.com/northwood-labs/terraform-provider-corefunc/issues/100)) (dependabot[bot])
* [`1ec1ac2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1ec1ac2488f25d39289247b882044a51335d0125): **deps**: Bump `github.com/cloudflare/circl` from 1.3.6 to 1.3.7 ([#101](https://github.com/northwood-labs/terraform-provider-corefunc/issues/101)) (dependabot[bot])
* [`7f40e69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7f40e69e02c6314e0728c776e3376492e9e1ccbe): **deps**: Bump `actions/go-dependency-submission` from 1.0.3 to 2.0.0 ([#105](https://github.com/northwood-labs/terraform-provider-corefunc/issues/105)) (dependabot[bot])
* [`1317bab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1317babb5379cd55528a2b932905acfbea9253de): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#106](https://github.com/northwood-labs/terraform-provider-corefunc/issues/106)) (dependabot[bot])
* [`8d22b24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d22b24cd8cc52a6ebdabc524a61f2f8ee785627): **deps**: Bump `github.com/gruntwork-io/terratest` in /terratest ([#110](https://github.com/northwood-labs/terraform-provider-corefunc/issues/110)) (dependabot[bot])
* [`104b541`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/104b541701e30608c39913de362524454f73a925): **deps**: Bump `github/codeql-action` from 3.22.12 to 3.23.1 ([#111](https://github.com/northwood-labs/terraform-provider-corefunc/issues/111)) (dependabot[bot])
* [`0063c0c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0063c0ce2b277b5201d14ab7cf6862c9963a3443): **deps**: Bump `actions/dependency-review-action` from 3.1.5 to 4.0.0 ([#113](https://github.com/northwood-labs/terraform-provider-corefunc/issues/113)) (dependabot[bot])
* [`2ca888e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2ca888ed39eae52a622fbe9049d6ecabdf2041f9): **deps**: Bump `actions/upload-artifact` from 4.0.0 to 4.3.0 ([#116](https://github.com/northwood-labs/terraform-provider-corefunc/issues/116)) (dependabot[bot])
* [`1750875`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1750875ff6005e48b8d8405c2b45ee486b116fab): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.7 to 3.64.0 ([#117](https://github.com/northwood-labs/terraform-provider-corefunc/issues/117)) (dependabot[bot])
* [`0844a2e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0844a2eb4e5d6ee49b69c5a6c9b5a8d7bc510c4a): **deps**: Bump `github.com/hashicorp/terraform-plugin-docs` ([#118](https://github.com/northwood-labs/terraform-provider-corefunc/issues/118)) (dependabot[bot])
* [`680a5d4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/680a5d4b3e437d1f21cc636995b924990cb1d6f3): **deps**: Bump `github/codeql-action` from 3.23.1 to 3.23.2 ([#120](https://github.com/northwood-labs/terraform-provider-corefunc/issues/120)) (dependabot[bot])
* [`771756b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/771756b2138312dadefebb7df8a651b6b9f6f93c): **deps**: Bump `github.com/hashicorp/terraform-plugin-go` ([#121](https://github.com/northwood-labs/terraform-provider-corefunc/issues/121)) (dependabot[bot])
* [`f35b806`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f35b806374aa31334d04075f0dd92e54e08b59ef): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#123](https://github.com/northwood-labs/terraform-provider-corefunc/issues/123)) (dependabot[bot])
* [`5396f2d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5396f2d698dfacaa6edc402d2c0e38b90fa34570): **deps**: Bump `step-security/harden-runner` from 2.6.1 to 2.7.0 ([#124](https://github.com/northwood-labs/terraform-provider-corefunc/issues/124)) (dependabot[bot])
* [`d0df49d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d0df49d95113e5fbe29ba592d3f51717798472e6): **deps**: Bump `trufflesecurity/trufflehog` from 3.64.0 to 3.66.3 ([#126](https://github.com/northwood-labs/terraform-provider-corefunc/issues/126)) (dependabot[bot])
* [`558c756`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/558c756c6a0b91d40233ee446d3dece24a1c3b53): **deps**: Bump `github.com/northwood-labs/terraform-provider-corefunc` ([#128](https://github.com/northwood-labs/terraform-provider-corefunc/issues/128)) (dependabot[bot])

### Documentation

* [`c30e191`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c30e1912d5e273e8238703d3f7bc67a492017e66): Added information about compatibility testing.
* [`b35d3f8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b35d3f85fd3cb8ed0a8355b6f631878484f3be98): Adjust the git-cliff format.
* [`6189ded`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6189ded118f391007e0efda00aafa90636caee3a): Resolved an issue with Markdown generation build more consistent docs.

### Linting

* [`f2c349f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f2c349f8c33e07c6fd2dcd6ac5e55825f9eaf53b): Ran the linter against the codebase after changes.

### Testing

* [`ea785fa`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ea785fa9dc992c8df52756501798240d63dacf91): Disable GoSec Sarif uploading until we know why it's failing.
* [`d2198f3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d2198f3152d878657f9bb77b3ab81affd15dc45e): Add Terraform 1.7 to the test matrix.
* [`73b83b1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73b83b1982933903b8a4e512a2af06f38607c624): Enable workflow_dispatch to trigger tests manually.
* [`4205631`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4205631e3b504b8448c31fde9a1603e1565128b1): Discovered issue in calling GoSec.
* [`9486823`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9486823382c8c36f3895ffaeb58bfb325c1cf344): Discovered issue in calling GoSec.
* [`472cb4c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/472cb4cb3ff48742d01a4d97a656e653d54dd835): Discovered issue in calling GoSec.
* [`8686958`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8686958d2d46f52438279d249c1455e4a038ea04): Disable GoSec workflow for the time being.
* [`bf75f94`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bf75f945ffdbc872d138c9f6effaf29d64639b7a): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`6672a01`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6672a0174d74b913a5462e1ea5856c3e37c900b2): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`834be43`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/834be43f30a1e68d1da701887e11ee3db623d124): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`40b77b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/40b77b88163ee6a711bd5770985f7a6fc5528ee4): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`e1e7b69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e1e7b695b4fc564b4b56a5ece28d674df3bb86a0): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`9749f52`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9749f529e519a8f6978e40f55eeb3b1f4fdbd308): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`6949371`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6949371a0332b730dad43c48c529dfd23714310c): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`157efbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/157efbe38e925b53deb4d5b729b650ecbfa014bc): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`9c09697`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9c096972030ce0d870d2df5100bab7bfb2d37a71): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`7ddd3e8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7ddd3e86d64d49528e310aa42456420b7f41d2a5): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`d7d77f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d7d77f292f2f8c2b56fd35419d9d9e01ad774dec): Auto-update `AUTHORS.md` and CHANGELOG on commit.
* [`2c282c8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2c282c8e863ab3527dcc605ac3a74ce3a28c026e): Work on adding automated OpenTofu testing.
* [`a20e1bf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a20e1bfaa2eea228abca619bc81840b4ee303ef1): Work on adding automated OpenTofu testing.
* [`82583e7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/82583e7bb9128268e702e65b5a9a90ed1f69b6f7): Work on adding automated OpenTofu testing.
* [`5d9cd26`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d9cd26136cebc600044b527874a97d6847ff641): Work on adding automated OpenTofu testing.
* [`2758ec4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2758ec4de8a6cfa8c99859786f43e526c83fa843): Work on adding automated OpenTofu testing.
* [`92c5e44`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92c5e440d9425fb83c11bf9ca403f2b85b74c6a5): Work on adding automated OpenTofu testing.
* [`ca1650f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca1650f2f2e9737a6635851bd19795e2d1b22908): Work on adding automated OpenTofu testing.
* [`ae243c3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ae243c36dba8429a3606b223495d2f94ec5b82e9): Work on adding automated OpenTofu testing.
* [`56501ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/56501ef5e2e26ec6ca6c1c06a1aa1eb7987cb280): Work on adding automated OpenTofu testing.
* [`4a1e899`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4a1e899979c9a3593fd741484bdf1e8b3febfe4c): Work on adding automated OpenTofu testing.
* [`a8b9724`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a8b97249bf45c2fc71fc6d32ffc6470a5b3dc5c3): Work on adding automated OpenTofu testing.
* [`d374a8c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d374a8c42fa2e3c20f410762f7369d41088bad6e): Work on adding automated OpenTofu testing.
* [`44c7648`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/44c76484bc48005d0484162a43bb5efa2170f724): Work on adding automated OpenTofu testing.
* [`b5f2b8e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b5f2b8e0382900ef4a85f0e5e0052c066b9da624): Work on adding automated OpenTofu testing.
* [`6be7398`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6be7398843e0a757ab4f0308a175843d8603c829): Work on adding automated OpenTofu testing.
* [`66abdf2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/66abdf213d2f88f3ee0e3d44b26dd0b557b25005): Work on adding automated OpenTofu testing.
* [`cc77ed9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cc77ed99085f9c81281dd7a6dbcc74298d4cdbe1): Improve CHANGELOG generation.
* [`0182c79`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0182c79239bb888e61e7affcd3ddfaeced41c0d2): Improve CHANGELOG generation.
* [`acee3ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/acee3ef09ce457f074bb25f8aada3d029e474027): Add auto-cancelation for multiple CI runs of the same test.
* [`85cfb6c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/85cfb6c8dfed8d61bf84d246ba58eacbf91cc650): Setup Terratest testing for OpenTofu from the Makefile.
* [`57b1e0e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/57b1e0eb5f91a87e5450dba5d6f9492fb2c0b374): Improve switching between Terraform/OpenTofu in Terratest.
* [`590c3bc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/590c3bce781674b2627e494b203570b142555844): Improve switching between Terraform/OpenTofu in Terratest.
## 1.1.1 — 2023-12-20

[Compare: v1.1.0 → v1.1.1](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.1.0..v1.1.1)

### Building and Dependencies

* [`0c431c2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0c431c21046806396bfe392f32fdfa5ba67743f3): Remove the restriction to pre-1.6 versions of Terraform in examples.
* [`27b9e60`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/27b9e60796e0bfdc3b972cae11bca284c4cfc8d2): Updated Go dependencies.
* [`bc39137`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bc391374137cb742fa719bd3790288cfa4a0c7cd): Pin the version of the GoSec action.
* [`aa4e217`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/aa4e2178e06786ed2a0adb21e0095326854904fb): Pin the version of the Trufflehog action.
* [`1d01f26`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1d01f2686604a931602bdd0880a640ac13d312ce): **deps**: Bump `golang.org/x/crypto` in /terratest ([#90](https://github.com/northwood-labs/terraform-provider-corefunc/issues/90)) (dependabot[bot])
* [`2f77df9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2f77df906e466c58ba3d4eb7d328684da4aab20e): **deps**: Bump `golang.org/x/crypto` from 0.3.0 to 0.17.0 in /generator ([#89](https://github.com/northwood-labs/terraform-provider-corefunc/issues/89)) (dependabot[bot])
* [`07df7b9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/07df7b911593bf1b65fbe060b005aae5b0e4faa6): **deps**: Bump `trufflesecurity/trufflehog` from 3.63.4 to 3.63.5 ([#91](https://github.com/northwood-labs/terraform-provider-corefunc/issues/91)) (dependabot[bot])
* [`956339c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/956339c396e8b199a35d2f4eeff058dd11175be3): **deps**: Bump `golang.org/x/crypto` from 0.16.0 to 0.17.0 ([#88](https://github.com/northwood-labs/terraform-provider-corefunc/issues/88)) (dependabot[bot])
* [`543eb0a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/543eb0ab95e3918f034c76a99099ff0f8fcdd490): Update the `go.sum` file.

### Documentation

* [`f3c07b2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f3c07b267eacc0cc4c3fabd81628112f8ea974f2): Updated copyright in all headers to include 2024.
* [`1d00129`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1d0012982e3c1760735a4fa53a6603781bad6930): Updated the `README.md` for the Terratest tests.
* [`1c2bc3c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1c2bc3c6f91b6326cd188320d8978972e2aa78dc): Updated the templates used by tfplugindocs.

### Testing

* [`e0915df`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e0915dfb90f9d2824be769dcc629501db99d1b8c): Added tests with Terratest to compare the Terraform provider with the Go library.
* [`e5a810a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e5a810a8f9a5ffe2737928429df0be3a9a0a25c3): Make the GoSec scanning more robust.
## 1.1.0 — 2023-12-18

[Compare: v1.0.3 → v1.1.0](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.3..v1.1.0)

### Building and Dependencies

* [`de4027c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/de4027ce4f48f09525b5f0ce4c4000ebd6c75de1): **deps**: Bump `github/codeql-action` from 2.22.7 to 2.22.8 ([#72](https://github.com/northwood-labs/terraform-provider-corefunc/issues/72)) (dependabot[bot])
* [`98d8612`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/98d86121b239c9d15a1a3f429e90b07cf750399f): **deps**: Bump `actions/dependency-review-action` from 3.1.3 to 3.1.4 ([#74](https://github.com/northwood-labs/terraform-provider-corefunc/issues/74)) (dependabot[bot])
* [`874d704`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/874d7046af06c46ed86ea818936e846188d612d6): **deps**: Bump `actions/setup-go` from 4.1.0 to 5.0.0 ([#75](https://github.com/northwood-labs/terraform-provider-corefunc/issues/75)) (dependabot[bot])
* [`ca8c440`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ca8c4409fb32db9b83647e085839b73923ded5ec): **deps**: Bump `github/codeql-action` from 2.22.8 to 2.22.9 ([#76](https://github.com/northwood-labs/terraform-provider-corefunc/issues/76)) (dependabot[bot])
* [`232c76e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/232c76e6ecaaa95e43e78182e02c259f2c8bb3b0): **deps**: Bump `github/codeql-action` from 2.22.9 to 3.22.11 ([#79](https://github.com/northwood-labs/terraform-provider-corefunc/issues/79)) (dependabot[bot])
* [`415d3d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/415d3d3bf5a648cb3168ab1015059613d53046f4): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#82](https://github.com/northwood-labs/terraform-provider-corefunc/issues/82)) (dependabot[bot])
* [`78ec578`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/78ec578eb53836a96aaf93ae8b78cb69217941b8): **deps**: Bump `actions/upload-artifact` from 3.1.3 to 4.0.0 ([#80](https://github.com/northwood-labs/terraform-provider-corefunc/issues/80)) (dependabot[bot])

### Documentation

* [`3b7d8ca`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b7d8cac9ddc38904505c5c9fbbd47fc15f54df5): Updated the `README.md`.
* [`d41ad2e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d41ad2e1dafbfac54dacf3ac9d49b50032122089): Updated the Terraform documentation to include the lockfile.
* [`00b7206`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/00b7206578b826912a4e7563d7d6caa2588432a3): Updated the CHANGELOG.
* [`92a73de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92a73de97de90a269afd08f4369ee4c5ad4aee59): Ran Markdownlint on the CHANGELOG.
* [`6b7b0a3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6b7b0a375e9b17be41055a4f8a2c2dce6401b674): Added a note about base10 to `corefunc_int_leftpad`.
* [`b83580e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b83580eeac1b986d8b0fd7d667da31356f7d8bd5): Regenerated the provider documentation.

### Features

* [`1135f94`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1135f9439b1966e3da849c2b9272f84eb89cfd36): Begin implementations of StrLeftPad and IntLeftPad.
* [`92fbf6f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/92fbf6f5e7f4fd17b1a33135a851143be1fd276e): Implemented the provider side of leftpad functions. ([#73](https://github.com/northwood-labs/terraform-provider-corefunc/issues/73))
* [`c047ee1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c047ee16b8db3378bcb76dc430d258be9730568c): Added support for `corefunc.Homedir()` and `corefunc_homedir_get`. ([#77](https://github.com/northwood-labs/terraform-provider-corefunc/issues/77))
* [`e3985ba`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e3985ba1d576e273a14fe930eefb81999a695c53): Added initial support for Homedir lookups.
* [`290e1cc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/290e1cc1a5138e533935d1d8160d86edecc5b065): Added `corefunc_homedir_expand`.

### Linting

* [`89b97d3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/89b97d37ea0495e56fefe181ef7f473680e73a0b): Updated the cache of dependency license information.

### Refactor

* [`2204ce8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2204ce86c8d36127cf84163559254549e819590f): Added wrapper corefunc functions for all caps functions.

### Testing

* [`57a7b85`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/57a7b85c17040f880dcaf6546c11cb9acd0a7615): Lookup and run fuzz tests programmatically.
## 1.0.3 — 2023-11-21

[Compare: v1.0.0 → v1.0.3](https://github.com/northwood-labs/terraform-provider-corefunc/compare/v1.0.0..v1.0.3)

### Building and Dependencies

* [`14759b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/14759b8792c77530a0538d2fdb3ed2cacddd8d73): Bump `the` version number to 1.0.3. No changes from 1.0.0.
## 1.0.0 — 2023-11-21

### Bug Fixes

* [`425bcd0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/425bcd0c60b57c99e8811bb91e29fa194904ca96): Fixed the last failing test for `EnvEnsure()`.
* [`52540b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/52540b8caf8794046afdf8f9bb135797e8764138): Removed appName from the build.
* [`edf0671`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/edf06712deb67a69d08764d92e87dcd480acb37d): Some cleanup in consistency.
* [`0ee237a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0ee237a94781673e3c34372048f55f5b8209c565): Allow Terraform 1.6+ after upgrading to terraform-plugin-framework v1.4.1.
* [`3b0c194`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3b0c19493005d66c3baad7bdf5c28b3109d23b3c): Update `SECURITY.md` to point to GitHub's tool.
* [`c7cba09`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c7cba0942031eab84b60e8fe3b52dfc94be908ea): Apply security best practices from StepSecurity. ([#32](https://github.com/northwood-labs/terraform-provider-corefunc/issues/32)) (StepSecurity Bot)

### Building and Dependencies

* [`d6f79f2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d6f79f2273c1935f2f2bc9a9b80123ccf87589c1): **deps**: Updated the Go dependencies.
* [`44c1b50`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/44c1b50f7b0811c2a77a731ea447556479b3eb83): **deps**: Added testify as a test dependency.
* [`0a66652`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0a66652c694115790b9af5228c81f0cae29e5377): **deps**: Updated GitHub project templates.
* [`6f1057e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6f1057ee91498774d943b15cec01509edf8b745e): Get licensing in order.
* [`5cefbd5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5cefbd569b23d4e95b1d6c371c8d861825a0fe14): **git**: Added a git hook to check commit message.
* [`0a3d40c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0a3d40c4f9296dd0a0c15400e1543b176450b53a): **deps**: Setup some basic project governance.
* [`ea42c5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ea42c5df77ba2c2a50f5c5964fe39b0c051cd4e9): **deps**: Added Trufflehog.
* [`d47d060`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d47d060ff8ed07db91dcc15f6fe38f307b8b5bcb): Introduce PGO.
* [`8d42454`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d4245463e3c787b17aef21061d638a7bc457143): **deps**: Updated Go dependencies.
* [`a0c8c7f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a0c8c7f1895165775c92f34ffb955a7957141a5d): **deps**: Integrate Dependabot.
* [`394d0b0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/394d0b03061787989880f14879954356ef2e8ef6): **deps**: Integrate OSSF Scorecard.
* [`6e12657`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6e12657fa729c3f80e7fd3ed4a703df39a17a66e): **deps**: Integrate CodeQL.
* [`f03f919`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f03f9198d695ee7503c26b8dc3ab54d121d4c1e8): **deps**: Bump `golang.org/x/net` from 0.16.0 to 0.17.0 ([#28](https://github.com/northwood-labs/terraform-provider-corefunc/issues/28)) (dependabot[bot])
* [`97822cf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/97822cf3a60e3b9901fb96ee036bde540c6bbb16): **deps**: Bump `github.com/google/go-cmp` from 0.5.9 to 0.6.0 ([#30](https://github.com/northwood-labs/terraform-provider-corefunc/issues/30)) (dependabot[bot])
* [`35b7b63`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/35b7b63c5679b752093e18ce2aa4492b1cffe5a8): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#29](https://github.com/northwood-labs/terraform-provider-corefunc/issues/29)) (dependabot[bot])
* [`ac6dce2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ac6dce25c3a615e2d292df0e516eb7fdc44a75eb): `Pkg.go.dev` having trouble detecting a license.
* [`41827d2`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/41827d2b0c83c0c4ed5a27c215ef2bf0ca477b1d): Add more security scanning.
* [`a93efad`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a93efadc92f392fb9f0e3214d3ad34eac0d61abe): **deps**: Bump `actions/upload-artifact` from 3.1.0 to 3.1.3 ([#33](https://github.com/northwood-labs/terraform-provider-corefunc/issues/33)) (dependabot[bot])
* [`f79b5da`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f79b5da624f6351b604bf38dfcf136fcba858042): **deps**: Bump `ossf/scorecard-action` from 2.1.2 to 2.3.0 ([#34](https://github.com/northwood-labs/terraform-provider-corefunc/issues/34)) (dependabot[bot])
* [`be9c9c3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/be9c9c333ef0cba5de5a0491ffd4f2ab8ca98833): **deps**: Bump `github/codeql-action` from 2.2.4 to 2.22.4 ([#35](https://github.com/northwood-labs/terraform-provider-corefunc/issues/35)) (dependabot[bot])
* [`62c5b97`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/62c5b97cee358f502ff95586ffe076360001ab0b): **deps**: Bump `actions/dependency-review-action` from 2.5.1 to 3.1.0 ([#36](https://github.com/northwood-labs/terraform-provider-corefunc/issues/36)) (dependabot[bot])
* [`a50fe00`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a50fe00c4d6cc183ed837a6b02ff3719eeb6590f): Update `AUTHORS.md`.
* [`3aad86b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3aad86b042e0145fce576487fd9ce250513499d7): Remove devcontainer for the time being.
* [`5307c33`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5307c33ef37bc8373fd87247332f5a1a655cb30b): Add golangci-lint to CI.
* [`6c752ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6c752aeb05727c68fe3be30878b2afca16d9439d): **deps**: Bump `google.golang.org/grpc` from 1.58.2 to 1.58.3 ([#40](https://github.com/northwood-labs/terraform-provider-corefunc/issues/40)) (dependabot[bot])
* [`ef529fb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ef529fb5b1da2e7f5223f1103296163c52a743eb): **deps**: Bump `github.com/hashicorp/terraform-plugin-framework` ([#39](https://github.com/northwood-labs/terraform-provider-corefunc/issues/39)) (dependabot[bot])
* [`0981387`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/09813878b0270e56aebb03f7aeb742931871d2f7): **deps**: Bump `ossf/scorecard-action` from 2.3.0 to 2.3.1 ([#38](https://github.com/northwood-labs/terraform-provider-corefunc/issues/38)) (dependabot[bot])
* [`3792e1c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3792e1c465f6c8ab36a6f806f256356ebe4c0752): **deps**: Bump `github/codeql-action` from 2.22.4 to 2.22.5 ([#53](https://github.com/northwood-labs/terraform-provider-corefunc/issues/53)) (dependabot[bot])
* [`7db0e5d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7db0e5ddc66ff91525ce08185046ee05aec303d7): Testing Sigstore Gitsign commit signing.
* [`b9846fd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b9846fd5dab730e513d106c48321eb66cf590721): Testing Sigstore Gitsign commit signing.
* [`bd005ed`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/bd005ed4bd4c1200000e6b6f456166e4362b3899): **deps**: Bump `hashicorp/setup-terraform` from 2.0.3 to 3.0.0 ([#56](https://github.com/northwood-labs/terraform-provider-corefunc/issues/56)) (dependabot[bot])
* [`a0cabe1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a0cabe1d4b2686ab6057c05c3733558528b5a0c2): **deps**: Bump `actions/dependency-review-action` from 3.1.0 to 3.1.3 ([#61](https://github.com/northwood-labs/terraform-provider-corefunc/issues/61)) (dependabot[bot])
* [`ee012ce`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ee012ce9473571640b7a1b76fe80a551db077bbe): **deps**: Bump `github.com/hashicorp/terraform-plugin-sdk/v2` ([#60](https://github.com/northwood-labs/terraform-provider-corefunc/issues/60)) (dependabot[bot])
* [`8d9004c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8d9004cc706a78935808621c70934ef0861febf4): **deps**: Bump `github.com/spf13/cobra` from 1.7.0 to 1.8.0 ([#57](https://github.com/northwood-labs/terraform-provider-corefunc/issues/57)) (dependabot[bot])
* [`b58ec24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b58ec2443a157beec3e68771ab91d1ca7108c2ae): Update the `go.sum` file.
* [`247c257`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/247c25731a9df60daa9dc86f850506090c36404e): Update Go dependencies.
* [`6ee61b7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6ee61b75a6224179ca4e71942f9626d5bd95e6ea): **deps**: Bump `github/codeql-action` from 2.22.5 to 2.22.6 ([#62](https://github.com/northwood-labs/terraform-provider-corefunc/issues/62)) (dependabot[bot])
* [`85864cd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/85864cd8b4ae635c548f02c9e7372e245e3457f1): Added editorconfig-checker as a Homebrew dependency.
* [`668be53`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/668be53fb2b9a2565a3dfcf2d5939bf9518031e4): Update the Go dependencies.
* [`f68f8f0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f68f8f00fef79acf3c7db82d81a0f7582f01616f): **deps**: Bump `github/codeql-action` from 2.22.6 to 2.22.7 ([#65](https://github.com/northwood-labs/terraform-provider-corefunc/issues/65)) (dependabot[bot])
* [`47ff449`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/47ff449b6e55e7406e1fde01f51b8170caebbe51): **deps**: Bump `step-security/harden-runner` from 2.6.0 to 2.6.1 ([#64](https://github.com/northwood-labs/terraform-provider-corefunc/issues/64)) (dependabot[bot])
* [`63775bf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/63775bf0368919d4f0ef403abd3293650f46fb47): Replaced much of Licensei with Trivy for licenses. ([#66](https://github.com/northwood-labs/terraform-provider-corefunc/issues/66))
* [`e0b5924`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e0b59240e4a8dce4f86499b668b038e8f478d55f): Updated license cache data.
* [`33557b7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/33557b7ca19fce19b28ded1b97c4cb0d62a3734b): Added git-cliff as a dependency for changelogs.
* [`d3d344e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d3d344e370a173613d7cab7bfa80a832c2215707): Remove default.pgo.
* [`b267f1e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b267f1ee74b109510ceb38ef0f0e9e9a49f88d59): Prepping a new test tag to trigger the build.
* [`5d52625`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5d52625a44197115de7e90b6a2bf3b9c14f901d8): Prepping a new test tag to trigger the build.
* [`dee279e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/dee279e1b75b230716ccc17517cd41c3b5d6e4e4): Prepping a new test tag to trigger the build.

### Documentation

* [`e9e2ad6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e9e2ad680320dedd9449bdce46343910399a7830): **terraform**: Updated the documentation.
* [`cf79e84`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cf79e84019cbcb7c854a1dbef4dfd751d85fb61f): **terraform**: More documentation edits.
* [`f4341db`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4341dba1db6bdbd3bdef234fdc81f08a41ddcb7): Updated the Contribution Guide.
* [`540584d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/540584ddf3f7fc6f0c420c237e79fba4347b1350): Updated the Go package documentation.
* [`4e9c027`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4e9c027a434c98501d9b7527336008da0eb9947f): Preparing support for the Terraform Registry.
* [`f4278de`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4278de00eac06cffb5d4b6112784312b3897f83): Added an example to `TruncateLabel()`.
* [`822bdf9`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/822bdf9d6b966034a00a76218690c3c569a1fcee): Updated the output of a make task.
* [`c12ccfc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c12ccfcf4f0209997a9e4f750ae66db3d655775d): Update the formatting of the provider examples.
* [`56b682a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/56b682aa6fc8143e874b734fabc07d82aa6cf9ec): Mark a paragraph as a note in the Terraform Registry docs.
* [`b2a5e4e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b2a5e4e82fcd75c094e6b9650672fd37e7089a4b): Update the list of planned features.
* [`16f8552`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/16f8552bff32fd211c3f90076c7e897e92050889): **terraform**: Updated examples.
* [`52efecc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/52efecc177335987ff16113c7cfe1b88f8ffa848): **readme**: Update `README.md`
* [`b163ca6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b163ca6f8086afc3582a443b7252c7d377af533d): Removed the list of to-dos from the `README.md`.
* [`a5b1aab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a5b1aab169ce6338525810178bc4670bd14a06ff): **provider**: Split the examples for env_ensure into separate code blocks.
* [`8670d50`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/8670d50ba1961e6a95e0d415e7ad4b964e1637b7): **provider**: Split the examples for env_ensure into separate code blocks.
* [`e204b4a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e204b4ac6f53579ae5eec37a624308d2f6a5c540): Updated `corefunc_str_snake` docs.
* [`918eb60`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/918eb60b43194bc93dbc72987e4c8879fc770eb5): Updated the provider example.
* [`6664868`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/666486822d8d220654c028ae7c739ba7ceb757b7): Clarified the `CONTRIBUTING.md` guide.
* [`ff5fa12`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ff5fa1291e92e13ed621db7d93877b57a2202683): Added Git LFS to the documentation.
* [`7325224`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/73252242e32ec5113248a2063584c8b6fda80fe9): Updated the `README.md`.
* [`c1e0640`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c1e064071c814286baa79dd9fe9be24a3e5921ec): Added documentation about BATS tests.
* [`0e1f3ec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0e1f3ecdcd4293ea263cafc20666071f432284af): Define a  security policy.
* [`ba01f24`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/ba01f2403d3892fc6d0e59391b8cad364c49061a): Renamed OpenTF → OpenTofu in the `README.md`.
* [`81dfde8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/81dfde8d7b22940fe4543d83b8de668d1a05ae4d): Add badges to `README.md`.
* [`5b173d1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5b173d15cea08cbc091738dca21d7a2201886171): Updated the generator docs.
* [`5a014fb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5a014fb0ec891a3ef782aa36c2ee38d28d6c48c4): Updated the `README.md` and `AUTHORS.md`.
* [`42e1e7d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/42e1e7d7f9274250178e71434e791c659a166802): Fixed issue link to `pkg.go.dev`.
* [`be02409`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/be02409405351f0619bba8f165510c2bb1acd356): Added link to chanced/caps documentation.
* [`1e7c19b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/1e7c19bbca8de371380a52d4ffcdcfa65786d7dd): Removing the Markdown version of the license.
* [`0be7a78`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0be7a78e18bd718a1eee181d8050de988f4a65e4): Trying to get both `pkg.go.dev` and GitHub to understand the license.
* [`21deae6`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/21deae652b7d6839c9c1069f019d3873308cf1da): Messing with the license again.

### Features

* [`d4a065d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d4a065d012cb4ad546054c774d614c6a3d5582e8): Initial commit of code. We have one function working.
* [`2f0a4ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2f0a4ae04d1b32b78c980dc4df03976ddb9abcbb): Added the Go code for the `EnvEnsure()` function.
* [`b69b67b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b69b67bff73b68fc4b389896ae3d48a712ab5612): Added the Terraform wrapper and docs for `EnvEnsure()`.
* [`5c75b92`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/5c75b922b13185dfee646b7bea6a7f839ed96703): Significant improvements to the Makefile.
* [`6f96ddf`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6f96ddf400090026fc38561323e5564d6776ab8e): Add support for regex patterns to `EnvEnsure()`.
* [`7f2e3b8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7f2e3b8de1e9875b2b32eea2a3f838d5b8ef6d50): Simplify Makefile by reading build data.
* [`27060e4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/27060e41a312bb2b1cc9cfd0850cfc443beca198): Adding some initial checking for conventional commits.
* [`3228245`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/32282453ad3e2d3f58d85d3623c9ed1d728f9f19): Initial attempt at a devcontainer.
* [`b763957`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b763957d7f4412de9f9258783fef50a5afc14fe1): Added `corefunc_str_snake`.
* [`3065ca0`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3065ca02b0331733eb417cf3a42a55229fa528a9): Added `corefunc_str_camel`.
* [`690c7f8`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/690c7f84d6f0d67f7cad4123fc7f90ac4669e887): Added `corefunc_str_pascal`.
* [`67c33fd`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/67c33fd2058f46f16568db3f15b5ca7222460857): Added `corefunc_str_kebab`.
* [`c6f0a0d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c6f0a0d243f450292bf82e4e611d9154e2906b71): Added `corefunc_str_constant`.
* [`02db190`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/02db1904110bf9b0950e97d63cc7034a73620056): Added code generation for new provider functions.
* [`194dde1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/194dde1db6bec67f8c4c82273071a4c73729c41d): Implemented support for `str_iterative_replace`. ([#31](https://github.com/northwood-labs/terraform-provider-corefunc/issues/31))
* [`280dc21`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/280dc217f860709fbeb82e20e7657cd63e5c91c3): Modernize issue creation.
* [`edd5210`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/edd521023634afd33a9efd38a137bea368aec392): Create issue template for bug reports.
* [`9080b41`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/9080b417f2cb4713da83f5e83578a078a5c5e49d): Added a template for feature requests.
* [`261aa0d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/261aa0dcab848277960a1a02e1231096d788d569): Update `feature.yml`
* [`65fa14c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/65fa14cd32b098464d28704b8a3ba88b0b359b67): Added generation for binsize.
* [`0e21531`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/0e21531d2ac18bfb0adfe36cb13e95f602dfca5d): Improve code coverage reporting.

### Linting

* [`6770816`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6770816190f2262d71c4f0f482c43ed59b8e09b2): Updated the linter definitions.
* [`83a1a67`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/83a1a674f6930ede51c3703735a4ea707904b4c6): Cleaned up the YAML files.
* [`a4085ee`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a4085ee2139ba9e737ec16791be7d07b634c79f9): Cleaned up the JSON files.
* [`cc22b71`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/cc22b718bf9f8e52cfd3039463735e19ebc8dab9): Cleaned up the Markdown files.
* [`e2bc42a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e2bc42a0328d5300fd6537c42668591277ded90d): Updated the Markdownlint rules.
* [`f33108f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f33108fdfeffe97d4efead1d87331fc2d3bbc65c): Removed unnecessary pre-commit rules.
* [`14356ab`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/14356abb4e1cfa1a5ed0934857c2e172ff8f439d): Ran the linter and performed a little cleanup.
* [`f4c2504`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f4c2504d3e31b70c2e800b3e8fd7f672ac35dc8a): Run yamlfmt.
* [`778d4cb`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/778d4cb0e637a6b898e5e232348a2f77e993ad41): Run yamlfmt.
* [`de6a00b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/de6a00bbe308feab45c5e2630dfe105a87de3462): Added and applied more linting.
* [`58555a5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/58555a5f351e6b65c773205a7f7eac574df9fa00): Resolved conflicts in linting rules.
* [`2be9512`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2be9512991c7ee4638a5c8f0a48bcee815e89a92): Add CodeQL analysis.
* [`15321a7`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/15321a74b62d07a2aa72d5198bb88c02deca9759): Improve CodeQL analysis.
* [`f941cf1`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f941cf1a6ca38bcc987298e1713b444944037b86): Remove CodeQL analysis YAML. Rely on built-in feature.
* [`31b5e9a`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/31b5e9a3497668bc7405af51223dc6cf3eb83c48): Additional linting.
* [`f0fe79b`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f0fe79b289b4d613bf2d9b4b3c5a6ac56029f143): Additional linting.
* [`58c6d30`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/58c6d308a1d334d02223451707fc7521cb6121cc): Fix Go version.
* [`480862d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/480862d2f01247acb8512d39f5b36f62315614be): Update PR review config.
* [`e5e2328`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e5e2328833d15644869aa78be79d5223b9c00608): Move test running into CI.
* [`f334dc4`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f334dc40d54c7056ca371fb69b38a911d90b6300): Verified that all files adhere to .editorconfig.

### Miscellaneous Tasks

* [`7e8e142`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/7e8e14253539a76669a1c251ec6f0d7918808db4): Cleaned up some mis-named jobs.
* [`993304e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/993304ed23f234cf5959c1ff236f9287180b07ea): Cleaned up some mis-named jobs.
* [`d1276b3`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d1276b3ffac3ab3d93d031d63f88c42de9393983): Cleaned up some mis-named jobs.
* [`a1dd8ef`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a1dd8efc0aef98d1269e4e3beccd709cfcec9003): Update the GoReleaser configuration.
* [`07dc031`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/07dc0315da588063861b54862b72e78d8f10fa74): Testing GoReleaser config changes.
* [`e61c254`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/e61c254b8008b5b9a8c2e18f7c317024a92be0e7): Update CHANGELOG configuration for GoReleaser.
* [`6901616`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/6901616b298e5be1bf159d4d6f7c54e6ec339894): Run GoReleaser on Git tag.
* [`b6447ec`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b6447ec9b6848f35c1d0449fe305bc2bc3d6afb9): Improving the security of the workflow.
* [`eb9f7ae`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/eb9f7aecfdaa7bda9087b9b58229d3943571a863): Update the GoReleaser configuration.
* [`a01b12e`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/a01b12e6870b821f561624bca79407e4c87c73cf): Trying to get the build to work.

### Refactor

* [`d24133f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d24133f28f580002231227dc247daa9a09a189f6): Fix naming of functions and data sources.
* [`b946b69`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/b946b693e5728b1cbd0d272485f2db880b2db6e5): Split the library code from the provider code.
* [`df2690f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/df2690f01db54e17ae8bbcb547168e133d336303): **make**: Split pre-commit out from lint.
* [`3c7a078`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/3c7a078b6376d8b9c76796e52949e3d7c445eb24): Remove line from Makefile; add it to GoReleaser later.

### Styling

* [`d451600`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d451600b1b601c242979cae99f7c915f86742ec4): Minor adjustments to the Makefile and author script.

### Testing

* [`c80d484`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/c80d484232cdd54545f7aa926c0bb6fd55cf0905): Add a debug mode for acc tests.
* [`15662fc`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/15662fc5ecb07db57eec6d7fbbe35866b1eb005a): Added mutation tests.
* [`4546390`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/4546390ca53c108cb424614c10ef1d493cb601db): Fix the tests running in CI.
* [`2100e0c`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/2100e0cf246193cdd808f1c7c32fa584894c8e21): Fix the tests running in CI.
* [`650e1c5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/650e1c5a89f284124d4306bb01dfd36210f8a094): Fix the tests running in CI.
* [`d7df64f`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d7df64f8533221b01b98feafe48cc6d1f190e71e): Fix the tests running in CI.
* [`d54db2d`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/d54db2d14880d822348a9f098828755f38bf5adb): Manually install latest Go for govulncheck.
* [`f965fbe`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/f965fbe6a0246138cf8d80e50ea705a54ddde548): Manually install latest Go for govulncheck.
* [`78966d5`](https://github.com/northwood-labs/terraform-provider-corefunc/commit/78966d5edc0e3c57f7ab86827f8d3595b2442072): Allow `storage.googleapis.com:443` from GitHub Actions.
