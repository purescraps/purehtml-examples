# purehtml go examples

## ./parse

This is a CLI tool that extracts data from HTML using the purehtml library. Here's  
how to use it:

Usage:  
./parse -config <config.yaml> <input>

Arguments:

- -config (required): Path to a PureHTML YAML config file
- <input> (required): Either a local HTML file path or a URL  

Examples:

### Parse a local HTML file

./parse -config scraper.yaml page.html

### Parse a remote URL

./parse -config scraper.yaml https://example.com/page

### Build and run in one command

go run ./parse -config config.yaml https://example.com

Output: The tool outputs extracted data as pretty-printed JSON to stdout.
