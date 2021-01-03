# gosetlocale

A simple [setlocale(3)](https://linux.die.net/man/3/setlocale) wrapper for Go.

## Table of Contents 

* [Install](#install)
* [Example](#example)
* [Build](#build)
* [License](#license)

## Install

To add this package in your project, run the following:

```
go get github.com/vsakkas/gosetlocale
```

## Usage

This package contains only one function, `SetLocale`, which receives two arguments, the locale and the category.

The available locale values can be found in the [gosetlocale.go](gosetlocale.go) file.


## Example

An example usage of this package is the following:

```go
package main

import (
    sl "github.com/vsakkas/gosetlocale"
)

func main() {
    // Make program portable to all locales.
    sl.SetLocale(sl.LcAll, "")
}
```

## Build

To build this project locally, first clone it:

```
git clone git@github.com:vsakkas/gosetlocale.git
```

Then, build it using the following command:

```
go build
```

It is recommended to use Go 1.15 or newer.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
