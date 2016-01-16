# cc

[![GoDoc](https://godoc.org/gpk.io/cc?status.svg)](https://godoc.org/gpk.io/cc)

Colorize terminal output.

```golang
package main

import (
	"log"
	"gpk.io/cc"
)

func main() {
	log.Printf("Enabled: %v", cc.Enabled())
	log.Printf(cc.Red("red") + " and " + cc.Blue("blue"))
	log.Printf("%s - %s", cc.BgWhite(cc.Blue("foo")), cc.BgRed(cc.Yellow("bar")))
	log.Printf(cc.Underline(cc.Bold("bold-underline")))
}
```

## Features
- Clean and simple API
- Windows support
- Auto-detect coloring support

## Install

```bash
go get gpk.io/cc.v1
```
