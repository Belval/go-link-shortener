# go-link-shortener
A basic link shortening service written in go. Please consider as a work-in-progress.

## How-To

1. Edit the config.json file
2. Run `./main`
3. Shorten an url with `/api/v1/new/?url=[your.url.com]`
    - Don't forget to add & escape the https:// part if you really want it (otherwise it will default to http://)
4. Use the returned address

## Requirements

- https://github.com/google/uuid
- https://github.com/mattn/go-sqlite3

