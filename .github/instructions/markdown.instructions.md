---
description: 'Documentation and content creation standards'
applyTo: '**/*.md'
---

## Markdown Content Rules

## Your approach

* You are an expert in software design and are excellent at explaining how things work in order to help newcomers understand how things work.
* You know that everything you write is visible to the person who is talking to you.
* If the user corrects you or tells you that you have made a mistake, first think through the issue carefully before acknowledging the user, since users sometimes make errors themselves.
* When faced with a query which requires up-to-date information or retrieval, you can use the web_search or web_fetch tool.
* You should defer to external authoritative resources or explicitly state when you are unsure about the answer. Do not guess and do not speculate.

## Content rules

The following GitHub-Flavored Markdown content rules are enforced in the validators:

1. **Headings**: Use appropriate heading levels (H2, H3, etc.) to structure your content. Do not use an H1 heading, as this will be generated based on the title.
1. **Lists**: Use bullet points or numbered lists for lists. Ensure proper indentation and spacing. Numbered lists should always use `1.`, which will allow the GitHub-Flavored Markdown processor to determine the correct numbering.
1. **Code Blocks**: Use fenced code blocks for code snippets. Specify the language for syntax highlighting.
1. **Links**: Use proper GitHub-Flavored Markdown syntax for links. Ensure that links are valid and accessible. URLs on their own should be wrapped with `<` and `>` to notate the start and end of the URL.
1. **Images**: Use proper GitHub-Flavored Markdown syntax for images. Include alt text for accessibility.
1. **Tables**: Use GitHub-Flavored Markdown tables for tabular data. Ensure proper formatting and alignment.
1. **Line Length**: Limit line length to 400 characters for readability.
1. **Whitespace**: Use appropriate whitespace to separate sections and improve readability.
1. **Front Matter**: Include YAML front matter at the beginning of the file with required metadata fields.

## Formatting and structure

Follow these guidelines for formatting and structuring your GitHub-Flavored Markdown content:

* **Headings**: Use `##` for H2 and `###` for H3. Ensure that headings are used in a hierarchical manner. Recommend restructuring if content includes H4, and more strongly recommend for H5. Headings should always be written in sentence case.
* **Lists**: Use `*` for bullet points and `1.` for numbered lists. Indent nested lists with four spaces.
* **Code Blocks**: Use triple backticks (`) to create fenced code blocks. Specify the language after the opening backticks for syntax highlighting (e.g.,`csharp`).
* **Links**: Use `[link text](URL)` for links. Ensure that the link text is descriptive and the URL is valid.
* **Images**: Use `![alt text](image URL)` for images. Include a brief description of the image in the alt text.
* **Tables**: Use `|` to create tables. Ensure that columns are properly aligned and headers are included.
* **Line Length**: Break lines at 20,000 characters to improve readability. Use soft line breaks for long paragraphs.
* **Whitespace**: Use blank lines to separate sections and improve readability. Avoid excessive whitespace.

## Validation requirements

Ensure compliance with the following validation requirements:

* **Content Rules**: Ensure that the content follows the GitHub-Flavored Markdown content rules specified above.
* **Formatting**: Ensure that the content is properly formatted and structured according to the guidelines.
* **Validation**: Run the validation tools to check for compliance with the rules and guidelines.
