# Shuff

Shuff is a simple Go-based command-line tool that shuffles lines in the input,
with support for files or standard input. It is created for educational purposes
as an alternative to the Unix `shuf` command, with a simplified interface and
implementation. Shuff can be installed easily with Go and used for various
shuffling tasks in scripts or command pipelines.

## Installation

To install `shuff`, you can use the following command in your terminal:

```bash
go install github.com/kamilturek/shuff/cmd/shuff@latest
```

This command will download the latest version of the shuff tool from this
repository, build an executable binary, and install it on your system.

Note that you need to have Go installed on your system for this command to work.
If you don't have Go installed, you can download it from the official website at
[https://go.dev/](https://go.dev/).

## Usage

The basic usage of `shuff` is:

```bash
shuff [FILE]
```

where `FILE` is the input file to shuffle. If no `FILE` is specified, `shuff` reads
from the standard input.

Here are some example usages of `shuff`:

```bash
cat file.txt | shuff
```

This command reads the contents of file.txt from the standard input and shuffles
the lines.

```bash
shuff < file.txt
```

This command redirects the contents of `file.txt` to the standard input of `shuff`
and shuffles the lines.

```bash
shuff file.txt
```

This command reads the contents of `file.txt` from the file and shuffles the lines.

## Contributing

Contributions to `shuff` are welcome! If you find a bug, have an idea for a
feature, or want to improve the documentation, feel free to submit a pull
request or open an issue.

## License

`shuff` is licensed under the MIT License. See [LICENSE](./LICENSE.md) for details.
