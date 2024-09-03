# go_crawler

- Clone the repo
- Run `go build`
- Run the following command `./go_crawler <base url> <maxConcurrency> <maxPages>` with base url being the url you'd like to crawl, maxConcurrency being the number of concurrent goroutines you'd like to run, and maxPages being the maximum number of pages you'd like to crawl in the specified base URL.

## Example

`./go_crawler https://www.google.com 1 100`

## Output

```
====================================================>
REPORT for https://www.google.com
====================================================>
Found 10 internal links to https://www.google.com/intl/en/about/
Found 10 internal links to https://www.google.com/intl/en/about/company/
Found 10 internal links to https://www.google.com/intl/en/about/products/
Found 10 internal links to https://www.google.com/intl/en/about/careers/
Found 10 internal links to https://www.google.com/intl/en/about/locations/
Found 10 internal links to https://www.google.com/intl/en/about/contact/
Found 10 internal links to https://www.google.com/intl/en/about/press/
Found 10 internal links to https://www.google.com/intl/en/about/careers/students/
Found 10 internal links to https://www.google.com/intl/en/about/careers/locations/
Found 10 internal links to https
```
