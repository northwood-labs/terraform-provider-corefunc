# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com), adheres to
[Semantic Versioning](https://semver.org), and uses
[Conventional Commit](https://www.conventionalcommits.org) syntax.

## Unreleased

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
