gokc
====
[![Build Status](https://travis-ci.org/yuuki/diamondb.svg?branch=master)](https://travis-ci.org/yuuki/diamondb)
[![Latest Version](http://img.shields.io/github/release/yuuki/gokc.svg?style=flat-square)](https://github.com/yuuki/gokc/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuuki/gokc)](https://goreportcard.com/report/github.com/yuuki/gokc)

Yet Another [Keepalived](http://keepalived.org/) syntax checker in golang.

## Usage

```bash
gokc -f /etc/keepalived/keepalived.conf
gokc: the configuration file /etc/keepalived/keepalived.conf syntax is ok
```

## Installation

### Homebrew
```bash
$ brew tap yuuki/gokc
$ brew install gokc
```

### Download binary from GitHub Releases
[Releases・yuuki/gokc - GitHub](https://github.com/yuuki/gokc/releases)

### Build from source
```bash
 $ go get github.com/yuuki/gokc
 $ go install github.com/yuuki/gokc
```

## Contribution

1. Fork ([https://github.com/yuuki/gokc/fork](https://github.com/yuuki/gokc/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## License

[The MIT License](./LICENSE).

## Author

[y_uuki](https://github.com/yuuki)
