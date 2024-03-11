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

<p>Generated on {{ now() | date(format="%Y-%m-%d") }}.</p>
