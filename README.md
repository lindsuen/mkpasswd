# mkpasswd

A command line tool for generating random passwords.

## Features

- Command line with parameters.
- Support for generating multiple passwords.
- Arabic numerals supported `0-9`.
- Special characters supported `!@#$%^&*+?`.
- `Makefile` is supported.

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
go build -o mkpasswd.exe .
```

```sh
# Linux or MacOS or FreeBSD
go build -o mkpasswd .
```

Build for all mainstream platforms. Please see the `Makefile` for details.

```sh
make all
```

## License

GPL-3.0 license
