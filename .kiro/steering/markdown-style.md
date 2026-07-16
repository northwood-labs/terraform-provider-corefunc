---
inclusion: fileMatch
fileMatchPattern: "**/*.md"
---

# Markdown Style Guide

This project enforces Markdown conventions via `.rumdl.toml` and `.editorconfig`. Follow these rules when writing or editing any `.md` file.

## Indentation and whitespace

<!-- @config-manager:start whitespace -->
* Use 2-space indentation for Markdown files (per `.editorconfig`).
* Use LF line endings, no trailing whitespace.
* End every file with a single trailing newline.
* Maximum one consecutive blank line (no double blanks).
* One blank line above and below headings, lists, tables, fenced code blocks, and horizontal rules.
<!-- @config-manager:end whitespace -->

## Headings

<!-- @config-manager:start headings -->
* Use ATX style (`#` prefix), not Setext (underlines).
* Start the document with a single `# H1` heading. Only one H1 per file.
* Increment heading levels by one — no skipping (e.g., `##` → `####` is invalid).
* Headings must start at column 0 (no indentation).
* No trailing punctuation (`.`, `,`, `;`, `:`) on headings.
* One space after `#` — no multiple spaces.
* The `# H1` heading should use Title Case formatting (conjunctions and other connector words should always be lowercase — `but`, `and`, `or`, `to`, `from`).
* All other headings (`## H2`, `### H3`, `#### H4`, `##### H5`, `###### H6`) should use sentence case formatting.
<!-- @config-manager:end headings -->

## Emphasis and strong

<!-- @config-manager:start em-strong -->
* Use _underscores_ for emphasis (italic): `_text_`, not `*text*`.
* Use **asterisks** for strong (bold): `**text**`, not `__text__`.
* No spaces inside emphasis or strong markers.
<!-- @config-manager:end em-strong -->

## Lists

<!-- @config-manager:start lists -->
* Use `*` (asterisk) for unordered list markers, not `-` or `+`.
* Use ordered numbering for ordered lists (`1.`, `2.`, `3.`), not all-ones.
* One space after the list marker.
* Indent nested unordered lists by 2 spaces (matching `ul-indent` config).
* Blank line before and after list blocks.
<!-- @config-manager:end lists -->

## Code

<!-- @config-manager:start code-pre -->
* Use fenced code blocks (triple backticks), not indented blocks.
* Use backtick fences, not tilde fences.
* Always specify a language identifier on fenced code blocks (e.g., ` ```go `, ` ```text `, ` ```bash `).
* Use ` ```text ` for plain-text diagrams, pipeline flows, and non-executable pseudocode.
* No spaces inside inline code spans.
<!-- @config-manager:end code-pre -->

## Tables

<!-- @config-manager:start tables -->
* Use leading and trailing pipes on every row.
* Align column delimiters (pad cells with spaces so pipes line up).
* Blank line before and after tables.

Example:

```markdown
| Column A | Column B |
|----------|----------|
| value    | value    |
```

<!-- @config-manager:end tables -->

## Links and images

<!-- @config-manager:start links-images -->
* Use descriptive link text. Prohibited: "click here", "here", "link", "more".
* No bare URLs — wrap in angle brackets or use `[text](url)` syntax.
* No reversed link syntax (`(url)[text]`).
* No spaces inside link brackets or parentheses.
* All images must have alt text.
<!-- @config-manager:end links-images -->

## Proper names

<!-- @config-manager:start proper-names -->
These terms must use exact casing wherever they appear in prose (not enforced inside code blocks or HTML):

* Commonmark
* CSS
* GFM
* GitHub-Flavored Markdown
* HTML
* JSON
* macOS
* Markdown
* Northwood Labs
* OpenTofu
* Terraform
* Terragrunt
* Terratest
* TOML
* XML
* YAML
<!-- @config-manager:end proper-names -->

## HTML

<!-- @config-manager:start htmltags -->
Inline HTML is restricted to these allowed elements: `a`, `b`, `br`, `code`, `details`, `div`, `img`, `li`, `nobr`, `ol`, `p`, `pre`, `span`, `summary`, `ul`. All other HTML elements will trigger a lint warning.
<!-- @config-manager:end htmltags -->

## Blockquotes

<!-- @config-manager:start blockquotes -->
* No blank lines inside blockquotes.
* No multiple spaces after `>`.
* Consecutive blockquotes should be separated Horizontal Rule.
<!-- @config-manager:end blockquotes -->

## Horizontal rules

<!-- @config-manager:start hr -->
* Use the `---` style.
* Should be used sparingly.
* Do not use them immediately before headings.
<!-- @config-manager:end hr -->

## Footnotes

<!-- @config-manager:start footnotes -->
* All footnote references must have a corresponding definition.
* Footnote definitions must appear in the order they are referenced.
* No empty footnote definitions.
<!-- @config-manager:end footnotes -->

## Front matter

<!-- @config-manager:start frontmatter -->
* Include a blank line after YAML front matter closing `---`.
<!-- @config-manager:end frontmatter -->
