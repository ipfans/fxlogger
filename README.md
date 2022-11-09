# fxlogger

[![MIT License](https://img.shields.io/github/license/ipfans/fxlogger)](https://github.com/ipfans/fxlogger/blob/main/LICENSE)
[![GoDoc](http://godoc.org/github.com/ipfans/fxlogger?status.svg)](http://godoc.org/github.com/ipfans/fxlogger)
[![GitHub release](https://img.shields.io/github/v/release/ipfans/fxlogger)](https://github.com/ipfans/fxlogger/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ipfans/fxlogger)](https://github.com/ipfans/fxlogger/blob/main/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/ipfans/fxlogger)](https://goreportcard.com/report/github.com/ipfans/fxlogger)
[![Issues](https://img.shields.io/github/issues/ipfans/fxlogger)](https://github.com/ipfans/fxlogger/issues)

fxevent zerolog adoptor for go-uber/fx 1.14.0+.

## Installation

```
go get github.com/ipfans/fxlogger@latest
```

## Usage

```go

fx.New(fx.WithLogger(fxlogger.Default()))

// or

fx.New(fx.WithLogger(fxlogger.WithZerolog(zerolog.New(os.Stdout))))
```

## Contributors

<a href="https://github.com/ipfans/fxlogger/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ipfans/fxlogger" />
</a>
