# mkpasswd

[![build](https://github.com/lindsuen/mkpasswd/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/lindsuen/mkpasswd/actions/workflows/build.yml)
![GitHub License](https://img.shields.io/github/license/lindsuen/mkpasswd)

A command line tool for generating random passwords.

## Features

- Support command line parameters.
- Support generating multiple passwords.
- Support Arabic numerals `0-9`.
- Support special characters `!@#$%^&*+?`.

## Usage

The command includes four optional parameters. You can type `mkpasswd -h` for help.

For example:

```sh
mkpasswd -N 5 -l 16 -n 4 -c 4
```

- `-N`: The quantity of created passwords.
- `-l`: The length of password.
- `-n`: The number of Arabic numerals in password.
- `-c`: The number of special characters in password.

## Build

Build separately in your operating system.

```sh
# Windows
go build -o mkpasswd.exe -ldflags "-s -w" .
```

```sh
# Linux or MacOS or FreeBSD
go build -o mkpasswd -ldflags "-s -w" .
```

Build for all mainstream platforms. Please see the `Makefile` for details.

```sh
make all
```

## License

GPL-3.0 license
