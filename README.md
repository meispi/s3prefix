# s3prefix

A golang tool inspired by [lazys3](https://github.com/nahamsec/lazys3) by [Nahamsec](https://twitter.com/nahamsec).

## Installation

```
go install github.com/meispi/s3prefix@latest
```
Create a directory `.config` on your root (or home) directory and copy the file `common_bucket_prefixes.txt` in that directory.
Go to your root directory:
```
cd
```

Check if you already have a `.config` directory
```
ls -a
```

If no then make a directory else skip this step:
```
mkdir .config
```

Go to `.config`
```
cd .config
```

Download the `common_bucket_prefixes.txt` on this directory:
```
curl https://raw.githubusercontent.com/meispi/s3prefix/main/common_bucket_prefixes.txt -o common_bucket_prefixes.txt
```

And now you are good to go!

## How to use

Make sure you have copied the file `common_bucket_prefixes.txt` in `.config` directory in your root (or home) directory
```
s3prefix <target_name>
```