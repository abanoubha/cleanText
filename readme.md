# cleanText

A simple program to clean up text. The goal of the program is to accept text from standard input then doing the clean process then output the result onto the standard output.

Basic usage:

```sh
$ echo "àbanôūb" | cleanText
abanoub
```

## Commands

```sh
# build
go mod tidy && go build -o cleantext .

# run the program
./cleantext -h

# detect race conditions & memory leaks
go run -race .
```

## Resources & references

- <https://www.charset.org/utf-8>

## Converted symbols

See source code for now. I will document them later.

## Kept symbols

symbols that are kept after cleaning

| symbol | meaning         |
|:------:|:----------------|
| ±      | Plus-minus Sign |
| °      | degree sign     |

## Skipped symbols

See source code for now. I will document them later.

## Further development

Let me know if you have any suggestion and/or problems by writing an [issue](https://github.com/abanoubha/cleanText/issues). If you have an addition or feature or a fix, make a [pull request](https://github.com/abanoubha/cleanText/pulls).
