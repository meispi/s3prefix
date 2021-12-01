# s3prefix

A golang tool inspired by [lazys3](https://github.com/nahamsec/lazys3) by [Nahamsec](https://twitter.com/nahamsec).

## Installation

```
go install github.com/meispi/s3prefix@latest
```
Create a directory `.config` on your root (or home) directory and copy the file `common_bucket_prefixes.txt` in that directory.

## How to use

Make sure you have copied the file `common_bucket_prefixes.txt` in `.config` directory in your root (or home) directory
```
s3prefix <target_name>
```