gokc
====
[![Latest Version](http://img.shields.io/github/release/yuuki1/gokc.svg?style=flat-square)][release]

[release]: https://github.com/yuuki1/gokc/releases

Yet Another [Keepalived](http://keepalived.org/) syntax checker in golang.

## Usage

```bash
gokc -f /etc/keepalived/keepalived.conf
gokc: the configuration file /etc/keepalived/keepalived.conf syntax is ok
```

## Installation

### Homebrew
```bash
$ brew tap yuuki1/gokc
$ brew install gokc
```

### Download binary from GitHub Releases
[Releases・yuuki1/gokc - GitHub](https://github.com/yuuki1/gokc/releases)

### Build from source
```bash
 $ go get github.com/yuuki1/gokc
 $ go install github.com/yuuki1/gokc/...
```

## Release

```bash
$ vim ./version.go
$ make cross
$ ghr -u yuuki1 -p 2 $VERSION snapshot/
```

## Contribution

1. Fork ([https://github.com/y_uuki/gokc/fork](https://github.com/y_uuki/gokc/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## License

[The MIT License](./LICENSE).

## Author

[y_uuki](https://github.com/yuuki1)
